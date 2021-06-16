package db

// 区块
type Block struct {
	Id                string       `json:"id" bson:"id"`
	Num               int64        `json:"block_num" bson:"block_num"`
	Timestamp         string       `json:"timestamp" bson:"timestamp"`
	Producer          string       `json:"producer" bson:"producer"`
	Confirmed         int          `json:"confirmed" bson:"confirmed"`
	Previous          string       `json:"previous" bson:"previous"`
	TransactionMroot  string       `json:"transaction_mroot" bson:"transaction_mroot"`
	ActionMroot       string       `json:"action_mroot" bson:"action_mroot"`
	ScheduleVersion   int          `json:"schedule_version" bson:"schedule_version"`
	ProducerSignature string       `json:"producer_signature" bson:"producer_signature"`
	RefBlockPrefix    int64        `json:"ref_block_prefix" bson:"ref_block_prefix"`
	NewProducers      NewProducers `json:"new_producers" bson:"new_producers"`
	TrxCount          int          `json:"trx_count" bson:"trx_count"`
	OriginData        []byte       `json:"origin_data" bson:"origin_data"`
}

type NewProducers struct {
	Version int       `json:"version" bson:"version"`
	Detail  PdrDetail `json:"detail" bson:"detail"`
}

type PdrDetail struct {
	ProducerName    string `json:"producer_name" bson:"producer_name"`
	BlockSigningKey string `json:"block_signing_key" bson:"block_signing_key"`
}

// 交易
type Transaction struct {
	Status        string   `json:"status" bson:"status"`
	CpuUsageUs    int64    `json:"cpu_usage_us" bson:"cpu_usage_us"`
	NetUsageWords int64    `json:"net_usage_words" bson:"net_usage_words"`
	TRX           Trx      `json:"trx" bson:"trx"`
	BlockNum      int64    `json:"block_num" bson:"block_num"` // block_num
	BlockId       string   `json:"block_id" bson:"block_id"`   // block_id
	Timestamp     string   `json:"timestamp" bson:"timestamp"` // 交易时间（区块时间）
	Accounts      []string `json:"accounts" bson:"accounts"`   // 交易过程中出现的账户名(合约名）
	TimeUnix      int64    `json:"time_unix" bson:"time_unix"` // 交易时间戳

}
type Trx struct {
	ID                    string   `json:"id" bson:"id"`
	Signatures            []string `json:"signatures" bson:"signatures"`
	Compression           string   `json:"compression" bson:"compression"`
	PackedContextFreeData string   `json:"packed_context_free_data" bson:"packed_context_free_data"`
	ContextFreeData       []byte   `json:"context_free_data" bson:"context_free_data"` //json wei []
	PackedTrx             string   `json:"packed_trx" bson:"packed_trx"`
	Transaction           TX       `json:"transaction" bson:"transaction"`
}
type TX struct {
	Expiration         string   `json:"expiration" bson:"expiration"`
	RefBlockNum        int64    `json:"ref_block_num" bson:"ref_block_num"`
	RefBlockPrefix     int64    `json:"ref_block_prefix" bson:"ref_block_prefix"`
	MaxNetUsageWords   int64    `json:"max_net_usage_words" bson:"max_net_usage_words"`
	MaxCpuUsageMs      int64    `json:"max_cpu_usage_ms" bson:"max_cpu_usage_ms"`
	DelaySec           int64    `json:"delay_sec" bson:"delay_sec"`
	ContextFreeActions []byte   `json:"context_free_actions" bson:"context_free_actions"`
	Actions            []Action `json:"max_cpu_usage_ms" bson:"actions"`
}
type Action struct {
	Account       string                 `json:"account" bson:"account"`
	Name          string                 `json:"name" bson:"name"`
	Authorization []Auth                 `json:"authorization" bson:"authorization"`
	Data          map[string]interface{} `json:"data" bson:"data"`
	HexData       string                 `json:"hex_data" bson:"hex_data"`
}

type Auth struct {
	Actor      string `json:"actor" bson:"actor"`
	Permission string `json:"permission" bson:"permission"`
}
type DataInfo struct {
	From     string `json:"from" bson:"from"`
	To       string `json:"to" bson:"to"`
	Quantity string `json:"quantity" bson:"quantity"`
	Memo     string `json:"memo" bson:"memo"`
}

// 合约
type Contract struct {
	Name       string   `json:"name" bson:"name"`
	TransCount int64    `json:"trans_count" bson:"trans_count"` // 交易数量 出现一次加1
	Actions    []string `json:"actions" bson:"actions"`
}

// 账户
type Account struct {
	Name       string `json:"name" bson:"name"`               // 用户名
	PublicKey  string `json:"public_key" bson:"public_key"`   // 公钥
	Creator    string `json:"creator" bson:"creator"`         // 创建者
	CreateTime string `json:"create_time" bson:"create_time"` // 创建时间
	TimeUnix   int64  `json:"time_unix" bson:"time_unix"`     // 交易时间戳
}

// 当前认为合约名就是账户名
type AccountAndContract struct {
	Name       string   `json:"name" bson:"name"`               // 账户名 合约名
	TransCount int64    `json:"trans_count" bson:"trans_count"` // 交易数量 action出现一次加-
	Creator    string   `json:"creator" bson:"creator"`         // 创建者
	CreateTime string   `json:"create_time" bson:"create_time"` // 创建时间
	Actions    []string `json:"actions" bson:"actions"`         // 已出现的交易动作
}
