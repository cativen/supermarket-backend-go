package common

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

var db *gorm.DB

func InitDataBase() {
	// 配置日志模式，启用打印所有SQL
	host := viper.GetString("datasource.host")
	port := viper.GetString("datasource.port")
	database := viper.GetString("datasource.database")
	username := viper.GetString("datasource.username")
	password := viper.GetString("datasource.password")
	charset := viper.GetString("datasource.charset")
	loc := viper.GetString("datasource.loc")

	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=%s",
		username,
		password,
		host,
		port,
		database,
		charset,
		url.QueryEscape(loc))
	DB, err := gorm.Open(mysql.Open(args), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db = DB
}

func GetDB() *gorm.DB {
	return db
}
