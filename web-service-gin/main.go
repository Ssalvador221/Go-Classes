package main

import (
	"encoding/json"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Pequeno Principe", Artist: "John Coltrane", Price: 70.00},
	{ID: "2", Title: "Little Dark Age", Artist: "MGMT", Price: 20.00},
	{ID: "3", Title: "Joãozinho", Artist: "Salvador", Price: 100.00},
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
	c.IndentedJSON(http.StatusFound, gin.H{"message": "New album created with succes"})
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumsById(c *gin.Context) {
	getId := c.Param("id")

	for _, a := range albums {
		if a.ID == getId {
			c.IndentedJSON(http.StatusOK, a)
			c.IndentedJSON(http.StatusFound, gin.H{"message": "Album succes found"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})

}

func saveApiData(c *gin.Context) {

	jsonFile, err := os.Create("db.json")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating file!"})
		return
	}
	defer jsonFile.Close()

	jsonData, err := json.MarshalIndent(albums, "", "  ")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Não foi possivel salvar"})
		return
	}

	_, err = jsonFile.Write(jsonData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Data error writer!"})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data saved to data.json"})

}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumsById)
	router.POST("/albums", postAlbums)
	router.GET("/save-data", saveApiData)

	router.Run("localhost:8080")
}
