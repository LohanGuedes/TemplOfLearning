package application

import "github.com/lohanguedes/templOfLearning/handler"

func (app *Application) BindRoutes() {
	v1 := app.echo.Group("/v1") // /v1
	v1.GET("/health", handler.HandleHealthCheck)

	app.user.Grouper(v1.Group("/users")) // /v1/users

	app.echo.GET("/", app.pages.HandleShowHome)
}
