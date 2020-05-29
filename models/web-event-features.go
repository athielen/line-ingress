package models

import "net/url"

type WebEventFeatures struct {
	Method          string
	URL             *url.URL
	Host            string
	Form            url.Values
	RemoteAddr      string
	RequestURI      string
	QueryParams     url.Values
	Origin          string
	Browser         string
	OperatingSystem string
	DeviceType      string
	VisitorCountry  string
	InitialReferrer string
}
