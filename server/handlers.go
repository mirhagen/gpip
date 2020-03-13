package server

import (
	"net/http"
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
		contentType := headerValues(r.Header.Get("accept"))[0]

		var response func(w http.ResponseWriter, s string)
		if len(contentType) == 0 || contentType == "application/json" || contentType == "*/*" {
			response = jsonResponse
		} else if contentType == "text/plain" {
			response = textResponse
		} else {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			return
		}

		ip := remoteIPAddress(r)

		w.Header().Set("Content-Type", contentType+"; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		response(w, ip)
	})
}
