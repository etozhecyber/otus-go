package httpapi

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/etozhecyber/otus-go/calsrv/internal/domain/models"
	"github.com/etozhecyber/otus-go/calsrv/internal/domain/services"
	"github.com/etozhecyber/otus-go/calsrv/utilities"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

var accessLog = log.New()

//HTTPServer HTTP server
type HTTPServer struct {
	EventUsecases *services.EventService
}

//Serve start http server
func (hs *HTTPServer) Serve(config utilities.Config) error {

	//init access_log
	logfile, err := os.OpenFile(config.AccessLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Error(err.Error())
	}
	accessLog.SetOutput(logfile)

	http.Handle("/create_event", accessLogger(validator(http.HandlerFunc(hs.createEventHandler))))
	http.Handle("/update_event", accessLogger(validator(http.HandlerFunc(hs.updateEventHandler))))
	http.Handle("/delete_event", accessLogger(http.HandlerFunc(hs.deleteEventHandler)))
	http.Handle("/events_for_day", accessLogger(http.HandlerFunc(hs.eventsForDayHandler)))
	http.Handle("/events_for_week", accessLogger(http.HandlerFunc(hs.eventsForWeekHandler)))
	http.Handle("/events_for_month", accessLogger(http.HandlerFunc(hs.eventsForMonthHandler)))
	log.Info("start http server on:", config.HTTPListen)
	err = http.ListenAndServe(config.HTTPListen, nil)
	return err
}

func (hs *HTTPServer) eventsForDayHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error := httpError{"Method not allowed"}
		error.jsonPrint(w, http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()
	now := time.Now()
	startday := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local) //get beginning of the current day
	endday := startday.Add(24 * time.Hour)                                            //get end of the current day
	events, err := hs.EventUsecases.GetEvents(ctx, startday, endday)
	if err != nil {
		log.Error(err.Error())
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusInternalServerError)
		return
	}
	httpEv := httpEvent{events}
	data, err := httpEv.toJSON()
	if err != nil {
		log.Error(err.Error())
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(data))
}

func (hs *HTTPServer) eventsForWeekHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error := httpError{"Method not allowed"}
		error.jsonPrint(w, http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()
	now := time.Now()
	nowdate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local) //get beginning of the current day
	startweek := nowdate.AddDate(0, 0, 1-int(nowdate.Weekday()))                     //get beginning of the current week
	endweek := startweek.Add(24 * 7 * time.Hour)                                     //get end of the current week
	events, err := hs.EventUsecases.GetEvents(ctx, startweek, endweek)
	if err != nil {
		log.Error(err.Error())
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusInternalServerError)
		return
	}
	httpEv := httpEvent{events}
	data, err := httpEv.toJSON()
	if err != nil {
		log.Error(err.Error())
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(data))
}

func (hs *HTTPServer) eventsForMonthHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		error := httpError{"Method not allowed"}
		error.jsonPrint(w, http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()
	now := time.Now()
	startweek := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local) //get beginning of the current month
	endweek := startweek.AddDate(0, 1, 0)                                      //get end of the current month
	events, err := hs.EventUsecases.GetEvents(ctx, startweek, endweek)
	if err != nil {
		log.Error(err.Error())
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusInternalServerError)
		return
	}
	httpEv := httpEvent{events}
	data, err := httpEv.toJSON()
	if err != nil {
		log.Error(err.Error())
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(data))
}

func (hs *HTTPServer) createEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		error := httpError{"Method not allowed"}
		error.jsonPrint(w, http.StatusMethodNotAllowed)
		return
	}
	user, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		log.Error("cannot convert user to int64:", err)
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}
	ctx, title, body := context.Background(), r.FormValue("title"), r.FormValue("body")

	starttime, err := time.Parse(time.RFC3339, r.FormValue("starttime"))
	if err != nil {
		log.Error("cannot convert starttime to time", err)
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}

	endtime, err := time.Parse(time.RFC3339, r.FormValue("endtime"))
	if err != nil {
		log.Error("cannot convert endtime to time", err)
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}

	err = hs.EventUsecases.CreateEvent(ctx, user, title, body, starttime, endtime)
	if err != nil {
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusOK)
		return
	}
	result := httpResponce{"OK"}
	result.jsonPrint(w, http.StatusCreated)
}

func (hs *HTTPServer) updateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		error := httpError{"Method not allowed"}
		error.jsonPrint(w, http.StatusMethodNotAllowed)
		return
	}
	user, err := strconv.Atoi(r.FormValue("user"))
	if err != nil {
		log.Error("cannot convert user to int64:", err)
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}
	ctx, title, body := context.Background(), r.FormValue("title"), r.FormValue("body")
	id, err := uuid.FromString(r.FormValue("id"))
	if err != nil {
		log.Error("cannot convert id to UUID:", err)
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}

	starttime, err := time.Parse(time.RFC3339, r.FormValue("starttime"))
	if err != nil {
		log.Error("cannot convert starttime to time:", err)
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}

	endtime, err := time.Parse(time.RFC3339, r.FormValue("endtime"))
	if err != nil {
		log.Error("cannot convert endtime to time:", err)
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusOK)
		return
	}

	event := models.Event{
		ID:        id,
		EndTime:   endtime,
		Owner:     user,
		StartTime: starttime,
		Text:      body,
		Title:     title,
	}
	err = hs.EventUsecases.UpdateEventbyID(ctx, id, event)
	if err != nil {
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusOK)
		return
	}
	result := httpResponce{"OK"}
	result.jsonPrint(w, http.StatusCreated)
	return
}

func (hs *HTTPServer) deleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		error := httpError{"Method not allowed"}
		error.jsonPrint(w, http.StatusMethodNotAllowed)
		return
	}
	ctx := context.Background()
	if r.FormValue("id") == "" {
		error := httpError{"parameter required: id"}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}
	id, err := uuid.FromString(r.FormValue("id"))
	if err != nil {
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusBadRequest)
		return
	}
	err = hs.EventUsecases.DelEventbyID(ctx, id)
	if err != nil {
		error := httpError{err.Error()}
		error.jsonPrint(w, http.StatusOK)
		return
	}
	result := httpResponce{"ok"}
	result.jsonPrint(w, http.StatusAccepted)
	return
}
