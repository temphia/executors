package wizard

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/executors/backend/wizard/wmodels"
	"github.com/thoas/go-funk"
)

func (sw *SimpleWizard) RunNext(ev *event.Request) (interface{}, error) {

	req := wmodels.RequestNext{}

	err := json.Unmarshal(ev.Data, &req)
	if err != nil {
		return nil, err
	}

	subData, err := sw.getSub(req.OpaqueData)
	if err != nil {
		return nil, err
	}

	stage := sw.model.Stages[subData.CurrentStage]
	if stage == nil {
		return nil, easyerr.NotFound()
	}

	nextStage := ""
	errors := make(map[string]string)
	skipChecks := make(map[string]struct{})
	eerr := ""

	if stage.BeforeValidate != "" {
		binds := map[string]interface{}{
			"_wizard_set_err": func(e string) {
				eerr = e
			},
			"_wizard_set_field_err": func(field, e string) {
				errors[field] = e
			},

			"_wizard_skip_field_check": func(field string) {
				skipChecks[field] = struct{}{}
			},

			"_wizard_set_next_stage": func(name string) {
				nextStage = name
			},
			"_wizard_set_shared_var": func(name string, data interface{}) {
				subData.SharedVars[name] = data
			},
			"_wizard_get_shared_var": func(name string) interface{} {
				return subData.SharedVars[name]
			},

			"_wizard_get_prev_data": func(name string) interface{} {
				return subData.Data[name]
			},
		}

		err := sw.execScript(stage.BeforeValidate, req.Data, binds)
		if err != nil {
			return nil, err
		}
	}

	if eerr != "" {
		return nil, easyerr.Error(eerr)
	}

	for _, field := range stage.Fields {
		if _, ok := skipChecks[field.Name]; ok {
			continue
		}

		ferr := sw.validateField(field, req.Data[field.Name])
		if ferr != "" {
			errors[field.Name] = ferr
		}
	}

	if len(errors) > 0 {
		return wmodels.ResponseNext{
			Errors: errors,
			Ok:     false,
		}, nil
	}

	_currentStage := subData.CurrentStage
	{
		subData.Data[subData.CurrentStage] = req.Data
		subData.VisitedStages = append(subData.VisitedStages, subData.CurrentStage)
	}

	group := sw.getStageGroup(subData.StageGroup)
	if group == nil {
		return nil, easyerr.NotFound()
	}

	if nextStage == "" {
		idx := sw.stageIndex(group, _currentStage)

		switch idx {
		case -1:
			pp.Println("@group", group)
			pp.Println("@subdata", subData)
			// this should not happen
			return nil, easyerr.NotFound()
		case len(group.Stages) - 1:
			return sw.endStageGroup(group, subData)
		default:
			nextStage = group.Stages[idx+1]
			subData.CurrentStage = nextStage
		}

	} else {
		if !funk.ContainsString(group.Stages, nextStage) {
			return nil, easyerr.NotFound()
		}
	}

	return nil, nil

}

func (sw *SimpleWizard) generate(sub *wmodels.Submission, group *wmodels.StageGroup, nStage string) (interface{}, error) {

	nstage := sw.model.Stages[nStage]
	if nstage == nil {
		pp.Println("@next_stage", nStage)
		pp.Println("@all_stage", sw.model.Stages)
		return nil, easyerr.NotFound()
	}

	eerr := ""

	if group.OnNext != "" {
		binds := map[string]interface{}{
			"_wizard_set_err": func(e string) {
				eerr = e
			},

			"_wizard_set_shared_var": func(name string, data interface{}) {
				sub.SharedVars[name] = data
			},
			"_wizard_get_shared_var": func(name string) interface{} {
				return sub.SharedVars[name]
			},
			"_wizard_get_stage_data": func(name string) interface{} {
				return sub.Data[name]
			},
		}

		err := sw.execScript(nstage.BeforeValidate, nil, binds)
		if err != nil {
			return nil, err
		}
	}

	pp.Println(eerr)

	resp := wmodels.ResponseNext{
		StageTitle:  nstage.Name,
		Fields:      nstage.Fields,
		DataSources: make(map[string]interface{}),
		Message:     nstage.Message,
		OpaqueData:  nil,
		Ok:          true,
		Final:       false,
	}

	if nstage.BeforeGenerate != "" {
		binds := map[string]interface{}{
			"_wizard_set_err": func(e string) {
				eerr = e
			},
			"_wizard_set_source": func(name string, data interface{}) {
				resp.DataSources[name] = data
			},
			"_wizard_set_shared_var": func(name string, data interface{}) {
				sub.SharedVars[name] = data
			},
			"_wizard_get_shared_var": func(name string) interface{} {
				return sub.SharedVars[name]
			},
			"_wizard_get_stage_data": func(name string) interface{} {
				return sub.Data[name]
			},
		}

		// fixme => pass proper ctx to method
		err := sw.execScript(nstage.BeforeValidate, nil, binds)
		if err != nil {
			return nil, err
		}
	}

	err := sw.genSource(nstage, sub, resp.DataSources)
	if err != nil {
		return nil, err
	}

	opdata, err := sw.updateSub(sub)
	if err != nil {
		return nil, err
	}

	resp.OpaqueData = opdata

	return resp, nil

}

func (sw *SimpleWizard) endStageGroup(group *wmodels.StageGroup, subData *wmodels.Submission) (interface{}, error) {
	// fixme => handle nested stage_group differently

	if group.BeforeEnd == "" {
		return wmodels.ResponseFinal{
			Ok:          true,
			LastMessage: group.LastMessage,
			Final:       true,
		}, nil
	}

	eerr := ""
	msg := group.LastMessage

	binds := map[string]interface{}{
		"_wizard_set_err": func(e string) {
			eerr = e
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			subData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return subData.SharedVars[name]
		},
		"_wizard_get_stage_data": func(name string) interface{} {
			return subData.Data[name]
		},
		"_wizard_set_message": func(m string) {
			msg = m
		},
	}

	// fixme => pass proper ctx to method
	err := sw.execScript(group.BeforeEnd, nil, binds)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	if eerr != "" {
		return wmodels.ResponseFinal{
			LastMessage: msg,
			Ok:          false,
			Final:       true,
		}, nil
	}

	return wmodels.ResponseFinal{
		Ok:          true,
		LastMessage: msg,
		Final:       true,
	}, nil
}
