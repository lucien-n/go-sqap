package main

import (
	"fmt"
	"sqap/internal/config"
	"sqap/internal/database"
	"sqap/internal/repositories"
	"sqap/internal/router"
	"sqap/internal/routes"
	"sqap/internal/services"

	_ "github.com/libsql/libsql-client-go/libsql"
)

func main() {
	cfg := config.LoadConfig("D:/go/go-sqap/.env")
	db := database.ConnectDb(cfg)
	app := router.CreateRouter()

	todoRepository := repositories.NewTodoRepository(db)
	todoService := services.NewTodoService(todoRepository)

	routes.RegisterTodo(app, todoService)

	app.Listen(fmt.Sprintf("%s:%s", cfg.APIHost, cfg.APIPort))
}
