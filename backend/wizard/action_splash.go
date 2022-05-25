package wizard

import (
	"encoding/json"

	"github.com/temphia/core/backend/server/btypes/easyerr"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
	"github.com/temphia/executors/backend/wizard/lifecycle"
	"github.com/temphia/executors/backend/wizard/wmodels"
)

func (sw *SimpleWizard) GetSplash(ev *event.Request, msg string) (interface{}, error) {
	req := wmodels.RequestSplash{}

	if ev.Data != nil {
		err := json.Unmarshal(ev.Data, &req)
		if err != nil {
			return nil, err
		}
	}

	return sw.getSplash(req.HasExecData, msg)
}

func (sw *SimpleWizard) getSplash(hasExecData bool, msg string) (interface{}, error) {

	dataSources := map[string]interface{}{}
	skipSplash := hasExecData

	if sw.model.Splash.OnLoad != "" {
		lf := lifecycle.OnSplashLoad{
			Models: &sw.model,
			SideEffect: lifecycle.OnSplashLoadSideEffect{
				FailErr:     "",
				SkipSplash:  skipSplash,
				DataSources: map[string]interface{}{},
			},
			HasExecData: skipSplash,
		}

		err := lf.Execute()
		if err != nil {
			return nil, err
		}

		if lf.SideEffect.FailErr != "" {
			return nil, easyerr.Error(lf.SideEffect.FailErr)
		}

		skipSplash = lf.SideEffect.SkipSplash
		dataSources = lf.SideEffect.DataSources
	}

	return wmodels.ResponseSplash{
		WizardTitle: sw.model.Title,
		Message:     msg,
		SkipSplash:  skipSplash,
		Fields:      sw.model.Splash.Fields,
		DataSources: dataSources,
	}, nil
}
