package product

type State int

const (
	StateDraft State = iota
	StateActive
	StateArchived
)
