package simplewizard2

import (
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/dop251/goja"
	"github.com/goccy/go-yaml"
	"github.com/temphia/core/backend/server/btypes"
	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes"
	"github.com/temphia/core/backend/server/registry"
	"github.com/ztrue/tracerr"
)

func init() {
	registry.SetExecutor("simple.wizard", &SWBuilder{})
}

type SWBuilder struct {
}

func (sd *SWBuilder) Instance(opts rtypes.ExecutorOption) (rtypes.Executor, error) {
	return New(opts)
}

func (sd *SWBuilder) ExecFile(file string) ([]byte, error) {
	if strings.HasSuffix(file, ".css") {
		return ioutil.ReadFile("frontend/public/build/plug_simplewizard.css")
	}

	if strings.HasSuffix(file, ".js") {
		return ioutil.ReadFile("frontend/public/build/plug_simplewizard.js")
	}

	return nil, easyerr.NotFound()
}

func New(opts rtypes.ExecutorOption) (rtypes.Executor, error) {

	if btypes.DevMode {
		return newDev(opts)
	}

	out, err := opts.Binder.GetSelfFile("wizard.yaml")
	if err != nil {
		return nil, tracerr.Wrap(err)
	}

	model := Wizard{}

	err = yaml.Unmarshal(out, &model)
	if err != nil {
		return nil, err
	}

	return &SimpleWizard{
		model:         model,
		binding:       opts.Binder,
		jsRuntime:     nil,
		nativeScripts: nil,
	}, nil
}

const fpath = "/backend/stdplugs/simplewizard/sample/"

func newDev(opts rtypes.ExecutorOption) (rtypes.Executor, error) {

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	out, err := os.ReadFile(path.Join(pwd, fpath, "_two.yaml"))
	if err != nil {
		return nil, err
	}

	hookjs, err := os.ReadFile(path.Join(pwd, fpath, "_one.js"))
	if err != nil {
		hookjs = []byte(``)
	}

	model := Wizard{}

	err = yaml.Unmarshal(out, &model)
	if err != nil {
		return nil, err
	}

	rt := goja.New()

	_, err = rt.RunString(string(hookjs))
	if err != nil {
		return nil, err
	}

	for skey, s := range model.Stages {
		s.Name = skey
	}

	for k, s := range model.Sources {
		s.Name = k
	}

	return &SimpleWizard{
		model:         model,
		binding:       opts.Binder,
		jsRuntime:     rt,
		nativeScripts: nil,
	}, nil

}
