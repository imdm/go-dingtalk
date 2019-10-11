package dingtalkTest

import (
	"os"

	"../src"
)

func GetCompanyDingTalkClient() *dingtalk.DingTalkClient {
	CorpID := os.Getenv("CorpId")
	CorpSecret := os.Getenv("AppSecret")
	AgentID := os.Getenv("AgentID")
	SSOSecret := os.Getenv("SSOSecret")
	SNSAppID := os.Getenv("SNSAppID")
	SNSSecret := os.Getenv("SNSSecret")
	config := &dingtalk.DTConfig{
		AppKey:    CorpID,
		AppSecret: CorpSecret,
		AgentID:   AgentID,
		SSOSecret: SSOSecret,
		SNSAppID:  SNSAppID,
		SNSSecret: SNSSecret,
	}
	c := dingtalk.NewDingTalkCompanyClient(config)
	return c
}
