package rolling

import (
	"fmt"
)

type Iterator struct {
	count         int
	iteratedCount int
	cur           *Bucket
}

func (i *Iterator) CanNext() bool {
	return i.iteratedCount != i.count
}

func (i *Iterator) Bucket() Bucket {
	if !i.CanNext() {
		panic(fmt.Errorf("rolling: iteration out of range %d/%d", i.iteratedCount, i.count))
	}
	bucket := *i.cur
	i.iteratedCount++
	i.cur = i.cur.Next()
	return bucket
}
