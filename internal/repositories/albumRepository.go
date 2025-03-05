package repositories

import (
	"gorm.io/gorm"

	models "example/web-service-gin/internal/models"
)

type AlbumRepository struct {
	DB *gorm.DB
}

func NewAlbumRepository(db *gorm.DB) *AlbumRepository {
	return &AlbumRepository{DB: db}
}

func (r *AlbumRepository) GetAlbums() ([]models.Album, error) {
	var albums []models.Album
	err := r.DB.Find(&albums).Error
	return albums, err
}
