package port

type (
	// To maps one source representation to a target representation.
	To[I, O any] interface {
		// To maps one input value to one output value.
		To(I) O
	}
	// ToMany maps multiple source representations to target representations.
	ToMany[I, O any] interface {
		// ToMany maps multiple input values to output values.
		ToMany(...I) []O
	}
)
