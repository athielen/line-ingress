package models

import (
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"time"
)

// Line corresponds to an incoming tracking request and contains the data
// to be persisted.
type Line struct {
	Payload []byte
	JSON string
	Time time.Time
	SourceType string
}

func (line Line) String() string {
	s, err := json.Marshal(line)
	if err == nil {
		return fmt.Sprintf("%s", string(s))
	} else {
		return "error"
	}
}

func MarshalRequestJSON(request *http.Request) ([]byte, error) {

	var queryParams, err = url.ParseQuery(request.URL.RawQuery)

	// add error message and prometheus counter
	if err != nil {
		log.Warn("Failed to marshal query parameters. Setting as blank.")
		queryParams = make(url.Values)
	}

	return json.Marshal(
		WebEventFeatures {
			Method:          request.Method,
			URL:             request.URL,
			Host:            request.Host,
			Form:            request.Form,
			RemoteAddr:      request.RemoteAddr,
			IpAddress:	     request.Header.Get("X-FORWARDED-FOR"),
			RequestURI:      request.RequestURI,
			QueryParams:     queryParams,
			Origin:          request.Header.Get("Origin"),
			Browser:         request.UserAgent(),
			OperatingSystem: "",
			DeviceType:      "",
			VisitorCountry:  "",
			InitialReferrer: request.Referer(),
		})

}
