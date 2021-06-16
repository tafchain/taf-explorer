package models

import (
	"github.com/globalsign/mgo/bson"
	"tafexplorer/api"
	"tafexplorer/db"
	"tafexplorer/taf_server"
)

func AccountInfo(req *api.AccountInfoReq) (*api.AccountInfoResp, error) {
	resp := &api.AccountInfoResp{}
	if len(req.AccountName) == 0 {
		// 	没有值默认查询最近的一个
		req.AccountName = findLatestAccount()
	}
	balance, err := taf_server.TafGetBalance(req.AccountName)
	if err != nil {
		return resp, err
	}
	result := []db.Account{}
	err = db.DB().C(db.Accounts).Find(bson.M{"name": req.AccountName}).All(&result)
	if err != nil {
		return resp, err
	}
	if len(result) == 0 {
		return resp, nil
	}
	account := result[0]
	resp.AccountName = account.Name
	resp.Balance = balance
	resp.CreateTime = account.CreateTime
	resp.PublicKey = account.PublicKey
	resp.Creator = account.Creator
	resp.TradeData = []*api.ActionData{}
	// 查询账户下最近的一条交易 todo redis中获取
	trans := []db.Transaction{}
	err = db.DB().C(db.Transactions).Find(bson.M{"accounts": req.AccountName}).Sort("-time_unix").Limit(1).All(&trans)
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
	return resp, nil
}

func findLatestAccount() string {
	var account string
	result := []db.Account{}
	err := db.DB().C(db.Accounts).Find(bson.M{}).Sort("-time_unix").Limit(1).All(&result)
	if err != nil {
		return account
	}
	accountInfo := result[0]
	return accountInfo.Name
}
