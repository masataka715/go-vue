package models

import (
	"gotrading/config"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3" //忘れずに！
)

type Person struct {
	gorm.Model
	Name string
	Age  int
}

func init() {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&Person{})
}

func Create(name string, age int) {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	db.Create(&Person{Name: name, Age: age})
}

func Get_all() []Person {
	db, err := gorm.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatalln(err)
	}
	var people []Person
	db.Find(&people)
	return people

}
