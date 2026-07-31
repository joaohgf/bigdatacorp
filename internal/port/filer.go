package port

import "context"

type Filer[I, O any] interface {
	Generate(context.Context, Sequence[I]) ([]O, error)
}
