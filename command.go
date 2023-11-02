package main

type commandID int

const (
	nick commandID = iota
	mesg
	quit
	priv
)

type command struct {
	id     commandID
	client *client
	args   []string
}