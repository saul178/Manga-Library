package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Black Water Park", Artist: "Opeth", Price: 56.99},
	{ID: "2", Title: "Inmazes(Deluxe Edition)", Artist: "VOLA", Price: 19.99},
	{ID: "3", Title: "Flower Boy", Artist: "Tyler the Creator", Price: 30.95},
}

func getAlbums(ptrToThis *gin.Context) {
	ptrToThis.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ptrToThis *gin.Context) {
	var newAlbum album
	err := ptrToThis.BindJSON(&newAlbum)

	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ptrToThis.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(ptrToThis *gin.Context) {
	id := ptrToThis.Param("id")

	for _, a := range albums {
		if a.ID == id {
			ptrToThis.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	ptrToThis.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func initialize() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}

func main() {
	initialize()
}
