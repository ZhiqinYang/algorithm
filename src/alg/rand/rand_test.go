package rand

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {

	for i := 0; i < 100; i++ {
		f := <-LCG
		fmt.Println(f)
	}
}

func Benchmark_Rand(b *testing.B) {
	a := <-LCG
	for i := 0; i < b.N; i++ {

		b := <-LCG
		if b == a {
			fmt.Println(a, b, true)
            break
		}
		fmt.Println(b)
	}
}
