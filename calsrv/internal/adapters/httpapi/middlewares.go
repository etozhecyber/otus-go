package httpapi

import (
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

func accessLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		accessLog.WithFields(log.Fields{
			"method":        r.Method,
			"URI":           r.RequestURI,
			"RemoteAddress": r.RemoteAddr,
			"Fields":        r.Form,
		}).Println()
		next.ServeHTTP(w, r)
	})
}

func validator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.FormValue("user")
		if user == "" {
			error := httpError{"parameter required: user"}
			error.jsonPrint(w, http.StatusBadRequest)
			return
		}

		title := r.FormValue("title")
		if title == "" {
			error := httpError{"parameter required: title"}
			error.jsonPrint(w, http.StatusBadRequest)
			return
		}

		if r.FormValue("starttime") == "" {
			error := httpError{"parameter required: starttime"}
			error.jsonPrint(w, http.StatusBadRequest)
			return
		}

		_, err := time.Parse(time.RFC3339, r.FormValue("starttime"))
		if err != nil {
			error := httpError{err.Error()}
			error.jsonPrint(w, http.StatusBadRequest)
			return
		}

		if r.FormValue("endtime") == "" {
			error := httpError{"parameter required: endtime"}
			error.jsonPrint(w, http.StatusBadRequest)
			return
		}
		_, err = time.Parse(time.RFC3339, r.FormValue("endtime"))
		if err != nil {
			error := httpError{err.Error()}
			error.jsonPrint(w, http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	})
}
