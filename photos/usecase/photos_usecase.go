package usecase

import (
	"context"
	"muhwyndham/gothos/hello/models"
	_photosRepo "muhwyndham/gothos/hello/photos/repository"
)

type Usecase interface {
	SavePhoto(ctx context.Context, fileName, contentType string, data []byte) (*models.Photo, error)
	FetchPhoto(ctx context.Context, id uint32) (*models.Photo, error)
}

type usecase struct {
	repo _photosRepo.Repository
}

// FetchPhoto implements Usecase
func (u usecase) FetchPhoto(ctx context.Context, id uint32) (*models.Photo, error) {
	return u.repo.FetchPhoto(ctx, id)
}

// SavePhoto implements Usecase
func (u usecase) SavePhoto(ctx context.Context, fileName, contentType string, data []byte) (*models.Photo, error) {
	photo := models.Photo{
		FileName:    fileName,
		ContentType: contentType,
		Data:        data,
	}

	if err := u.repo.SavePhoto(ctx, &photo); err != nil {
		return nil, err
	}
	return &photo, nil
}

func NewUsecase(repo _photosRepo.Repository) Usecase {
	return usecase{
		repo: repo,
	}
}
