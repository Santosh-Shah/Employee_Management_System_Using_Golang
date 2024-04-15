package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []Album{
	{ID: "1", Title: "Nepali Album1", Artist: "Kamal khatri", Price: 56000.99},
	{ID: "2", Title: "Nepali Album2", Artist: "Kamal Chhetri", Price: 17000.99},
	{ID: "3", Title: "Nepali Album3", Artist: "Unknown Singer", Price: 390000.99},
}

func main() {
	router := gin.Default()

	// Define routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", addAlbum)

	// Start server
	router.Run("localhost:8080")
}

// Handler to get all albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Handler to get a specific album by ID
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}

	// If album with specified ID is not found
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// Handler to add a new album
func addAlbum(c *gin.Context) {
	var newAlbum Album

	// Bind the JSON request body to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	// Add the new album to the albums slice
	albums = append(albums, newAlbum)

	// Return the newly added album
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
