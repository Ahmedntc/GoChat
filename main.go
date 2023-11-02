package main

import (
	"log"
	"net"
)

func main() {
	address:= "0.0.0.0:8888"
	s := newServer()


	l, err := net.Listen("tcp", address)
	bot := bot()
	go s.run(bot)

	if err != nil {
		log.Fatalf("falha em subir servidor: %s", err.Error())
	}

	defer l.Close()
	log.Printf("server aberto em %s\n", address)

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Printf("falha ao conectar: %s", err.Error())
			continue
		}
		c := s.newClient(conn)

		go c.command()
	}
}






