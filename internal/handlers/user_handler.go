package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

var nextId = 1
var users = make(map[int]*User)

func HealthCheck(c echo.Context) error {
	return c.String(http.StatusOK, "Hello")
}

func GetAllUsers(c echo.Context) error {
	userList := make([]*User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}
	return c.JSON(http.StatusOK, userList)
}

func GetUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid User ID"})
	}
	user, found := users[id]
	if !found {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User Not Found"})
	}
	return c.JSON(http.StatusOK, user)
}

func AddUser(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request Payload"})
	}
	user.Id = nextId
	users[nextId] = user
	nextId++
	return c.JSON(http.StatusCreated, user)
}

func DeleteUserById(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid User ID"})
	}
	delete(users, id)
	return c.NoContent(http.StatusNoContent)
}

func UpdateUserById(c echo.Context) error {
	user := &User{}
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid Request Payload"})
	}
	//user contains id and new name
	existingUser, found := users[user.Id]
	if !found {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User Not Found"})
	}
	existingUser.Name = user.Name
	return c.JSON(http.StatusOK, existingUser)
}
