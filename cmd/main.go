package main

import (
	"flag"
	"fmt"
	"runtime"

	"github.com/labstack/echo/v4/middleware"
	application "github.com/lohanguedes/templOfLearning/cmd/app"
	"github.com/lohanguedes/templOfLearning/handler"
)

func main() {
	var addr string
	port := flag.String("port", ":8080", "HTTP port default :8080")
	flag.Parse()

	userHandler := handler.UserHandler{}
	pageHandler := handler.PageHandler{}
	app := application.New(
		application.WithUserhandler(userHandler),
		application.WithPageHandler(pageHandler),
	)

	app.WithGlobalMiddleware(
		middleware.Logger(),
		middleware.Recover(),
	)

	app.BindRoutes()

	if runtime.GOOS == "darwin" {
		addr = "localhost"
	}

	fmt.Printf("%s\n", addr+*port)
	app.Start(addr + *port)
}
