package api

type TransListReq struct {
	PageIndex int `json:"page_index"`
	PageSize  int `json:"page_size"`
}

type TransListResp struct {
	PageNum   int      `json:"page_num"`
	PageSize  int      `json:"page_size"`
	PageIndex int      `json:"page_index"`
	Total     int      `json:"total"`
	Data      []*Trans `json:"data"`
}

type Trans struct {
	Id        string `json:"id"`
	Direction string `json:"direction"`
	BlockNum  int64  `json:"block_num"`
	Timestamp string `json:"timestamp"`
	Amount    string `json:"amount"`
}

type TransInfoReq struct {
	TransId string `json:"trans_id"`
}

type TransInfoResp struct {
	Status    int           `json:"status"`    // 执行状态 执行成功 0 - executed > 0 失败
	Hash      string        `json:"hash"`      // 交易哈希
	BlockNum  int64         `json:"block_num"` // 区块号
	Timestamp string        `json:"timestamp"` // 交易时间
	Actions   []*ActionData `json:"actions"`   // 整个交易的动作信息
}

// 交易中的动作
type ActionData struct {
	Id           string `json:"id"`            // 整个交易的id
	Direction    string `json:"direction"`     // 方向
	Amount       string `json:"amount"`        // 金额
	ContractName string `json:"contract_name"` // 合约
	OriginData   string `json:"origin_data"`   // 整个动作的原始数据
	ActionName   string `json:"action_name"`   // 动作名
}
