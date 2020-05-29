package main

import (
	"line-ingress/backend"
	"line-ingress/handler"
)

func main() {

	// Copnfigure and build Backends and their Manager
	bman := backend.BuildBackendManager()
	bman.StartBackends()

	// Configure Handlers and pass them Backend Manager
	handler.ConfigureHandlers(&bman)
}
