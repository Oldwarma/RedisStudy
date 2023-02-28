package main

import (
	"github.com/go-redis/redis/v9"
)

func main() {
	// 创建 Redis 客户端
	conn := redis.NewClient(&redis.Options{
		Addr:     "120.76.96.94:6379", // Redis 服务器地址
		Password: "",                  // Redis 密码
		DB:       0,                   // 使用的 Redis 数据库编号
	})

}
