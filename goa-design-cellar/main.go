//go:generate goagen bootstrap -d github.com/nurali-techie/play-go-web/goa-design-cellar/design

package main

import (
	"net/http"

	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/nurali-techie/play-go-web/goa-design-cellar/app"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// Create service
	service := goa.New("cellar")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "bottle" controller
	c := NewBottleController(service)
	app.MountBottleController(service, c)

	http.Handle("/api/", service.Mux)

	// Enable prometheus metrics
	// http.Handle("/metrics", prometheus.Handler())
	http.Handle("/metrics", promhttp.Handler())
	// service.Use(mw.MetricRecorder())

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
