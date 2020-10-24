package middlewares

// ctxKey defines context key to avoid key collisions
type ctxKey int

const (
	// TestCtxKey for getting test with id
	TestCtxKey ctxKey = iota
)
