package model

import (
	"errors"
	"strconv"
	"strings"
	"wechat-server/common"
)

type Option struct {
	Key   string `json:"key" gorm:"primaryKey"`
	Value string `json:"value"`
}

func AllOption() ([]*Option, error) {
	var options []*Option
	var err error
	err = DB.Find(&options).Error
	return options, err
}

func InitOptionMap() {
	common.OptionMapRWMutex.Lock()
	common.OptionMap = make(map[string]string)
	common.OptionMap["FileUploadPermission"] = strconv.Itoa(common.FileUploadPermission)
	common.OptionMap["FileDownloadPermission"] = strconv.Itoa(common.FileDownloadPermission)
	common.OptionMap["ImageUploadPermission"] = strconv.Itoa(common.ImageUploadPermission)
	common.OptionMap["ImageDownloadPermission"] = strconv.Itoa(common.ImageDownloadPermission)
	common.OptionMap["PasswordLoginEnabled"] = strconv.FormatBool(common.PasswordLoginEnabled)
	common.OptionMap["RegisterEnabled"] = strconv.FormatBool(common.RegisterEnabled)
	common.OptionMap["EmailVerificationEnabled"] = strconv.FormatBool(common.EmailVerificationEnabled)
	common.OptionMap["GitHubOAuthEnabled"] = strconv.FormatBool(common.GitHubOAuthEnabled)
	common.OptionMap["SMTPServer"] = ""
	common.OptionMap["SMTPAccount"] = ""
	common.OptionMap["SMTPToken"] = ""
	common.OptionMap["Notice"] = ""
	common.OptionMap["FooterHTML"] = common.FooterHTML
	common.OptionMap["ServerAddress"] = ""
	common.OptionMap["GitHubClientId"] = ""
	common.OptionMap["GitHubClientSecret"] = ""
	common.OptionMap["WeChatToken"] = ""
	common.OptionMap["WeChatAppID"] = ""
	common.OptionMap["WeChatAppSecret"] = ""
	common.OptionMap["WeChatEncodingAESKey"] = ""
	common.OptionMap["WeChatOwnerID"] = ""
	common.OptionMap["WeChatMenu"] = common.WeChatMenu
	common.OptionMapRWMutex.Unlock()
	options, _ := AllOption()
	for _, option := range options {
		updateOptionMap(option.Key, option.Value)
	}
}

func UpdateOption(key string, value string) error {
	if key == "StatEnabled" && value == "true" && !common.RedisEnabled {
		return errors.New("未启用 Redis，无法启用统计功能")
	}

	// Save to database first
	option := Option{
		Key:   key,
		Value: value,
	}
	// When updating with struct it will only update non-zero fields by default
	// So we have to use Select here
	if DB.Model(&option).Where("key = ?", key).Update("value", option.Value).RowsAffected == 0 {
		DB.Create(&option)
	}
	// Update OptionMap
	updateOptionMap(key, value)
	return nil
}

func updateOptionMap(key string, value string) {
	common.OptionMapRWMutex.Lock()
	defer common.OptionMapRWMutex.Unlock()
	common.OptionMap[key] = value
	if strings.HasSuffix(key, "Permission") {
		intValue, _ := strconv.Atoi(value)
		switch key {
		case "FileUploadPermission":
			common.FileUploadPermission = intValue
		case "FileDownloadPermission":
			common.FileDownloadPermission = intValue
		case "ImageUploadPermission":
			common.ImageUploadPermission = intValue
		case "ImageDownloadPermission":
			common.ImageDownloadPermission = intValue
		}
	}
	boolValue := value == "true"
	switch key {
	case "RegisterEnabled":
		common.RegisterEnabled = boolValue
	case "PasswordLoginEnabled":
		common.PasswordLoginEnabled = boolValue
	case "EmailVerificationEnabled":
		common.EmailVerificationEnabled = boolValue
	case "GitHubOAuthEnabled":
		common.GitHubOAuthEnabled = boolValue
	case "SMTPServer":
		common.SMTPServer = value
	case "SMTPAccount":
		common.SMTPAccount = value
	case "SMTPToken":
		common.SMTPToken = value
	case "ServerAddress":
		common.ServerAddress = value
	case "GitHubClientId":
		common.GitHubClientId = value
	case "GitHubClientSecret":
		common.GitHubClientSecret = value
	case "FooterHTML":
		common.FooterHTML = value
	case "WeChatToken":
		common.WeChatToken = value
	case "WeChatAppID":
		common.WeChatAppID = value
	case "WeChatAppSecret":
		common.WeChatAppSecret = value
	case "WeChatEncodingAESKey":
		common.WeChatEncodingAESKey = value
	case "WeChatOwnerID":
		common.WeChatOwnerID = value
	case "WeChatMenu":
		common.WeChatMenu = value
	}
}
