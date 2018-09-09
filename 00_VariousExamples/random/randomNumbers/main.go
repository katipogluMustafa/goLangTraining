package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
	 * rand.Intn returns a random int n, 0 <= n < 100
	 */

	fmt.Println(rand.Intn(99))
	fmt.Println(rand.Float64())

	fmt.Println()

	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	fmt.Println(r1.Intn(100))

	/*
	 * If you seed a source with the same number, it produces the same sequence of random numbers.
	 */
	s2 := rand.NewSource(42)
	r2 := rand.New(s2)
	fmt.Println(r2.Intn(100))

	s3 := rand.NewSource(42)
	r3 := rand.New(s3)
	fmt.Println(r3.Intn(100))
}

/*
 * The default number generator is deterministic, so it’ll produce the same sequence of numbers each time by default. To produce varying sequences, give it a seed that changes. Note that this is not safe to use for random numbers you intend to be secret, use crypto/rand for those.
 */
