package handler

import (
	"net/http"
)

type UpHandler struct {
}

func (handler *UpHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	// Change the response depending on the method being requested
	switch request.Method {
	case "GET":
		SendSuccessUp(writer)
	default:
		SendStatusNotfound(writer, "Can't find method requested")
	}

}
