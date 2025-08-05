package requests

import (
	"net/http"

	"github.com/tirlochanarora16/go_reverse_proxy/internal/config"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/lb"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/middleware"
	"github.com/tirlochanarora16/go_reverse_proxy/internal/proxy"
	"go.uber.org/zap"
)

type responseRecorder struct {
	http.ResponseWriter
	status int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.status = code
	rr.ResponseWriter.WriteHeader(code)
}

func HandleMuxRoutes() {
	mux := http.NewServeMux()

	// req from localhost:8080 will be transferred to localhost:3000
	for _, route := range config.ConfigFileData.Routes {
		route := route

		reverseProxy := proxy.CreateReverseProxy(route.Target)

		handler := func(w http.ResponseWriter, r *http.Request) {
			rr := &responseRecorder{ResponseWriter: w, status: 200}
			reverseProxy.ServeHTTP(rr, r)

			method := zap.String("method", r.Method)
			url := zap.String("url", r.URL.String())
			status := zap.Int("status", rr.status)

			isError := rr.status >= 400 && rr.status < 600

			if isError {
				middleware.Logger.Error("Response <-", method, url, status)
				return
			}

			middleware.Logger.Info("Response <- ", method, url, status)
		}

		var finalHandler http.Handler = http.HandlerFunc(handler)

		if route.RateLimit != nil {
			rps := route.RateLimit.Rate
			burst := route.RateLimit.Burst
			finalHandler = middleware.RateLimitMiddleware(finalHandler, rps, burst)
		}

		mux.Handle(route.Path, finalHandler)
	}

	lb.ActiveMutex.Store(mux)
}
