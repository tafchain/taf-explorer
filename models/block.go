package models

import (
	"encoding/json"
	"errors"
	"github.com/globalsign/mgo/bson"
	"tafexplorer/api"
	"tafexplorer/db"
)

// 区块列表
func BlockList(req *api.BlockListReq) (*api.BlockListResp, error) {
	resp := &api.BlockListResp{}
	skip := (req.PageIndex - 1) * req.PageSize
	result := make([]db.Block, 0)
	err := db.DB().C(db.BlockName).Find(bson.M{}).Sort("-block_num").Skip(skip).Limit(req.PageSize).All(&result)
	if err != nil {
		return resp, err
	}
	for _, one := range result {
		block := &api.Block{
			BlockNum:  one.Num,
			Packer:    one.Producer,
			Verier:    one.Producer,
			TrxCount:  one.TrxCount,
			Timestamp: one.Timestamp,
		}
		resp.Data = append(resp.Data, block)
	}
	count, _ := db.DB().C(db.BlockName).Count()
	resp.Total = count
	resp.PageSize = req.PageSize
	resp.PageIndex = req.PageIndex
	resp.PageNum = count / req.PageSize
	return resp, nil
}

func BlockInfo(req *api.BlockInfoReq) (*api.BlockInfoResp, error) {
	resp := &api.BlockInfoResp{}
	result := []db.Block{}
	err := db.DB().C(db.BlockName).Find(bson.M{"block_num": req.BlockNum}).All(&result)
	if err != nil {
		return resp, err
	}
	if len(result) == 0 {
		return resp, errors.New("not find block")
	}
	block := result[0]
	resp.Packer = block.Producer
	resp.PackTime = block.Timestamp
	resp.Verier = block.Producer
	resp.VerifyTime = block.Timestamp
	resp.HexData = block.Id

	// 解析原始数据
	originData := &map[string]interface{}{}
	err = json.Unmarshal(block.OriginData, originData)

	bytes, _ := json.Marshal(originData)

	resp.OriginData = string(bytes)
	return resp, nil
}
