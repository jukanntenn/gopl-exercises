package popcount

import (
	"testing"
)

// BenchmarkPopcount-8             1000000000               0.3068 ns/op
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

// BenchmarkPopCountShift-8        43614860                25.94 ns/op
func BenchmarkPopCountShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountShift(0x1234567890ABCDEF)
	}
}

func TestCorrectness(t *testing.T) {
	tests := []uint64{0, 1, 21845, 1185589760, 0x1234567890ABCDEF}

	for _, test := range tests {
		if PopCount(test) != PopCountShift(test) {
			t.Errorf("PopCount(%v) != PopCountShift(%v)", test, test)
		}
	}
}
