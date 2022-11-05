package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-template/common"
	"gin-template/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type GitHubOAuthResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

type GitHubUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func GitHubOAuth(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "非法的参数",
		})
		return
	}
	values := map[string]string{"client_id": common.GitHubClientId, "client_secret": common.GitHubClientSecret, "code": code}
	jsonData, err := json.Marshal(values)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	defer res.Body.Close()
	var oAuthResponse GitHubOAuthResponse
	err = json.NewDecoder(res.Body).Decode(&oAuthResponse)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	req, err = http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oAuthResponse.AccessToken))
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	res2, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	defer res2.Body.Close()
	var githubUser GitHubUser
	err = json.NewDecoder(res2.Body).Decode(&githubUser)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	user := model.User{
		Email: githubUser.Email,
	}
	if githubUser.Email != "" && model.IsEmailAlreadyTaken(githubUser.Email) {
		user.FillUserByEmail()
	} else {
		if githubUser.Login == "" {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": "返回值非法，用户字段为空",
			})
			return
		}
		user.Username = "github_" + githubUser.Login
		if model.IsUsernameAlreadyTaken(user.Username) {
			user.FillUserByUsername()
		} else {
			user.DisplayName = githubUser.Name
			user.Role = common.RoleCommonUser
			user.Status = common.UserStatusEnabled

			if !common.RegisterEnabled {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": "管理员关闭了新用户注册",
				})
				return
			}

			if err := user.Insert(); err != nil {
				c.JSON(http.StatusOK, gin.H{
					"success": false,
					"message": err.Error(),
				})
				return
			}
		}
	}
	if user.Status != common.UserStatusEnabled {
		c.JSON(http.StatusOK, gin.H{
			"message": "用户已被封禁",
			"success": false,
		})
		return
	}
	setupLogin(&user, c)
}
