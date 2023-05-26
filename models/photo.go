package models

import "gorm.io/gorm"

type Photo struct {
	gorm.Model
	FileName    string `json:"file_name"`
	ContentType string `json:"content_type"`
	Data        []byte `json:"data"`
}

func (Photo) TableName() string {
	return "photos"
}
