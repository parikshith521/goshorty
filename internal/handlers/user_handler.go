package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func GetAllUsers(c echo.Context) error {
	return c.String(http.StatusOK, "getAllUsers")
}

func GetUserById(c echo.Context) error {
	return c.String(http.StatusOK, "getUserById: "+c.Param("id"))
}

func AddUser(c echo.Context) error {
	return c.String(http.StatusCreated, "addUser")
}

func DeleteUserById(c echo.Context) error {
	return c.String(http.StatusOK, "deleteUserById")
}

func UpdateUserById(c echo.Context) error {
	return c.String(http.StatusOK, "updateUserById")
}
