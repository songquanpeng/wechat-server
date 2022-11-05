package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type accessTokenStore struct {
	AccessToken       string
	Mutex             sync.RWMutex
	ExpirationSeconds int
}

type response struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

var s accessTokenStore

func InitAccessTokenStore() {
	go func() {
		for {
			RefreshAccessToken()
			s.Mutex.RLock()
			sleepDuration := Max(s.ExpirationSeconds, 60)
			s.Mutex.RUnlock()
			time.Sleep(time.Duration(sleepDuration) * time.Second)
		}
	}()
}

func RefreshAccessToken() {
	// https://developers.weixin.qq.com/doc/offiaccount/Basic_Information/Get_access_token.html
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", WeChatAppID, WeChatAppSecret), nil)
	if err != nil {
		SysError(err.Error())
		return
	}
	responseData, err := client.Do(req)
	if err != nil {
		SysError("failed to refresh access token: " + err.Error())
		return
	}
	defer responseData.Body.Close()
	var res response
	err = json.NewDecoder(responseData.Body).Decode(&res)
	if err != nil {
		SysError("failed to decode response: " + err.Error())
		return
	}
	s.Mutex.Lock()
	s.AccessToken = res.AccessToken
	s.ExpirationSeconds = res.ExpiresIn
	s.Mutex.Unlock()
	SysLog("access token refreshed")
}

func GetAccessTokenAndExpirationSeconds() (string, int) {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	return s.AccessToken, s.ExpirationSeconds
}

func GetAccessToken() string {
	s.Mutex.RLock()
	defer s.Mutex.RUnlock()
	return s.AccessToken
}
