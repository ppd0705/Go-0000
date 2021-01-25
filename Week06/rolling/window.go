package rolling

import "fmt"

type Bucket struct {
	Points []float64
	Count  int64
	next   *Bucket
}

func (b *Bucket) Append(val float64) {
	b.Points = append(b.Points, val)
	b.Count++
}

func (b *Bucket) Add(offset int, val float64) {
	b.Points[offset] += val
	b.Count++
}

func (b *Bucket) Reset() {
	b.Points = b.Points[:0]
	b.Count = 0
}

func (b *Bucket) Next() *Bucket {
	return b.next
}

type Window struct {
	window []Bucket
	size   int
}

type WindowOpt struct {
	Size int
}

func NewWindow(opt WindowOpt) *Window {
	buckets := make([]Bucket, opt.Size)
	for i := range buckets {
		buckets[i] = Bucket{Points: []float64{}}
		nextOffset := i + 1
		if nextOffset == opt.Size {
			nextOffset = 0
		}
		buckets[i].next = &buckets[nextOffset]
	}
	return &Window{window: buckets, size: opt.Size}
}

func (w *Window) ResetWindow() {
	for i := range w.window {
		w.ResetBucket(i)
	}
}

func (w *Window) ResetBucket(offset int) {
	w.window[offset].Reset()
}

func (w *Window) ResetBuckets(offsets []int) {
	for _, offset := range offsets {
		w.ResetBucket(offset)
	}
}

func (w *Window) Append(offset int, val float64) {
	w.window[offset].Append(val)
}

func (w *Window) Add(offset int, val float64) {
	if w.window[offset].Count == 0 {
		w.window[offset].Append(val)
		return
	}
	w.window[offset].Add(0, val)
}

func (w *Window) Size() int {
	return w.size
}

func (w *Window) Iterator(offset int, count int) Iterator {
	return Iterator{count: count, cur: &w.window[offset]}
}

func (w *Window) Stats() {
	for i:=0; i < w.size; i++ {
		fmt.Printf("bucket %d: %v\n", i, w.window[i].Points)
	}
}