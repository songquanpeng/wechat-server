package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"wechat-server/common"
	"wechat-server/model"
)

func GetOptions(c *gin.Context) {
	var options []*model.Option
	common.OptionMapRWMutex.Lock()
	for k, v := range common.OptionMap {
		if strings.Contains(k, "Token") || strings.Contains(k, "Secret") || strings.Contains(k, "Key") {
			continue
		}
		options = append(options, &model.Option{
			Key:   k,
			Value: common.Interface2String(v),
		})
	}
	common.OptionMapRWMutex.Unlock()
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
		"data":    options,
	})
	return
}

func UpdateOption(c *gin.Context) {
	var option model.Option
	err := json.NewDecoder(c.Request.Body).Decode(&option)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "无效的参数",
		})
		return
	}
	if option.Key == "GitHubOAuthEnabled" && option.Value == "true" && common.GitHubClientId == "" {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": "无法启用 GitHub OAuth，请先填入 GitHub Client ID 以及 GitHub Client Secret！",
		})
		return
	}
	err = model.UpdateOption(option.Key, option.Value)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	if option.Key == "WeChatMenu" {
		httpResponse, err := http.Post(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/menu/create?access_token=%s", common.GetAccessToken()), "application/json", bytes.NewBuffer([]byte(option.Value)))
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		defer httpResponse.Body.Close()
		var res wechatResponse
		err = json.NewDecoder(httpResponse.Body).Decode(&res)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": err.Error(),
			})
			return
		}
		if res.ErrCode != 0 {
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"message": res.ErrMsg,
			})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "",
	})
	return
}
