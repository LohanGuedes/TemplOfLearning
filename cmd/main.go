package main

import (
	"flag"
	"runtime"

	"github.com/labstack/echo/v4/middleware"
	application "github.com/lohanguedes/templOfLearning/cmd/app"
	"github.com/lohanguedes/templOfLearning/handler"
)

func main() {
	port := flag.String("port", "8080", "HTTP port")
	flag.Parse()

	userHandler := handler.UserHandler{}
	app := application.New(
		application.WithUserhandler(userHandler),
	)

	app.WithGlobalMiddleware(
		middleware.Logger(),
		middleware.Recover(),
	)

	app.BindRoutes()

	if runtime.GOOS == "darwin" {
		app.Start("localhost:" + *port)
		return
	}

	app.Start(":" + *port)
}
