package encoding

type Encoder interface {
	Encode([]byte) ([]byte, error)
}

type Decoder interface {
	Decode([]byte) ([]byte, error)
}
