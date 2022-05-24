package sloader

import (
	"fmt"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/executors/backend/wizard/wmodels"
)

type SLoader struct {
	Binding      rtypes.Bindings
	Event        *event.Request
	Model        *wmodels.Wizard
	Field        string
	Stage        string
	SharedVars   map[string]interface{}
	PreviousData map[string]map[string]interface{}
	Source       *wmodels.Source
}

func (s *SLoader) Process() (interface{}, error) {
	switch s.Source.Type {
	case "static":
		return s.Source.Data, nil
	case "js_script":
		return s.jsScript()
	default:
		return nil, easyerr.Error((fmt.Sprint("Skipping field, source not implemented", s.Field)))
	}
}

func (s *SLoader) jsScript() (interface{}, error) {
	pp.Println("Executing =>", s.Source.Target, s.bindings())
	return nil, nil
}

func (s *SLoader) bindings() map[string]interface{} {
	return map[string]interface{}{
		"_wizard_set_shared_var": func(name string, data interface{}) {
			s.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return s.SharedVars[name]
		},
	}
}
