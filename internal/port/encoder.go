package port

import "context"

// TODO check
type Encoder[D any] interface {
	Encode(context.Context, ...D) (string, error)
}
