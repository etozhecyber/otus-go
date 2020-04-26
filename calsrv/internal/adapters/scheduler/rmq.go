package scheduler

import (
	"database/sql"
	"encoding/json"
	"os"
	"os/signal"
	"syscall"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/jackc/pgx/stdlib" //psql
	log "github.com/sirupsen/logrus"

	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	"github.com/streadway/amqp"
)

// NotifyScheduler rmq scheduler struct
type NotifyScheduler struct {
	channel *amqp.Channel
	queue   amqp.Queue
	db      *sql.DB
}

// NewNotifySheduler create new sheduler
func NewNotifySheduler(config utilities.Config) (*NotifyScheduler, error) {
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

	db, err := sql.Open("pgx", config.DBDSN)
	if err != nil {
		log.Error("Cannot connect to DB:", err)
		return nil, err
	}

	return &NotifyScheduler{
		channel: ch,
		queue:   queue,
		db:      db,
	}, nil
}

// Start notify scheduler
func (n *NotifyScheduler) Start() error {

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

	done := make(chan struct{}, 1)
	query := "SELECT id,owner_id,title,text,starttime,endtime from events WHERE starttime > current_timestamp and is_notified = 'false'"
	go func() {
		ticker := time.NewTicker(time.Second * time.Duration(5))
		for {
			select {
			case <-done:
				log.Infoln("shutdown sheduler")
				return
			case <-ticker.C:
				{
					rows, err := n.db.Query(query)
					if err != nil {
						log.Error("failed get event from DB:", err)
					}
					for rows.Next() {
						var event models.Event
						if err = rows.Scan(&event.ID, &event.Owner, &event.Title, &event.Text, &event.StartTime, &event.EndTime); err != nil {
							log.Error("failed scan event:", err)
						}
						jsonEvent, err := json.Marshal(event)
						if err != nil {
							log.Error("failed create json:", err)
						}
						err = n.channel.Publish("", n.queue.Name, false, false, amqp.Publishing{
							ContentType: "application/json",
							Body:        jsonEvent,
						})
						if err != nil {
							log.Error("failed to publish:", err)
						}
						_, err = sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
							Update("events").
							Set("is_notified", true).
							Where(sq.Eq{"id": event.ID.String()}).
							RunWith(n.db).Exec()
						if err != nil {
							log.Error("failed to update:", err)
						}
					}
					rows.Close()
				}
			}
		}
	}()
	log.Infof("Got signal from OS: %v. Exit.", <-osSignals)
	close(done)
	n.db.Close()
	n.channel.Close()
	return nil
}
