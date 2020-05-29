package handler

import (
	"fmt"
	"net/http"
)

// Send generic success with 200 response code
func SendGenericSuccess(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"status":"success"}`))

}

//
func SendSuccessUp(writer http.ResponseWriter) {
	// Set the return Content-Type as JSON like before
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte(`{"status":"up"}`))
}

// Send generic success with 404
func SendStatusNotfound(writer http.ResponseWriter, message string) {
	writer.WriteHeader(http.StatusNotFound)
	writer.Write([]byte(fmt.Sprintf(`{"message": "%s"}`, message)))
}
