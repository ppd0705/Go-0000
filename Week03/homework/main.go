package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func signalHandle(ctx context.Context) error {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	select {
	case sig := <-quit:
		return fmt.Errorf("signal exit: %s", sig)
	case <-ctx.Done():
		return nil
	}
}

func httpHandle(ctx context.Context, addr string) error {
	var err error
	mux := http.NewServeMux()
	done := make(chan int)
	mux.HandleFunc("/done", func(w http.ResponseWriter, req *http.Request) {
		_, _ = fmt.Fprintf(w, "exit!")
		done <- 1
	})
	svc := &http.Server{Addr: addr, Handler: mux}
	go func() {
		select {
		case <-ctx.Done():
			_ = svc.Shutdown(context.Background())
		case <-done:
			log.Printf("server[%s] done", addr)
			_ = svc.Shutdown(context.Background())
		}
	}()
	err = svc.ListenAndServe()
	return err
}
func main() {
	g, ctx := errgroup.WithContext(context.Background())
	g.Go(func() error {
		return signalHandle(ctx)
	})
	g.Go(func() error {
		return httpHandle(ctx, ":8080")
	})
	g.Go(func() error {
		return httpHandle(ctx, ":8081")
	})
	if err := g.Wait(); err != nil {
		log.Fatalf("err: %v", err)
	}
}
