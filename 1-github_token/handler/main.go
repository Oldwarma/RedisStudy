package handler

import (
	"RedisStudy/1-github_token/model"
	"RedisStudy/global"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"math/rand"
	"net/http"
	"time"
)

const (
	userToken     = "userToken:"
	clientId      = "6d1212170052bf137266"
	clientSecrets = "18ec858fd085bdc02035dfea2ca7f9913852d0b6"
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
func GetGitHubUser(c *gin.Context) {
	userId := c.Query("userid")
	tokenPrefix := fmt.Sprintf("%s%s", userToken, userId)
	//取值
	result, err := global.GoRedisClient.Get(context.Background(), tokenPrefix).Result()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "参数错误"})
		return
	}
	var token model.GithubToken
	err = json.Unmarshal([]byte(result), &token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"mag": "转换错误"})
		return
	}
	client := http.Client{}
	//开始获取用户信息
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Add("Authorization", "Bearer"+token.AccessToken)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
		return
	}
	if res.StatusCode != 200 {
		fmt.Println("using github token to fetch User Info failed with not 200 error")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "内部错误"})
		return
	}
	defer res.Body.Close()
	bs, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	user := &model.GithubUser{}
	err = json.Unmarshal(bs, user)
	if err != nil {
		fmt.Println(err)
		return
	}
	user.Token = token.AccessToken
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
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

	//随机数
	userId := rand.Intn(10000)
	//存值
	result, err := global.GoRedisClient.Set(context.Background(), fmt.Sprintf("%s%d", userToken, userId),
		string(resp), 8*time.Hour).Result()
	if err != nil {
		return nil, err
	}
	fmt.Println(result)
	return &token, nil
}
