package lb

import (
	"net/http"
	"sync/atomic"
)

var ActiveMutex atomic.Value

func GetActiveMutex() http.Handler {
	return ActiveMutex.Load().(http.Handler)
}
