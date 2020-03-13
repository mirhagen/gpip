package config

import (
	"os"
)

// Configuration defines the configuration for the server.
type Configuration struct {
	Host string
	Port string
}

// Options options defines the incoming
// options for Configure().
type Options struct {
	Host string
	Port string
}

// Configure takes incoming parameters and determines
// if parameter should be used or environment variables.
func Configure(opts Options) Configuration {
	var h string
	var p string

	if len(opts.Host) > 0 {
		h = opts.Host
	} else if len(os.Getenv("GPIP_LISTEN_HOST")) > 0 {
		h = os.Getenv("GPIP_LISTEN_HOST")
	} else {
		h = "0.0.0.0"
	}

	if len(opts.Port) > 0 {
		p = opts.Port
	} else if len(os.Getenv("GPIP_LISTEN_PORT")) > 0 {
		p = os.Getenv("GPIP_LISTEN_PORT")
	} else {
		p = "5050"
	}

	return Configuration{Host: h, Port: p}
}
