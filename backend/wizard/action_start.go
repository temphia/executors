package simplewizard2

import (
	"encoding/json"
	"fmt"

	"github.com/rs/xid"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
)

func (sw *SimpleWizard) RunStart(ev *event.Request) (interface{}, error) {

	data := RequestStart{}
	err := json.Unmarshal(ev.Data, &data)
	if err != nil {
		return nil, err
	}

	nextStageGroup := ""
	nextStage := ""
	eerr := ""
	skipValidation := false

	if sw.model.Splash.BeforeValidate != "" {

		binds := map[string]interface{}{
			"_wizard_set_next_stage_group": func(name string) {
				nextStageGroup = name
			},
			"_wizard_set_next_stage": func(name string) {
				nextStage = name
			},
			"_wizard_set_err": func(name string) {
				eerr = name
			},
			"_wizard_skip_validation": func() {
				skipValidation = true
			},
		}

		err = sw.execScript(sw.model.Splash.BeforeValidate, data, binds)
		if err != nil {
			return nil, err
		}
	}

	if eerr != "" {
		return sw.GetSplash(ev, eerr)
	}

	if !skipValidation {
		for _, field := range sw.model.Splash.Fields {
			_, ok := data.Data[field.Name]
			if !ok && field.Optional {
				continue
			}
			// fixme actually check data
			if !ok {
				return sw.GetSplash(ev, fmt.Sprintf("Empty required field: %s", field.Name))
			}
		}
	}

	var sg *StageGroup
	if nextStageGroup != "" {
		for _, _sg := range sw.model.StageGroups {
			if _sg.Name == nextStageGroup {
				sg = &_sg
			}
		}
	} else {
		sg = &sw.model.StageGroups[0]
	}

	if sg == nil {
		ev.Data = nil
		return sw.GetSplash(ev, fmt.Sprintf("Stage Group not found: %s", nextStageGroup))
	}

	if sg.BeforeStart != "" {

		binds := map[string]interface{}{
			"_wizard_set_next_stage": func(name string) {
				nextStage = name
			},
			"_wizard_set_err": func(name string) {
				eerr = name
			},
		}

		err := sw.execScript(sg.BeforeStart, data.Data, binds)
		if err != nil {
			return nil, err
		}
	}

	if eerr != "" {
		ev.Data = nil
		return sw.GetSplash(ev, eerr)
	}

	var stage *Stage

	if nextStage != "" {
		_stage, ok := sw.model.Stages[nextStage]
		if !ok {
			return sw.GetSplash(ev, fmt.Sprintf("Stage not found: %s", nextStage))
		}
		stage = _stage
	} else {
		stage = sw.model.Stages[sg.Stages[0]]
	}

	subData := Submission{
		Id:               xid.New().String(),
		StageGroup:       sg.Name,
		CurrentStage:     stage.Name,
		Data:             make(map[string]map[string]interface{}),
		SharedVars:       make(map[string]interface{}),
		ParentStageGroup: "",
		ParentStage:      "",
		PrevStages:       []string{},
	}

	resp := &ResponseStart{
		StartStage:  true,
		StageTitle:  stage.Name,
		Message:     stage.Message,
		Fields:      stage.Fields,
		DataSources: make(map[string]interface{}),
		OpaqueData:  nil,
		Ok:          true,
	}

	if stage.BeforeGenerate != "" {
		binds := map[string]interface{}{
			"_wizard_err": func(e string) {
				eerr = e
			},
			"_wizard_set_source_data": func(source string, data interface{}) {
				resp.DataSources[source] = data
			},
		}

		err := sw.execScript(sg.BeforeStart, data.Data, binds)
		if err != nil {
			return nil, err
		}

		if eerr != "" {
			ev.Data = nil
			return sw.GetSplash(ev, eerr)
		}
	}

	err = sw.genSource(stage, &subData, resp.DataSources)
	if err != nil {
		return nil, err
	}

	opData, err := json.Marshal(&subData)
	if err != nil {
		return nil, err
	}

	resp.OpaqueData = opData

	return resp, nil
}
