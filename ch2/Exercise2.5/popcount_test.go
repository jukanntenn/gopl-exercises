package popcount

import (
	"testing"
)

// BenchmarkPopCountClean-8        46215768                26.19 ns/op
func BenchmarkPopCountClean(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountClean(0x1234567890ABCDEF)
	}
}

func TestCorrectness(t *testing.T) {
	tests := []uint64{0, 1, 21845, 1185589760, 0x1234567890ABCDEF}

	for _, test := range tests {
		if PopCount(test) != PopCountClean(test) {
			t.Errorf("PopCount(%v) != PopCountShift(%v)", test, test)
		}
	}
}
