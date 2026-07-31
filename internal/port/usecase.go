package port

import "context"

// Usecase coordinates business processing for a stream of input values.
type Usecase[I, O any] interface {
	// Generate processes an input stream and returns generated artifacts.
	Generate(context.Context, Sequence[I]) ([]O, error)
}
