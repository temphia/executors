package wizard

import (
	"encoding/json"

	"github.com/k0kubun/pp"
	"github.com/temphia/core/backend/server/btypes/rtypes/event"
)

func (sw *SimpleWizard) RunBack(ev *event.Request) (interface{}, error) {

	req := RequestBack{}

	err := json.Unmarshal(ev.Data, &req)
	if err != nil {
		return nil, err
	}

	sub, err := sw.getSub(req.OpaqueData)
	if err != nil {
		return nil, err
	}

	if len(sub.VisitedStages) == 0 {
		panic("cannot back further")
	}

	pp.Println(sub)

	// json.Unmarshal()

	return nil, nil
}
