package port

import "iter"

type (
	// Sequence streams decoded values together with terminal decoding errors.
	Sequence[T any] = iter.Seq2[T, error]
	// Decoder streams values decoded from a source path.
	Decoder[D any] interface {
		// Decode returns a lazy sequence for the supplied source path.
		Decode(string) Sequence[D]
	}
)
