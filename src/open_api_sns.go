package dingtalk

import (
	"fmt"
	"net/url"
	"time"
)

type SNSGetPersistentCodeResponse struct {
	OpenAPIResponse
	OpenID         string `json:"openid"`
	PersistentCode string `json:"persistent_code"`
	UnionID        string `json:"unionid"`
}

type SNSGetSNSTokenResponse struct {
	OpenAPIResponse
	Expires  int    `json:"expires_in"`
	SnsToken string `json:"sns_token"`
}

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

func (dtc *Client) SNSGetPersistentCode(code string) (SNSGetPersistentCodeResponse, error) {
	var data SNSGetPersistentCodeResponse
	requestData := map[string]string{
		"tmp_auth_code": code,
	}
	err := dtc.httpSNS("sns/get_persistent_code", nil, requestData, &data)
	return data, err
}

func (dtc *Client) SNSGetSNSToken(openID string, persistentCode string) (SNSGetSNSTokenResponse, error) {
	var data SNSGetSNSTokenResponse
	requestData := map[string]string{
		"openid":          openID,
		"persistent_code": persistentCode,
	}
	err := dtc.httpSNS("sns/get_sns_token", nil, requestData, &data)
	return data, err
}

func (dtc *Client) SNSGetUserInfo(snsToken string) (SNSGetUserInfoResponse, error) {
	var data SNSGetUserInfoResponse
	params := url.Values{}
	params.Add("sns_token", snsToken)
	err := dtc.httpSNS("sns/getuserinfo", params, nil, &data)
	return data, err
}

func (dtc *Client) SNSGetUserInfoByCode(code string) (SNSGetUserInfoResponse, error) {
	var data SNSGetUserInfoResponse
	params := url.Values{}
	timestamp := fmt.Sprintf("%v", time.Now().Unix()*1000)
	params.Add("accessKey", dtc.DTConfig.SNSAppID)
	params.Add("timestamp", timestamp)
	signature := hamcSha256Sign(timestamp, dtc.DTConfig.SNSSecret)
	params.Add("signature", signature)
	rd := url.Values{}
	rd.Add("tmp_auth_code", code)
	err := dtc.httpSNS("sns/getuserinfo_bycode", params, rd, &data)
	return data, err
}
