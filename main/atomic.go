package main

import (
	"sync/atomic"
	"time"
)

var count1 int32
var count2 int32

func race() {
	count1++
	atomic.AddInt32(&count2, 1)
}

func main() {
	for i := 0; i < 1000; i++ {
		go race()
	}
	time.Sleep(2 * time.Second)
	println(count1, count2)

}