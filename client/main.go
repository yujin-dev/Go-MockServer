package main

import (
	"log"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8000")
	if err != nil {
		log.Println(err)
	}

	go connHanlder(conn)

	_, Cerr := conn.Write([]byte("config"))
	if Cerr == nil {
		log.Println("Successfully sent")
	}
	defer conn.Close()
}

func connHanlder(conn net.Conn) {
	data := make([]byte, 4096)
	for {
		count, err := conn.Read(data)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Server send : ", string(data[:count]))
		time.Sleep(time.Duration(3) * time.Second)
	}
}
