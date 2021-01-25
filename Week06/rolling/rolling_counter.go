package rolling

import (
	"fmt"
	"time"
)

type CounterOpts struct {
	Size           int
	BucketDuration time.Duration
}

type Counter struct {
	policy *Policy
}

func NewCounter(opt CounterOpts) *Counter {
	window := NewWindow(WindowOpt{Size: opt.Size})
	policy := NewPolicy(window, PolicyOpt{bucketDuration: opt.BucketDuration})
	return &Counter{policy: policy}
}

func (r *Counter) Add(val int64) {
	if val < 0 {
		panic(fmt.Errorf("rolling: cannot decease in value. val: %d", val))
	}
	r.policy.Add(float64(val))
}

func (r *Counter) Reduce(f func(Iterator) float64) float64 {
	return r.policy.Reduce(f)
}

func (r *Counter) Avg() float64 {
	return r.policy.Reduce(Avg)
}

func (r *Counter) Min() float64 {
	return r.policy.Reduce(Min)
}

func (r *Counter) Max() float64 {
	return r.policy.Reduce(Max)
}

func (r *Counter) Sum() float64 {
	return r.policy.Reduce(Sum)
}

func (r *Counter) Value() int64 {
	return int64(r.Sum())
}

func (r *Counter) TimeSpan() int {
	return r.policy.timespan()
}

func (r *Counter) Stats(){
	r.policy.Stats()
}
