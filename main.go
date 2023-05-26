package main

import (
	"fmt"
	_photosHandler "muhwyndham/gothos/hello/photos/handler/http"
	_photosRepo "muhwyndham/gothos/hello/photos/repository"
	_photosUsecase "muhwyndham/gothos/hello/photos/usecase"

	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	e := echo.New()

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
		"localhost",
		"5432",
		"postgres",
		"gothos_db",
		"postgres",
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	repository := _photosRepo.NewRepository(db)
	usecase := _photosUsecase.NewUsecase(repository)

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{
			"*://localhost:*",
		},
	}))

	e.GET("/", hello)

	api := e.Group("/api")
	v1 := api.Group("/v1")

	v1.File("/upload", "gothos/static.html")
	_photosHandler.NewPhotosHandler(v1, usecase)

	e.Logger.Fatal(e.Start(":1323"))
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
