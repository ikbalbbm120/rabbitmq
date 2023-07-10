package main

import (
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("go rabbitmq tutorial")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("succes")

	ch, err := conn.Channel() 
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"testqueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println(q)
	
	err = ch.Publish(
		"",
		"testqueue",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte("hello world"),
		},
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("success")

}

