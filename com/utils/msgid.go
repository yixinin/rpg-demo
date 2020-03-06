package utils

func BytesToUint16(bs []byte) uint16 {
	return uint16(bs[0])<<8 | uint16(bs[1])
}

// 字符串转为16位整形哈希
func StringHash(s string) (hash uint16) {
	for _, c := range s {
		ch := uint16(c)
		hash = hash + ((hash) << 5) + ch + (ch << 7)
	}
	return
}
