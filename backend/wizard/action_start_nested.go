package wizard

import (
	"encoding/json"
	"fmt"

	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/thoas/go-funk"
)

func (sw *SimpleWizard) RunNestedStart(ev *event.Request) (interface{}, error) {

	// fixme => handle prevdata properly / edit through nested

	data := RequestStartNested{}

	err := json.Unmarshal(ev.Data, &data)
	if err != nil {
		return nil, err
	}

	sub, err := sw.getSub(data.ParentOpData)
	if err != nil {
		return nil, err
	}

	if sub.ParentStageGroup != "" {
		panic("cannot have double nested stages")
	}

	pgroup := sw.getStageGroup(sub.StageGroup)
	if pgroup == nil {
		panic("Empty parent group to start nested stage")
	}

	if !funk.ContainsString(pgroup.Stages, sub.CurrentStage) {
		panic("stage is not in current group")
	}

	pstage := sw.model.Stages[sub.CurrentStage]
	if pstage == nil {
		panic("Empty parent stage to start nested stage")
	}

	field := pstage.getField(data.Field)
	if field == nil {
		panic("field not found, to start nested stage")
	}

	ngroup := sw.getStageGroup(field.Attrs["nested_stage_group"].(string))
	if ngroup == nil {
		panic("Nesed stage group not found")
	}

	nextStage := ""
	eerr := ""

	if ngroup.BeforeStart != "" {
		binds := map[string]interface{}{
			"_wizard_set_next_stage": func(name string) {
				nextStage = name
			},
			"_wizard_set_err": func(name string) {
				eerr = name
			},
			// fixme => add get_stage, get_group, get_parent_data, get_parent_ctx_var
		}

		type nextedBeforeCtx struct {
			IsNested    bool   `json:"is_nested,omitempty"`
			Field       string `json:"field,omitempty"`
			ParentGroup string `json:"parent_group,omitempty"`
			ParentStage string `json:"parent_stage,omitempty"`
		}

		err := sw.execScript(ngroup.BeforeStart, &nextedBeforeCtx{
			IsNested:    true,
			Field:       data.Field,
			ParentGroup: sub.StageGroup,
			ParentStage: sub.CurrentStage,
		}, binds)
		if err != nil {
			return nil, err
		}
	}

	if eerr != "" {
		return nil, easyerr.Error(eerr)
	}

	var stage *Stage

	if nextStage != "" {
		_stage, ok := sw.model.Stages[nextStage]
		if !ok {
			return sw.GetSplash(ev, fmt.Sprintf("Stage not found: %s", nextStage))
		}
		stage = _stage
	} else {
		stage = sw.model.Stages[ngroup.Stages[0]]
	}

	return sw.startStage(ngroup, stage, sub.StageGroup, sub.CurrentStage)
}
