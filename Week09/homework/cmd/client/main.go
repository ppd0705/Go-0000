package main

import (
	"bufio"
	"flag"
	"homework/api/procotol"
	"log"
	"net"
	"os"
	"strconv"
	"strings"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 16010, "listen port")
	flag.Parse()
}

func main() {
	addr := "0.0.0.0:" + strconv.Itoa(port)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		info := strings.Trim(input, "\n")
		if strings.ToUpper(info) == "QUIT" {
			conn.Write(procotol.Encode(&procotol.Message{}))
			return
		}
		log.Printf("send msg: %s\n", info)
		msg := procotol.Message{Content: []byte(info)}
		buff := procotol.Encode(&msg)
		_, err := conn.Write(buff)
		if err != nil {
			log.Fatal(err)
		}
	}
}
