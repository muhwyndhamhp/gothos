package gothos

import (
	"io"
	"net/http"
	"strconv"

	_photosUsecase "muhwyndham/gothos/hello/photos/usecase"

	"github.com/labstack/echo/v4"
)

type PhotosHandler struct {
	usecase _photosUsecase.Usecase
}

func NewPhotosHandler(g *echo.Group, usecase _photosUsecase.Usecase) {
	service := &PhotosHandler{
		usecase: usecase,
	}

	g.GET("/gothos", service.HelloService)
	g.POST("/photos", service.UploadPhoto)
	g.GET("/download/:id", service.Download)
}

func (s *PhotosHandler) Download(c echo.Context) error {
	photoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	ctx := c.Request().Context()

	photo, err := s.usecase.FetchPhoto(ctx, uint32(photoID))
	if err != nil {
		return err
	}

	if photo == nil {
		return c.JSON(http.StatusBadRequest, "photo not found")
	}

	return c.Blob(http.StatusOK, photo.ContentType, photo.Data)
}

func (s *PhotosHandler) UploadPhoto(c echo.Context) error {
	ctx := c.Request().Context()

	file, err := c.FormFile("data")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	contentType := file.Header.Values("Content-Type")

	byteContainer, err := io.ReadAll(src)
	if err != nil {
		return err
	}

	photo, err := s.usecase.SavePhoto(ctx, file.Filename, contentType[0], byteContainer)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, photo)
}
func (s *PhotosHandler) HelloService(c echo.Context) error {
	return c.JSON(http.StatusOK, struct {
		Data string `json:"data"`
	}{
		Data: "Hello Gothos Service!",
	})
}
