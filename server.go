package main

import (
	"fmt"
	"log"
	"net"
	"strings"
	//"strings"
)
var Clients = make(map[string]chan []byte)

type Message struct{
	from string
	payload []byte
}
type server struct {
	commands chan command

	message chan Message
}

func newServer() *server {
	return &server{
		commands: make(chan command),
	}
}

func (s *server) run() {
	for cmd := range s.commands {
		switch cmd.id {
		case nick:
			s.nick(cmd.client, cmd.args)
		case mesg:
			s.msg(cmd.client, cmd.args)
		case quit:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("nova conexão: %s", conn.RemoteAddr().String())

	return &client{
		conn:     conn,
		nick:     " ",
		commands: s.commands,
	}
}

func (s *server) nick(c *client, args []string) {
	if len(args) < 2 {
		c.conn.Write([]byte("> Nick não pode ser vazio\n"))
		return
	}
	if c.nick != " "{
		c.nick = args[1]
		c.conn.Write([]byte(fmt.Sprintf("> %s é seu novo nick\n",c.nick)))
	}else{
		c.nick = args[1]
		c.conn.Write([]byte(fmt.Sprintf("> %s conectou\n",c.nick)))
	}


}




func (s *server) msg(c *client, args []string) {
	
	msg := strings.Join(args[1:], " ")

	c.conn.Write([]byte(fmt.Sprintf("%s: %s", c.nick, msg)))
	log.Printf("User %s enviou: %s\n", c.nick, msg)

	msgRev := reverse(msg)
	c.conn.Write([]byte(fmt.Sprintf("\nServer replied : %s",msgRev)))

}


func(s *server) bot(){
	conn, err := net.Dial("tcp", ":8888")
	if err != nil{
		return
	}
	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')
		
	}
	
	
	
}

func (s *server) quit(c *client) {
	log.Printf("conexão terminada: %s", c.conn.RemoteAddr().String())
	c.conn.Close()
}

func reverse(s string) string { 
    rns := []rune(s)
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  

        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    return string(rns) 
} 



