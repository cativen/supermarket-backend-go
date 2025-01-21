package main

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"os"
	"superMarket-backend/common"
	"superMarket-backend/route"
)

func main() {
	initConfig()
	common.InitDataBase()
	common.InitRedis()
	common.InitLog()
	common.InitPay()
	routes := route.Routers()
	port := viper.GetString("server.port")
	routes.Run(":" + port)
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}
