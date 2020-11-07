package main

import (
	"log"
	"os"
	"os/signal"
	"time"

	nats "github.com/nats-io/nats.go"
)

func printMsg(m *nats.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'", i, m.Subject, string(m.Data))
}

func main() {

	opts := []nats.Option{nats.Name("NATS Subscriber")}
	opts = setupConnOptions(opts)

	conn, err := nats.Connect(nats.DefaultURL, opts...)
	if err != nil {
		log.Fatal(err)
	}

	conn.Subscribe("test", func(msg *nats.Msg) {
		log.Printf("Received on [%s]: '%s'", msg.Subject, string(msg.Data))
	})
	conn.Flush()

	if err := conn.LastError(); err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Printf("Draining...")
	conn.Drain()
	log.Fatalf("Exiting")
}

func setupConnOptions(opts []nats.Option) []nats.Option {
	opts = append(opts, nats.ReconnectWait(time.Second))
	opts = append(opts, nats.MaxReconnects(3))
	opts = append(opts, nats.DisconnectErrHandler(func(conn *nats.Conn, err error) {
		log.Printf("Disconnected due to: %s ", err)
	}))
	opts = append(opts, nats.ReconnectHandler(func(conn *nats.Conn) {
		log.Printf("Reconnected to [%s]", conn.ConnectedUrl())
	}))
	opts = append(opts, nats.ClosedHandler(func(conn *nats.Conn) {
		log.Fatalf("Exiting: %v", conn.LastError())
	}))
	return opts
}
