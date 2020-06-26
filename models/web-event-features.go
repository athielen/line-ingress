package models

import "net/url"

type WebEventFeatures struct {
	Method                 string
	URL                    *url.URL
	Host                   string
	Form                   url.Values
	RemoteAddr             string
	IpAddress              string
	RequestURI             string
	QueryParams            url.Values
	Origin                 string
	Browser                string
	BrowserVersion         string
	OperatingSystem        string
	OperatingSystemVersion string
	DeviceType             string
	Device                 string
	VisitorCountry         string
	InitialReferrer        string
}
