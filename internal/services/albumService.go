package services

import (
	models "example/web-service-gin/internal/models"
	repositories "example/web-service-gin/internal/repositories"
	"log"

	database "example/web-service-gin/internal/database"
)

type AlbumService struct {
	Repo *repositories.AlbumRepository
}

func NewAlbumService(repo *repositories.AlbumRepository) *AlbumService {
	return &AlbumService{Repo: repo}
}

// GetAlbums retrieves all albums from the database
func GetAlbums() ([]models.Album, error) {
	rows, err := database.DB.Query("SELECT id, title, artist, price FROM [Album]")
	if err != nil {
		log.Println("Error fetching albums:", err)
		return nil, err
	}
	defer rows.Close()

	var albums []models.Album
	for rows.Next() {
		var album models.Album
		if err := rows.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}

	return albums, nil
}
