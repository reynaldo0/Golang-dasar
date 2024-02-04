package main

import (
	"daytwo/mappings"
	"daytwo/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = []models.Albumsiswa{
	{ID: "1", Title: "EL PARADE", Artist: "Michael", Price: 23.00},
	{ID: "2", Title: "JO SIKUS", Artist: "Joseph", Price: 57.00},
	{ID: "3", Title: "EREREEH", Artist: "Febrian", Price: 45.00},
}

func main() {
	fmt.Println("============", models.Nama)

	// router := gin.Default()
	// router.GET("/albums", getAlbums)

	// router.Run("localhost:8080")

	mappings.CreateUrlMappings()
	mappings.Router.Run(":8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
