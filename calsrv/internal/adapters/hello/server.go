package hello

import (
	"fmt"
	"net/http"

	"github.com/etozhecyber/otus-go/calsrv/utilities"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func helloServer(w http.ResponseWriter, r *http.Request) {
	log.WithFields(logrus.Fields{
		"method":        r.Method,
		"URI":           r.RequestURI,
		"RemoteAddress": r.RemoteAddr,
	}).Debug()
	fmt.Fprintf(w, "Hello!")
}

//Serve hello server
func Serve(config utilities.Config) error {
	http.HandleFunc("/hello", helloServer)
	log.Info("start hello server on:", config.HTTPListen)
	http.ListenAndServe(config.HTTPListen, nil)
	return nil
}
