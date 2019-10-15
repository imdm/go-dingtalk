package dingtalk

import (
	"fmt"
	"net/url"
	"time"
)

type SNSGetUserInfoResponse struct {
	OpenAPIResponse
	UserInfo SNSGetUserInfo `json:"user_info"`
}

type SNSGetUserInfo struct {
	MaskedMobile string
	Nick         string
	OpenID       string
	UnionID      string
	DingID       string
}

func (dtc *Client) SNSGetUserInfoByCode(code string) (*SNSGetUserInfoResponse, error) {
	var data SNSGetUserInfoResponse
	params := url.Values{}
	timestamp := fmt.Sprintf("%v", time.Now().Unix()*1000)
	params.Add("accessKey", dtc.DTConfig.SNSAppID)
	params.Add("timestamp", timestamp)
	signature := hmacSha256Sign(timestamp, dtc.DTConfig.SNSSecret)
	params.Add("signature", signature)
	rd := make(map[string]string)
	rd["tmp_auth_code"] = code
	err := dtc.httpSNS("sns/getuserinfo_bycode", params, rd, &data)
	return &data, err
}

