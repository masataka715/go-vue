package twitter

import (
	"gortfolio/config"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //忘れずに！
)

type Twitter struct {
	ID        int       `json:"id"`
	Content   string    `json:"content"`
	NiceCount int       `json:"nice_count"`
	CreatedAt time.Time `json:"created_at"`
}

//DB初期化
func Init() {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Println(err)
	}
	db.AutoMigrate(&Twitter{})
	defer db.Close()
}

//DB追加
func Insert(content string) {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Println(err)
	}
	db.Create(&Twitter{Content: content})
	defer db.Close()
}

// //DB更新
// func Update(id int, text string, status string) {
// 	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
// 	if err != nil {
// 		panic("データベース開けず！（dbUpdate)")
// 	}
// 	var todo Todo
// 	db.First(&todo, id)
// 	todo.Text = text
// 	todo.Status = status
// 	db.Save(&todo)
// 	db.Close()
// }

//DB削除
func Delete(id int) {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Println(err)
	}
	var t Twitter
	db.First(&t, id)
	db.Delete(&t)
	db.Close()
}

//DB全取得
func GetAll() []Twitter {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Println(err)
	}
	var twitters []Twitter
	db.Order("created_at desc").Find(&twitters)
	db.Close()
	return twitters
}

// //DB一つ取得
// func GetOne(id int) Todo {
// 	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
// 	if err != nil {
// 		panic("データベース開けず！(dbGetOne())")
// 	}
// 	var todo Todo
// 	db.First(&todo, id)
// 	db.Close()
// 	return todo
// }

func AddNice(id int) {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Println(err)
	}
	var tweet Twitter
	db.First(&tweet, id)
	tweet.NiceCount++
	db.Save(&tweet)
	defer db.Close()
}
