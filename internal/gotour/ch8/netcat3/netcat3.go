package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")

	if err != nil {
		log.Fatal(err)
	}

	done := make(chan struct{})
	go func() {
		fmt.Println("start")
		io.Copy(os.Stdout, conn)
		log.Println("done")
		fmt.Println("done")
		done <- struct{}{}
	}()

	mustCopy(conn, os.Stdin)
	conn.Close()
	<-done
}

func mustCopy(dst io.Writer, src io.Reader) {
	for _, err := io.Copy(dst, src); err != nil; {
		log.Fatal(err)
	}
}
