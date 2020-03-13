package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupServer() Server {
	srv := Server{
		router: http.NewServeMux(),
	}
	srv.routes()
	return srv
}

func TestGetIP(t *testing.T) {
	expected := "192.168.0.1"
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = expected + ":57000"
	res := httptest.NewRecorder()
	SetupServer().router.ServeHTTP(res, req)

	expectedCode := 200
	if res.Code != expectedCode {
		t.Errorf("status code incorrect, got: %d, want: %d", res.Code, expectedCode)
	}

	var rb responseBody
	b, err := ioutil.ReadAll(res.Body)
	if err = json.Unmarshal(b, &rb); err != nil {
		t.Errorf("error in test: %v", err)
	}

	if rb.IP != expected {
		t.Errorf("response incorrect, got: %s, want: %s", rb.IP, expected)
	}
}

func TestGetIPWrongMethod(t *testing.T) {
	expected := "192.168.0.1"
	req, _ := http.NewRequest("POST", "/", nil)
	req.RemoteAddr = expected + ":57000"
	res := httptest.NewRecorder()
	SetupServer().router.ServeHTTP(res, req)

	expectedCode := 404
	if res.Code != expectedCode {
		t.Errorf("status code incorrect, got: %d, want: %d", res.Code, expectedCode)
	}
}

func TestGetIPForwardHeaders(t *testing.T) {
	expected := "192.168.0.1"
	var tests = []struct {
		header   string
		value    string
		want     string
		wantCode int
	}{
		{header: "X-Forwarded-For", value: "192.168.0.1, 192.168.0.2", want: expected, wantCode: 200},
		{header: "Forwarded", value: "proto=https; for=192.168.0.1, for=192.168.0.2; host=somehost", want: expected, wantCode: 200},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = "192.168.0.5" + ":57000"
		req.Header.Set(test.header, test.value)
		res := httptest.NewRecorder()
		SetupServer().router.ServeHTTP(res, req)

		if res.Code != test.wantCode {
			t.Errorf("status code incorrect, got: %d, want: %d", res.Code, test.wantCode)
		}

		var rb responseBody
		b, err := ioutil.ReadAll(res.Body)
		if err = json.Unmarshal(b, &rb); err != nil {
			t.Errorf("error in test: %v", err)
		}

		if rb.IP != test.want {
			t.Errorf("response incorrect, got: %s, want: %s", rb.IP, test.want)
		}
	}
}

func TestGetIPAcceptJSON(t *testing.T) {
	expected := "192.168.0.1"

	var tests = []struct {
		header   string
		value    string
		want     string
		wantCode int
	}{
		{header: "", value: "", want: expected, wantCode: 200},
		{header: "Accept", value: "application/json", want: expected, wantCode: 200},
		{header: "Accept", value: "*/*", want: expected, wantCode: 200},
	}

	for _, test := range tests {
		req, _ := http.NewRequest("GET", "/", nil)
		req.RemoteAddr = expected + ":57000"
		if len(test.header) > 0 {
			req.Header.Set(test.header, test.value)
		}

		res := httptest.NewRecorder()
		SetupServer().router.ServeHTTP(res, req)

		if res.Code != test.wantCode {
			t.Errorf("status code incorrect, got: %d, want: %d", res.Code, test.wantCode)
		}

		var rb responseBody
		b, err := ioutil.ReadAll(res.Body)
		if err = json.Unmarshal(b, &rb); err != nil {
			t.Errorf("error in test: %v", err)
		}

		if rb.IP != expected {
			t.Errorf("response incorrect, got: %s, want: %s", rb.IP, test.want)
		}
	}
}

func TestGetIPAcceptText(t *testing.T) {
	expected := "192.168.0.1"
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = expected + ":57000"
	req.Header.Set("Accept", "text/plain")
	res := httptest.NewRecorder()
	SetupServer().router.ServeHTTP(res, req)

	expectedCode := 200
	if res.Code != expectedCode {
		t.Errorf("status code incorrect, got: %d, want: %d", res.Code, expectedCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error in test: %v", err)
	}

	if string(b) != expected {
		t.Errorf("response incorrect, got: %s, want: %s", string(b), expected)
	}
}

func TestGetIPWrongAccept(t *testing.T) {
	expected := "192.168.0.1"
	req, _ := http.NewRequest("GET", "/", nil)
	req.RemoteAddr = expected + ":57000"
	req.Header.Set("Accept", "text/html")
	res := httptest.NewRecorder()
	SetupServer().router.ServeHTTP(res, req)

	expectedCode := 415
	if res.Code != expectedCode {
		t.Errorf("status code incorrect, got: %d, want: %d", res.Code, expectedCode)
	}
}
