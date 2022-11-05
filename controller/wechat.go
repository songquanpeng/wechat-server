package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
	"time"
	"wechat-server/common"
)

type wechatResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func WeChatVerification(c *gin.Context) {
	// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Access_Overview.html
	signature := c.Query("signature")
	timestamp := c.Query("timestamp")
	nonce := c.Query("nonce")
	echoStr := c.Query("echostr")
	arr := []string{common.WeChatToken, timestamp, nonce}
	sort.Strings(arr)
	str := strings.Join(arr, "")
	hash := sha1.Sum([]byte(str))
	hexStr := hex.EncodeToString(hash[:])
	if signature == hexStr {
		c.String(http.StatusOK, echoStr)
	} else {
		c.Status(http.StatusForbidden)
	}
}

func ProcessWeChatMessage(c *gin.Context) {
	var req common.WeChatMessageRequest
	err := xml.NewDecoder(c.Request.Body).Decode(&req)
	if err != nil {
		common.SysError(err.Error())
		c.Abort()
		return
	}
	res := common.WeChatMessageResponse{
		ToUserName:   req.FromUserName,
		FromUserName: req.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      "",
	}
	common.ProcessWeChatMessage(&req, &res)
	if res.Content == "" {
		c.String(http.StatusOK, "")
		return
	}
	c.XML(http.StatusOK, &res)
}

func GetUserIDByCode(c *gin.Context) {
	code := c.Query("code")
	if code == "" {
		c.JSON(http.StatusOK, gin.H{
			"message": "无效的参数",
			"success": false,
		})
		return
	}
	id := common.GetWeChatIDByCode(code)
	c.JSON(http.StatusOK, gin.H{
		"message": "",
		"success": true,
		"data":    id,
	})
	return
}

func GetAccessToken(c *gin.Context) {
	accessToken, expiration := common.GetAccessTokenAndExpirationSeconds()
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "",
		"access_token": accessToken,
		"expiration":   expiration,
	})
}
