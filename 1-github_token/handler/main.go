package handler

import (
	"RedisStudy/1-github_token/model"
	"RedisStudy/global"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
)

const (
	userToKen     = "userToken:"
	clientId      = ""
	clientSecrets = ""
)

func CodeHandler(c *gin.Context) {
	code := c.Query("code")
	//验证
	token, err := getToken(code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "内部错误",
		})
	}
	c.JSON(http.StatusOK, token)
}
func getToken(code string) (*model.GithubToken, error) {
	client := http.Client{}
	params := fmt.Sprintf(`{"client_id":"%s"},"client_secret":"%s","code":"%s"`,
		clientId, clientSecrets, code)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBufferString(params))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	resp, err := io.ReadAll(res.Body)
	fmt.Println(string(resp))
	token := model.GithubToken{}
	json.Unmarshal(resp, &token)

	userId := rand.Intn(10000)
	global.GoRedisClient.Set()
}
