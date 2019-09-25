package note

import (
	"gortfolio/common"
	"gortfolio/common/database"
	"time"

	"github.com/gin-gonic/gin"
)

func Router(router *gin.Engine) {
	database.Migrate(Note{})

	router.POST("/note/new", func(ctx *gin.Context) {
		var note Note
		common.BindJSON(ctx, &note)
		NoteInsert(note)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/all/:id", func(ctx *gin.Context) {
		auth_id := common.GetQueryID(ctx)
		// 新規ノート保存待ち
		time.Sleep(10 * time.Millisecond)
		note := GetAllNotes(auth_id)
		common.SubmitJson(ctx, note)
	})

	router.GET("/note/details/:id", func(ctx *gin.Context) {
		note_id := common.GetQueryID(ctx)
		note := GetOneNote(note_id)
		common.SubmitJson(ctx, note)
	})

	router.POST("/note/details/", func(ctx *gin.Context) {
		var note Note
		common.BindJSON(ctx, &note)
		NoteUpdate(note)
		common.SubmitJson(ctx, note)
	})

	database.Migrate(Garbage{})

	router.GET("/note/garbage_all/:id", func(ctx *gin.Context) {
		auth_id := common.GetQueryID(ctx)
		garbages := GetGarbageAll(auth_id)
		common.SubmitJson(ctx, garbages)
	})

	router.GET("/note/garbage/:id", func(ctx *gin.Context) {
		note_id := common.GetQueryID(ctx)
		note := NoteDel(note_id)
		GarbageInsert(note)
		notes := GetAllNotes(note.AuthID)
		common.SubmitJson(ctx, notes)
	})
}
