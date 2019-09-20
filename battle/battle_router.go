package battle

import (
	"gortfolio/config"
	"log"

	"github.com/gin-gonic/gin"
)

type Member struct {
	Name string `json:"name"`
	Hp   int    `json:"hp"`
	Mp   int    `json:"mp"`
}

func Router(router *gin.Engine) {
	router.POST("/battle", func(ctx *gin.Context) {
		var m Member
		err := ctx.BindJSON(&m)
		if err != nil {
			log.Println(err)
		}
		m.Hp = 80
		ctx.Header("Access-Control-Allow-Origin", config.Config.VueUrl)
		ctx.JSON(200, m)
	})
}
