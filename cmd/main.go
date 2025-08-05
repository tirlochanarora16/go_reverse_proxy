package main

import (
	"log"
	"net/http"

	"github.com/tirlochanarora16/go_reverse_proxy/internal/config"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/lb"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/middleware"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/watcher"
)

func main() {
	lb.CheckConfigFlag()
	lb.ReadConfigFile()
	config.ParseConfigFile()

	middleware.InitLogger()
	go middleware.InitRateLimiter()
	go watcher.StartConfigFileWatcher()
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lb.GetActiveMutex().ServeHTTP(w, r)
	})))
}
