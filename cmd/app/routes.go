package application

func (app *Application) BindRoutes() {
	app.user.Grouper(app.echo.Group("/users"))
	app.echo.GET("/", app.pages.HandleShowHome)
}
