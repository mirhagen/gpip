package server

import (
	"net/http"
	"regexp"
	"strings"
)

// remoteIPAddress parses incoming request and looks for header
// Forwarded, X-Forwarded-For or X-Real-IP and if none of those exists
// uses RemoteAddr.
func remoteIPAddress(r *http.Request) string {
	var rAddr string
	if len(r.Header.Get("forwarded")) > 0 {
		rAddr = forwardedFor(r.Header.Get("forwarded"))[0]
	} else if len(r.Header.Get("x-forwarded-for")) > 0 {
		rAddr = headerValues(r.Header.Get("x-forwarded-for"))[0]
	} else if len(r.Header.Get("x-real-ip")) > 0 {
		rAddr = r.Header.Get("x-real-ip")
	} else {
		rAddr = r.RemoteAddr
	}

	reg := regexp.MustCompile(`:\d+$`)
	ip := reg.ReplaceAllLiteralString(rAddr, "")
	return ip
}

// headerValues takes a header string and returns
// a slice containing the previously comma separated
// values.
func headerValues(header string) []string {
	return strings.Split(strings.Replace(header, " ", "", -1), ",")
}

// forwardedFor parses the content from a Forwarded header.
// and returns a slice containing the entries containing 'for='.
// Example: proty=https; for=<ip1>, for=<ip2> will return
// a slice containing ip1, ip2.
func forwardedFor(header string) []string {
	h := strings.Replace(header, " ", "", -1)
	parts := strings.Split(h, ";")

	var forwarded []string
	for _, part := range parts {
		if strings.HasPrefix(part, "for=") {
			forwarded = append(forwarded, strings.Split(strings.Split(part, ",")[0], "=")[1])
		}
	}
	return forwarded
}
