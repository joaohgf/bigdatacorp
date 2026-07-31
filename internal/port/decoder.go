package port

import "iter"

type (
	Sequence[T any] = iter.Seq2[T, error]
	Decoder[D any]  interface {
		Decode(string) Sequence[D]
	}
)
