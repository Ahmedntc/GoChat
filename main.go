package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("falha em subir servidor: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("server aberto em :8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("falha ao conectar: %s", err.Error())
			continue
		}

		c := s.newClient(conn)
		go c.command()
	}
}






