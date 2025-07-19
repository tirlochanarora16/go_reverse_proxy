package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func CreateReverseProxy(target string) *httputil.ReverseProxy {
	url, err := url.Parse(target)

	if err != nil {
		log.Fatalf("Unable to parse the URL %s", target)
	}

	proxy := httputil.NewSingleHostReverseProxy(url)
	originalDirector := proxy.Director

	proxy.Director = func(r *http.Request) {
		originalDirector(r)

		log.Printf("Forwarding %s reqeuest to %s", r.Method, r.URL.String())
	}

	return proxy
}
