package main

import (
	"log"

	nats "github.com/nats-io/nats.go"
)

func main() {

	opts := []nats.Option{nats.Name("NATS Publisher")}

	conn, error := nats.Connect(nats.DefaultURL, opts...)
	if error != nil {
		log.Fatal(error)
	}
	defer conn.Close()

	subject, msg := "test", "this is a test message."

	conn.Publish(subject, []byte(msg))

	conn.Flush()

	if error := conn.LastError(); error != nil {
		log.Fatal(error)
	} else {
		log.Printf("Published to [%s]: %s \n", subject, msg)
	}
}
