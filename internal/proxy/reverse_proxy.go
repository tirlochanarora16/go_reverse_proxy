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

		middleware.Logger.Info("Request -> ", zap.String("method", r.Method), zap.String("URL", r.URL.String()), zap.String("host", r.Host))
	}

	return proxy
}
