package models

import (
	"encoding/json"
	"fmt"
	uaParser "github.com/mileusna/useragent"
	log "github.com/sirupsen/logrus"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Line corresponds to an incoming tracking request and contains the data
// to be persisted.
type Line struct {
	Payload    []byte
	JSON       string
	Time       time.Time
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

	queryParams, err := url.ParseQuery(request.URL.RawQuery)

	// add error message and prometheus counter
	if err != nil {
		log.Warn("Failed to marshal query parameters. Setting as blank.")
		queryParams = make(url.Values)
	}

	ipAddress, err := getIP(request)
	if err != nil {
		log.Warn(err.Error())
	}

	ua, dt := parseUserAgent(request)

	fmt.Println(request)

	return json.Marshal(
		WebEventFeatures{
			Method:                 request.Method,
			URL:                    request.URL,
			Host:                   request.Host,
			Form:                   request.Form,
			RemoteAddr:             request.RemoteAddr,
			IpAddress:              ipAddress,
			RequestURI:             request.RequestURI,
			QueryParams:            queryParams,
			Origin:                 request.Header.Get("Origin"),
			Browser:                ua.Name,
			BrowserVersion:         ua.Version,
			OperatingSystem:        ua.OS,
			OperatingSystemVersion: ua.OSVersion,
			DeviceType:             dt,
			Device:                 ua.Device,
			VisitorCountry:         "",
			InitialReferrer:        request.Referer(),
		})

}

func parseUserAgent(r *http.Request) (uaParser.UserAgent, string) {

	ua := uaParser.Parse(r.UserAgent())

	var deviceType = ""

	if ua.Mobile {
		deviceType = "mobile"
	} else if ua.Tablet {
		deviceType = "tablet"
	} else if ua.Desktop {
		deviceType = "desktop"
	} else if ua.Bot {
		deviceType = "bot"
	}

	return ua, deviceType
}

func getIP(r *http.Request) (string, error) {

	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}
	return "", fmt.Errorf("No valid ip found")
}
