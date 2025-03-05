package handlers

import (
	albumService "example/web-service-gin/internal/services"

	"net/http"

	"github.com/gin-gonic/gin"
)

type AlbumHandler struct {
	service albumService.AlbumService
}

func NewAlbumHandler(s albumService.AlbumService) *AlbumHandler {
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
	albums, err := albumService.GetAlbums()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	c.JSON(http.StatusOK, albums)
}

// func (h *AlbumHandler) GetAlbums(c *gin.Context) {
// 	albums := h.service.GetAll()
// 	c.JSON(http.StatusOK, albums)
// }

// func (h *AlbumHandler) CreateAlbum(c *gin.Context) {
// 	var newAlbum services.Album
// 	if err := c.ShouldBindJSON(&newAlbum); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	created := h.service.Create(newAlbum)
// 	c.JSON(http.StatusCreated, created)
// }

// func (h *AlbumHandler) GetAlbumByID(c *gin.Context) {
// 	id := c.Param("id")
// 	album, err := h.service.GetByID(id)
// 	if err != nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, album)
// }
