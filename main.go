package main

import (
	"GoGin_Film_Arsiv/controllers"
	"GoGin_Film_Arsiv/models"

	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDatabase()

	r := gin.Default()

	r.Static("/static", "./public")

	r.LoadHTMLFiles(
		"views/index.html",
		"views/create.html",
		"views/update.html",
		"views/layouts/header.html",
		"views/layouts/footer.html",
	)

	r.GET("/", controllers.RenderIndex)
	r.GET("/create-page", controllers.RenderCreatePage)
	r.GET("/update-page", controllers.RenderUpdatePage)

	r.GET("/movies", controllers.GetFilms)
	r.GET("/movies/:id", controllers.GetFilm)
	r.POST("/movies", controllers.CreateFilm)
	r.PUT("/movies/:id", controllers.UpdateFilm)
	r.DELETE("/movies/:id", controllers.DeleteFilm)

	r.Run(":8085")
}
