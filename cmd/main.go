package main

import (
	"log"
	"net/http"

	"github.com/tirlochanarora16/go_reverse_proxy/internal/config"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/lb"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/middleware"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/requests"
)

func main() {
	lb.CheckConfigFlag()
	lb.ReadConfigFile()
	config.ParseConfigFile()

	go middleware.InitRateLimiter()
	middleware.InitLogger()
	mux := http.NewServeMux()
	requests.HandleMuxRoutes(mux)
	log.Fatal(http.ListenAndServe(":8080", mux))
}
