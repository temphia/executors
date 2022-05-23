package lifecycle

import "github.com/temphia/executors/backend/wizard/wmodels"

// before

type StageBeforeGenerate struct {
	Models     *wmodels.Wizard
	SideEffect StageBeforeGenerateEffect
	SubData    *wmodels.Submission
}

type StageBeforeGenerateEffect struct {
	FailErr     string
	DataSources map[string]interface{}
}

type StageBeforeGenerateCtx struct {
	Type string
}

func (s *StageBeforeGenerate) Execute() error {
	return nil
}

func (s *StageBeforeGenerate) Bindings() map[string]interface{} {

	return map[string]interface{}{
		"_wizard_set_err": func(e string) {
			s.SideEffect.FailErr = e
		},
		"_wizard_set_shared_var": func(name string, data interface{}) {
			s.SubData.SharedVars[name] = data
		},
		"_wizard_get_shared_var": func(name string) interface{} {
			return s.SubData.SharedVars[name]
		},
		"_wizard_get_stage_data": func(name string) interface{} {
			return s.SubData.Data[name]
		},
		"_wizard_set_data_source": func(name string, data interface{}) {
			s.SideEffect.DataSources[name] = data
		},
	}

}

// after

type StageAfterGenerate struct {
	Models     *wmodels.Wizard
	SideEffect StageAfterGenerateEffect
}

type StageAfterGenerateEffect struct {
	FailErr string
}

type StageAfterGenerateCtx struct {
	Type string
}

func (s *StageAfterGenerate) Execute() error {
	return nil
}

func (s *StageAfterGenerate) Bindings() map[string]interface{} {

	return nil
}
