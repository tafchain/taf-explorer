package taf_server

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/globalsign/mgo/bson"
	"net/http"
	"tafexplorer/db"
	"tafexplorer/framework/cache"
	"tafexplorer/framework/loop_task"
	"time"
)

const (
	ContentType  = "application/json"
	TfSetCode    = "tfsetcode" // 部署合约的动作
	TfSetAbi     = "tfsetabi"  // 部署合约的动作
	TfNewAccount = "tfnew"     // 创建新账号
	TafSys       = "tafsys"    // 内置系统账号 默认部署合约
)

var client *http.Client

// taf-server地址
var tafServer, chainApi, blockApi, accountApi, balanceApi, countApi, currencyApi, minersApi, votersApi string

func InitTafServer() {
	tafServer = beego.AppConfig.String("chain_server")
	chainApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("chain_info_api"))
	blockApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("block_info_api"))
	accountApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("deal_account_info_api"))
	balanceApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("get_balance_api"))
	countApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("count_api"))
	currencyApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("currency_api"))
	minersApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("miners_api"))
	votersApi = fmt.Sprintf("%s%s", tafServer, beego.AppConfig.String("voters_api"))

	client = &http.Client{Timeout: 3 * time.Second}
	initCache()
	initRedis()
	// 循环查询
	//go AllRun(TafBlockInfo)
	//go AllRun(TafChainInfo)
	//go AllRun(InsertContract)
	chainTask := loop_task.NewBaseLoopTask(10, 1, TafChainInfo)
	chainTask.Run()
	blockTask := loop_task.NewBaseLoopTask(10, 0, TafBlockInfo)
	blockTask.Run()
}

func Client() *http.Client {
	return client
}

// 初始化cache
func initCache() {
	count, err := db.DB().C(db.QueriedBlock).Find(bson.M{}).Count()
	if err != nil {
		logs.Info("init cache err: ", err)
	}

	cache.Set(cache.QueriedMaxBlock, count, 0)
}

func initRedis() {
	// 内置系统账号，默认部署合约
	db.RedisSession().HSet(db.Contracts, TafSys, true)
}

func AllRun(f func()) {
	for {
		safeRunFunc(f)
		time.Sleep(1 * time.Second)
	}
}

func safeRunFunc(f func()) {
	defer func() {
		err := recover()
		if err != nil {
			logs.Warn("AllRunErr: ", err)
		}
	}()
	f()
}
