package sender

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
)

// NotifySender rmq Sender struct
type NotifySender struct {
	channel *amqp.Channel
	queue   amqp.Queue
}

// NewNotifySender create new Sender
func NewNotifySender(config utilities.Config) (*NotifySender, error) {
	conn, err := amqp.Dial(config.AMQPURI)
	if err != nil {
		log.Error("Cannot connect to RMQ:", err)
		return nil, err
	}
	ch, err := conn.Channel()
	if err != nil {
		log.Error("Cannot connect create channel:", err)
		return nil, err
	}

	queue, err := ch.QueueDeclare("event", false, false, false, false, nil)
	if err != nil {
		log.Error("Cannot connect declare queue:", err)
		return nil, err
	}

	return &NotifySender{
		channel: ch,
		queue:   queue,
	}, nil
}

// Start notify sender
func (n *NotifySender) Start() error {

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	msgs, err := n.channel.Consume(
		n.queue.Name, // queue
		"",           // consumer
		true,         // auto-ack
		false,        // exclusive
		false,        // no-local
		false,        // no-wait
		nil,          // args
	)
	if err != nil {
		log.Error("Cannot connect declare queue:", err)
	}
	go func() {
		for d := range msgs {
			var event models.Event
			err = json.Unmarshal(d.Body, &event)
			if err != nil {
				log.Error("Cannot decode json:", err)
			}
			fmt.Printf("Received a message:\n\tID: %s\n\tOwnerID: %v\n\tTitle: %s\n\tBody:%s", event.ID.String(), event.Owner, event.Title, event.Text)
		}
	}()
	fmt.Printf("Got signal from OS: %v. Exit.", <-osSignals)
	return nil
}
