package wizard

type BeforeNext struct {
	SideEffects BeforeNextSideEffects
}

type BeforeNextSideEffects struct {
	FailErr   string
	NextStage string
}

type BeforeNextCtx struct {
	Type         string
	CurrentStage string
}

func (b *BeforeNext) Execute() error { return nil }

func (b *BeforeNext) Bindings() map[string]interface{} { return nil }
