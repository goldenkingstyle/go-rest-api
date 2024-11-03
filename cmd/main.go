package main

import (
	"flag"
	"fmt"
	"go-rest-api/internal/config"
	"go-rest-api/internal/user"
	"log"
	"net/http"
)

var flagConfig = flag.String("config", "./", "Path to config file")

func main() {

	flag.Parse()

	cfg, err := config.Load(*flagConfig)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()

	registerHandlers(mux)

	address := fmt.Sprintf(":%v", cfg.ServerPort)

	_, err = fmt.Printf("Server is running on %s port\n", address)
	if err != nil {
		return
	}
	if err := http.ListenAndServe(address, mux); err != nil {
		log.Fatal(err)
	}

}

func registerHandlers(m *http.ServeMux) {
	user.RegisterHandlers(m)
}
