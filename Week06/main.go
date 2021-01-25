package main

import (
	"Week06/rolling"
	"fmt"
	"time"
)

func main() {
	opt := rolling.CounterOpts{Size: 5, BucketDuration: 100 * time.Millisecond}
	rollingCounter := rolling.NewCounter(opt)
	rollingCounter.Add(2)
	time.Sleep(300 * time.Millisecond)
	rollingCounter.Add(5)
	rollingCounter.Add(5)
	fmt.Printf("rolling counter avg:%f,Value:%d\n", rollingCounter.Avg(), rollingCounter.Value())
	rollingCounter.Stats()
	time.Sleep(100 * time.Millisecond)
	rollingCounter.Add(1)
	//time.Sleep(200 * time.Millisecond)
	rollingCounter.Add(3)
	fmt.Printf("rolling counter avg:%f,Value:%d\n", rollingCounter.Avg(), rollingCounter.Value())
	rollingCounter.Stats()
}
