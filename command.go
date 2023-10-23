package main

type commandID int

const (
	CMD_NICK commandID = iota
	CMD_CHN
	CMD_MSG
	CMD_QUIT
)

type command struct {
	id     commandID
	client *client
	args   []string
}