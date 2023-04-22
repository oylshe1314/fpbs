package fpbs

import (
	"math/rand"
	"testing"
)

func TestVerify(t *testing.T) {

	var p = make([]byte, 8)

	for i := 0; i < 7; i++ {
		p[i] = byte(rand.Intn(256))
	}

	p[7] = Verification(p[:7])

	t.Log(p)

	t.Log(Verify(p))
}
