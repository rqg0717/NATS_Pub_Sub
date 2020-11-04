package main

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

func main() {

	opts := []nats.Option{nats.Name("NATS Publisher")}

	conn, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	subject, msg := "test", "this is a test message."

	conn.Publish(subject, []byte(msg))

	conn.Flush()

	if err := conn.LastError(); err != nil {
		log.Fatal(err)
	} else {
		log.Printf("Published to [%s]: %s \n", subject, msg)
	}
}
