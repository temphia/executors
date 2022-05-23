package wizard

import "github.com/temphia/executors/backend/wizard/wmodels"

type OnSplashSubmit struct {
	Models     *wmodels.Wizard
	SideEffect OnSplashSubmitSideEffect
}

type OnSplashSubmitSideEffect struct {
	Err        string
	SkipSplash bool
	NextStage  string
	NextGroup  string
}

type OnSplashSubmitCtx struct {
	Type       string
	SubmitData map[string]interface{}
	ExecData   interface{}
}

func (s *OnSplashSubmit) Execute() error {

	return nil
}
