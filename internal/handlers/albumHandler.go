package handlers

import (
	models "example/web-service-gin/internal/models"
	albumService "example/web-service-gin/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	service *albumService.AlbumService
}

func NewAlbumHandler(s *albumService.AlbumService) *AlbumHandler {
	return &AlbumHandler{service: s}
}

// GetAlbums returns all albums as JSON.
// @Summary Get Albums
// @Description Retrieves a list of all albums
// @Tags albums
// @Produce json
// @Success 200
// @Router /dbalbums [get]
func (h *AlbumHandler) GetAlbums(c *gin.Context) {

	albums, err := h.service.GetAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, albums)
}

// InsertAlbum insert an album to db.
// @Summary Insert Albums
// @Description Retrieves a list of all albums
// @Tags albums
// @Produce json
// @Param album body models.InsertAlbum true "Album details"
// @Success 200
// @Router /insertalbum [post]
func (h *AlbumHandler) InsertAlbum(c *gin.Context) {
	var albumInsert models.InsertAlbum

	if err := c.ShouldBindJSON(&albumInsert); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input format"})
		return
	}

	id, err := h.service.InsertAlbum(albumInsert)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, id)
}
