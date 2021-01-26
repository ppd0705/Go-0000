package server

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	xerrors "github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"homework/api/procotol"
	"log"
	"net"
	"strconv"
)

var (
	port    int
	timeout int
)

func init() {
	flag.IntVar(&port, "port", 16010, "listen port")
	flag.IntVar(&timeout, "timeout", 5, "timeout(second)")
	flag.Parse()
}

func readHandle(ctx context.Context, conn net.Conn, ch chan *procotol.Message) error {
	reader := bufio.NewReader(conn)
	defer conn.Close()
	for {
		msg, err := procotol.Decode(reader)
		if err != nil {
			return xerrors.Wrap(err, "read error")
		}
		if msg.Empty() {
			return fmt.Errorf("remote closed")
		}
		ch <- msg
	}
}

func writeHandle(ctx context.Context, conn net.Conn, ch chan *procotol.Message) error {
	//writer := bufio.NewWriter(conn)
	log.Printf("start writer")
	for {
		select {
		case msg := <- ch:
			log.Printf("recieve msg: %s\n", msg.Content)
			//data := procotol.Encode(msg)
			//_, err := writer.Write(data)
			//if err != nil {
			//	return xerrors.Wrap(err, "write error")
			//}
		case <-ctx.Done():
			return nil
		}
	}
}

func Handle(conn net.Conn) {
	g, ctx := errgroup.WithContext(context.Background())
	ch := make(chan *procotol.Message)
	g.Go(func() error {
		return readHandle(ctx, conn, ch)
	})
	g.Go(func() error {
		return writeHandle(ctx, conn, ch)
	})
	if err := g.Wait(); err != nil {
		log.Printf("hadle conn(%v), err: %v\n", conn.RemoteAddr(), err)
	}
}

func StartServer() error {
	addr := "0.0.0.0:" + strconv.Itoa(port)
	fmt.Printf("listen on %v\n", addr)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go Handle(conn)
	}
}
