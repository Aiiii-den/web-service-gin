package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
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
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", postAlbums)
	router.PUT("/albums/:id", updateAlbumByID)
	router.DELETE("albums/:id", deleteAlbumByID)
	router.DELETE("/albums", deleteAllAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		fmt.Print("Couldn't run service")
		return
	}
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func updateAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Create a new album object
	var updatedAlbum album

	// Bind the JSON request body to the updatedAlbum variable
	if err := c.ShouldBindJSON(&updatedAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Check if id exists, if yes, update album
	found := false
	for i, a := range albums {
		if a.ID == id {
			// Update the existing album with the new data
			albums[i] = updatedAlbum
			found = true
			break
		}
	}

	// If the album with the given ID was not found, return a 404 response
	if !found {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Album updated successfully"})
}

// postAlbums adds an album from JSON received in the request body.
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

func deleteAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Find the index of the album with the given ID
	index := -1
	for i, a := range albums {
		if a.ID == id {
			index = i
			break
		}
	}

	// If the album with the given ID was not found, return a 404 response
	if index == -1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Album not found"})
		return
	}

	// Remove the album at the specified index
	albums = append(albums[:index], albums[index+1:]...)

	// Return a success response
	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}

func deleteAllAlbums(c *gin.Context) {
	albums = albums[:0]
	c.JSON(http.StatusOK, gin.H{"message": "All albums removed successfully"})
}
