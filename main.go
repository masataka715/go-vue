package main

import (
	"gortfolio/config"
	"gortfolio/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/", func(c *gin.Context) {
		people := models.Get_all()
		c.HTML(200, "index.html", gin.H{"people": people})
	})

	router.POST("/new", func(c *gin.Context) {
		name := c.PostForm("name")
		age, _ := strconv.Atoi(c.PostForm("age"))
		models.Create(name, age)
		c.Redirect(302, "/")
	})

	_ = router.Run(":" + config.Config.Port)
}
