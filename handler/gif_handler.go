package handler

import (
	"encoding/base64"
	"line-ingress/backend"
	"line-ingress/models"
	"net/http"
)

type GifHandler struct {
	bman *backend.BackendManager
}

func (h *GifHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case "GET":
		SendSuccessfulGifPixel(writer)
		json, err := models.MarshalRequestJSON(request)

		if err == nil {
			h.bman.SendLineToChannels(json)
		}

	default:
		SendStatusNotfound(writer, "Can't find method requested")
	}

}

func SendSuccessfulGifPixel(writer http.ResponseWriter) {
	// headers to prevent caching
	writer.Header().Set("Content-Type", "image/gif")
	writer.Header().Set("Expires", "Mon, 01 Jan 1990 00:00:00 GMT")
	writer.Header().Set("Cache-Control", "no-store")
	writer.Header().Set("Pragma", "no-cache")

	// response, 1x1 px transparent GIF
	writer.WriteHeader(http.StatusOK)
	b, _ := base64.StdEncoding.DecodeString("R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7")
	writer.Write(b)
}
