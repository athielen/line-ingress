package models

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func MarshalRequestJSON(request *http.Request) ([]byte, error) {
	var queryParams, err = url.ParseQuery(request.URL.RawQuery)

	if err != nil {
		queryParams = make(url.Values)
	}

	return json.Marshal(
		WebEventFeatures{
			Method:          request.Method,
			URL:             request.URL,
			Host:            request.Host,
			Form:            request.Form,
			RemoteAddr:      request.RemoteAddr,
			RequestURI:      request.RequestURI,
			QueryParams:     queryParams,
			Origin:          request.Header.Get("Origin"),
			Browser:         request.UserAgent(),
			OperatingSystem: "",
			DeviceType:      "",
			VisitorCountry:  "840",
			InitialReferrer: request.Referer(),
		})
}
