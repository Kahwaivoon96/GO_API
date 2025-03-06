package services

import (
	models "example/web-service-gin/internal/models"
	repositories "example/web-service-gin/internal/repositories"
)

type AlbumService struct {
	Repo *repositories.AlbumRepository
}

func NewAlbumService(repo *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{Repo: repo}
}

// GetAlbums calls the repository to retrieve album data
func (s *AlbumService) GetAlbums() ([]models.Album, error) {
	return s.Repo.GetAllAlbums()
}
func (s *AlbumService) InsertAlbum(album models.InsertAlbum) (int32, error) {
	return s.Repo.InsertAlbum(album)
}
