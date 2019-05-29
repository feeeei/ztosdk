package common

type ZTORequest interface {
	Sign(key []byte) (string, error)
	EncodeBody() (string, error)
}
