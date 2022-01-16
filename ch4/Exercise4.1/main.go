package main

import (
	"crypto/sha256"
	"fmt"

	popcount "github.com/jukanntenn/gopl-exercises/ch2/Exercise2.3"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := SHA256BitDiffCount(c1, c2)
	fmt.Println(count)
}

func SHA256BitDiffCount(c1, c2 [32]byte) int {
	cnt := 0
	for i := 0; i < 32; i++ {
		r := uint8(c1[i]) ^ uint8(c2[i])
		cnt += popcount.PopCount(uint64(r))
	}
	return cnt
}
