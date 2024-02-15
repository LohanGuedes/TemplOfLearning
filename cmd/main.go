package main

import (
	"context"
	"flag"
	"log"
	"runtime"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	application "github.com/lohanguedes/templOfLearning/cmd/app"
	"github.com/lohanguedes/templOfLearning/handler"
	"github.com/lohanguedes/templOfLearning/internal/database"
	userservice "github.com/lohanguedes/templOfLearning/internal/user-service"
)

func main() {
	var addr string
	port := flag.String("port", ":8080", "HTTP port default :8080")
	envFile := flag.String("env", ".env", "Enviroment file")
	flag.Parse()

	env, err := godotenv.Read(*envFile)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	conn, err := pgx.Connect(ctx, env["DB_CONN_STRING"])
	if err != nil {
	}
	defer conn.Close(ctx)

	db := database.New(conn)

	us := userservice.New(db)
	userHandler := handler.UserHandler{
		UserService: us,
	}
	pageHandler := handler.PageHandler{}

	app := application.New(
		application.WithUserhandler(userHandler),
		application.WithPageHandler(pageHandler),
	)

	app.WithGlobalMiddleware(
		middleware.Logger(), // TODO: Make this logger better for developmemnt.
		middleware.Recover(),
	)

	app.BindRoutes()

	if runtime.GOOS == "darwin" {
		addr = "localhost"
	}

	app.Start(addr + *port)
}
