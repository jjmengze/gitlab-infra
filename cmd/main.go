package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	server := &hook.Server{}
	// Return 200 on / for health checks.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {})
	// For /hook, handle a webhook normally.
	http.Handle("/hook", server)
	httpServer := &http.Server{Addr: ":" + "8080"}

	logrus.WithError(httpServer.ListenAndServe()).Info("Server exited.")
}
