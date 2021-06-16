package api

type VoteCountReq struct {
}

type VoteCountResp struct {
	Votes     int64 `json:"votes"`
	Peoples   int64 `json:"peoples"`
	VoteRatio int64 `json:"vote_ratio"`
}

type VoteListReq struct {
}

type VoteListResp struct {
	Data []*VoteDetail `json:"data"`
}

type VoteDetail struct {
	Miners  string `json:"miners"`   // 矿工
	VoteNum int    `json:"vote_num"` // 得票数
	Share   int    `json:"share"`    // 分红比例
	Bonus   int64  `json:"bonus"`    // 奖金
}
