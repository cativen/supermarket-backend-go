package common

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func InitRedis() *redis.Client {
	// 创建日志记录器

	// 设置 go-redis 的日志输出
	// 配置日志模式，启用打印所有SQL
	host := viper.GetString("redis.host")
	port := viper.GetString("redis.port")
	database := viper.GetInt("redis.database")

	reds := redis.NewClient(&redis.Options{
		Addr:     host + ":" + port,
		DB:       database, // 数据库
		PoolSize: 20,       // 连接池大小
	})

	rdb = reds
	return reds
}

func GetRDB() *redis.Client {
	return rdb
}
