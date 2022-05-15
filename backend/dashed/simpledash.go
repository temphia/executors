package simpledash

import (
	"github.com/temphia/temphia/backend/server/btypes/easyerr"
	"github.com/temphia/temphia/backend/server/btypes/rtypes"
	"github.com/temphia/temphia/backend/server/btypes/rtypes/event"
	"github.com/temphia/temphia/backend/stdplugs/simpledash/dashmodels"
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
