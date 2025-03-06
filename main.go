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
	repositories "example/web-service-gin/internal/repositories"

	"log"
)

func main() {
	// Load configuration
	cfg, err := loadConfig()

	// Initialize Database
	db, err := database.Connect(cfg)

	if err != nil {
		log.Fatalf("❌ Error connecting to database: %v", err)
	}
	defer database.DB.Close() // automatically closes the database connection when the function exits.

	albumRepo := repositories.NewAlbumRepository(db)
	albumService := services.NewAlbumService(albumRepo)
	albumsHandler := handlers.NewAlbumHandler(albumService)

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.GET("/dbalbums", albumsHandler.GetAlbums)
	router.POST("/insertalbum", albumsHandler.InsertAlbum)

	fmt.Println("Registered Routes:")
	for _, route := range router.Routes() {
		fmt.Println(route.Method, route.Path)
	}
	router.Run("localhost:8080")
}

func loadConfig() (*config.Config, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("❌ Failed to load config: %v", err)
	}
	return cfg, err
}
