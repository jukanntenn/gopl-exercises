package popcount

import (
	"testing"
)

// BenchmarkPopcount-8             1000000000               0.3038 ns/op
func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1000)
	}
}

// BenchmarkPopcountLoop-8         197738172                6.075 ns/op
func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(1000)
	}
}

func TestCorrectness(t *testing.T) {
	tests := []uint64{0, 1, 21845, 1185589760, 0x1234567890ABCDEF}

	for _, test := range tests {
		if PopCount(test) != PopCountLoop(test) {
			t.Errorf("PopCount(%v) != PopCountLoop(%v)", test, test)
		}
	}
}
