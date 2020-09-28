package ip

import (
	"net/http"
	"regexp"
	"strings"
)

// Resolve parses incoming request and resolves real IP
// by looking for headers Forwarded, X-Forwarded-For or
// X-Real-IP and if none of thos exists uses RemoteAddr.
func Resolve(r *http.Request) string {
	var rAddr string
	if len(r.Header.Get("forwarded")) > 0 {
		rAddr = forwardedFor(r.Header.Get("forwarded"))[0]
	} else if len(r.Header.Get("x-forwarded-for")) > 0 {
		rAddr = strings.Split(strings.Replace(r.Header.Get("x-forwarded-for"), " ", "", -1), ",")[0]
	} else if len(r.Header.Get("x-real-ip")) > 0 {
		rAddr = r.Header.Get("x-real-ip")
	} else {
		rAddr = r.RemoteAddr
	}

	reg := regexp.MustCompile(`:\d+$`)
	ip := reg.ReplaceAllLiteralString(rAddr, "")
	return ip
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
