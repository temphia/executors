package wizard

import "github.com/temphia/executors/backend/wizard/wmodels"

type OnSplashSubmit struct {
	Models     *wmodels.Wizard
	SideEffect OnSplashSubmitSideEffect
}

type OnSplashSubmitSideEffect struct {
	FailErr        string
	NextGroup      string
	SkipValidation bool
}

type OnSplashSubmitCtx struct {
	Type       string
	SubmitData map[string]interface{}
	ExecData   interface{}
}

func (s *OnSplashSubmit) Execute() error {

	return nil
}

func (s *OnSplashSubmit) Bindings() map[string]interface{} {

	b := map[string]interface{}{
		"_wizard_set_next_stage_group": func(name string) {
			s.SideEffect.NextGroup = name
		},
		"_wizard_set_err": func(err string) {
			s.SideEffect.FailErr = err
		},

		"_wizard_set_skip_validation": func(skip bool) {
			s.SideEffect.SkipValidation = skip
		},
	}

	return b

}
