package main

import (
	"RedisStudy/1-github_token/handler"
	"RedisStudy/global"
	"github.com/gin-gonic/gin"
)

func init() {
	global.InitGoRedisClient()
}
func main() {
	// https://github.com/login/oauth/authorize?client_id=f6e64e86547e9ff5e98c
	r := gin.Default()
	r.GET("/oauth2/redirect", handler.CodeHandler)
	r.GET("/user", handler.GetGitHubUser)
	r.Run(":9099")
}
