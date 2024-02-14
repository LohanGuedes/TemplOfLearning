package application

import (
	"github.com/labstack/echo/v4"
	"github.com/lohanguedes/templOfLearning/handler"
	"github.com/lohanguedes/templOfLearning/internal/database"
)

type Application struct {
	DB    *database.Queries
	echo  *echo.Echo
	user  handler.UserHandler
	pages handler.PageHandler
}

func New(options ...func(*Application)) *Application {
	app := &Application{echo: echo.New()}

	for _, opt := range options {
		opt(app)
	}
	return app
}

func WithUserhandler(h handler.UserHandler) func(*Application) {
	return func(app *Application) {
		app.user = h
	}
}

func WithPageHandler(ph handler.PageHandler) func(*Application) {
	return func(app *Application) {
		app.pages = ph
	}
}

func (app *Application) WithGlobalMiddleware(middlewares ...echo.MiddlewareFunc) {
	for _, m := range middlewares {
		app.echo.Use(m)
	}
}

func (app *Application) Start(addr string) {
	app.echo.Start(addr)
}
