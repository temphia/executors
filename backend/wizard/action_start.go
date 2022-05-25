package wizard

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/executors/backend/wizard/lifecycle"
	"github.com/temphia/executors/backend/wizard/sloader"
	"github.com/temphia/executors/backend/wizard/wmodels"
)

func (sw *SimpleWizard) RunStart(ev *event.Request) (interface{}, error) {

	req := wmodels.RequestStart{}
	err := json.Unmarshal(ev.Data, &req)
	if err != nil {
		return nil, err
	}

	nextgroup := ""

	if sw.model.Splash.OnSubmit != "" {

		lf := lifecycle.OnSplashSubmit{
			Models:     &sw.model,
			SideEffect: lifecycle.OnSplashSubmitSideEffect{},
			SubmitData: req.SplashData,
			ExecData:   req.StartRawData,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}

		if !lf.SideEffect.SkipValidation {
			for _, field := range sw.model.Splash.Fields {
				_, ok := req.SplashData[field.Name]
				if !ok && field.Optional {
					continue
				}
				// fixme actually check data
				if !ok {
					return sw.getSplash(req.StartRawData != nil, fmt.Sprintf("Empty required field: %s", field.Name))
				}
			}
		}
	}
	return sw.runStart("", "", "", nextgroup, req.StartRawData)
}

func (sw *SimpleWizard) runStart(pgroup, pstage, psubid, nextgroup string, execData interface{}) (interface{}, error) {

	var sg *wmodels.StageGroup
	if nextgroup != "" {
		for _, _sg := range sw.model.StageGroups {
			if _sg.Name == nextgroup {
				sg = &_sg
			}
		}
	} else {
		sg = &sw.model.StageGroups[0]
	}

	if sg == nil {
		return nil, easyerr.Error("Stage Group not found")
	}

	nextStage := ""

	if sg.BeforeStart != "" {
		lf := lifecycle.BeforeStart{
			SideEffects: lifecycle.BeforeStartSideEffects{},
			ParentGroup: pgroup,
			ParentStage: pstage,
			ExecData:    execData,
		}

		err := lf.Execute()

		if err != nil {
			return nil, err
		}

		nextStage = lf.SideEffects.NextStage
	}

	var stage *wmodels.Stage
	if nextStage != "" {
		_stage, ok := sw.model.Stages[nextStage]
		if !ok {
			return nil, easyerr.Error("stage not found")
		}
		stage = _stage
	} else {
		stage = sw.model.Stages[sg.Stages[0]]
	}

	if stage == nil {
		return nil, easyerr.Error("Stage not found")
	}

	subData := wmodels.NewSub(pgroup, pstage, psubid, sg.Name, stage.Name)
	datasources := make(map[string]interface{})

	if sg.AfterStart != "" {
		lf := lifecycle.AfterStart{
			SubData: &subData,
			SideEffects: lifecycle.AfterStartSideEffects{
				DataSources: datasources,
			},
			ExecData: execData,
		}

		err := lf.Execute()

		if err != nil {
			return nil, err
		}

		if lf.SideEffects.PrevData != nil {
			subData.Data = lf.SideEffects.PrevData
		}
	}

	if stage.BeforeGenerate != "" {
		lf := lifecycle.StageBeforeGenerate{
			Models:  &sw.model,
			SubData: &subData,
			SideEffect: lifecycle.StageBeforeGenerateEffect{
				FailErr:     "",
				DataSources: datasources,
			},
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}
	}

	loader := sloader.SLoader{
		Binding:     sw.binding,
		Model:       &sw.model,
		SubData:     &subData,
		Stage:       stage,
		Group:       sg,
		DataSources: datasources,
	}

	err := loader.Process()
	if err != nil {
		return nil, err
	}

	if stage.AfterGenerate != "" {
		lf := lifecycle.StageAfterGenerate{
			Models:     &sw.model,
			SideEffect: lifecycle.StageAfterGenerateEffect{},
			SubData:    &subData,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}
	}

	return &wmodels.ResponseStart{
		StartStage:  true,
		StageTitle:  stage.Name,
		Message:     stage.Message,
		Fields:      stage.Fields,
		DataSources: datasources,
		OpaqueData:  nil,
		Ok:          true,
		PrevData:    subData.Data[stage.Name],
	}, nil

}
