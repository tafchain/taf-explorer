package taf_server

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"tafexplorer/framework/cache"
)

//========================链的基本信息
type ChainInfo struct {
	ServerVersion            string `json:"server_version"`
	ChainID                  string `json:"chain_id"`
	HeadBlockNum             int64  `json:"head_block_num"`
	LastIrreversibleBlockNum int64  `json:"last_irreversible_block_num"`
	LastIrreversibleBlockID  string `json:"last_irreversible_block_id"`
	HeadBlockID              string `json:"head_block_id"`
	HeadBlockTime            string `json:"head_block_time"`
	HeadBlockProducer        string `json:"head_block_producer"`
	VirtualBlockCpuLimit     int64  `json:"virtual_block_net_limit"`
	BlockCpuLimit            int64  `json:"block_cpu_limit"`
	BlockNetLimit            int64  `json:"block_net_limit"`
	ServerVersionString      string `json:"server_version_string"`
	ForkDbHeadBlockNum       int64  `json:"fork_db_head_block_num"`
	ForkDbHeadBlockID        string `json:"fork_db_head_block_id"`
	ServerFullVersionString  string `json:"server_full_version_string"`
}

func TafChainInfo() error {
	resp, err := Client().Post(chainApi, ContentType, nil)
	if err != nil || resp == nil {
		logs.Warn("post chain err or resp is nil", err)
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	chainInfo := &ChainInfo{}
	err = json.Unmarshal(body, chainInfo)
	if err != nil {
		logs.Warn("unmarshal respBody err", err)
		return err
	}
	// 缓存当前区块最大高度
	cache.Set(cache.HeadBlockNum, chainInfo.HeadBlockNum, 0)
	return nil
}
