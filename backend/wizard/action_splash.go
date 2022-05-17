package wizard

import (
	"encoding/json"

	"github.com/temphia/core/backend/server/btypes/rtypes/event"
)

func (sw *SimpleWizard) GetSplash(ev *event.Request, msg string) (interface{}, error) {
	req := RequestSplash{}

	if ev.Data != nil {
		err := json.Unmarshal(ev.Data, &req)
		if err != nil {
			return nil, err
		}
	}

	if req.HasExecData {
		return ResponseSplash{
			WizardTitle: sw.model.Title,
			Message:     msg,
			SkipSplash:  true,
		}, nil
	}

	// fixme => before_generate

	return ResponseSplash{
		WizardTitle: sw.model.Title,
		Message:     msg,
		Fields:      sw.model.Splash.Fields,
		DataSources: make(map[string]interface{}),
		SkipSplash:  sw.model.Splash.Skip,
	}, nil
}
