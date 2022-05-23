package wizard

import (
	"encoding/json"

	"github.com/temphia/executors/backend/wizard/wmodels"
	"github.com/thoas/go-funk"
)

func (sw *SimpleWizard) getStageGroup(group string) *wmodels.StageGroup {
	for _, grp := range sw.model.StageGroups {
		if group != grp.Name {
			continue
		}
		return &grp
	}

	return nil
}

func (sw *SimpleWizard) genSource(stage *wmodels.Stage, subData *wmodels.Submission, psData map[string]interface{}) error {

	/*

			for _, field := range stage.Fields {
			if field.Source == "" {
				continue
			}

			_, ok := psData[field.Source]
			if ok {
				continue
			}

			source := sw.model.Sources[field.Source]

			if source == nil {
				pp.Println("@not found ", field.Source, "@", sw.model.Sources)
				return easyerr.NotFound()
			}

			ctx := SourceCtx{
				Binding:      sw.binding,
				Model:        &sw.model,
				Field:        field.Name,
				Stage:        stage.Name,
				SharedVars:   subData.SharedVars,
				PreviousData: subData.Data,
				Source:       source,
			}

			sdata, err := ctx.Process()
			if err != nil {
				return err
			}

			psData[field.Source] = sdata
		}

	*/

	return nil

}

func (sw *SimpleWizard) stageIndex(group *wmodels.StageGroup, stage string) int {
	return funk.IndexOfString(group.Stages, stage)
}

func (sw *SimpleWizard) getSub(opData []byte) (*wmodels.Submission, error) {
	subData := wmodels.Submission{}
	err := json.Unmarshal(opData, &subData)
	if err != nil {
		return nil, err
	}

	if subData.Data == nil {
		subData.Data = make(map[string]map[string]interface{})
	}

	if subData.SharedVars == nil {
		subData.SharedVars = make(map[string]interface{})
	}

	return &subData, nil
}

func (sw *SimpleWizard) updateSub(sdata *wmodels.Submission) ([]byte, error) {
	return json.Marshal(sdata)
}
