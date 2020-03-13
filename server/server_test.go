package server

import (
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/RedeployAB/gpip/config"
)

func TestNew(t *testing.T) {
	conf := config.Configuration{Host: "0.0.0.0", Port: "5050"}
	r := http.NewServeMux()
	srv := New(conf, r)

	expected := "0.0.0.0:5050"
	if srv.httpServer.Addr != expected {
		t.Errorf("incorrect value, got: %s, want: %s", srv.httpServer.Addr, expected)
	}
}

func TestStartAndShutdown(t *testing.T) {
	conf := config.Configuration{Host: "0.0.0.0", Port: "5050"}
	r := http.NewServeMux()
	srv := New(conf, r)

	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Fatal("error in getting running process")
	}

	var result *os.ProcessState

	go func() {
		srv.Start()
		result, _ = proc.Wait()
	}()

	time.Sleep(3 * time.Second)
	proc.Signal(os.Interrupt)

	exitCode := result.ExitCode()
	expected := -1
	if exitCode != expected {
		t.Errorf("incorrect value, got: %d, want: %d", exitCode, expected)
	}
}
