package repository

import (
	"context"
	"muhwyndham/gothos/hello/models"

	"gorm.io/gorm"
)

type Repository interface {
	SavePhoto(ctx context.Context, value *models.Photo) error
	FetchPhoto(ctx context.Context, id uint32) (*models.Photo, error)
}

type repository struct {
	db *gorm.DB
}

// FetchPhoto implements Repository
func (m repository) FetchPhoto(ctx context.Context, id uint32) (*models.Photo, error) {
	photo := models.Photo{}

	err := m.db.WithContext(ctx).Where("id", id).First(&photo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &photo, nil
}

// SavePhoto implements Repository
func (m repository) SavePhoto(ctx context.Context, value *models.Photo) error {
	err := m.db.WithContext(ctx).Save(value).Error
	if err != nil {
		return err
	}

	return nil
}

func NewRepository(db *gorm.DB) Repository {
	return repository{
		db: db,
	}
}
