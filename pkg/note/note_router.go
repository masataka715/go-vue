package note

import (
	"go_vue/pkg/common"
	"go_vue/pkg/common/database"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	database.Migrate(Note{})
	router.Use(common.AccessDenyMiddleware())
	noteRouter(router)

	database.Migrate(Garbage{})
	garbageRouter(router)

	database.Migrate(Tag{})
	tagRouter(router)

}

func noteRouter(router *gin.Engine) {
	router.POST("/note/new", func(ctx *gin.Context) {
		var note Note
		common.BindJSON(ctx, &note)
		NoteInsert(note)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/all/:auth_id", func(ctx *gin.Context) {
		auth_id := common.GetQueryID(ctx, "auth_id")
		// 新規ノート保存待ち
		time.Sleep(10 * time.Millisecond)
		note := GetAllNotes(auth_id)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/details/:note_id", func(ctx *gin.Context) {
		note_id := common.GetQueryID(ctx, "note_id")
		note := GetOneNote(note_id)
		common.SubmitJson(ctx, note)
	})

	router.POST("/note/details/", func(ctx *gin.Context) {
		var note Note
		common.BindJSON(ctx, &note)
		NoteUpdate(note)
		common.SubmitJson(ctx, note)
	})
}
