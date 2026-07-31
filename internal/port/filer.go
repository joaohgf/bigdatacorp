package port

import "context"

// Filer generates output artifacts from a stream of input values.
type Filer[I, O any] interface {
	// Generate consumes a value stream and returns generated artifacts.
	Generate(context.Context, Sequence[I]) ([]O, error)
}
