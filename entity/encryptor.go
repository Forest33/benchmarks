package entity

type Decoder interface {
	Decrypt(data any) ([]byte, error)
}

type Encoder interface {
	Encrypt(data []byte) ([]byte, error)
	GetLength(int) int
}

type Encryptor interface {
	Decoder
	Encoder
}
