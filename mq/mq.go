package mq

import (
    // "fmt"
    "log"
    "github.com/streadway/amqp"
    "logSearcher/lib"
)

var conn *amqp.Connection;
var channel *amqp.Channel;

func init() {
}

func InitConnection(url *string) () {
    var err error
    conn, err = amqp.Dial(*url)
    lib.FailOnError(err, "Failed to connect to RabbitMQ")
    // defer conn.Close()
    channel, err = conn.Channel()
    lib.FailOnError(err, "Failed to open a channel")
    // defer ch.Close()
}

func Get(name string) {
    q, err := channel.QueueDeclare(
        name, // name
        false,   // durable
        false,   // delete when unused
        false,   // exclusive
        false,   // no-wait
        nil,     // arguments
    )
    lib.FailOnError(err, "Failed to declare a queue")

    msgs, err := channel.Consume(
        q.Name, // queue
        "",     // consumer
        true,   // auto-ack
        false,  // exclusive
        false,  // no-local
        false,  // no-wait
        nil,    // args
    )
    lib.FailOnError(err, "Failed to register a consumer")

    // forever := make(chan bool)

    go func() {
        for d := range msgs {
            log.Printf("Received a message, length: %d", len(d.Body))
            factory := MessageProcesssorFactory{}
            processor := factory.GetProcessor(name)
            processor.ProcessMessage(d.Body)
        }
    }()

    // log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
    // <-forever
}
