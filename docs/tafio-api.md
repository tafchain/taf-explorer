### TAFChain API
1. url: http://192.168.0.234:8888
2. Content-type: application/json
3. response 200

<br/>

* taf_getInfo
> 获取链的基本信息

```
POST
/taf/taf_getInfo

body:
{
}

response:
{
    "server_version": "String",
    "chain_id": "String",
    "head_block_num": Number,
    "last_irreversible_block_num": Number,
    "last_irreversible_block_id": "String",
    "head_block_id": "String",
    "head_block_time": "String",
    "head_block_producer": "String",
    "virtual_block_cpu_limit": Number,
    "virtual_block_net_limit": Number,
    "block_cpu_limit": Number,
    "block_net_limit": Number,
    "server_version_string": "String",
    "fork_db_head_block_num": Number,
    "fork_db_head_block_id": "String",
    "server_full_version_string": "String"
}
```


* taf_getAccount
> 获取账户信息

```
POST
/taf/taf_getAccount

body:
{
    "account_name": "String"
}

response:
{
    "account_name": "String",
    "head_block_num": Number,
    "head_block_time": "String",
    "privileged": Boolean,
    "last_code_update": "String",
    "created": "String",
    "ram_quota": Number,
    "net_weight": "String",
    "cpu_weight": "String",
    "net_limit": {
        "used": Number,
        "available": "String",
        "max": "String"
    },
    "cpu_limit": {
        "used": Number,
        "available": "String",
        "max": "String"
    },
    "ram_usage": Number,
    "permissions": [
        {
            "perm_name": "String",
            "parent": "String",
            "required_auth": {
                "threshold": Number,
                "keys": [
                    {
                        "key": "String",
                        "weight": Number
                    }
                ],
                "accounts": [],
                "waits": []
            }
        },
        {
            "perm_name": "String",
            "parent": "String",
            "required_auth": {
                "threshold": Number,
                "keys": [
                    {
                        "key": "String",
                        "weight": Number
                    }
                ],
                "accounts": [],
                "waits": []
            }
        }
    ],
    "total_resources": {
        "owner": "String",
        "net_weight": "String",
        "cpu_weight": "String",
        "ram_bytes": Number
    },
    "self_delegated_bandwidth": {
        "from": "String",
        "to": "String",
        "net_weight": "String",
        "cpu_weight": "String"
    },
    "refund_request": {
        "owner": "String",
        "request_time": "String",
        "net_amount": "String",
        "cpu_amount": "String"
    },
    "voter_info": {
        "owner": "String",
        "proxy": "String",
        "producers": [
            "String",
            "String"
        ],
        "staked": "String",
        "last_vote_weight": "String",
        "proxied_vote_weight": "String",
        "is_proxy": Number,
        "flags1": Number,
        "reserved2": Number,
        "reserved3": "String"
    },
    "rex_info": Null
}
```


* taf_getBlock
> 获取区块信息

```
POST
/taf/taf_getBlock

body:
{
    "block_num_or_id": "String"
}

response:
{
    "timestamp": "String",
    "producer": "String",
    "confirmed": Number,
    "previous": "String",
    "transaction_mroot": "String",
    "action_mroot": "String",
    "schedule_version": Number,
    "new_producers": {
        "version": Number,
        "producers": [
            {
                "producer_name": "String",
                "block_signing_key": "String"
            }
        ]
    },
    "producer_signature": "String",
    "transactions": [
        {
            "status": "String",
            "cpu_usage_us": Number,
            "net_usage_words": Number,
            "trx": {
                "id": "String",
                "signatures": [
                    "String"
                ],
                "compression": "String",
                "packed_context_free_data": "String",
                "context_free_data": [],
                "packed_trx": "String",
                "transaction": {
                    "expiration": "String",
                    "ref_block_num": Number,
                    "ref_block_prefix": Number,
                    "max_net_usage_words": Number,
                    "max_cpu_usage_ms": Number,
                    "delay_sec": Number,
                    "context_free_actions": [],
                    "actions": [
                        {
                            "account": "String",
                            "name": "String",
                            "authorization": [
                                {
                                    "actor": "String",
                                    "permission": "String"
                                }
                            ],
                            "data": {
                                "from": "String",
                                "to": "String",
                                "quantity": "String",
                                "memo": "String"
                            },
                            "hex_data": "String"
                        }
                    ]
                }
            }
        }
    ],
    "id": "String",
    "block_num": Number,
    "ref_block_prefix": Number
}
```


* taf_getBalance
> 获取账户余额

```
POST
/taf/taf_getBalance

body:
{
    "code": "String",                               // default: "tafio.token"
    "account": "String",
    "symbol": "String"                              // e.g. "SYS"
}

response:
[
  "String"
]
```


* taf_getMiners
> 获取生产者列表

```
POST
/taf/taf_getMiners

body:
{
    "limit": Number,                              // e.g. 50
    "lower_bound": "String",                        // e.g. ""
    "json": Boolean                                 // true
}

response:
{
    "rows": [
        {
            "owner": "String",
            "total_votes": "String",
            "producer_key": "String",
            "is_active": Number,
            "url": "String",
            "unpaid_blocks": Number,
            "last_claim_time": "String",
            "location": Number,
            "producer_authority": [
                "String",
                {
                    "threshold": Number,
                    "keys": [
                        {
                            "key": "String",
                            "weight": Number
                        }
                    ]
                }
            ]
        },
        ...
    ],
    "total_producer_vote_weight": "String",
    "more": "String"
}
```


* get_transaction
> 根据交易ID获取交易信息

```
POST
/history/get_transaction

body:
{
    "id": "String",             // e.g. "9b99c4f97e90d9863874b23f3094de86f37f61e1bea965983658212912f352d8"
    "block_num_hint": Number    // e.g. 149146
}

response:
{
    "id": "String",
    "trx": {
        "receipt": {
            "status": "String",
            "cpu_usage_us": Number,
            "net_usage_words": Number,
            "trx": [
                1,
                {
                    "signatures": [
                        "String"
                    ],
                    "compression": "String",
                    "packed_context_free_data": "",
                    "packed_trx": "String"
                }
            ]
        },
        "trx": {
            "expiration": "String",
            "ref_block_num": Number,
            "ref_block_prefix": Number,
            "max_net_usage_words": Number,
            "max_cpu_usage_ms": Number,
            "delay_sec": Number,
            "context_free_actions": [],
            "actions": [
                {
                    "account": "String",
                    "name": "String",
                    "authorization": [
                        {
                            "actor": "String",
                            "permission": "String"
                        }
                    ],
                    "data": {
                        "from": "String",
                        "to": "String",
                        "quantity": "String",
                        "memo": ""
                    },
                    "hex_data": "String"
                }
            ],
            "transaction_extensions": [],
            "signatures": [
                "String"
            ],
            "context_free_data": []
        }
    },
    "block_time": "String",
    "block_num": Number,
    "last_irreversible_block": Number,
    "traces": []
}
```


* /taf/taf_getTableRows
> 获取质押总量

```
POST
/taf/taf_getTableRows

body:
{
    "json": true,
    "code": "tafio.token",
    "scope": "tafio.stake",
    "table": "accounts",
    "table_key": "",
    "lower_bound": "",
    "upper_bound": "",
    "limit": 10,
    "key_type": "",
    "index_position": "",
    "encode_type": "dec",
    "reverse": false,
    "show_payer": false
}

response:
{
    "rows": [
        {
            "balance": "String"
        }
    ],
    "more": Boolean,
    "next_key": ""
}
```


* /taf/taf_getTableRows
> 获取投票人数与数量

```
POST
/taf/taf_getTableRows

body:
{
    "json": true,
    "code": "tafio",
    "scope": "tafio",
    "table": "voters",
    "table_key": "",
    "lower_bound": "",
    "upper_bound": "",
    "limit": 10,
    "key_type": "",
    "index_position": "",
    "encode_type": "dec",
    "reverse": false,
    "show_payer": false
}

response:
{
    "rows": [
        {
            "owner": "String",
            "proxy": "",
            "producers": [
                "String",
                "String",
                "String",
                "String",
                "String",
                "String",
                "String"
            ],
            "staked": "String",
            "last_vote_weight": "String",
            "proxied_vote_weight": "String",
            "is_proxy": Number,
            "flags1": Number,
            "reserved2": Number,
            "reserved3": "String"
        },
        ...
    ],
    "more": Boolean,
    "next_key": ""
}
```