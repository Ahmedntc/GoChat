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
		case CMD_NICK:
			s.nick(cmd.client, cmd.args)
		case CMD_CHN:
			s.change(cmd.client, cmd.args)
		case CMD_MSG:
			s.msg(cmd.client, cmd.args)
		case CMD_QUIT:
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
	c.nick = args[1]
	c.msg(fmt.Sprintf("%s conectou", c.nick))
}

func (s *server) change(c *client, args []string) {
	c.nick = args[1]
	c.msg(fmt.Sprintf("%s é seu novo nick", c.nick))
}


func (s *server) msg(c *client, args []string) {
	
	msg := strings.Join(args[1:], " ")

	c.conn.Write([]byte(fmt.Sprintf("%s: %s", c.nick, msg)))
	log.Printf("User %s enviou: %s", c.nick, msg)

	msgRev := reverse(msg)
	c.conn.Write([]byte(fmt.Sprintf("\nServer replied : %s",msgRev)))

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



