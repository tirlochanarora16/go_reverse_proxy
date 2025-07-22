package proxy

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/tirlochanarora16/go_reverse_proxy/internal/middleware"
	"go.uber.org/zap"
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

		method := zap.String("method", r.Method)
		url := zap.String("url", r.URL.String())
		path := zap.String("path", r.URL.Path)
		host := zap.String("host", r.Host)

		middleware.Logger.Info("Request -> ", method, path, url, host)
	}

	return proxy
}
