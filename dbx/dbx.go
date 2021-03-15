package dbx

import (
	"dapan/config"
	"dapan/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetMysqlDb() {
	database := config.NewDefaultConf()

	con := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local", database.User, database.Password, database.Host, database.Port,database.DbName, database.Charset)

	db, err := gorm.Open(mysql.Open(con), &gorm.Config{})

	if (err != nil) {
		panic(err)
	}

	db.AutoMigrate(
		&model.UserInfo{},
		&model.MemoInfo{},
	)

	DB = db
}