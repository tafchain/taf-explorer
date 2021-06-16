package cache

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

var c *gocache.Cache

const (
	HeadBlockNum    = "headBlockNum"    // 区块最大高度
	QueriedMaxBlock = "queriedMaxBlock" // 已查询的最大区块
)

func init() {
	c = gocache.New(15*time.Minute, 30*time.Minute)
}

func Set(key string, value interface{}, d time.Duration) {
	c.Set(key, value, d)
}

func Get(key string) (interface{}, bool) {
	return c.Get(key)
}

// 定时保存到文件
func Save() {
	err := c.SaveFile("./save")
	if err != nil {
	}
}