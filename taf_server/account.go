package taf_server

import (
"bytes"
"encoding/json"
"github.com/astaxie/beego/logs"
"io/ioutil"
)

//----------账户信息
type AccountResp struct {
	AccountName            string                   `json:"account_name"`
	HeadBlockNum           int64                    `json:"head_block_num"`
	Privileged             bool                     `json:"privileged"`
	LastCodeUpdate         string                   `json:"last_code_update"`
	Created                string                   `json:"created"`
	CoreLiquidBalance      string                   `json:"core_liquid_balance"`
	RamQuota               int                      `json:"ram_quota"`
	NetWeight              int                      `json:"net_weight"`
	NetLimit               map[string]interface{}   `json:"net_limit"`
	CpuLimit               map[string]interface{}   `json:"cpu_limit"`
	RamUsage               int64                    `json:"ram_usage"`
	Permissions            []map[string]interface{} `json:"permissions"`
	TotalResources         map[string]interface{}   `json:"total_resources"`
	SelfDelegatedBandwidth map[string]interface{}   `json:"self_delegated_bandwidth"`
	RefundRequest          map[string]interface{}   `json:"refund_request"`
	VoteInfo               map[string]interface{}   `json:"vote_info"`
}

type BalanceResp struct {
}

type getAccount struct {
	AccountName string `json:"account_name"`
}

type getBalance struct {
	Code    string `json:"code"`
	Account string `json:"account"`
	Symbol  string `json:"symbol"`
}

func TafGetAccount(accountName string) (*AccountResp, error) {
	param := &getAccount{AccountName: accountName}
	paramBytes, _ := json.Marshal(param)
	resp, err := Client().Post(accountApi, ContentType, bytes.NewBuffer(paramBytes))
	if err != nil {
		logs.Warn("req account err: ", err)
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	account := &AccountResp{}
	err = json.Unmarshal(body, account)
	if err != nil {
		logs.Warn("unmarshal account resp err: ", err)
	}
	return account, err
}

func TafGetBalance(name string) (string, error) {
	param := &getBalance{
		Code:    "tafsys.token",
		Account: name,
		Symbol:  "TAFT",
	}
	paramBytes, _ := json.Marshal(param)
	resp, err := Client().Post(balanceApi, ContentType, bytes.NewBuffer(paramBytes))
	if err != nil {
		logs.Warn("req balance err: ", err)
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	result := []string{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		logs.Warn("unmarshal balance resp err: ", err)
	}
	if len(result) == 0 {
		return "", nil
	}
	return result[0], nil
}
