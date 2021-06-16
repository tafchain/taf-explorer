package api

type AccountInfoReq struct {
	AccountName string `json:"account_name"` // 账号名
}

type AccountInfoResp struct {
	AccountName string        `json:"account_name"` // 账号名
	PublicKey   string        `json:"public_key"`   // 公钥
	Balance     string        `json:"balance"`      // 余额
	Creator     string        `json:"creator"`      // 创建者
	CreateTime  string        `json:"create_time"`  // 创建时间
	TradeData   []*ActionData `json:"trade_data"`   // 交易数据   单个交易里面的详细动作
}
