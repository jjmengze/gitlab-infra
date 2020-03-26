package main

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

func main() {
	httpServer := &http.Server{Addr: ":" + "8080"}

	logrus.WithError(httpServer.ListenAndServe()).Info("Server exited.")
}
