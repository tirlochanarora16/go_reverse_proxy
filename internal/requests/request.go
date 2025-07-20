package requests

import (
	"net/http"

	"github.com/tirlochanarora16/go_reverse_proxy/internal/proxy"
)

func HandleMuxRoutes(mux *http.ServeMux) {
	// req from localhost:8080 will be transferred to localhost:3000
	reverseProxy := proxy.CreateReverseProxy("http://localhost:3000")

	mux.Handle("/", reverseProxy)
}
