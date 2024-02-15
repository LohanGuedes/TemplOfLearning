package handler

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	userservice "github.com/lohanguedes/templOfLearning/internal/user-service"
	"github.com/lohanguedes/templOfLearning/view/user"
)

type UserHandler struct {
	UserService userservice.UserService
}

func (h UserHandler) handleShowUser(c echo.Context) error {
	return nil
}

func (h UserHandler) handleCreateUser(c echo.Context) error {
	userData, err := h.UserService.Create(c.Request().Context(), c.Request().Body)
	if err != nil {
		return c.JSON(http.StatusNoContent, err)
	}
	return c.JSON(http.StatusCreated, userData)
}

func (h UserHandler) handleGetUserByID(c echo.Context) error {
	userData, err := h.UserService.GetUserByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Errorf("user with id: %s not found: %w", c.Param("id"), err))
	}

	return Render(c, http.StatusOK, user.ShowUser(userData))
}

func (h UserHandler) Grouper(g *echo.Group) {
	g.GET("/", h.handleShowUser)
	g.POST("/", h.handleCreateUser)
	g.GET("/:id", h.handleGetUserByID)
}
