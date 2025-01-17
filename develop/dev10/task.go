package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

var (
	flagTimeout = flag.Duration("timeout", 10*time.Second, "dial connection timeout")
)

func main() {
	flag.Parse()
	args := flag.Args()

	if len(args) < 2 {
		log.Fatal("Usage: my-telnet [--timeout=10s] host port")
	}
	conn, err := net.DialTimeout("tcp", net.JoinHostPort(args[0], args[1]), *flagTimeout)
	if err != nil {
		log.Fatal(err)
	}
	defer func() { _ = conn.Close() }()

	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("connected to", conn.RemoteAddr())

	go read(ctx, conn)
	go write(ctx, conn)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt
	cancel()
}

func write(ctx context.Context, conn net.Conn) {
	sc := bufio.NewScanner(os.Stdin)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			if !sc.Scan() {
				if sc.Err() != nil && errors.Is(sc.Err(), io.EOF) {
					log.Fatal("scan error", sc.Err())
				}
			}
			_, err := conn.Write([]byte(sc.Text() + "\n"))
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func read(ctx context.Context, conn net.Conn) {
	sc := bufio.NewReader(conn)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			s, err := sc.ReadString('\n')
			if err != nil {
				log.Fatal("read error:", err)
			}
			fmt.Print(s)
		}
	}
}

// simple tcp echo server
//func main() {
//	listener, err := net.Listen("tcp", "127.0.0.1:8080")
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer func() { _ = listener.Close() }()
//
//	fmt.Println("tcp listen on 127.0.0.1:8080")
//
//	for {
//		conn, err := listener.Accept()
//		if err != nil {
//			fmt.Println("accept conn error:", err)
//			continue
//		}
//		go accept(conn)
//	}
//}
//
//func accept(conn net.Conn) {
//	defer func() { _ = conn.Close() }()
//	fmt.Println("new conn:", conn.RemoteAddr())
//
//	reader := bufio.NewReader(conn)
//	for {
//		data, err := reader.ReadString('\n')
//		if err != nil {
//			fmt.Println("client disconnect:", conn.RemoteAddr())
//			return
//		}
//		fmt.Printf("received data: %s", data)
//		_, _ = conn.Write([]byte("echo: " + data))
//	}
//}
