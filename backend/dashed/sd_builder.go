package dashed

import (
	"io/ioutil"
	"strings"

	"github.com/goccy/go-yaml"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/registry"

	"github.com/temphia/executors/backend/dashed/dashmodels"
	"github.com/ztrue/tracerr"
)

func init() {
	registry.SetExecutor("simple.dash", &SDBuilder{})
}

type SDBuilder struct{}

func (sd *SDBuilder) Instance(opts rtypes.ExecutorOption) (rtypes.Executor, error) {
	return New(opts)
}

func (sd *SDBuilder) ExecFile(file string) ([]byte, error) {
	if strings.HasSuffix(file, ".js") {
		return ioutil.ReadFile("frontend/public/build/dashed.js")
	}

	if strings.HasSuffix(file, ".css") {
		return ioutil.ReadFile("frontend/public/build/dashed.css")
	}

	if strings.HasSuffix(file, ".js.map") {
		return ioutil.ReadFile("frontend/public/build/dashed.js.map")
	}

	return nil, easyerr.NotFound()
}

func New(opts rtypes.ExecutorOption) (*SimpleDash, error) {

	out, err := opts.Binder.GetSelfFile("dash.yaml")
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	model := dashmodels.Dashboard{}

	err = yaml.Unmarshal(out, &model)
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	return &SimpleDash{
		model:    model,
		bindings: opts.Binder,
	}, nil
}
