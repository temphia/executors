package wizard

import (
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/core/backend/server/registry"

	"github.com/ztrue/tracerr"

	_ "embed"
)

var (
	//go:embed embed/wizard.loader.js
	loaderJS []byte

	//go:embed embed/wizard.loader.css
	loaderCSS []byte

	DevPath = "../executors/frontend/public/build/"
)

type SimpleWizard struct {
	model         Wizard
	binding       rtypes.Bindings
	jsRuntime     *goja.Runtime
	nativeScripts map[string]registry.DynamicScript
}

func (s *SimpleWizard) Process(ev *event.Request) (*event.Response, error) {

	var resp interface{}
	var err error

	switch ev.Name {
	case "get_splash":
		resp, err = s.GetSplash(ev, "")
	case "run_start":
		resp, err = s.RunStart(ev)
	case "run_nested_start":
		resp, err = s.RunNestedStart(ev)
	case "run_back":
		resp, err = s.RunBack(ev)
	case "run_next":
		resp, err = s.RunNext(ev)
	default:
		return nil, easyerr.NotImpl()
	}

	if err != nil {
		return nil, err
	}

	out, err := json.Marshal(resp)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &event.Response{
		Vars:    map[string]interface{}{},
		Payload: out,
	}, nil

}
