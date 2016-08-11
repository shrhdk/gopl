package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	printDigest([]byte("x"), []byte("X"))
}

func printDigest(m1, m2 []byte) {
	c1 := sha256.Sum256(m1)
	c2 := sha256.Sum256(m2)

	fmt.Printf("%x\n", c1)
	fmt.Printf("%x\n", c2)
	fmt.Printf("%t\n", c1 == c2)
	fmt.Printf("%d\n", popCountArray(xor(&c1, &c2)))
	fmt.Printf("%T\n", c1)
}

func xor(c1, c2 *[32]byte) [32]byte {
	var d [32]byte

	for i := 0; i < len(d); i++ {
		d[i] = c1[i] ^ c2[i]
	}

	return d
}

func popCountArray(array [32]byte) (count int) {
	for _, v := range array {
		count += popCount(v)
	}
	return
}

func popCount(v byte) (count int) {
	for v != 0 {
		v &= (v - 1)
		count++
	}
	return
}
