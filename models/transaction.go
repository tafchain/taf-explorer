package models

import (
	"encoding/json"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/cast"
	"tafexplorer/api"
	"tafexplorer/db"
)

func TransactionList(req *api.TransListReq) (*api.TransListResp, error) {
	resp := &api.TransListResp{}
	skip := (req.PageIndex - 1) * req.PageSize
	result := make([]db.Transaction, 0)
	err := db.DB().C(db.Transactions).Find(bson.M{}).Sort("-block_num").Skip(skip).Limit(req.PageSize).All(&result)
	if err != nil {
		return resp, err
	}
	for _, one := range result {
		trans := &api.Trans{
			Id:        one.TRX.ID,
			BlockNum:  one.BlockNum,
			Timestamp: one.TRX.Transaction.Expiration,
		}
		// 从交易的actions中获取方向、金额
		resp.Data = append(resp.Data, trans)
	}
	count, _ := db.DB().C(db.Transactions).Count()
	resp.Total = count
	resp.PageSize = req.PageSize
	resp.PageIndex = req.PageIndex
	resp.PageNum = count / req.PageSize
	return resp, nil
}

func TransactionInfo(req *api.TransInfoReq) (*api.TransInfoResp, error) {
	resp := &api.TransInfoResp{}
	result := []db.Transaction{}
	err := db.DB().C(db.Transactions).Find(bson.M{"trx.id": req.TransId}).All(&result)
	if err != nil {
		return resp, err
	}
	if len(result) == 0 {
		return resp, nil
	}
	trans := result[0]
	resp.BlockNum = trans.BlockNum
	resp.Timestamp = trans.TRX.Transaction.Expiration
	resp.Hash = trans.TRX.ID
	resp.Status = 0
	if trans.Status != "executed" {
		resp.Status = 1
	}
	// actions
	for _, action := range trans.TRX.Transaction.Actions {
		one := &api.ActionData{}
		one.Direction, one.Amount = parseAction(action)
		one.ContractName = action.Account
		one.ActionName = action.Name
		originData, _ := json.Marshal(action.Data)
		one.OriginData = string(originData)
		resp.Actions = append(resp.Actions, one)
	}
	return resp, nil
}

func parseAction(action db.Action) (string, string) {
	var direction, from, to, amount string
	if action.Name == "transfer" {
		// 目前交易方向和金额只从转账中获取
		if v, ok := action.Data["from"]; ok {
			from = cast.ToString(v)
		}
		if v, ok := action.Data["to"]; ok {
			to = cast.ToString(v)
		}
		if v, ok := action.Data["amount"]; ok {
			amount = cast.ToString(v)
		}
	}

	if len(from) != 0 {
		direction = from + "->" + to
	} else {
		direction = action.Account
	}
	return direction, amount
}
