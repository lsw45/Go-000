package Week06

import (
	"fmt"
	"testing"
	"time"
)

// 滑动窗口计数有很多使用场景，比如说限流防止系统雪崩。相比计数实现，滑动窗口实现会更加平滑，能自动消除毛刺。
// 滑动窗口原理是在每次有访问进来时，先判断前N个单位时间内的总访问量是否超过了设置的阈值，并对当前时间片上的请求数+1。

const (
	// 窗口请求数阈值
	reqLimit = 100

	// 单位时间
	unitTime = 1 * time.Millisecond

	// 窗口前N个单位时间——10个
	WinSize = 10 * time.Millisecond
)

type SlideWin struct {
	reqCount int
	bucket   []*Bucket
}

type Bucket struct {
	count int
	time  int64
}

func NewSlideWin() *SlideWin {
	return &SlideWin{
		reqCount: 0,
		bucket:   []*Bucket{},
	}
}

// 加入新的请求
func (s *SlideWin) Add(b *Bucket) {
	s.bucket = append(s.bucket, b)
	s.reqCount += b.count
}

// 删除时间窗前面的请求
func (s *SlideWin) Delete() {
	now := time.Now().Unix()

	for k, v := range s.bucket {
		if v.time < now-int64(WinSize)/10000000 {
			if k == 0 {
				s.bucket = nil
				s.reqCount = 0
				return
			}
			s.bucket = append(s.bucket[:k], s.bucket[k+1:]...)
			s.reqCount -= v.count
			fmt.Println("delete:", k)
		}
	}
}

func (s *SlideWin) Average() int {
	return s.reqCount / int(WinSize)
}

func TestReq(t *testing.T) {
	slide := NewSlideWin()
	reqList := []int{100, 200, 140, 160, 20, 50, 60, 20, 80, 1, 1, 1, 1, 1, 1}

	for _, v := range reqList {
		// 模拟请求数
		if slide.reqCount < reqLimit {
			now := time.Now().Unix()
			slide.Add(&Bucket{count: v, time: now})
			fmt.Printf("time:%v, reqCount:%v\n", now, slide.reqCount)
		} else {
			time.Sleep(1 * time.Second)
			slide.Delete()
			fmt.Println("sleep one second")
		}
	}
}
