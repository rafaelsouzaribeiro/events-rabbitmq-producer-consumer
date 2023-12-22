package main

import (
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/rafaelsouzaribeiro/fcutils/pkg/rabbitmq"
)

func main() {
	ch, err := rabbitmq.OpenChannel()

	if err != nil {
		panic(err)
	}

	defer ch.Close()

	// tudo que vai chegando vai jogando no channel e no for
	msgs := make(chan amqp.Delivery)
	go rabbitmq.Consume(ch, msgs, "minhafila")
	for msgs := range msgs {
		fmt.Println(string(msgs.Body))
		// mensagem ja foi lida e não colocar na linha
		msgs.Ack(false)
	}

}
