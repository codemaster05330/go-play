package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/goadesign/goa"
)

// MetricRecorder record metrics
func MetricRecorder() goa.Middleware {
	return func(h goa.Handler) goa.Handler {
		return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
			log.Println("In MetricRecorder")

			action := goa.ContextAction(ctx)
			ctrl := goa.ContextController(ctx)

			log.Println("action=", action)
			log.Println("ctrl=", ctrl)

			err := h(ctx, rw, req)
			log.Println("Out MetricRecorder")
			return err
		}
	}
}
