package wizard

type BeforeBack struct {
	SideEffects BeforeBackSideEffects
}

type BeforeBackSideEffects struct {
	FailErr   string
	NextStage string
}

type BeforeBackCtx struct {
	Type         string
	CurrentStage string
}

func (b *BeforeBack) Execute() error { return nil }

func (b *BeforeBack) Bindings() map[string]interface{} { return nil }
