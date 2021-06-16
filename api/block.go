package api

type BlockListReq struct {
	PageIndex int `json:"page_index"` // 第几页
	PageSize  int `json:"page_size"`  // 页大小
}

type BlockListResp struct {
	PageNum   int      `json:"page_num"`   // 总页数
	PageSize  int      `json:"page_size"`  // 每页大小
	PageIndex int      `json:"page_index"` // 第几页
	Total     int      `json:"total"`      // 数据总量
	Data      []*Block `json:"data"`       // 当前查询返回的数据
}

type Block struct {
	BlockNum  int64  `json:"block_num"` // 区块高度
	Packer    string `json:"packer"`    // 打包者
	Verier    string `json:"verier"`    // 验证者
	TrxCount  int    `json:"trx_count"` // 交易数量
	Timestamp string `json:"timestamp"` // 时间戳
}

type BlockInfoReq struct {
	BlockNum int64 `json:"block_num"` // 区块高度
}

// verify
type BlockInfoResp struct {
	Packer     string `json:"packer"`
	PackTime   string `json:"pack_time"`
	Verier     string `json:"verier"`
	VerifyTime string `json:"verify_time"`
	HexData    string `json:"hex_data"`
	OriginData string `json:"origin_data"`
}
