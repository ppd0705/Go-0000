package rolling

import (
	"fmt"
	"sync"
	"time"
)

type Policy struct {
	mu     sync.RWMutex
	size   int
	window *Window
	offset int

	bucketDuration time.Duration
	lastAppendTime time.Time
}

type PolicyOpt struct {
	bucketDuration time.Duration
}


func NewPolicy(w *Window, opt PolicyOpt) *Policy {
	return &Policy{
		window:         w,
		size:           w.size,
		bucketDuration: opt.bucketDuration,
		lastAppendTime: time.Now(),
	}
}

func (p *Policy) timespan() int {
	v := int(time.Since(p.lastAppendTime) / p.bucketDuration)
	if v > -1 {
		return v
	}
	return p.size
}

func (p *Policy) add(f func(offset int, val float64), val float64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	timespan := p.timespan()
	fmt.Printf("current timespan: %v\n", timespan)
	if timespan > 0 {
		p.lastAppendTime = p.lastAppendTime.Add(time.Duration(timespan * int(p.bucketDuration)))
		offset := p.offset
		s := offset + 1
		if timespan > p.size {
			timespan = p.size
		}
		e, e1 := s+timespan, 0
		if e > p.size {
			e1 = e - p.size
			e = p.size
		}
		for i := s; i < e; i++ {
			p.window.ResetBucket(i)
			offset = i
		}
		for i := 0; i < e1; i++ {
			p.window.ResetBucket(i)
			offset = i
		}
		fmt.Printf("[0,%d), [%d, %d) adjust offset: %d -> %d\n", e1, s, e,p.offset, offset)
		p.offset = offset
	}
	f(p.offset, val)
}

func (p *Policy) Append(val float64) {
	p.add(p.window.Append, val)
}

func (p *Policy) Add(val float64) {
	p.add(p.window.Add, val)
}

func (p *Policy) Reduce(f func(Iterator) float64) (val float64) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	timespan := p.timespan()
	if count := p.size - timespan; count > 0 {
		offset := p.offset + timespan + 1
		if offset >= p.size {
			offset = offset - p.size
		}
		val = f(p.window.Iterator(offset, count))
	}
	return val
}

func (p *Policy) Stats() {
	p.window.Stats()
}
