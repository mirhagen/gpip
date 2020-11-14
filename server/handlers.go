package server

import (
	"net/http"
	"strings"

	"github.com/RedeployAB/gpip/ip"
)

// getIP handles incoming requests. With no provided headers, Accept: */*,
// and Accept: application/json it will send the response body in JSON.
// Response body: {"ip":"<ip-address>"}.
// With Accept: text/plain it will send the response body in plain text.
// Response body: "<ip-address>".
func (s *Server) getIP() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// Check accept header.
		contentType := "application/json"
		accepts := strings.Split(strings.Replace(r.Header.Get("accept"), " ", "", -1), ",")[0]
		var response func(w http.ResponseWriter, s string)
		switch accepts {
		case "application/json", "*/*", "":
			response = jsonResponse
		case "text/plain":
			response = textResponse
			contentType = accepts
		default:
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		w.Header().Set("Content-Type", contentType+"; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response(w, ip.Resolve(r))
	})
}
