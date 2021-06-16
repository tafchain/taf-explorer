package models

import (
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/cast"
	"strings"
	"tafexplorer/api"
	"tafexplorer/db"
	"tafexplorer/framework/cache"
	"tafexplorer/taf_server"
)

func HomepageCount(req *api.HomepageCountReq) (*api.HomepageCountResp, error) {
	resp := &api.HomepageCountResp{}
	headBlockNum, _ := cache.Get(cache.HeadBlockNum)
	resp.HeadBlockNum = cast.ToInt64(headBlockNum)
	// 交易数量
	resp.TradeNum, _ = db.DB().C(db.Transactions).Count()
	// 账户数量
	resp.AccountNum, _ = db.DB().C(db.Accounts).Count()
	resp.ContractNum, _ = db.DB().C(db.Contracts).Count()
	resp.NodeNum = 1000
	resp.DoVotes = 1000
	// TAFT质押总量
	str, _ := taf_server.GetTotalBalance()
	resp.TaftNum = cast.ToInt64(strings.Split(str, ".")[0])

	// total_activated_stake
	str, _ = taf_server.GetTotalActivatedStake()
	totalActivatedStake := cast.ToFloat64(str)
	resp.VoteNum = int64(totalActivatedStake / 10000)
	// currency
	str, _ = taf_server.GetCurrencyInfo()
	currency := cast.ToFloat64(strings.Split(str, ".")[0])

	resp.VoteRatio = int64(totalActivatedStake / 10000 / currency * 100)
	resp.DoVotes, _ = taf_server.GetVoters()

	// 矿工数据
	miners, _ := taf_server.GetMiners()
	resp.NodeNum = int64(len(miners.Rows))

	return resp, nil
}

func HomepageSearch(req *api.HomepageSearchReq) (*api.HomepageSearchResp, error) {
	resp := &api.HomepageSearchResp{}
	err := searchBlock(req, resp)
	if err == nil {
		resp.DataType = 1
	}
	err = searchTrans(req, resp)
	if err == nil {
		resp.DataType = 2
	}
	err = searchAccount(req, resp)
	if err == nil {
		resp.DataType = 3
	}
	return resp, nil
}

// 查询区块
func searchBlock(req *api.HomepageSearchReq, resp *api.HomepageSearchResp) error {
	result := []db.Block{}
	blockNum := cast.ToInt64(req.SearchContent)
	if blockNum == 0 {
		return errors.New("not blockNum")
	}
	err := db.DB().C(db.BlockName).Find(bson.M{"block_num": blockNum}).All(&result)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("not find block")
	}

	block := result[0]
	resp.BlockInfo = &api.Block{
		BlockNum:  block.Num,
		Packer:    "",
		Verier:    "",
		TrxCount:  block.TrxCount,
		Timestamp: block.Timestamp,
	}
	return nil
}

// 查询交易
func searchTrans(req *api.HomepageSearchReq, resp *api.HomepageSearchResp) error {
	result := make([]db.Transaction, 0)
	err := db.DB().C(db.Transactions).Find(bson.M{"trx.id": req.SearchContent}).All(&result)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("not find trans")
	}
	resp.TransInfo = &api.TransInfoResp{}

	trans := result[0]
	resp.TransInfo.BlockNum = trans.BlockNum
	resp.TransInfo.Timestamp = trans.TRX.Transaction.Expiration
	resp.TransInfo.Hash = trans.TRX.ID
	resp.TransInfo.Status = 0
	if trans.Status != "executed" {
		resp.TransInfo.Status = 1
	}
	// actions
	for _, action := range trans.TRX.Transaction.Actions {
		one := &api.ActionData{}
		one.Direction, one.Amount = parseAction(action)
		one.ContractName = action.Account
		one.ActionName = action.Name
		originData, _ := json.Marshal(action.Data)
		one.OriginData = string(originData)
		resp.TransInfo.Actions = append(resp.TransInfo.Actions, one)
	}
	return nil
}

// 查询账户
func searchAccount(req *api.HomepageSearchReq, resp *api.HomepageSearchResp) error {
	result := []db.Account{}
	err := db.DB().C(db.Accounts).Find(bson.M{"$or": []bson.M{bson.M{"name": req.SearchContent}, bson.M{"public_key": req.SearchContent}}}).All(&result)
	if err != nil {
		return err
	}
	if len(result) == 0 {
		return errors.New("not find account")
	}
	account := result[0]
	resp.AccountInfo = &api.AccountInfoResp{
		AccountName: account.Name,
		PublicKey:   account.PublicKey,
		Balance:     "",
		Creator:     account.Creator,
		CreateTime:  account.CreateTime,
		TradeData:   nil,
	}
	resp.AccountInfo.TradeData = []*api.ActionData{}

	// 查询账户下最近的一条交易 todo redis中获取
	trans := []db.Transaction{}
	err = db.DB().C(db.Transactions).Find(bson.M{"trx.transaction.actions": bson.M{"$elemMatch": bson.M{"name": req.SearchContent}}}).Sort("-block_num").All(&trans)
	if err == nil && len(trans) > 0 {
		for _, action := range trans[0].TRX.Transaction.Actions {
			one := &api.ActionData{}
			one.Direction, one.Amount = parseAction(action)
			one.ContractName = action.Account
			one.ActionName = action.Name
			one.Id = trans[0].TRX.ID
			resp.AccountInfo.TradeData = append(resp.AccountInfo.TradeData, one)
		}
	}

	// 查询金额
	resp.AccountInfo.Balance, _ = taf_server.TafGetBalance(resp.AccountInfo.AccountName)
	return nil
}
