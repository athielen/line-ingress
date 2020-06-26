package handler

import (
	log "github.com/sirupsen/logrus"
	"line-ingress/backend"
	"net/http"
)

// Configures handles for
func ConfigureHandlers(bman *backend.BackendManager) {

	log.Debug("Initializing handlers...")

	http.Handle("/up", &UpHandler{})
	http.Handle("/line", &LineHandler{bman})
	http.Handle("/example.gif", &GifHandler{bman})
	http.Handle("/api/event", &LineHandler{bman})

	log.Info("Listening and serving on port :9090")
	http.ListenAndServe(":9090", nil)
}
