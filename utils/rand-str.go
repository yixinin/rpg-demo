package utils

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func RandStr(length int) string {
	var str = UUID()
	if len(str) >= length {
		return str[:length]
	}
	for len(str) < length {
		str += UUID()
	}

	return str
}

func UUID() string {
	return strings.ReplaceAll(uuid.NewV4().String(), "-", "")
}
