package port

type (
	To[I, O any] interface {
		To(I) O
	}
	ToMany[I, O any] interface {
		ToMany(...I) []O
	}
)
