package ip

import (
	"net/http"
	"testing"
)

func TestResolve(t *testing.T) {
	var tests = []struct {
		header string
		value  string
		want   string
	}{
		{"Forwarded", "proto=https; for=192.168.0.1, for=192.168.0.2; host=somehost", "192.168.0.1"},
		{"Forwarded", "for=192.168.0.2", "192.168.0.2"},
		{"X-Forwarded-For", "192.168.0.2", "192.168.0.2"},
		{"X-Forwarded-For", "192.168.0.10, 192.168.0.11", "192.168.0.10"},
		{"X-Real-IP", "192.168.0.5", "192.168.0.5"},
		{"", "", "192.168.0.1"},
	}

	header := http.Header{}
	req := &http.Request{RemoteAddr: "192.168.0.1:55999", Header: header}
	for _, test := range tests {
		req.Header.Set(test.header, test.value)
		got := Resolve(req)
		if got != test.want {
			t.Errorf("incorrect value, got: %s, want: %s", got, test.want)
		}
		req.Header.Set(test.header, "")
	}
}

func TestForwardedFor(t *testing.T) {
	expected := "192.168.0.1"

	var tests = []struct {
		input string
		want  string
	}{
		{input: "proto=https; for=192.168.0.1, for=192.168.0.2; host=somehost", want: expected},
		{input: "for=192.168.0.1", want: expected},
	}

	for _, test := range tests {
		got := forwardedFor(test.input)
		if got[0] != test.want {
			t.Errorf("incorrect value, got: %s, want: %s", got[0], test.want)
		}
	}
}
