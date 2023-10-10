package entity

type CompressionLevel uint8

func (c CompressionLevel) Byte() byte {
	return byte(c)
}
