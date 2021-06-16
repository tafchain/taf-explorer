package api

type ContractListReq struct {
	PageIndex int `json:"page_index"` // 第几页
	PageSize  int `json:"page_size"`  // 页大小
}

type ContractListResp struct {
	PageNum   int         `json:"page_num"`   // 总页数
	PageSize  int         `json:"page_size"`  // 每页大小
	PageIndex int         `json:"page_index"` // 第几页
	Total     int         `json:"total"`      // 数据总量  合约总数
	Data      []*Contract `json:"data"`       // 当次查询的列表数据
}

type Contract struct {
	Name     string `json:"name"`      // 合约名
	TradeNum int64  `json:"trade_num"` // 交易数量
	Actions  string `json:"actions"`   // 动作
}

type ContractInfoReq struct {
	ContractName string `json:"contract_name"`
}

type ContractInfoResp struct {
	AccountName string        `json:"account_name"` // 账号名
	PublicKey   string        `json:"public_key"`   // 公钥
	Balance     string        `json:"balance"`      // 余额
	Creator     string        `json:"creator"`      // 创建者
	CreateTime  string        `json:"create_time"`  // 创建时间
	TradeData   []*ActionData `json:"trade_data"`   // 交易数据
}
