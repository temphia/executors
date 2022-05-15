package dashed

import (
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/stdplugs/backend/dashed/dashmodels"
)

var _ rtypes.Executor = (*SimpleDash)(nil)

type SimpleDash struct {
	bindings rtypes.Bindings
	model    dashmodels.Dashboard
}

func (s *SimpleDash) Process(ev *event.Request) (*event.Response, error) {

	switch ev.Name {
	case "generate":
		return s.generate(ev)
	default:
		return nil, easyerr.NotImpl()
	}
}
