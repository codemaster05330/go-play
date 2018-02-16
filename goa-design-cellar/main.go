//go:generate goagen bootstrap -d github.com/nurali-techie/play-go-web/goa-design-cellar/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/nurali-techie/play-go-web/goa-design-cellar/app"
	mw "github.com/nurali-techie/play-go-web/goa-design-cellar/middleware"
)

func main() {
	// Create service
	service := goa.New("cellar")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	service.Use(mw.MetricRecorder())

	// Mount "bottle" controller
	c := NewBottleController(service)
	app.MountBottleController(service, c)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
