package auth

import (
	"gortfolio/common"
	"gortfolio/common/database"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	database.Migrate(Auth{})
	router.POST("/register", func(ctx *gin.Context) {
		var auth Auth
		common.BindJSON(ctx, &auth)
		auth.UserName = "test"
		common.SubmitJson(ctx, auth)
	})
}
