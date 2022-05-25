package lifecycle

import "github.com/temphia/executors/backend/wizard/wmodels"

type OnSplashLoad struct {
	Models      *wmodels.Wizard
	SideEffect  OnSplashLoadSideEffect
	HasExecData bool
}

type OnSplashLoadSideEffect struct {
	FailErr     string
	SkipSplash  bool
	DataSources map[string]interface{}
}

type OnSplashLoadCtx struct {
	Type        string
	HasExecData bool
}

func (s *OnSplashLoad) Execute() error {
	return nil
}

func (s *OnSplashLoad) Bindings() map[string]interface{} {
	return map[string]interface{}{
		"_wizard_set_err": func(err string) {
			s.SideEffect.FailErr = err
		},
		"_wizard_set_skip_splash": func(skip bool) {
			s.SideEffect.SkipSplash = skip
		},
		"_wizard_set_data_source": func(name string, data interface{}) {
			s.SideEffect.DataSources[name] = data
		},
	}
}
