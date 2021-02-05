package socket

import (
	"io"
	"io/ioutil"
	"log"
	"net"
	"os"

	"github.com/YJ-dev/go-server/util"
)

func Run() {
	lister, err := net.Listen("tcp", ":8000") // port bind -> listen
	util.CheckErr(err)
	defer lister.Close()

	for {
		connection, err := lister.Accept() // accept
		if err != nil {
			log.Println("Connection Failed. ", err)
			continue
		} else {
			log.Println("Connection Succeeded. ", connection.RemoteAddr())
		}
		defer connection.Close() // close
		go sendRecv(connection)  // send , recv

	}
}

func sendRecv(conn net.Conn) {
	recvBuf := make([]byte, 4096)
	for {
		count, err := conn.Read(recvBuf)
		if err != nil {
			if err == io.EOF { // error by Read( no more input is available )
				log.Println(err)
				return
			}
			log.Println(err)
			return
		}
		if count > 0 {
			req := recvBuf[:count]
			reqType(string(req), conn)
		}
	}
}

func reqType(req string, conn net.Conn) {

	switch req {
	case "config":
		jsonFile, err := os.Open("handlers/config.json")
		if err != nil {
			panic(err)
		}
		defer jsonFile.Close()
		byteValue, _ := ioutil.ReadAll(jsonFile)
		_, Werr := conn.Write(byteValue)
		if Werr != nil {
			log.Println(Werr)
			return
		}
	}
}
