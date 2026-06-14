package controllers

import (
	"GoGin_Film_Arsiv/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetFilms(c *gin.Context) {
	var films []models.Film
	models.DB.Find(&films)
	c.JSON(http.StatusOK, films)
}

func CreateFilm(c *gin.Context) {
	var input models.Film

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, input)
}

func GetFilm(c *gin.Context) {
	var film models.Film
	if err := models.DB.Where("id = ?", c.Param("id")).First(&film).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Film bulunamadi!"})
		return
	}
	c.JSON(http.StatusOK, film)
}

func UpdateFilm(c *gin.Context) {
	var film models.Film
	if err := models.DB.Where("id = ?", c.Param("id")).First(&film).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Film bulunamadi!"})
		return
	}

	var input models.Film
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&film).Updates(input)
	c.JSON(http.StatusOK, film)
}

func DeleteFilm(c *gin.Context) {
	var film models.Film
	if err := models.DB.Where("id = ?", c.Param("id")).First(&film).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Film bulunamadi!"})
		return
	}

	models.DB.Delete(&film)
	c.JSON(http.StatusOK, gin.H{"data": true})
}

func RenderIndex(c *gin.Context) {
	var films []models.Film
	searchQuery := c.Query("q")

	if searchQuery != "" {
		likeQuery := "%" + searchQuery + "%"
		models.DB.Where("LOWER(title) LIKE LOWER(?) OR LOWER(genre) LIKE LOWER(?) OR LOWER(director) LIKE LOWER(?)", likeQuery, likeQuery, likeQuery).Find(&films)
	} else {
		models.DB.Find(&films)
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"films": films,
		"query": searchQuery,
	})
}

func RenderCreatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

func RenderUpdatePage(c *gin.Context) {
	c.HTML(http.StatusOK, "update.html", nil)
}
