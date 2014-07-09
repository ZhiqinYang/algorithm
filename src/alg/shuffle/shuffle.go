package shuffle

import (
	"math/rand"
)

//--------------------------------------------------------- 简单的洗牌算法
func Shuffle(data []int32, length int) {
	for i, j := length-1, 0; i > 0; i-- {
		j = int(rand.Int31() % int32(i+1))
		data[i], data[j] = data[j], data[i]
	}
}
