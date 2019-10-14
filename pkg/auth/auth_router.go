package auth

import (
	"go_vue/pkg/common"
	"go_vue/pkg/common/database"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	database.Migrate(Auth{})
	// セッションの設定
	// store := cookie.NewStore([]byte("QT4wXnGSuq3i"))
	// router.Use(sessions.Sessions("gortfolio", store))

	router.POST("/register", func(ctx *gin.Context) {
		var auth Auth
		common.BindJSON(ctx, &auth)
		Insert(auth.EmailAddress, auth.Password)
		auth = GetMatchingAuth(auth)
		common.SubmitJson(ctx, auth)
	})

	router.POST("/login", func(ctx *gin.Context) {
		var auth Auth
		common.BindJSON(ctx, &auth)
		auth = GetMatchingAuth(auth)
		common.SubmitJson(ctx, auth)
	})
}
