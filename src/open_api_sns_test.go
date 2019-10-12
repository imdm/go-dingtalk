package dingtalk

import (
	"fmt"
	"testing"
)

func TestClient_SNSGetUserInfoByCode(t *testing.T) {
	config := &DTConfig{
		AppKey:    "dingmktfpzxp7wkdfepf",
		AppSecret: "wy_oZGaU_Gzv_O6vbElQ6O9pCvwAzBYEtiNvCdvcuZwJdCFMtvTn9fD7TYSD4Y_o",
		SNSAppID:  "dingoapviyh1rlkjz5eheg",
		SNSSecret: "RmD9nRRv83HyxOr-_rRBKIAUeAhB9XGawBfbMx8UUcTn1m-Oyx1V0MO2O-tRV1fH",
	}
	c := NewDTClient(config)
	tmpCode := "bc5496371d793a09b7791098f8218b05"
	info, err := c.SNSGetUserInfoByCode(tmpCode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(info)
}
