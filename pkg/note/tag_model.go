package note

import (
	"gortfolio/pkg/common/database"
	"time"
)

type Tag struct {
	ID        int       `json:"id"`
	AuthID    int       `json:"auth_id"`
	TagName   string    `json:"tag_name"`
	CreatedAt time.Time `json:"created_at"`
}

func TagInsert(tag Tag) {
	db := database.Open()
	db.Create(&Tag{AuthID: tag.AuthID, TagName: tag.TagName})
	db.Close()
}

func GetAllTags(auth_id int) []Tag {
	db := database.Open()
	var tags []Tag
	db.Where("auth_id = ?", auth_id).Order("created_at desc").Find(&tags)
	db.Close()
	return tags
}
