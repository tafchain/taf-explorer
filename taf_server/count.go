package taf_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

// 获取一些统计值 比如质押总量

// 质押总量
type totalBalanceReq struct {
	Json          bool   `json:"json"`
	Code          string `json:"code"`
	Scope         string `json:"scope"`
	Table         string `json:"table"`
	TableKey      string `json:"table_key"`
	LowerBound    string `json:"lower_bound"`
	UpperBound    string `json:"upper_bound"`
	Limit         int    `json:"limit"`
	KeyType       string `json:"key_type"`
	IndexPosition string `json:"index_position"`
	EncodeType    string `json:"encode_type"`
	Reverse       bool   `json:"reverse"`
	ShowPayer     bool   `json:"show_payer"`
}

type totalBalanceResp struct {
	Rows    []map[string]interface{} `json:"rows"`
	More    bool                     `json:"more"`
	NextKey string                   `json:"next_key"`
}

// 货币总量
type currencyInfoReq struct {
	Json   bool   `json:"json"`
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
}

// 矿工信息
type miner struct {
	Owner         string        `json:"owner"`
	TotalVotes    string        `json:"total_votes"`
	ProducerKey   string        `json:"producer_key"`
	IsActive      bool          `json:"is_active"`
	Url           string        `json:"url"`
	UnpaiedBlocks int           `json:"unpaied_blocks"`
	LastClaimTime string        `json:"last_claim_time"`
	Location      int           `json:"location"`
	Authority     []interface{} `json:"authority"`
}

type minersInfoReq struct {
	Limit      int    `json:"limit"` // -1
	LowerBound string `json:"lower_bound"`
	Json       bool   `json:"json"`
}

type minersInfoResp struct {
	Rows                    []miner `json:"rows"`
	TotalProducerVoteWeight string  `json:"total_producer_vote_weight"`
	More                    string  `json:"more"`
}

type votersInfoReq struct {
	Json          bool   `json:"json"`
	Code          string `json:"code"`
	Scope         string `json:"scope"`
	Table         string `json:"table"`
	TableKey      string `json:"table_key"`
	LowerBound    string `json:"lower_bound"`
	UpperBound    string `json:"upper_bound"`
	Limit         int    `json:"limit"`
	KeyTypes      string `json:"key_types"`
	IndexPosition string `json:"index_position"`
	EncodeType    string `json:"encode_type"`
	Reverse       bool   `json:"reverse"`
	ShowPayer     bool   `json:"show_payer"`
}

// 获取投票人数
type votersInfo struct {
	Rows    []voterInfo `json:"rows"`
	More    bool        `json:"more"`
	NextKey string      `json:"next_key"`
}

// 单个投票信息
type voterInfo struct {
	Owner             string      `json:"owner"`
	Proxy             string      `json:"proxy"`
	Producers         []string    `json:"producers"`
	Staked            interface{} `json:"staked"`
	LastVoteWeight    string      `json:"last_vote_weight"`
	ProxiedVoteWeight string      `json:"proxied_vote_weight"`
	IsProxy           int         `json:"is_proxy"`
	Flags1            int         `json:"flags1"`
	Reserved2         int         `json:"reserved2"`
	Reserved3         string      `json:"reserved3"`
}

// TAFT质押总量
func GetTotalBalance() (string, error) {
	param := &totalBalanceReq{
		Json:          true,
		Code:          "tafsys.token",
		Scope:         "tafsys.stake",
		Table:         "accounts",
		TableKey:      "",
		LowerBound:    "",
		UpperBound:    "",
		Limit:         1,
		KeyType:       "",
		IndexPosition: "",
		EncodeType:    "dec",
		Reverse:       false,
		ShowPayer:     false,
	}
	paramBytes, _ := json.Marshal(param)
	resp, err := Client().Post(countApi, ContentType, bytes.NewBuffer(paramBytes))
	if err != nil {
		logs.Warn("req balance err: ", err)
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := &totalBalanceResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return "", err
	}
	if len(result.Rows) == 0 {
		return "", nil
	}
	rows := result.Rows
	if _, ok := rows[0]["balance"]; !ok {
		return "", errors.New("no total TAFT balance")
	}

	return rows[0]["balance"].(string), nil
}

// total_activated_stake string
func GetTotalActivatedStake() (string, error) {
	param := &totalBalanceReq{
		Json:          true,
		Code:          "tafsys",
		Scope:         "tafsys",
		Table:         "global",
		TableKey:      "",
		LowerBound:    "",
		UpperBound:    "",
		Limit:         1,
		KeyType:       "",
		IndexPosition: "",
		EncodeType:    "dec",
		Reverse:       false,
		ShowPayer:     false,
	}
	paramBytes, _ := json.Marshal(param)
	resp, err := Client().Post(countApi, ContentType, bytes.NewBuffer(paramBytes))
	if err != nil {
		logs.Warn("req total activated stake err: ", err)
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	result := &totalBalanceResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return "", err
	}
	if len(result.Rows) == 0 {
		return "", nil
	}
	rows := result.Rows
	if _, ok := rows[0]["total_activated_stake"]; !ok {
		return "", errors.New("no total activated stake")
	}
	return rows[0]["total_activated_stake"].(string), nil
}

func GetCurrencyInfo() (string, error) {
	param := &currencyInfoReq{
		Json:   false,
		Code:   "tafsys.token",
		Symbol: "TAFT",
	}
	paramBytes, _ := json.Marshal(param)
	resp, err := Client().Post(currencyApi, ContentType, bytes.NewBuffer(paramBytes))
	if err != nil {
		logs.Warn("req total activated stake err: ", err)
		return "", err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]interface{}{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}
	sysInfo := result["TAFT"].(map[string]interface{})
	if _, ok := sysInfo["supply"]; !ok {
		return "", errors.New("no supply")
	}
	return sysInfo["supply"].(string), nil
}

// 获取矿工信息
func GetMiners() (*minersInfoResp, error) {
	param := &minersInfoReq{
		Limit:      -1,
		LowerBound: "",
		Json:       true,
	}
	paramBytes, _ := json.Marshal(param)
	resp, err := Client().Post(minersApi, ContentType, bytes.NewBuffer(paramBytes))
	if err != nil {
		logs.Warn("req total activated stake err: ", err)
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := &minersInfoResp{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// 获取投票人数
func GetVoters() (int64, error) {
	param := &votersInfoReq{
		Json:          true,
		Code:          "tafsys",
		Scope:         "tafsys",
		Table:         "voters",
		TableKey:      "",
		LowerBound:    "",
		UpperBound:    "",
		Limit:         -1,
		KeyTypes:      "",
		IndexPosition: "",
		EncodeType:    "dec",
		Reverse:       false,
		ShowPayer:     false,
	}
	paramBytes, _ := json.Marshal(param)
	resp, err := Client().Post(votersApi, ContentType, bytes.NewBuffer(paramBytes))
	if err != nil {
		logs.Warn("req voters err: ", err)
		return 0, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	result := &votersInfo{}
	err = json.Unmarshal(body, result)
	if err != nil {
		return 0, err
	}

	return int64(len(result.Rows)), nil
}
