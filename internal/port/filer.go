package port

import "context"

type Filer[I, O any] interface {
	Generate(context.Context, ...I) ([]O, error)
}
