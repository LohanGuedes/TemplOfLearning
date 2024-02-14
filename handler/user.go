package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/lohanguedes/templOfLearning/internal/database"
	"github.com/lohanguedes/templOfLearning/view/user"
)

type UserHandler struct {
	DB database.Queries
}

func (h UserHandler) handleShowUser(c echo.Context) error {
	return nil
}

func (h UserHandler) handleCreateUser(c echo.Context) error {
	var data database.CreateUserParams

	err := json.NewDecoder(c.Request().Body).Decode(&data)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}
	fmt.Printf("Data parsed with success")

	userData, err := h.DB.CreateUser(c.Request().Context(), database.CreateUserParams{
		Username: data.Username,
		Email:    data.Email,
		Password: data.Password,
		Role:     data.Role,
	})
	fmt.Printf("User Created with success\n")
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, nil)
	}
	return c.JSON(http.StatusCreated, userData)
}

func (h UserHandler) handleGetUserByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Errorf("invalid id: %w. Id must be of an UUID", err))
	}

	userData, err := h.DB.GetUserByID(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return Render(c, http.StatusOK, user.ShowUser(userData))
}

func (h UserHandler) Grouper(g *echo.Group) {
	g.GET("/", h.handleShowUser)
	g.POST("/", h.handleCreateUser)
	g.GET("/:id", h.handleGetUserByID)
}
