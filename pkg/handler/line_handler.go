package handler

import (
	"athielen.com/line-ingress/pkg/backend"
	"athielen.com/line-ingress/pkg/models"
	"net/http"
)

type LineHandler struct {
	bman *backend.BackendManager
}

func (h *LineHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "POST":

		json, err := models.MarshalRequestJSON(request)
		if err == nil {
			h.bman.SendLineToChannels(json)
		}
		SendGenericSuccess(writer)

	default:
		SendStatusNotfound(writer,"Can't find method requested")
	}
}