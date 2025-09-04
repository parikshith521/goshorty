package server

import (
	"github.com/labstack/echo/v4"

	"github.com/parikshith521/goshorty/internal/handlers"
)

func Run(port string) {
	e := echo.New()
	// set up server routes here
	e.GET("/", handlers.HealthCheck)
	e.GET("/users", handlers.GetAllUsers)
	e.GET("users/:id", handlers.GetUserById)
	e.POST("users", handlers.AddUser)
	e.PUT("users/:id", handlers.UpdateUserById)
	e.DELETE("users/:id", handlers.DeleteUserById)

	e.Logger.Fatal(e.Start(port))
}
