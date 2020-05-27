package main

import (
	"athielen.com/line-ingress/pkg/backend"
	"athielen.com/line-ingress/pkg/handler"
)

func main() {

	// Copnfigure and build Backends and their Manager
	bman := backend.BuildBackendManager()
	bman.StartBackends()

	// Configure Handlers and pass them Backend Manager
	handler.ConfigureHandlers(&bman)
}