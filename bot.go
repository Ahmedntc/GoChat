package main

import (
	"fmt"
	"log"

	"net"

)

type b struct {
	conne     net.Conn
	nick     string
}



func bot() *b {
	address:= "0.0.0.0:8888"
	conn, err := net.Dial("tcp", address)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	fmt.Println("Bot On", conn.RemoteAddr().String())


	return &b{
		conne:     conn,
		nick:     "bot",
	}
}

func reverse(s string, b *b) string { 

	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
		rns[i], rns[j] = rns[j], rns[i] 
	} 
	return string(rns) 
} 
