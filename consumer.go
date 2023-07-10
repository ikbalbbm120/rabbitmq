package main

import(
	"fmt"
	"github.com/streadway/amqp"
)

func main() {
	fmt.Println("consumer application")
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"testqueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	forever := make(chan bool )
	go func() {
		for d := range msgs {
			fmt.Println("recieved message: %s\n", d.Body)
		}
	}()
	fmt.Println("success")
	fmt.Println("[*]- waiting for massage")
	<-forever
}