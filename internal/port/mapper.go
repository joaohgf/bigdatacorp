package port

type (
	From[I, O any] interface {
		From(O) I
	}
	ToMany[I, O any] interface {
		ToMany(...I) []O
	}
)
