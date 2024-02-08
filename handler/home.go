package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lohanguedes/templOfLearning/view/pages"
)

type PageHandler struct{}

func (ph PageHandler) HandleShowHome(c echo.Context) error {
	return Render(c, http.StatusOK, pages.Home())
}
