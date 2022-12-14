package common

import (
	"github.com/google/uuid"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"sync"
	"time"
)

type verificationValue struct {
	code string
	time time.Time
}

const (
	EmailVerificationPurpose  = "v"
	PasswordResetPurpose      = "r"
	WeChatVerificationPurpose = "w"
)

var verificationMutex sync.Mutex
var verificationMap map[string]verificationValue
var verificationMapMaxSize = 20
var VerificationValidMinutes = 3

func GenerateVerificationCode(length int) string {
	code := uuid.New().String()
	code = strings.Replace(code, "-", "", -1)
	if length == 0 {
		return code
	}
	return code[:length]
}

func GenerateAllNumberVerificationCode(length int) string {
	min := math.Pow10(length - 1)
	max := math.Pow10(length) - 1
	code := strconv.Itoa(rand.Intn(int(max-min)) + int(min))
	if GetWeChatIDByCode(code) != "" {
		SysError("repeated verification code detected")
		return GenerateAllNumberVerificationCode(length + 1)
	}
	return code
}

func RegisterWeChatCodeAndID(code string, id string) {
	RegisterVerificationCodeWithKey(code, id, WeChatVerificationPurpose)
}

func GetWeChatIDByCode(code string) string {
	verificationMutex.Lock()
	defer verificationMutex.Unlock()
	value, okay := verificationMap[WeChatVerificationPurpose+code]
	if okay {
		delete(verificationMap, WeChatVerificationPurpose+code)
	}
	now := time.Now()
	if !okay || int(now.Sub(value.time).Seconds()) >= VerificationValidMinutes*60 {
		return ""
	}
	return value.code
}

func RegisterVerificationCodeWithKey(key string, code string, purpose string) {
	verificationMutex.Lock()
	defer verificationMutex.Unlock()
	verificationMap[purpose+key] = verificationValue{
		code: code,
		time: time.Now(),
	}
	if len(verificationMap) > verificationMapMaxSize {
		removeExpiredPairs()
	}
}

func VerifyCodeWithKey(key string, code string, purpose string) bool {
	verificationMutex.Lock()
	defer verificationMutex.Unlock()
	value, okay := verificationMap[purpose+key]
	now := time.Now()
	if okay {
		delete(verificationMap, purpose+key)
	}
	if !okay || int(now.Sub(value.time).Seconds()) >= VerificationValidMinutes*60 {
		return false
	}
	return code == value.code
}

func DeleteKey(key string, purpose string) {
	verificationMutex.Lock()
	defer verificationMutex.Unlock()
	delete(verificationMap, purpose+key)
}

// no lock inside!
func removeExpiredPairs() {
	now := time.Now()
	for key := range verificationMap {
		if int(now.Sub(verificationMap[key].time).Seconds()) >= VerificationValidMinutes*60 {
			delete(verificationMap, key)
		}
	}
}

func init() {
	verificationMutex.Lock()
	defer verificationMutex.Unlock()
	verificationMap = make(map[string]verificationValue)
}
