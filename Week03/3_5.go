package main

import (
	"context"
	"fmt"
	"time"
)

type Tracker struct {
	ch chan string
	stop chan struct{}
}

func (t  *Tracker) Event(ctx context.Context, data string) error {
	select {
	case t.ch <- data:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *Tracker) Run(){
	for data := range t.ch {
		time.Sleep(time.Second)
		fmt.Printf("data: %s\n", data)
	}
	t.stop <- struct{}{}
}

func (t *Tracker) Shutdown (ctx context.Context) {
	close(t.ch)
	select {
	case <- t.stop:
	case <- ctx.Done():
	}
}

func NewTracker() *Tracker {
	return &Tracker{
		ch: make(chan string, 10),
	}
}
func main() {
	tr := NewTracker()
	go tr.Run()
	_ = tr.Event(context.Background(), "test1")
	_ = tr.Event(context.Background(), "test2")
	_ = tr.Event(context.Background(), "test3")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	tr.Shutdown(ctx)
}