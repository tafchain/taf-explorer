package api

type HomepageCountReq struct {
}

// 主页上的统计内容
type HomepageCountResp struct {
	HeadBlockNum int64 `json:"head_block_num"` // 最新区块高度
	TradeNum     int   `json:"trade_num"`      // 交易总数量
	AccountNum   int   `json:"account_num"`    // 账户总数量
	VoteNum      int64 `json:"vote_num"`       // 投票数量
	ContractNum  int   `json:"contract_num"`   // 合约总数
	NodeNum      int64 `json:"node_num"`       // 节点数量
	TaftNum      int64 `json:"taft_num"`       // 质押TAFT总量
	DoVotes      int64 `json:"do_votes"`       // 参与投票人数
	VoteRatio    int64 `json:"vote_ratio"`     // 投票比例
}

type HomepageSearchReq struct {
	SearchContent string `json:"search_content"`
}

type HomepageSearchResp struct {
	DataType    int              `json:"data_type"`    // 0: 没查到结果 1:区块 2：交易 3:账户信息
	BlockInfo   *Block           `json:"block_info"`   // 查询的是区块
	TransInfo   *TransInfoResp   `json:"trans_info"`   // 查询的是交易
	AccountInfo *AccountInfoResp `json:"account_info"` // 查询的是账户信息
}
