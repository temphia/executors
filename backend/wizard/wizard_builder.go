package wizard

import (
	"errors"
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

type SWBuilder struct{}

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

const fpath = "../executors/backend/wizard/sample/"

func newDev(opts rtypes.ExecutorOption) (rtypes.Executor, error) {

	pwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	out, err := os.ReadFile(path.Join(pwd, fpath, "_all.yaml"))
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

func (sd *SWBuilder) Instance(opts rtypes.ExecutorOption) (rtypes.Executor, error) {
	return New(opts)
}

func (sd *SWBuilder) ExecFile(file string) ([]byte, error) {

	if strings.HasSuffix(file, ".css") {
		if sd.fileExists(DevPath, "wizard.css") {
			return sd.serveFile(DevPath, "wizard.css")
		}
		return loaderCSS, nil
	}

	if strings.HasSuffix(file, ".js") {
		if sd.fileExists(DevPath, "wizard.js") {
			return sd.serveFile(DevPath, "wizard.js")
		}
		return loaderJS, nil
	}

	if strings.HasSuffix(file, ".js.map") {
		if sd.fileExists(DevPath, "wizard.js.map") {
			return sd.serveFile(DevPath, "wizard.js.map")
		}
	}

	return nil, easyerr.NotFound()
}

func (sd *SWBuilder) serveFile(dpath, file string) ([]byte, error) {
	return ioutil.ReadFile(path.Join("../executors/frontend/public/build/", file))
}

func (sd *SWBuilder) fileExists(dpath, file string) bool {
	_, err := os.Stat(path.Join(dpath, file))
	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}
