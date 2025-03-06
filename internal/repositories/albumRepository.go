package repositories

import (
	"database/sql"
	models "example/web-service-gin/internal/models"
	"fmt"
	"log"
)

type AlbumRepository struct {
	DB *sql.DB
}

func NewAlbumRepository(db *sql.DB) *AlbumRepository {
	return &AlbumRepository{DB: db}
}

// func (r *AlbumRepository) GetAlbums() ([]models.Album, error) {
// 	var albums []models.Album
// 	err := r.DB.Find(&albums).Error
// 	return albums, err
// }

func (repo *AlbumRepository) GetAllAlbums() ([]models.Album, error) {
	if repo.DB == nil {
		log.Fatal("Database connection is nil!")
	}
	rows, err := repo.DB.Query("SELECT id, title, artist, price FROM Album")

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

func (repo *AlbumRepository) InsertAlbum(album models.InsertAlbum) (int32, error) {
	if repo.DB == nil {
		log.Fatal("Database connection is nil!")
	}
	var insertedID int32
	tx, err := repo.DB.Begin()

	func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	query := "EXEC sp_InsertAlbum @Title=@p1, @Artist=@p2, @Price=@p3, @InsertedID=@p4 OUTPUT"

	// Execute the stored procedure
	_, err = tx.Exec(
		query,
		sql.Named("p1", album.Title),
		sql.Named("p2", album.Artist),
		sql.Named("p3", album.Price),
		sql.Named("p4", sql.Out{Dest: &insertedID}),
	)
	if err != nil {
		return 0, fmt.Errorf("failed to execute stored procedure: %w", err)
	}

	return insertedID, nil

}
