package main

import (
	"bufio"
	//"fmt"
	"net"
	"strings"
)

type client struct {
	conn     net.Conn
	nick     string
	commands chan<- command
}

func (c *client) command() {
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n')
		if err != nil {
			return
		}
		msg = strings.Trim(msg, "\r\n")
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0])
		switch cmd {
		case "/nick":
			c.commands <- command{
				id:    nick,
				client: c,
				args:   args,
			}
		case "/msg":
			c.commands <- command{
				id:     mesg,
				client: c,
				args:   args,
			}
		case "/priv":
			c.commands <- command{
				id:     priv,
				client: c,
				args:   args,
			}
		case "/quit":
			c.commands <- command{
				id:     quit,
				client: c,
			}
		default:
			c.conn.Write([]byte(">Comando Desconhecido\n"))
		}
	}
}






