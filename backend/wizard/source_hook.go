package wizard

import (
	"fmt"

	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/executors/backend/wizard/wmodels"
)

type SourceCtx struct {
	Binding      rtypes.Bindings
	Event        *event.Request
	Model        *wmodels.Wizard
	Field        string
	Stage        string
	SharedVars   map[string]interface{}
	PreviousData map[string]map[string]interface{}
	Source       *wmodels.Source
}

func (ctx *SourceCtx) Process() (interface{}, error) {
	switch ctx.Source.Type {
	case "static":
		return ctx.Source.Data, nil
	default:
		return nil, easyerr.Error((fmt.Sprint("Skipping field, source not implemented", ctx.Field)))
	}
}
