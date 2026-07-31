package port

import "context"

type Usecase[I, O any] interface {
	Generate(context.Context, Sequence[I]) ([]O, error)
}
