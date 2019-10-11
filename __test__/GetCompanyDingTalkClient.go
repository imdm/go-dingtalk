package dingtalkTest

import (
	"os"

	"../src"
)

func GetDTClient() *dingtalk.Client {
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
	c := dingtalk.NewDTClient(config)
	return c
}
