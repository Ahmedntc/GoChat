package main

type commandID int

const (
	nick commandID = iota
	mesg
	quit
)

type command struct {
	id     commandID
	client *client
	args   []string
}