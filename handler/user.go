package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/lohanguedes/templOfLearning/types"
	"github.com/lohanguedes/templOfLearning/view/user"
)

type UserHandler struct{}

func (h UserHandler) handleShowUser(c echo.Context) error {
	mockUser := types.User{Name: "Lohan", Age: 21, Email: "mail@lguedes.dev"}
	return Render(c, http.StatusOK, user.ShowUser(mockUser))
}

func (h UserHandler) handleGetUserByID(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity, fmt.Errorf("invalid id: %w", err))
	}
	mockUser := types.User{ID: id, Name: "Lohan", Age: 21, Email: "mail@lguedes.dev"}
	return Render(c, http.StatusOK, user.ShowUser(mockUser))
}

func (h UserHandler) Grouper(g *echo.Group) {
	g.GET("/", h.handleShowUser)
	g.GET("/:id", h.handleGetUserByID)
}
