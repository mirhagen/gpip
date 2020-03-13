package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/RedeployAB/gpip/config"
	"github.com/RedeployAB/gpip/server"
)

func main() {
	host := flag.String("host", "", "Host/IP Address to listen on")
	port := flag.String("port", "", "Port to listen on")
	flag.Parse()

	r := http.NewServeMux()
	conf := config.Configure(config.Options{Host: *host, Port: *port})
	srv := server.New(conf, r)
	srv.Start()
	log.Println("server has been stopped")
}
