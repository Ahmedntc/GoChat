package main

import (
	//
	//"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	//"strings"
)
var Clients = make(map[net.Conn]string)

//global bot declaration
var bt *b
type Message struct{
	from string
	payload []byte
}
type server struct {
	commands chan command

	message chan Message
}

func newServer() *server {
	s := &server{
		commands: make(chan command),
	}
	return s
}

func (s *server) run(botk *b) {
	bt = botk

	for cmd := range s.commands {
		switch cmd.id {
		case nick:
			s.nick(cmd.client, cmd.args)
		case mesg:
			s.msg(cmd.client, cmd.args)
		case priv:
			s.priv_msg(cmd.client, cmd.args)
		case quit:
			s.quit(cmd.client)
		}
	}
}

func (s *server) newClient(conn net.Conn) *client {
	log.Printf("nova conexão: %s ", conn.RemoteAddr().String(), )
	return &client{
		conn:     conn,
		nick:     " ",
		commands: s.commands,
	}

}

func (s *server) nick(c *client, args []string) {
	Clients[c.conn] = args[1]
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
	for client := range Clients {
			_, err := client.Write([]byte("\n> "+c.nick+" enviou: "+ msg))
			if err != nil {
				log.Printf("Erro")
			}
			break
	}

	rmsg := reverse(msg, bt)
	c.conn.Write([]byte(fmt.Sprintf("\nBot inverteu: %s\n", rmsg)))
}

func (s *server) priv_msg(c *client, args []string) {
	msg := strings.Join(args[1:], " ")
	
	parts := strings.SplitN(msg, ":", 2)
	if len(parts) == 2 {
	to := strings.TrimSpace(parts[0])
	payload := []byte(strings.TrimSpace(parts[1]))
	for client := range Clients {
		if Clients[client] == to {
			_, err := client.Write([]byte("> Você recebeu uma mensagem privada de "+c.nick+": "+ string(payload)))
			if err != nil {
				log.Printf("Erro")
			}
			break
		}
	}
}
	rmsg := reverse(msg, bt)
	c.conn.Write([]byte(fmt.Sprintf("\nBot inverteu: %s", rmsg)))
}




func (s *server) quit(c *client) {
	log.Printf("conexão terminada: %s", c.conn.RemoteAddr().String())
	c.conn.Close()
}





