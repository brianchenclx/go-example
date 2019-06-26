package goroutines

import (
	"io"
	"log"
	"net"
	"time"
)

func Clock1() {
	listener, err := net.Listen("tcp", "localhost:9000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05.000\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Millisecond)
	}
}
