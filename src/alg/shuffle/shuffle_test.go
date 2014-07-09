package shuffle

import (
	"fmt"
	"testing"
)

func TestShuffle(t *testing.T) {
	data := []int32{2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < 10000; i++ {
		Shuffle(data, len(data))
		fmt.Println(data)

	}
}
