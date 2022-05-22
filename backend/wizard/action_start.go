package wizard

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
)

func (sw *SimpleWizard) RunStart(ev *event.Request) (interface{}, error) {

	data1 := RequestStart{}
	err := json.Unmarshal(ev.Data, &data1)
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

		err = sw.execScript(sw.model.Splash.BeforeValidate, data1, binds)
		if err != nil {
			return nil, err
		}
	}

	if eerr != "" {
		return sw.GetSplash(ev, eerr)
	}

	if !skipValidation {
		for _, field := range sw.model.Splash.Fields {
			_, ok := data1.SplashData[field.Name]
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

		err := sw.execScript(sg.BeforeStart, data1.SplashData, binds)
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

	subData := newSub(sg.Name, stage.Name)

	resp := &ResponseStart{
		StartStage:  true,
		StageTitle:  stage.Name,
		Message:     stage.Message,
		Fields:      stage.Fields,
		DataSources: make(map[string]interface{}),
		OpaqueData:  nil,
		Ok:          true,
	}

	eerr = ""
	if stage.BeforeGenerate != "" {
		binds := map[string]interface{}{
			"_wizard_err": func(e string) {
				eerr = e
			},
			"_wizard_set_source_data": func(source string, data interface{}) {
				resp.DataSources[source] = data
			},
		}

		err := sw.execScript(sg.BeforeStart, nil, binds)
		if err != nil {
			return nil, err
		}

		if eerr != "" {

			return nil, easyerr.Error(eerr)
		}
	}

	err = sw.genSource(stage, &subData, resp.DataSources)
	if err != nil {
		return nil, err
	}

	opData, err := sw.updateSub(&subData)
	if err != nil {
		return nil, err
	}

	resp.OpaqueData = opData

	return resp, nil
}
