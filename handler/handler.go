package handler

import (
	"line-ingress/backend"
	"net/http"
)

// Configures handles for
func ConfigureHandlers(bman *backend.BackendManager) {

	http.Handle("/up", &UpHandler{})
	http.Handle("/line", &LineHandler{bman})
	http.Handle("/example.gif", &GifHandler{bman})
	http.Handle("/api/event", &LineHandler{bman})

	http.ListenAndServe(":9090", nil)
}
