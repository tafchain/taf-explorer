package taf_server

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo/bson"
	"github.com/spf13/cast"
	"io/ioutil"
	"tafexplorer/db"
	"tafexplorer/framework/cache"
	"time"
)

// taf-server api 请求返回值
// 区块信息

type BlockInfoResp struct {
	Id                string       `json:"id"`
	Num               int64        `json:"block_num"`
	Timestamp         string       `json:"timestamp"`
	Producer          string       `json:"producer"`
	Confirmed         int          `json:"confirmed"`
	Previous          string       `json:"previous"`
	TransactionMroot  string       `json:"transaction_mroot"`
	ActionMroot       string       `json:"action_mroot"`
	ScheduleVersion   int          `json:"schedule_version"`
	ProducerSignature string       `json:"producer_signature"`
	RefBlockPrefix    int64        `json:"ref_block_prefix"`
	NewProducers      NewProducers `json:"new_producers"`
	Transactions      []Txs        `json:"transactions"`
}

type NewProducers struct {
	Version int       `json:"version"`
	Detail  PdrDetail `json:"detail"`
}

type PdrDetail struct {
	ProducerName    string `json:"producer_name"`
	BlockSigningKey string `json:"block_signing_key"`
}

//======================交易信息
type Txs struct {
	Status        string `json:"status"`
	CpuUsageUs    int64  `json:"cpu_usage_us"`
	NetUsageWords int64  `json:"net_usage_words"`
	TRX           Trx    `json:"trx"`
}
type Trx struct {
	ID                    string   `json:"id"`
	Signatures            []string `json:"signatures"`
	Compression           string   `json:"compression"`
	PackedContextFreeData string   `json:"packed_context_free_data"`
	ContextFreeData       []byte   `json:"context_free_data"` //json wei []
	PackedTrx             string   `json:"packed_trx"`
	Transaction           TX       `json:"transaction"`
}
type TX struct {
	Expiration         string   `json:"expiration"`
	RefBlockNum        int64    `json:"ref_block_num"`
	RefBlockPrefix     int64    `json:"ref_block_prefix"`
	MaxNetUsageWords   int64    `json:"max_net_usage_words"`
	MaxCpuUsageMs      int64    `json:"max_cpu_usage_ms"`
	DelaySec           int64    `json:"delay_sec"`
	ContextFreeActions []byte   `json:"context_free_actions"`
	Actions            []Action `json:"actions,omitempty"`
}
type Action struct {
	Account       string      `json:"account"`
	Name          string      `json:"name"`
	Authorization []Auth      `json:"authorization"`
	Data          interface{} `json:"data,omitempty"`
	HexData       string      `json:"hex_data"`
}

type Auth struct {
	Actor      string `json:"actor"`
	Permission string `json:"permission"`
}
type DataInfo struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
	Memo     string `json:"memo"`
}

//---------------------------------------------------------------

type QueriedBlock struct {
	BlockNum int64  `json:"block_num" bson:"block_num"`
	BlockID  string `json:"block_id" bson:"block_id"`
}

type blockParam struct {
	BlockNumOrId string `json:"block_num_or_id"`
}

// 请求taf-server获取区块信息
func TafBlockInfo() error {
	// 已处理id
	queriedBlockNum, _ := cache.Get(cache.QueriedMaxBlock)
	// 当前最大高度
	headBlockNum, ok := cache.Get(cache.HeadBlockNum)
	if !ok {
		return errors.New("not get headBlockNum cache")
	}
	//testNumOrId := "1514990"
	queriedBlockNumInt := cast.ToInt64(queriedBlockNum)
	headBlockNumInt := cast.ToInt64(headBlockNum)
	if queriedBlockNumInt >= headBlockNumInt && queriedBlockNumInt > 0 {
		time.Sleep(1 * time.Second)
		logs.Warn("block query over")
		return nil
	}

	queriedBlockNumInt++
	// 测试
	//param := &blockParam{BlockNumOrId: testNumOrId}

	numOrId := cast.ToString(queriedBlockNumInt)
	param := &blockParam{BlockNumOrId: numOrId}
	str, _ := json.Marshal(param)
	resp, err := Client().Post(blockApi, ContentType, bytes.NewBuffer(str))
	if err != nil || resp == nil {
		logs.Warn("req block err: ", err)
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	block := &BlockInfoResp{}
	err = json.Unmarshal(body, block)
	if err != nil {
		logs.Warn("block %s unmarshal block resp err %v: ", numOrId, err)
		return err
	}

	// 区块
	err = saveBlock(block, body)
	if err != nil {
		logs.Warn("save block err: ", err)
		return err
	}

	// 交易
	err = saveTransactions(block)
	if err != nil {
		logs.Warn("save transaction err: ", err)
		return err
	}

	// 账户
	err = saveAccount(block)
	if err != nil {
		logs.Warn("save account err: ", err)
		return err
	}

	// 合约
	err = saveContract(block)
	if err != nil {
		logs.Warn("save contract err: ", err)
		return err
	}

	err = insertQueriedBlock(block)
	if err != nil {
		logs.Warn("insert queried block err: ", err)
	}
	cache.Set(cache.QueriedMaxBlock, queriedBlockNumInt, 0)
	return nil
}

// block 入库
func saveBlock(block *BlockInfoResp, originData []byte) error {
	dbBlocks := []interface{}{}
	dbBlock := db.Block{
		Num:               block.Num,
		Id:                block.Id,
		Timestamp:         block.Timestamp,
		Producer:          block.Producer,
		Confirmed:         block.Confirmed,
		Previous:          block.Previous,
		TransactionMroot:  block.TransactionMroot,
		ActionMroot:       block.ActionMroot,
		ScheduleVersion:   block.ScheduleVersion,
		ProducerSignature: block.ProducerSignature,
		TrxCount:          len(block.Transactions), // 区块交易总数
		OriginData:        originData,              // 原始数据
	}
	dbBlock.NewProducers = db.NewProducers{
		Version: block.NewProducers.Version,
		Detail: db.PdrDetail{
			ProducerName:    block.NewProducers.Detail.ProducerName,
			BlockSigningKey: block.NewProducers.Detail.BlockSigningKey,
		},
	}
	//_, err := db.DB().C(db.BlockName).Upsert(bson.M{"id": block.Id}, dbBlock)
	dbBlocks = append(dbBlocks, dbBlock)
	err := db.DB().C(db.BlockName).Insert(dbBlocks...)
	if err != nil {
		return err
	}
	return nil
}

//
func saveTransactions(block *BlockInfoResp) error {
	if len(block.Transactions) == 0 {
		return nil
	}

	dbTransactions := []interface{}{}

	for _, trans := range block.Transactions {
		// 记录这个交易过程中出现的账户信息
		accountMap := make(map[string]struct{})
		dbTransaction := db.Transaction{
			Status:        trans.Status,
			CpuUsageUs:    trans.CpuUsageUs,
			NetUsageWords: trans.NetUsageWords,
			TRX:           db.Trx{},
			BlockNum:      block.Num,
			BlockId:       block.Id,
			Timestamp:     block.Timestamp,
		}

		// 时间戳
		t, err := time.Parse("2006-01-02T15:04:05", block.Timestamp)
		if err != nil {
			return err
		}
		dbTransaction.TimeUnix = t.Unix()

		dbTransaction.TRX.ID = trans.TRX.ID
		dbTransaction.TRX.Signatures = trans.TRX.Signatures
		dbTransaction.TRX.Compression = trans.TRX.Compression
		dbTransaction.TRX.ContextFreeData = trans.TRX.ContextFreeData
		dbTransaction.TRX.PackedTrx = trans.TRX.PackedTrx

		//
		dbTransaction.TRX.Transaction.Expiration = trans.TRX.Transaction.Expiration
		dbTransaction.TRX.Transaction.RefBlockNum = trans.TRX.Transaction.RefBlockNum
		dbTransaction.TRX.Transaction.RefBlockPrefix = trans.TRX.Transaction.RefBlockPrefix
		dbTransaction.TRX.Transaction.MaxNetUsageWords = trans.TRX.Transaction.MaxNetUsageWords
		dbTransaction.TRX.Transaction.MaxCpuUsageMs = trans.TRX.Transaction.MaxCpuUsageMs
		dbTransaction.TRX.Transaction.ContextFreeActions = trans.TRX.Transaction.ContextFreeActions
		dbTransaction.TRX.Transaction.DelaySec = trans.TRX.Transaction.DelaySec

		// actions
		dbTransaction.TRX.Transaction.Actions = make([]db.Action, 0)
		for _, action := range trans.TRX.Transaction.Actions {
			oneAction := db.Action{}
			oneAction.Account = action.Account
			oneAction.Name = action.Name
			oneAction.HexData = action.HexData
			oneAction.Data = cast.ToStringMap(action.Data)
			// 账户名 合约名
			accountMap[action.Account] = struct{}{}
			for _, auth := range action.Authorization {
				oneAuth := db.Auth{
					Actor:      auth.Actor,
					Permission: auth.Permission,
				}
				oneAction.Authorization = append(oneAction.Authorization, oneAuth)
			}
			dbTransaction.TRX.Transaction.Actions = append(dbTransaction.TRX.Transaction.Actions, oneAction)
		}
		for key, _ := range accountMap {
			dbTransaction.Accounts = append(dbTransaction.Accounts, key)
		}
		dbTransactions = append(dbTransactions, dbTransaction)
		//_, err = db.DB().C(db.Transactions).Upsert(bson.M{"trx.id": dbTransaction.TRX.ID}, dbTransaction)
		//err = db.DB().C(db.Transactions).Insert([]{}{dbTransaction})
		//if err != nil {
		//	return err
		//}
	}
	err := db.DB().C(db.Transactions).Insert(dbTransactions...)
	if err != nil {
		return err
	}
	return nil
}

// 账户
func saveAccount(block *BlockInfoResp) error {
	for _, trans := range block.Transactions {
		for _, action := range trans.TRX.Transaction.Actions {
			if action.Name == TfNewAccount {
				account := &db.Account{
					Name:       "",
					PublicKey:  "",
					Creator:    "",
					CreateTime: block.Timestamp,
				}

				// 时间戳
				t, err := time.Parse("2006-01-02T15:04:05", block.Timestamp)
				if err != nil {
					return err
				}
				account.TimeUnix = t.Unix()

				data := cast.ToStringMap(action.Data)
				if name, ok := data["name"]; ok {
					account.Name = cast.ToString(name)
				}
				if creator, ok := data["creator"]; ok {
					account.Creator = cast.ToString(creator)
				}
				if owner, ok := data["owner"]; ok {
					ownerMap := owner.(map[string]interface{})
					keys := ownerMap["keys"].([]interface{})
					if len(keys) > 0 {
						account.PublicKey = cast.ToString(keys[0].(map[string]interface{})["key"])
					}
				}
				_, err = db.DB().C(db.Accounts).Upsert(bson.M{"name": account.Name}, account)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// 统计合约交易信息
func saveContract(block *BlockInfoResp) error {
	// 记录出现的合约(账户)名
	updateAccounts := []string{}
	for _, trans := range block.Transactions {
		for _, action := range trans.TRX.Transaction.Actions {
			// 账户 合约名
			account := action.Account
			// 交易动作
			actionName := action.Name

			// 是否是部署合约的动作
			if actionName == TfSetCode || actionName == TfSetAbi {
				if !isContract(account) {
					db.RedisSession().HSet(db.Contracts, account, true)
					//// 更新账户信息
					//db.DB().C(db.Accounts).Update(bson.M{"name": account}, bson.M{"$set": bson.M{"is_contract": true}})
					// 不在的话从数据库加载
				}
			}

			// 账户是否已部署合约
			if !isContract(account) {
				continue
			}

			// 账户的交易次数加1，记录交易的动作
			db.RedisSession().Incr(account)
			db.RedisSession().SAdd(accountActionsKey(account), actionName)
			updateAccounts = append(updateAccounts, account)
		}
	}
	// 更新合约信息
	for _, account := range updateAccounts {
		contract := db.Contract{
			Name:       account,
			TransCount: 0,
			Actions:    nil,
		}
		contract.TransCount, _ = db.RedisSession().Get(account).Int64()
		contract.Actions = db.RedisSession().SMembers(accountActionsKey(account)).Val()
		_, err := db.DB().C(db.Contracts).Upsert(bson.M{"name": contract.Name}, contract)
		if err != nil {
			return err
		}
	}
	return nil
}

func insertQueriedBlock(block *BlockInfoResp) error {
	// 记录查询的id
	c := &QueriedBlock{}
	c.BlockID = block.Id
	c.BlockNum = block.Num
	_, err := db.DB().C(db.QueriedBlock).Upsert(bson.M{"block_num": c.BlockNum}, c)
	if err != nil {
		return err
	}
	return nil
}

// 生成存储账户(合约)交易动作的key
func accountActionsKey(account string) string {
	return "ac" + "_" + account
}

// 判断账户是否存在和是否部署合约
func isContract(account string) bool {
	if db.RedisSession().HGet(db.Contracts, account) == nil {
		// 从db加载
		dbResult := []db.Contract{}
		err := db.DB().C(db.Contracts).Find(bson.M{"name": account}).All(&dbResult)
		if err != nil || len(dbResult) == 0 {
			db.RedisSession().HSet(db.Contracts, account, false)
		}
		contractInfo := dbResult[0]
		db.RedisSession().HSet(db.Contracts, account, true)
		// 交易次数
		db.RedisSession().Set(account, contractInfo.TransCount, 0)
		// 交易动作
		db.RedisSession().SAdd(accountActionsKey(account), contractInfo.Actions)
	}

	return cast.ToBool(db.RedisSession().HGet(db.Contracts, account).Val())
}
