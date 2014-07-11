package rand

import (
	"time"
)

const (
	m = uint32(1 << 31)
	a = uint32(1103515245)
	c = uint32(12345)
)

// 原子操作
var LCG = make(chan uint32)
var X = uint32(time.Now().Unix())

func init() {
	go func() {
		for {
			X = (a*X + c) % m
			LCG <- X
		}
	}()
}
