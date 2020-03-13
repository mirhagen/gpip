package config

import (
	"os"
	"testing"
)

func TestConfigure(t *testing.T) {
	conf := Configure(Options{Host: "", Port: ""})
	expectedHost := "0.0.0.0"
	expectedPort := "5050"
	if conf.Host != expectedHost {
		t.Errorf("incorrect value, got: %s, want: %s", conf.Host, expectedHost)
	}
	if conf.Port != expectedPort {
		t.Errorf("incorrect value, got: %s, want: %s", conf.Port, expectedPort)
	}

	conf = Configure(Options{Host: "127.0.0.1", Port: "5060"})
	expectedHost = "127.0.0.1"
	expectedPort = "5060"
	if conf.Host != expectedHost {
		t.Errorf("incorrect value, got: %s, want: %s", conf.Host, expectedHost)
	}
	if conf.Port != expectedPort {
		t.Errorf("incorrect value, got: %s, want: %s", conf.Port, expectedPort)
	}

	os.Setenv("GPIP_LISTEN_HOST", "127.0.0.2")
	os.Setenv("GPIP_LISTEN_PORT", "5070")
	conf = Configure(Options{Host: "", Port: ""})
	expectedHost = "127.0.0.2"
	expectedPort = "5070"
	if conf.Host != expectedHost {
		t.Errorf("incorrect value, got: %s, want: %s", conf.Host, expectedHost)
	}
	if conf.Port != expectedPort {
		t.Errorf("incorrect value, got: %s, want: %s", conf.Port, expectedPort)
	}
	os.Setenv("GPIP_LISTEN_HOST", "")
	os.Setenv("GPIP_LISTEN_PORT", "")
}
