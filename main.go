package main

import (
	"line-ingress/logging"
	"line-ingress/backend"
	"line-ingress/handler"
	log "github.com/sirupsen/logrus"
)

func main() {

	log.Info("Starting line ingress application...")
	logging.Init()

	// Copnfigure and build Backends and their Manager
	bman := backend.BuildBackendManager()
	bman.StartBackends()

	// Configure Handlers and pass them Backend Manager
	handler.ConfigureHandlers(&bman)
}
