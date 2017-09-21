package base64

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestEncode(t *testing.T) {
	cost1 := 0
	cost2 := 0
	for i := 0; i < 1000; i++ {
		l := rand.Int31n(int32(1024 * 1024 * 10))
		buf := make([]byte, l, l)
		rand.Read(buf)

		begin := time.Now()
		str1 := base64.StdEncoding.EncodeToString(buf)
		cost1 = cost1 + int(time.Now().Sub(begin)/time.Millisecond)

		begin = time.Now()
		str2 := Encode(buf)
		cost2 = cost2 + int(time.Now().Sub(begin)/time.Millisecond)

		if str1 != str2 {
			t.Fatalf("bad encode")
		}
	}
	fmt.Printf("cost1=%d,cost2=%d\n", cost1, cost2)
}
