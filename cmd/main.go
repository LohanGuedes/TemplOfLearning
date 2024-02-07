package main

import (
	"flag"

	"github.com/labstack/echo/v4/middleware"
	application "github.com/lohanguedes/templOfLearning/cmd/app"
	"github.com/lohanguedes/templOfLearning/handler"
)

func main() {
	port := flag.String("port", "3000", "HTTP port")
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

	app.Start("localhost:" + *port)
}
