package buildname

import (
	"encoding/binary"
	"encoding/hex"
	"math/rand"
	"testing"
)

func randStrings(n, k, seed int, res []string) []string {
	s := rand.NewSource(int64(seed))
	b := make([]byte, k)

	for i := 0; i < n; i++ {
		for j := 0; j < k/8; j++ {
			binary.BigEndian.PutUint64(b[8*j:], uint64(s.Int63()))
		}
		res = append(res, hex.EncodeToString(b))
	}

	return res
}

func TestFromVersion(t *testing.T) {
	tests := randStrings(200, 24, 420, []string{
		"",
		"3.20",
	})

	for _, test := range tests {
		name := FromVersion(test)
		if name == "" {
			t.Fatalf("name missing for %s", test)
		}
	}
}
