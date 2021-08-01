package utils

func MebibyteToBytes(mb int64) int64 {
	return mb << 20
}
