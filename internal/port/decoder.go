package port

type Decoder[D any] interface {
	Decode(string) ([]D, error)
}
