package wizard

// before

type BeforeEnd struct {
	SideEffects BeforeEndSideEffects
}

type BeforeEndSideEffects struct {
	FailErr     string
	GotoStage   string
	GotoMessage string
}

type BeforeEndCtx struct {
	Type        string
	ParentGroup string
	ParentStage string
}

func (b *BeforeEnd) Execute() error { return nil }

// after

type AfterEnd struct {
	SideEffects AfterEndSideEffects
}

type AfterEndSideEffects struct {
	FailErr     string
	GotoStage   string
	GotoMessage string
}

type AfterEndCtx struct {
	Type        string
	ParentGroup string
	ParentStage string
}

func (b *AfterEnd) Execute() error { return nil }
