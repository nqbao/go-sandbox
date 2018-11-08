package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	c, err := net.Dial("tcp", "localhost:3333")

	if err != nil {
		log.Fatal(err)
	}

	defer c.Close()

	// make some noise
	go func() {
		for {
			c.Write([]byte("hello\n"))
			time.Sleep(5 * time.Second)
		}
	}()

	for {
		// c.SetReadDeadline(time.Now().Add(2 * time.Second))
		l := 10
		b := make([]byte, l)
		n, err := c.Read(b)

		if err != nil {
			log.Printf("%v", err)
		}

		fmt.Printf("%s %v\n", b, n)
	}
}
