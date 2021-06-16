package models

import (
	"errors"
	"github.com/globalsign/mgo/bson"
	"strings"
	"tafexplorer/api"
	"tafexplorer/db"
	"tafexplorer/taf_server"
)

func ContractList(req *api.ContractListReq) (*api.ContractListResp, error) {
	resp := &api.ContractListResp{}
	skip := (req.PageIndex - 1) * req.PageSize
	result := make([]db.Contract, 0)
	err := db.DB().C(db.Contracts).Find(bson.M{}).Skip(skip).Limit(req.PageSize).All(&result)
	if err != nil {
		return resp, err
	}
	for _, one := range result {
		n := &api.Contract{
			Name:     one.Name,
			TradeNum: one.TransCount,
			Actions:  strings.Join(one.Actions, ","),
		}
		resp.Data = append(resp.Data, n)
	}
	count, _ := db.DB().C(db.Contracts).Count()
	resp.Total = count
	resp.PageSize = req.PageSize
	resp.PageIndex = req.PageIndex
	resp.PageNum = count / req.PageSize
	return resp, nil
}

func ContractInfo(req *api.ContractInfoReq) (*api.ContractInfoResp, error) {
	resp := &api.ContractInfoResp{}

	result := []db.Account{}
	err := db.DB().C(db.Accounts).Find(bson.M{"name": req.ContractName}).All(&result)
	if err != nil {
		return resp, errors.New("not find account")
	}
	if len(result) == 0 {
		return resp, errors.New("not find account")
	}
	account := result[0]
	resp = &api.ContractInfoResp{
		AccountName: account.Name,
		PublicKey:   account.PublicKey,
		Balance:     "",
		Creator:     account.Creator,
		CreateTime:  account.CreateTime,
		TradeData:   nil,
	}
	resp.TradeData = []*api.ActionData{}
	// 查询账户下最近的一条交易 todo redis中获取
	trans := []db.Transaction{}
	err = db.DB().C(db.Transactions).Find(bson.M{"trx.transaction.actions": bson.M{"$elemMatch": bson.M{"name": req.ContractName}}}).Sort("-block_num").All(&trans)
	if err == nil && len(trans) > 0 {
		for _, action := range trans[0].TRX.Transaction.Actions {
			one := &api.ActionData{}
			one.Direction, one.Amount = parseAction(action)
			one.ContractName = action.Account
			one.ActionName = action.Name
			one.Id = trans[0].TRX.ID
			resp.TradeData = append(resp.TradeData, one)
		}
	}

	// 查询金额
	resp.Balance, _ = taf_server.TafGetBalance(resp.AccountName)
	return resp, nil
}
