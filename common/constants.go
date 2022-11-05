package common

import (
	"github.com/google/uuid"
	"sync"
	"time"
)

var StartTime = time.Now().Unix() // unit: second
var Version = "v0.0.0"
var SystemName = "微信服务器"
var ServerAddress = "http://localhost:3000"
var FooterHTML = ""

// Any options with "Secret", "Token", "Key" in its key won't be return by GetOptions

var WeChatToken = ""
var WeChatAppID = ""
var WeChatAppSecret = ""
var WeChatEncodingAESKey = ""
var WeChatOwnerID = ""

var SessionSecret = uuid.New().String()
var SQLitePath = ".wechat-server.db"

var OptionMap map[string]string
var OptionMapRWMutex sync.RWMutex

var ItemsPerPage = 10

var PasswordLoginEnabled = true
var RegisterEnabled = false
var EmailVerificationEnabled = false
var GitHubOAuthEnabled = false

var SMTPServer = ""
var SMTPAccount = ""
var SMTPToken = ""

var GitHubClientId = ""
var GitHubClientSecret = ""

const (
	RoleGuestUser  = 0
	RoleCommonUser = 1
	RoleAdminUser  = 10
	RoleRootUser   = 100
)

var (
	FileUploadPermission    = RoleGuestUser
	FileDownloadPermission  = RoleGuestUser
	ImageUploadPermission   = RoleGuestUser
	ImageDownloadPermission = RoleGuestUser
)

// All duration's unit is seconds
// Shouldn't larger then RateLimitKeyExpirationDuration
var (
	GlobalApiRateLimitNum            = 20
	GlobalApiRateLimitDuration int64 = 60

	GlobalWebRateLimitNum            = 60
	GlobalWebRateLimitDuration int64 = 3 * 60

	UploadRateLimitNum            = 10
	UploadRateLimitDuration int64 = 60

	DownloadRateLimitNum            = 10
	DownloadRateLimitDuration int64 = 60

	CriticalRateLimitNum            = 5
	CriticalRateLimitDuration int64 = 10 * 60
)

var RateLimitKeyExpirationDuration = 20 * time.Minute

const (
	UserStatusEnabled  = 1 // don't use 0, 0 is the default value!
	UserStatusDisabled = 2 // also don't use 0
)
