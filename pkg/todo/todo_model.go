package todo

import (
	"go_vue/pkg/common/database"
	"time"

	_ "github.com/mattn/go-sqlite3" //忘れずに！
)

type Todo struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

//DB追加
func Insert(text string, status string) {
	db := database.Open()
	db.Create(&Todo{Text: text, Status: status})
	defer db.Close()
}

//DB更新
func Update(id int, text string, status string) {
	db := database.Open()
	var todo Todo
	db.First(&todo, id)
	todo.Text = text
	todo.Status = status
	db.Save(&todo)
	db.Close()
}

//DB削除
func Delete(id int) {
	db := database.Open()
	var todo Todo
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

//DB全取得
func GetAll() []Todo {
	db := database.Open()
	var todos []Todo
	db.Order("created_at desc").Find(&todos)
	db.Close()
	return todos
}

//DB一つ取得
func GetOne(id int) Todo {
	db := database.Open()
	var todo Todo
	db.First(&todo, id)
	db.Close()
	return todo
}
