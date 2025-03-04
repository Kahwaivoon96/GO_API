package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"fmt"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	_ "example/web-service-gin/docs"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.POST("/albumswitID", postAlbumsWitID)

	fmt.Println("Registered Routes:")
	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}
	router.Run("localhost:8080")
}

// getAlbums 获取所有专辑
// @Summary 获取专辑列表
// @Description 返回所有专辑
// @Tags albums
// @Accept  json
// @Produce  json
// @Success 200 {array} album
// @Router /albums [get]swag init
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
// @Summary 添加专辑
// @Description 添加一个新的专辑到列表
// @Tags albums
// @Accept  json
// @Produce  json
// @Param album body album true "专辑信息"
// @Success 201 {object} album
// @Router /albums [post]
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
func postAlbumsWitID(c *gin.Context) {

	albumID := c.Query("albumID")
	// Call BindJSON to bind the received JSON to
	// newAlbum.

	for _, a := range albums {
		if a.ID == albumID {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusBadGateway, nil)

}

// gin context = httpcontext in net core
// IndentedJSON = formatted json
// BindJSON  = Body For API
// If is Form data
// 		username := c.PostForm("username")  ** C = gin context
//		password := c.PostForm("password")
// If is query param
// 		query := c.Query("query")
// If is route param
// 		id := c.Param("id")
