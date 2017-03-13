package main

import (
	"fca/libs"
	"fca/libs/bufio"
)

type Channel struct {
	signal chan *libs.Proto
	Writer bufio.Writer
	Reader bufio.Reader
	UUID   string
	UDType uint16
}

func NewChannel(uuid string) *Channel {
	channel := &Channel{}
	channel.UUID = uuid
	channel.signal = make(chan *libs.Proto, 10000)
	return channel
}
