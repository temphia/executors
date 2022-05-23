package wizard

type BeforeStart struct {
	SideEffects BeforeStartSideEffects
}

type BeforeStartContext struct {
}

type BeforeStartSideEffects struct {
	Err       string
	NextStage string
}
