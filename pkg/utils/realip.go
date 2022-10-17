package utils

import (
	"github.com/aesoper101/kratos-utils/internal/constants"
	"net"
	"net/http"
	"strings"
)

var trueClientIP = http.CanonicalHeaderKey(constants.HeaderTrueClientIP)
var xForwardedFor = http.CanonicalHeaderKey(constants.HeaderXForwardedFor)
var xRealIP = http.CanonicalHeaderKey(constants.HeaderXRealIP)

func RealIP(r *http.Request) string {
	var ip string

	if tcip := r.Header.Get(trueClientIP); tcip != "" {
		ip = tcip
	} else if xrip := r.Header.Get(xRealIP); xrip != "" {
		ip = xrip
	} else if xff := r.Header.Get(xForwardedFor); xff != "" {
		i := strings.Index(xff, ",")
		if i == -1 {
			i = len(xff)
		}
		ip = xff[:i]
	}
	if ip == "" || net.ParseIP(ip) == nil {
		return ""
	}
	return ip
}
