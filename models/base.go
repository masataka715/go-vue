package models

import (
	"database/sql"
	"fmt"
	"gotrading/config"
	"log"

	_ "github.com/mattn/go-sqlite3" //忘れずに！
)

const (
	tableName = "todo"
)

// データベース操作に必要なのは、データコネクションとコマンド
var DbConnection *sql.DB

func init() {
	var err error
	// データコネクションに必要なのは、ドライバー名とSQLファイル名
	DbConnection, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	// テーブル名をプレースホルダ（代替物）で書くときは、fmt.Sprintfを使う。sが付くのは、文字列を返すという意味
	cmd := fmt.Sprintf(`
			CREATE TABLE IF NOT EXISTS %s (
				id INTEGER PRIMARY KEY NOT NULL,
				content STRING)`, tableName)
	_, _ = DbConnection.Exec(cmd)

}
