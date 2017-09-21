package base64

import (
	"encoding/base64"
	"math/rand"
	"testing"
)

func TestEncode(t *testing.T) {
	for i := 0; i < 1000; i++ {
		l := rand.Int31n(int32(1024 * 1024 * 10))
		buf := make([]byte, l, l)
		rand.Read(buf)

		str1 := base64.StdEncoding.EncodeToString(buf)
		str2 := Encode(buf)

		if str1 != str2 {
			t.Fatalf("bad encode")
		}
	}
}
