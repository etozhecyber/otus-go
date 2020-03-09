package api

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"

	"github.com/etozhecyber/otus-go/calsrv/internal/adapters/grpc/api"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/services"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var accessLog = log.New()

func eventMapper(event models.Event) api.Event {
	starttime, _ := ptypes.TimestampProto(event.StartTime)
	endtime, _ := ptypes.TimestampProto(event.EndTime)
	return api.Event{
		ID:        event.ID.String(),
		Owner:     event.Owner,
		Title:     event.Title,
		Text:      event.Text,
		StartTime: starttime,
		EndTime:   endtime,
	}
}

//CalendarServer ...
type CalendarServer struct {
	EventUsecases *services.EventService
}

// CreateNewEvent ...
func (cs *CalendarServer) CreateNewEvent(ctx context.Context, in *api.CreateEventRequest) (*api.Result, error) {
	user, title, body := in.GetOwner(), in.GetTitle(), in.GetText()
	// starttime := time.Now()
	// endtime := time.Now().Add(1 * time.Hour)
	starttime, err := ptypes.Timestamp(in.GetStartTime())
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	endtime, err := ptypes.Timestamp(in.GetStartTime())
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	err = cs.EventUsecases.CreateEvent(ctx, user, title, body, starttime, endtime)
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	accessLog.WithFields(log.Fields{
		"operation": "CreateNewEvent",
		"user":      user,
		"title":     title,
		"body":      body,
	}).Println()

	return &api.Result{Result: "ok"}, nil
}

// UpdateEvent by ID
func (cs *CalendarServer) UpdateEvent(ctx context.Context, in *api.UpdateEventRequest) (*api.Result, error) {
	user, title, body := in.GetOwner(), in.GetTitle(), in.GetText()
	// starttime := time.Now()
	// endtime := time.Now().Add(1 * time.Hour)
	starttime, err := ptypes.Timestamp(in.GetStartTime())
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	endtime, err := ptypes.Timestamp(in.GetStartTime())
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	id, err := uuid.FromString(in.GetID())
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	event := models.Event{
		ID:        id,
		EndTime:   endtime,
		Owner:     user,
		StartTime: starttime,
		Text:      body,
		Title:     title,
	}
	err = cs.EventUsecases.UpdateEventbyID(ctx, id, event)
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	accessLog.WithFields(log.Fields{
		"operation": "UpdateEvent",
		"id":        id,
		"user":      user,
		"title":     title,
		"body":      body,
	}).Println()

	return &api.Result{Result: "ok"}, nil
}

// DeleteEvent by ID
func (cs *CalendarServer) DeleteEvent(ctx context.Context, req *api.Id) (*api.Result, error) {
	id, err := uuid.FromString(req.GetUUID())
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	err = cs.EventUsecases.DelEventbyID(ctx, id)
	if err != nil {
		log.Error(err)
		return &api.Result{Result: err.Error()}, err
	}
	accessLog.WithFields(log.Fields{
		"operation": "DeleteEvent",
		"id":        id,
	}).Println()
	return &api.Result{Result: "ok"}, nil
}

// EventForMonth get all event of current month
func (cs *CalendarServer) EventForMonth(_ *empty.Empty, stream api.Calserv_EventForMonthServer) error {
	now := time.Now()
	startmonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local) //get beginning of the current month
	endmonth := startmonth.AddDate(0, 1, 0)                                     //get end of the current month
	ctx := context.Background()
	events, err := cs.EventUsecases.GetEvents(ctx, startmonth, endmonth)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	for _, event := range events {
		gevent := eventMapper(event)
		stream.Send(&gevent)
	}
	accessLog.WithFields(log.Fields{
		"operation": "EventForMonth",
	}).Println()
	return nil
}

// EventForDay get all event of current day
func (cs *CalendarServer) EventForDay(_ *empty.Empty, stream api.Calserv_EventForDayServer) error {
	ctx := context.Background()
	now := time.Now()
	startday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local) //get beginning of the current day
	endday := startday.Add(24 * time.Hour)                                            //get end of the current day
	events, err := cs.EventUsecases.GetEvents(ctx, startday, endday)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	for _, event := range events {
		gevent := eventMapper(event)
		stream.Send(&gevent)
	}
	accessLog.WithFields(log.Fields{
		"operation": "EventForDay",
	}).Println()
	return nil
}

// EventForWeek get all event of current week
func (cs *CalendarServer) EventForWeek(_ *empty.Empty, stream api.Calserv_EventForWeekServer) error {

	ctx := context.Background()
	now := time.Now()
	nowdate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local) //get beginning of the current day
	startweek := nowdate.AddDate(0, 0, 1-int(nowdate.Weekday()))                     //get beginning of the current week
	endweek := startweek.Add(24 * 7 * time.Hour)
	events, err := cs.EventUsecases.GetEvents(ctx, startweek, endweek)
	if err != nil {
		log.Error(err.Error())
		return err
	}
	for _, event := range events {
		gevent := eventMapper(event)
		stream.Send(&gevent)
	}
	accessLog.WithFields(log.Fields{
		"operation": "EventForWeek",
	}).Println()
	return nil
}

//Serve grpc server
func (cs *CalendarServer) Serve(config utilities.Config) error {
	//init access_log
	logfile, err := os.OpenFile(config.GRPCAccessLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err.Error())
	}
	accessLog.SetOutput(logfile)

	fmt.Println("start server on:", config.GRPCListen)
	lis, err := net.Listen("tcp", config.GRPCListen)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	api.RegisterCalservServer(grpcServer, cs)
	grpcServer.Serve(lis)

	return nil
}
