/*
题目：参考 Hystrix 实现一个滑动窗口计数器。
*/

/*
答：滑动窗口的原理还在学习中，这里是参照别人的实现写的
*/

package homework

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type rollingWindow struct {
	bucket []int64
	index  int64
	sum    int64
	limit  int64
	size   int64
	ch     chan struct{}
	mu     sync.Mutex
}

func NewRollingWindow(limit int64, size int64) *rollingWindow {
	r := &rollingWindow{
		limit: limit,
		size:  size,
	}

	r.bucket = make([]int64, size)
	return r
}

func (rw *rollingWindow) Do() {
	length := 1000 / rw.size
	ticker := time.NewTicker(time.Duration(length) * time.Millisecond)
	go func() {
		for {
			select {
			case <-rw.ch:
				ticker.Stop()
				return
			case <-ticker.C:
				rw.roll()
			}
		}
	}()
}

func (rw *rollingWindow) roll() {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	rw.index = (rw.index + 1) % rw.size
	rw.sum = rw.sum - rw.bucket[rw.index]
	fmt.Println("rolling! sum: ", rw.GetSum(), " index: ", rw.GetIndex())
}

func (rw *rollingWindow) AddEvent() bool {
	rw.mu.Lock()
	defer rw.mu.Unlock()

	atomic.AddInt64(&rw.bucket[rw.index], 1)
	atomic.AddInt64(&rw.sum, 1)

	return rw.sum >= rw.limit
}

func (rw *rollingWindow) GetSum() int64 {
	return atomic.LoadInt64(&rw.sum)
}

func (rw *rollingWindow) GetIndex() int64 {
	return atomic.LoadInt64(&rw.index)
}

func main() {
	rw := NewRollingWindow(10, 10)
	rw.Do()

	for {
		if rw.AddEvent() {
			fmt.Println("Over limit!")
			time.Sleep(100 * time.Millisecond)
		} else {
			fmt.Println("current sum: ", rw.GetSum(), " index: ", rw.GetIndex())
		}
	}
}
