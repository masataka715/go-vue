package twitter

import (
	"go_vue/pkg/common/database"
	"time"
)

type Twitter struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	NiceCount int       `json:"nice_count"`
	CreatedAt time.Time `json:"created_at"`
}

//DB追加
func Insert(content string) {
	db := database.Open()
	db.Create(&Twitter{Content: content})
	defer db.Close()
}

//DB削除
func Delete(id int) {
	db := database.Open()
	var t Twitter
	db.First(&t, id)
	db.Delete(&t)
	db.Close()
}

//DB全取得
func GetAll() []Twitter {
	db := database.Open()
	var twitters []Twitter
	db.Order("created_at desc").Find(&twitters)
	db.Close()
	return twitters
}

func AddNice(id int) {
	db := database.Open()
	var tweet Twitter
	db.First(&tweet, id)
	tweet.NiceCount++
	db.Save(&tweet)
	defer db.Close()
}
