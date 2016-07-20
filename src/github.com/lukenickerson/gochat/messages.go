package main

import (
	//"log"
)

var Messages = []Message{
	{
		Name: "tester",
		Content: "test message",
	},{
		Name: "tester2",
		Content: "life is like a box of chocolates",
	},
}

type Message struct {
	Name string // IP address for now
	Content string
	// Add dateTime?
}

func addMessage(m string) {
	var msg = new(Message) // return a pointer to msg (type *msg)
	msg.Name = "You"
	msg.Content = m
	Messages = append(Messages, *msg) // use the value pointed by mgs
	//log.Println()
}

