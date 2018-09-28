package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/goadesign/goa"
	"github.com/prometheus/client_golang/prometheus"
)

var reqLabels = []string{"entity", "action"}

// MetricRecorder record metrics
func MetricRecorder() goa.Middleware {
	reqCnt := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "goa_req_total",
			Help: "Total number of goa requests.",
		}, reqLabels)
	prometheus.Register(reqCnt)

	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			log.Println("In MetricRecorder")

			err := h(ctx, rw, req)

			action := goa.ContextAction(ctx)
			ctrl := goa.ContextController(ctx)
			log.Println("action=", action)
			log.Println("ctrl=", ctrl)
			reqCnt.WithLabelValues(ctrl, action).Inc()

			log.Println("Out MetricRecorder")
			return err
		}
	}
}
