package models

import (
	"github.com/spf13/cast"
	"strings"
	"tafexplorer/api"
	"tafexplorer/taf_server"
)

func VoteCount(req *api.VoteCountReq) (*api.VoteCountResp, error) {
	resp := &api.VoteCountResp{}
	resp.Peoples = 0

	// 投票数量
	str, _ := taf_server.GetTotalActivatedStake()
	totalActivatedStake := cast.ToFloat64(str)
	resp.Votes = int64(totalActivatedStake / 10000)

	// 货币总量
	str, _ = taf_server.GetCurrencyInfo()
	currency := cast.ToFloat64(strings.Split(str, ".")[0])
	resp.VoteRatio = int64(totalActivatedStake / 10000 / currency * 100)

	// 获取参与投票人数
	resp.Peoples, _ = taf_server.GetVoters()

	return resp, nil
}

func VoteList(req *api.VoteListReq) (*api.VoteListResp, error) {
	resp := &api.VoteListResp{}
	miners, err := taf_server.GetMiners()
	if err != nil {
		return resp, err
	}
	// 临时获取5个
	limit := 4
	for i, row := range miners.Rows {
		one := &api.VoteDetail{
			Miners:  row.Owner,
			VoteNum: 0,
			Share:   0,
			Bonus:   0,
		}
		resp.Data = append(resp.Data, one)
		if i >= limit {
			break
		}
	}
	return resp, nil
}
