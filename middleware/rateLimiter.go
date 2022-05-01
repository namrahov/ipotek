package middleware

import (
	"github.com/PB-Digital/ms-retail-products-info/properties"
	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type request struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

var requests = make(map[string]*request)
var mu sync.Mutex

func getRequest(ip string) *rate.Limiter {
	log.Info("ActionLog.getRequest.start for ip ", ip)
	mu.Lock()
	defer mu.Unlock()

	v, exists := requests[ip]
	if !exists {
		rt := rate.Every(5 * time.Minute / 1)
		limiter := rate.NewLimiter(rt, 5)
		requests[ip] = &request{limiter, time.Now()}
		return limiter
	}

	v.lastSeen = time.Now()
	log.Info("ActionLog.getRequest.end for ip ", ip)
	return v.limiter
}

func Throttle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.RequestURI, properties.RootPath) {
			return
		}
		log.Info("ActionLog.Throttle.start")
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		limiter := getRequest(ip)
		if limiter.Allow() == false {
			http.Error(w, http.StatusText(http.StatusTooManyRequests), http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
		log.Info("ActionLog.Throttle.end")
	})
}
