package db

import (
	"github.com/astaxie/beego"
	"github.com/globalsign/mgo"
)

var session *mgo.Session
var dataBase *mgo.Database

const (
	baseDbName   = "taf_explorer"
	BlockName    = "block"         // 区块信息表
	QueriedBlock = "queried_block" // 已查询的区块
	Transactions = "transactions"  // 交易信息表
	Contracts    = "contracts"     // 合约表
	Accounts     = "accounts"      // 账户表
)

func InitMgo() error {
	username := beego.AppConfig.String("username")
	password := beego.AppConfig.String("password")
	dbAuth, _ := beego.AppConfig.Bool("dbAuth")
	url := beego.AppConfig.String("mongodb")
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	session = session.Clone()
	dataBase = session.DB(baseDbName)
	if dbAuth {
		err := dataBase.Login(username, password)
		if err != nil {
			return err
		}
	}
	return nil
}

// 会话
func Session() *mgo.Session {
	return session.Clone()
}

func DB() *mgo.Database {
	return dataBase
}

func Close() {
	if session == nil {
		return
	}
	session.Clone()
	session = nil
	dataBase = nil
}
