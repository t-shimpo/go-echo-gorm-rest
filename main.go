package main

import (
	"github.com/labstack/echo/v4"
	"github.com/t-shimpo/go-echo-gorm-rest/controller"
	"github.com/t-shimpo/go-echo-gorm-rest/model"
)

func main() {
	e := echo.New()
	db, _ := model.DB.DB()
	defer db.Close()

	e.POST("/users", controller.CreateUser)
	e.Logger.Fatal(e.Start(":8080"))
}
