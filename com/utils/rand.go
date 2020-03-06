package utils

import (
	"math"
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

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

func RandInt64(min int64) int64 {
	n := rand.Int63n(math.MaxInt64 - min)
	return n + min
}
