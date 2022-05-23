package wizard

import "github.com/temphia/executors/backend/wizard/wmodels"

type OnSplashLoad struct {
	Models *wmodels.Wizard
}

type OnSplashLoadSideEffect struct {
	Err        string
	SkipSplash bool
}

type OnSplashLoadCtx struct {
	Type        string
	HasExecData bool
}

func (s *OnSplashLoad) Execute() error {
	return nil
}

// private
