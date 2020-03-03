package utils

func BytesToUint16(bs []byte) uint16 {
	return uint16(bs[0])<<8 | uint16(bs[1])
}
