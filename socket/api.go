package socket

import (
	"io"
	"log"
	"net"

	"github.com/YJ-dev/go-server/util"
)

func Run() {
	hello()
}

func hello() {
	lister, err := net.Listen("tcp", ":8000")
	util.CheckErr(err)
	defer lister.Close()

	for { // 무한 루프
		conn, err := lister.Accept()
		if err != nil {
			log.Println("Connection failed: ", err)
			continue
		} else {
			log.Println("Connection Succeeded: ", conn.RemoteAddr())
		}
		defer conn.Close()
		go connHanlder(conn)
	}
}

func connHanlder(conn net.Conn) {
	recvBuf := make([]byte, 4096) // 값을 저장할 버퍼( 정해진 byte 까지만 읽을 수 있음 : 4KB)
	for {                         // 무한 루프
		count, err := conn.Read(recvBuf)
		if err != nil {
			if err == io.EOF {
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		if 0 < count {
			data := recvBuf[:count]
			log.Println(string(data))
			_, err := conn.Write(data[:count])
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
