package wizard

import "github.com/temphia/executors/backend/wizard/wmodels"

type OnSplashLoad struct {
	Models     *wmodels.Wizard
	SideEffect OnSplashLoadSideEffect
}

type OnSplashLoadSideEffect struct {
	FailErr    string
	SkipSplash bool
}

type OnSplashLoadCtx struct {
	Type        string
	HasExecData bool
}

func (s *OnSplashLoad) Execute() error {
	return nil
}

func (s *OnSplashLoad) Bindings() map[string]interface{} {
	b := map[string]interface{}{
		"_wizard_set_err": func(err string) {
			s.SideEffect.FailErr = err
		},

		"_wizard_set_skip_splash": func(skip bool) {
			s.SideEffect.SkipSplash = skip
		},
	}
	return b
}
