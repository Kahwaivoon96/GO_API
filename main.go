package main

import (
	"github.com/gin-gonic/gin"

	"fmt"

	swaggerFiles "github.com/swaggo/files"

	ginSwagger "github.com/swaggo/gin-swagger"

	"example/web-service-gin/config"
	_ "example/web-service-gin/docs"

	database "example/web-service-gin/internal/database"
	"example/web-service-gin/internal/services"

	handlers "example/web-service-gin/internal/handlers"

	models "example/web-service-gin/internal/models"
	"log"
)

// albums slice to seed record album data.
var albums = []models.Album{
	{ID: 1, Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: 2, Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: 3, Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	// Initialize Database
	err = database.Connect(cfg)

	if err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}
	defer database.DB.Close() // automatically closes the database connection when the function exits.

	albumsHandler := handlers.NewAlbumHandler(services.AlbumService{})

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/dbalbums", albumsHandler.GetAlbums)
	// router.GET("/albums", getAlbums)
	// router.POST("/albums", postAlbums)
	// router.POST("/albumswitID", postAlbumsWitID)

	fmt.Println("Registered Routes:")
	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}
	router.Run("localhost:8080")
}

// // getAlbums 获取所有专辑
// // @Summary 获取专辑列表
// // @Description 返回所有专辑
// // @Tags albums
// // @Accept  json
// // @Produce  json
// // @Success 200 {array} album
// // @Router /albums [get]swag init
// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// // postAlbums adds an album from JSON received in the request body.
// // @Summary 添加专辑
// // @Description 添加一个新的专辑到列表
// // @Tags albums
// // @Accept  json
// // @Produce  json
// // @Param album body album true "专辑信息"
// // @Success 201 {object} album
// // @Router /albums [post]
// func postAlbums(c *gin.Context) {
// 	var newAlbum models.Album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }
// func postAlbumsWitID(c *gin.Context) {

// 	albumID := c.Query("albumID")

// 	// Convert albumID to uint
// 	num, err := strconv.ParseUint(albumID, 10, 32)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid album ID"})
// 		return
// 	}
// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.

// 	for _, a := range albums {
// 		if a.ID == uint(num) {
// 			c.JSON(http.StatusOK, a)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusBadGateway, nil)

// }

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
