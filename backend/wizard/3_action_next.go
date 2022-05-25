package wizard

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/executors/backend/wizard/lifecycle"
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

	currstage := sw.model.Stages[subData.CurrentStage]
	if currstage == nil {
		return nil, easyerr.NotFound()
	}

	currgroup := sw.getStageGroup(subData.StageGroup)
	if currgroup == nil {
		return nil, easyerr.NotFound()
	}

	var nstage string
	var skipChecks []string

	if currgroup.BeforeNext != "" {
		lf := lifecycle.BeforeNext{
			CurrentData: req.Data,
			SubData:     subData,
			SideEffects: lifecycle.BeforeNextSideEffects{},
		}
		err := lf.Execute()
		if err != nil {
			return nil, err
		}
		if len(lf.SideEffects.Errors) > 0 {
			return wmodels.ResponseNext{
				Errors: lf.SideEffects.Errors,
				Ok:     false,
			}, nil
		}
		if lf.SideEffects.NextStage != "" {
			nstage = lf.SideEffects.NextStage
		}

		if lf.SideEffects.SkipCheck != nil {
			skipChecks = lf.SideEffects.SkipCheck
		}
	}

	if currstage.BeforeVerify != "" {

		lf := lifecycle.StageBeforeVerify{
			Models: &sw.model,
			SideEffects: lifecycle.StageBeforeVerifyEffect{
				SkipCheck: skipChecks,
			},
			SubData: subData,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}
		if lf.SideEffects.NextStage != "" {
			nstage = lf.SideEffects.NextStage
		}

		if len(lf.SideEffects.Errors) > 0 {
			return wmodels.ResponseNext{
				Errors: lf.SideEffects.Errors,
				Ok:     false,
			}, nil
		}

		if lf.SideEffects.SkipCheck != nil {
			skipChecks = lf.SideEffects.SkipCheck
		}

		if lf.SideEffects.NextStage != "" {
			nstage = lf.SideEffects.NextStage
		}
	}

	pp.Println(nstage)

	errors := make(map[string]string)
	for _, field := range currstage.Fields {
		if funk.ContainsString(skipChecks, field.Name) {
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

	{
		subData.Data[subData.CurrentStage] = req.Data
		subData.VisitedStages = append(subData.VisitedStages, subData.CurrentStage)
	}

	if nstage == "" {
		idx := sw.stageIndex(currgroup, currstage.Name)

		switch idx {
		case -1:
			pp.Println("@group", currgroup)
			pp.Println("@subdata", subData)
			// this should not happen
			return nil, easyerr.NotFound()
		case len(currgroup.Stages) - 1:
			return sw.endStageGroup(currgroup, subData)
		default:
			nstage = currgroup.Stages[idx+1]
			subData.CurrentStage = nstage
		}

	} else {
		if !funk.ContainsString(currgroup.Stages, nstage) {
			return nil, easyerr.NotFound()
		}
	}

	return sw.generate(subData, currgroup, nstage)
}

func (sw *SimpleWizard) generate(sub *wmodels.Submission, group *wmodels.StageGroup, nStage string) (interface{}, error) {
	datasources := make(map[string]interface{})

	stage := sw.model.Stages[nStage]
	if stage == nil {
		pp.Println("@next_stage", nStage)
		pp.Println("@all_stage", sw.model.Stages)
		return nil, easyerr.NotFound()
	}

	if stage.BeforeGenerate != "" {
		lf := lifecycle.StageBeforeGenerate{
			Models: &sw.model,
			SideEffects: lifecycle.StageBeforeGenerateEffects{
				DataSources: datasources,
			},
			SubData: sub,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (sw *SimpleWizard) endStageGroup(group *wmodels.StageGroup, subData *wmodels.Submission) (interface{}, error) {
	return nil, nil
}

/*


func (sw *SimpleWizard) generate(sub *wmodels.Submission, group *wmodels.StageGroup, nStage string) (interface{}, error) {


		nstage := sw.model.Stages[nStage]
		if nstage == nil {
			pp.Println("@next_stage", nStage)
			pp.Println("@all_stage", sw.model.Stages)
			return nil, easyerr.NotFound()
		}

		eerr := ""

		if group.BeforeNext != "" {
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

			err := sw.execScript(nstage.BeforeVerify, nil, binds)
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
			err := sw.execScript(nstage.BeforeVerify, nil, binds)
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


	return nil, nil

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

*/
