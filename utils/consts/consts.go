package consts

type ContextKey int

const (
	ContextKeyUser ContextKey = iota + 1
)

type Entity string
type Action string

const (
	User Entity = "user"

	Create Action = "create"
)
