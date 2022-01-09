package popcount

import (
	"testing"
)

// BenchmarkPopcount-8             1000000000               0.3038 ns/op
func BenchmarkPopcount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(1000)
	}
}

// BenchmarkPopcountLoop-8         197738172                6.075 ns/op
func BenchmarkPopcountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(1000)
	}
}
