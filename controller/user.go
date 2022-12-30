package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/t-shimpo/go-echo-gorm-rest/model"
)

func CreateUser(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}
	model.DB.Create(&user)
	return c.JSON(http.StatusCreated, user)
}
