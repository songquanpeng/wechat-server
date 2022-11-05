package controller

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
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

func GetAccessToken(c *gin.Context) {
	accessToken, expiration := common.GetAccessTokenAndExpirationSeconds()
	c.JSON(http.StatusOK, gin.H{
		"success":      true,
		"message":      "",
		"access_token": accessToken,
		"expiration":   expiration,
	})
}
