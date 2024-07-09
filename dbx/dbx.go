package dbx

import (
	"dapan/config"
	"dapan/model"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitMysqlDb() {
	database := config.NewDefaultConf()

	dns := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local&timeout=%s", database.User, database.Password, database.Host, database.Port, database.DbName, database.Charset, database.Timeout)

	fmt.Println(dns)
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	db.AutoMigrate(
		&model.UserInfo{},
		// &model.MemoInfo{},
		&model.MenuClassify{},
		&model.Menu{},
		&model.Order{},
		&model.OrderDetail{},
	)

	DB = db
}
