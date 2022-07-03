package database

import (
	"fmt"
	"github.com/graphql-be-demo/database/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var _db *gorm.DB

func InitDB() {
	if _db != nil {
		return
	}
	var dsn = formatConfigToDSN(getConfigByEnv())
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("开启数据库失败")
		log.Fatal(err)
		return
	}

	db.AutoMigrate(&model.User{})
	_db = db

}

func registerModels(v ...interface{}) {

}

func UseDB() *gorm.DB {
	if _db == nil {
		InitDB()
	}
	return _db
}
