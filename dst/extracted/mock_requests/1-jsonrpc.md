### eth_blockNumber.json echo -e 
```json
{
  "name": "eth_blockNumber",
  "summary": "Returns the number of most recent block.",
  "params": [],
  "result": {
    "name": "Block number",
    "schema": {
      "title": "hex encoded unsigned integer",
      "type": "string",
      "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
    }
  }
}
```


### eth_call.json echo -e 
```json
{
  "name": "eth_call",
  "summary": "Executes a new message call immediately without creating a transaction on the block chain.",
  "params": [
    {
      "name": "Transaction",
      "required": true,
      "schema": {
        "title": "Transaction object with sender",
        "type": "object",
        "required": [
          "from"
        ],
        "oneOf": [
          {
            "type": "object",
            "title": "EIP-1559 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "maxFeePerGas",
              "maxPriorityFeePerGas",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "maxPriorityFeePerGas": {
                "title": "max priority fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
              },
              "maxFeePerGas": {
                "title": "max fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "EIP-2930 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "Legacy transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          }
        ],
        "properties": {
          "from": {
            "title": "from",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          }
        }
      }
    },
    {
      "name": "Block",
      "required": false,
      "schema": {
        "title": "Block number or tag",
        "oneOf": [
          {
            "title": "Block number",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          {
            "title": "Block tag",
            "type": "string",
            "enum": [
              "earliest",
              "finalized",
              "safe",
              "latest",
              "pending"
            ],
            "description": "`earliest`: The lowest numbered block the client has available; `finalized`: The most recent crypto-economically secure block, cannot be re-orged outside of manual intervention driven by community coordination; `safe`: The most recent block that is safe from re-orgs under honest majority and certain synchronicity assumptions; `latest`: The most recent block in the canonical chain observed by the client, this block may be re-orged out of the canonical chain even under healthy/normal conditions; `pending`: A sample next block built by the client on top of `latest` and containing the set of transactions usually taken from local mempool. Before the merge transition is finalized, any call querying for `finalized` or `safe` block MUST be responded to with `-39001: Unknown block` error"
          }
        ]
      }
    }
  ],
  "result": {
    "name": "Return data",
    "schema": {
      "title": "hex encoded bytes",
      "type": "string",
      "pattern": "^0x[0-9a-f]*$"
    }
  }
}
```


### eth_chainId.json echo -e 
```json
{
  "name": "eth_chainId",
  "summary": "Returns the chain ID of the current network.",
  "params": [],
  "result": {
    "name": "Chain ID",
    "schema": {
      "title": "hex encoded unsigned integer",
      "type": "string",
      "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
    }
  }
}
```


### eth_estimateGas.json echo -e 
```json
{
  "name": "eth_estimateGas",
  "summary": "Generates and returns an estimate of how much gas is necessary to allow the transaction to complete.",
  "params": [
    {
      "name": "Transaction",
      "required": true,
      "schema": {
        "title": "Transaction object with sender",
        "type": "object",
        "required": [
          "from"
        ],
        "oneOf": [
          {
            "type": "object",
            "title": "EIP-1559 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "maxFeePerGas",
              "maxPriorityFeePerGas",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "maxPriorityFeePerGas": {
                "title": "max priority fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
              },
              "maxFeePerGas": {
                "title": "max fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "EIP-2930 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "Legacy transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          }
        ],
        "properties": {
          "from": {
            "title": "from",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          }
        }
      }
    },
    {
      "name": "Block",
      "required": false,
      "schema": {
        "title": "Block number or tag",
        "oneOf": [
          {
            "title": "Block number",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          {
            "title": "Block tag",
            "type": "string",
            "enum": [
              "earliest",
              "finalized",
              "safe",
              "latest",
              "pending"
            ],
            "description": "`earliest`: The lowest numbered block the client has available; `finalized`: The most recent crypto-economically secure block, cannot be re-orged outside of manual intervention driven by community coordination; `safe`: The most recent block that is safe from re-orgs under honest majority and certain synchronicity assumptions; `latest`: The most recent block in the canonical chain observed by the client, this block may be re-orged out of the canonical chain even under healthy/normal conditions; `pending`: A sample next block built by the client on top of `latest` and containing the set of transactions usually taken from local mempool. Before the merge transition is finalized, any call querying for `finalized` or `safe` block MUST be responded to with `-39001: Unknown block` error"
          }
        ]
      }
    }
  ],
  "result": {
    "name": "Gas used",
    "schema": {
      "title": "hex encoded unsigned integer",
      "type": "string",
      "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
    }
  }
}
```


### eth_feeHistory.json echo -e 
```json
{
  "name": "eth_feeHistory",
  "summary": "Transaction fee history",
  "description": "Returns transaction base fee per gas and effective priority fee per gas for the requested/supported block range.",
  "params": [
    {
      "name": "blockCount",
      "description": "Requested range of blocks. Clients will return less than the requested range if not all blocks are available.",
      "required": true,
      "schema": {
        "title": "hex encoded unsigned integer",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      }
    },
    {
      "name": "newestBlock",
      "description": "Highest block of the requested range.",
      "required": true,
      "schema": {
        "title": "Block number or tag",
        "oneOf": [
          {
            "title": "Block number",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          {
            "title": "Block tag",
            "type": "string",
            "enum": [
              "earliest",
              "finalized",
              "safe",
              "latest",
              "pending"
            ],
            "description": "`earliest`: The lowest numbered block the client has available; `finalized`: The most recent crypto-economically secure block, cannot be re-orged outside of manual intervention driven by community coordination; `safe`: The most recent block that is safe from re-orgs under honest majority and certain synchronicity assumptions; `latest`: The most recent block in the canonical chain observed by the client, this block may be re-orged out of the canonical chain even under healthy/normal conditions; `pending`: A sample next block built by the client on top of `latest` and containing the set of transactions usually taken from local mempool. Before the merge transition is finalized, any call querying for `finalized` or `safe` block MUST be responded to with `-39001: Unknown block` error"
          }
        ]
      }
    },
    {
      "name": "rewardPercentiles",
      "description": "A monotonically increasing list of percentile values. For each block in the requested range, the transactions will be sorted in ascending order by effective tip per gas and the coresponding effective tip for the percentile will be determined, accounting for gas consumed.",
      "required": true,
      "schema": {
        "title": "rewardPercentiles",
        "type": "array",
        "items": {
          "title": "rewardPercentile",
          "description": "Floating point value between 0 and 100.",
          "type": "number"
        }
      }
    }
  ],
  "result": {
    "name": "feeHistoryResult",
    "description": "Fee history for the returned block range. This can be a subsection of the requested range if not all blocks are available.",
    "schema": {
      "title": "feeHistoryResults",
      "description": "Fee history results.",
      "type": "object",
      "required": [
        "firstBlock",
        "baseFeePerGas",
        "gasUsedRatio"
      ],
      "properties": {
        "oldestBlock": {
          "title": "oldestBlock",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
          "description": "Lowest number block of returned range."
        },
        "baseFeePerGas": {
          "title": "baseFeePerGasArray",
          "description": "An array of block base fees per gas. This includes the next block after the newest of the returned range, because this value can be derived from the newest block. Zeroes are returned for pre-EIP-1559 blocks.",
          "type": "array",
          "items": {
            "title": "hex encoded unsigned integer",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          }
        },
        "reward": {
          "title": "rewardArray",
          "description": "A two-dimensional array of effective priority fees per gas at the requested block percentiles.",
          "type": "array",
          "items": {
            "title": "rewardPercentile",
            "description": "An array of effective priority fee per gas data points from a single block. All zeroes are returned if the block is empty.",
            "type": "array",
            "items": {
              "title": "rewardPercentile",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "A given percentile sample of effective priority fees per gas from a single block in ascending order, weighted by gas used. Zeroes are returned if the block is empty."
            }
          }
        }
      }
    }
  }
}
```


### eth_gasPrice.json echo -e 
```json
{
  "name": "eth_gasPrice",
  "summary": "Returns the current price per gas in wei.",
  "params": [],
  "result": {
    "name": "Gas price",
    "schema": {
      "title": "Gas price",
      "type": "string",
      "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
    }
  }
}
```


### eth_getBlockTransactionCountByNumber.json echo -e 
```json
{
  "name": "eth_getBlockTransactionCountByNumber",
  "summary": "Returns the number of transactions in a block matching the given block number.",
  "params": [
    {
      "name": "Block",
      "schema": {
        "title": "Block number or tag",
        "oneOf": [
          {
            "title": "Block number",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          {
            "title": "Block tag",
            "type": "string",
            "enum": [
              "earliest",
              "finalized",
              "safe",
              "latest",
              "pending"
            ],
            "description": "`earliest`: The lowest numbered block the client has available; `finalized`: The most recent crypto-economically secure block, cannot be re-orged outside of manual intervention driven by community coordination; `safe`: The most recent block that is safe from re-orgs under honest majority and certain synchronicity assumptions; `latest`: The most recent block in the canonical chain observed by the client, this block may be re-orged out of the canonical chain even under healthy/normal conditions; `pending`: A sample next block built by the client on top of `latest` and containing the set of transactions usually taken from local mempool. Before the merge transition is finalized, any call querying for `finalized` or `safe` block MUST be responded to with `-39001: Unknown block` error"
          }
        ]
      }
    }
  ],
  "result": {
    "name": "Transaction count",
    "schema": {
      "title": "Transaction count",
      "type": "array",
      "items": {
        "title": "hex encoded unsigned integer",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      }
    }
  }
}
```


### eth_getFilterLogs.json echo -e 
```json
{
  "name": "eth_getFilterLogs",
  "summary": "Returns an array of all logs matching filter with given id.",
  "params": [
    {
      "name": "Filter Identifier",
      "schema": {
        "title": "hex encoded unsigned integer",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      }
    }
  ],
  "result": {
    "name": "Log objects",
    "schema": {
      "title": "Filter results",
      "oneOf": [
        {
          "title": "new block hashes",
          "type": "array",
          "items": {
            "title": "32 byte hex value",
            "type": "string",
            "pattern": "^0x[0-9a-f]{64}$"
          }
        },
        {
          "title": "new transaction hashes",
          "type": "array",
          "items": {
            "title": "32 byte hex value",
            "type": "string",
            "pattern": "^0x[0-9a-f]{64}$"
          }
        },
        {
          "title": "new logs",
          "type": "array",
          "items": {
            "title": "log",
            "type": "object",
            "required": [
              "transactionHash"
            ],
            "properties": {
              "removed": {
                "title": "removed",
                "type": "boolean"
              },
              "logIndex": {
                "title": "log index",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "transactionIndex": {
                "title": "transaction index",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "transactionHash": {
                "title": "transaction hash",
                "type": "string",
                "pattern": "^0x[0-9a-f]{64}$"
              },
              "blockHash": {
                "title": "block hash",
                "type": "string",
                "pattern": "^0x[0-9a-f]{64}$"
              },
              "blockNumber": {
                "title": "block number",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "address": {
                "title": "address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "data": {
                "title": "data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "topics": {
                "title": "topics",
                "type": "array",
                "items": {
                  "title": "32 hex encoded bytes",
                  "type": "string",
                  "pattern": "^0x[0-9a-f]{64}$"
                }
              }
            }
          }
        }
      ]
    }
  }
}
```


### eth_getTransactionByBlockHashAndIndex.json echo -e 
```json
{
  "name": "eth_getTransactionByBlockHashAndIndex",
  "summary": "Returns information about a transaction by block hash and transaction index position.",
  "params": [
    {
      "name": "Block hash",
      "required": true,
      "schema": {
        "title": "32 byte hex value",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      }
    },
    {
      "name": "Transaction index",
      "required": true,
      "schema": {
        "title": "hex encoded unsigned integer",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      }
    }
  ],
  "result": {
    "name": "Transaction information",
    "schema": {
      "type": "object",
      "title": "Transaction information",
      "required": [
        "blockHash",
        "blockNumber",
        "from",
        "hash",
        "transactionIndex"
      ],
      "oneOf": [
        {
          "title": "Signed 1559 Transaction",
          "type": "object",
          "required": [
            "accessList",
            "chainId",
            "gas",
            "input",
            "maxFeePerGas",
            "maxPriorityFeePerGas",
            "nonce",
            "r",
            "s",
            "type",
            "value",
            "yParity"
          ],
          "properties": {
            "type": {
              "title": "type",
              "type": "string",
              "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
            },
            "nonce": {
              "title": "nonce",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "to": {
              "title": "to address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "gas": {
              "title": "gas limit",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "value": {
              "title": "value",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "input": {
              "title": "input data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "maxPriorityFeePerGas": {
              "title": "max priority fee per gas",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
            },
            "maxFeePerGas": {
              "title": "max fee per gas",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
            },
            "accessList": {
              "title": "accessList",
              "type": "array",
              "description": "EIP-2930 access list",
              "items": {
                "title": "Access list entry",
                "type": "object",
                "properties": {
                  "address": {
                    "title": "hex encoded address",
                    "type": "string",
                    "pattern": "^0x[0-9,a-f,A-F]{40}$"
                  },
                  "storageKeys": {
                    "type": "array",
                    "items": {
                      "title": "32 byte hex value",
                      "type": "string",
                      "pattern": "^0x[0-9a-f]{64}$"
                    }
                  }
                }
              }
            },
            "chainId": {
              "title": "chainId",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Chain ID that this transaction is valid on."
            },
            "yParity": {
              "title": "yParity",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
            },
            "r": {
              "title": "r",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "s": {
              "title": "s",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            }
          }
        },
        {
          "title": "Signed 2930 Transaction",
          "type": "object",
          "required": [
            "accessList",
            "chainId",
            "gas",
            "gasPrice",
            "input",
            "nonce",
            "r",
            "s",
            "type",
            "value",
            "yParity"
          ],
          "properties": {
            "type": {
              "title": "type",
              "type": "string",
              "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
            },
            "nonce": {
              "title": "nonce",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "to": {
              "title": "to address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "gas": {
              "title": "gas limit",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "value": {
              "title": "value",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "input": {
              "title": "input data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "gasPrice": {
              "title": "gas price",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The gas price willing to be paid by the sender in wei"
            },
            "accessList": {
              "title": "accessList",
              "type": "array",
              "description": "EIP-2930 access list",
              "items": {
                "title": "Access list entry",
                "type": "object",
                "properties": {
                  "address": {
                    "title": "hex encoded address",
                    "type": "string",
                    "pattern": "^0x[0-9,a-f,A-F]{40}$"
                  },
                  "storageKeys": {
                    "type": "array",
                    "items": {
                      "title": "32 byte hex value",
                      "type": "string",
                      "pattern": "^0x[0-9a-f]{64}$"
                    }
                  }
                }
              }
            },
            "chainId": {
              "title": "chainId",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Chain ID that this transaction is valid on."
            },
            "yParity": {
              "title": "yParity",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
            },
            "r": {
              "title": "r",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "s": {
              "title": "s",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            }
          }
        },
        {
          "title": "Signed Legacy Transaction",
          "type": "object",
          "required": [
            "gas",
            "gasPrice",
            "input",
            "nonce",
            "r",
            "s",
            "type",
            "v",
            "value"
          ],
          "properties": {
            "type": {
              "title": "type",
              "type": "string",
              "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
            },
            "nonce": {
              "title": "nonce",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "to": {
              "title": "to address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "gas": {
              "title": "gas limit",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "value": {
              "title": "value",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "input": {
              "title": "input data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "gasPrice": {
              "title": "gas price",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The gas price willing to be paid by the sender in wei"
            },
            "chainId": {
              "title": "chainId",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Chain ID that this transaction is valid on."
            },
            "v": {
              "title": "v",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "r": {
              "title": "r",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "s": {
              "title": "s",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            }
          }
        }
      ],
      "properties": {
        "blockHash": {
          "title": "block hash",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        },
        "blockNumber": {
          "title": "block number",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
        },
        "from": {
          "title": "from address",
          "type": "string",
          "pattern": "^0x[0-9,a-f,A-F]{40}$"
        },
        "hash": {
          "title": "transaction hash",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        },
        "transactionIndex": {
          "title": "transaction index",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
        }
      }
    }
  }
}
```


### eth_getTransactionByHash.json echo -e 
```json
{
  "name": "eth_getTransactionByHash",
  "summary": "Returns the information about a transaction requested by transaction hash.",
  "params": [
    {
      "name": "Transaction hash",
      "required": true,
      "schema": {
        "title": "32 byte hex value",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      }
    }
  ],
  "result": {
    "name": "Transaction information",
    "schema": {
      "type": "object",
      "title": "Transaction information",
      "required": [
        "blockHash",
        "blockNumber",
        "from",
        "hash",
        "transactionIndex"
      ],
      "oneOf": [
        {
          "title": "Signed 1559 Transaction",
          "type": "object",
          "required": [
            "accessList",
            "chainId",
            "gas",
            "input",
            "maxFeePerGas",
            "maxPriorityFeePerGas",
            "nonce",
            "r",
            "s",
            "type",
            "value",
            "yParity"
          ],
          "properties": {
            "type": {
              "title": "type",
              "type": "string",
              "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
            },
            "nonce": {
              "title": "nonce",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "to": {
              "title": "to address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "gas": {
              "title": "gas limit",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "value": {
              "title": "value",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "input": {
              "title": "input data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "maxPriorityFeePerGas": {
              "title": "max priority fee per gas",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
            },
            "maxFeePerGas": {
              "title": "max fee per gas",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
            },
            "accessList": {
              "title": "accessList",
              "type": "array",
              "description": "EIP-2930 access list",
              "items": {
                "title": "Access list entry",
                "type": "object",
                "properties": {
                  "address": {
                    "title": "hex encoded address",
                    "type": "string",
                    "pattern": "^0x[0-9,a-f,A-F]{40}$"
                  },
                  "storageKeys": {
                    "type": "array",
                    "items": {
                      "title": "32 byte hex value",
                      "type": "string",
                      "pattern": "^0x[0-9a-f]{64}$"
                    }
                  }
                }
              }
            },
            "chainId": {
              "title": "chainId",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Chain ID that this transaction is valid on."
            },
            "yParity": {
              "title": "yParity",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
            },
            "r": {
              "title": "r",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "s": {
              "title": "s",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            }
          }
        },
        {
          "title": "Signed 2930 Transaction",
          "type": "object",
          "required": [
            "accessList",
            "chainId",
            "gas",
            "gasPrice",
            "input",
            "nonce",
            "r",
            "s",
            "type",
            "value",
            "yParity"
          ],
          "properties": {
            "type": {
              "title": "type",
              "type": "string",
              "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
            },
            "nonce": {
              "title": "nonce",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "to": {
              "title": "to address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "gas": {
              "title": "gas limit",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "value": {
              "title": "value",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "input": {
              "title": "input data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "gasPrice": {
              "title": "gas price",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The gas price willing to be paid by the sender in wei"
            },
            "accessList": {
              "title": "accessList",
              "type": "array",
              "description": "EIP-2930 access list",
              "items": {
                "title": "Access list entry",
                "type": "object",
                "properties": {
                  "address": {
                    "title": "hex encoded address",
                    "type": "string",
                    "pattern": "^0x[0-9,a-f,A-F]{40}$"
                  },
                  "storageKeys": {
                    "type": "array",
                    "items": {
                      "title": "32 byte hex value",
                      "type": "string",
                      "pattern": "^0x[0-9a-f]{64}$"
                    }
                  }
                }
              }
            },
            "chainId": {
              "title": "chainId",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Chain ID that this transaction is valid on."
            },
            "yParity": {
              "title": "yParity",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
            },
            "r": {
              "title": "r",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "s": {
              "title": "s",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            }
          }
        },
        {
          "title": "Signed Legacy Transaction",
          "type": "object",
          "required": [
            "gas",
            "gasPrice",
            "input",
            "nonce",
            "r",
            "s",
            "type",
            "v",
            "value"
          ],
          "properties": {
            "type": {
              "title": "type",
              "type": "string",
              "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
            },
            "nonce": {
              "title": "nonce",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "to": {
              "title": "to address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "gas": {
              "title": "gas limit",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "value": {
              "title": "value",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "input": {
              "title": "input data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "gasPrice": {
              "title": "gas price",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "The gas price willing to be paid by the sender in wei"
            },
            "chainId": {
              "title": "chainId",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
              "description": "Chain ID that this transaction is valid on."
            },
            "v": {
              "title": "v",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "r": {
              "title": "r",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "s": {
              "title": "s",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            }
          }
        }
      ],
      "properties": {
        "blockHash": {
          "title": "block hash",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        },
        "blockNumber": {
          "title": "block number",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
        },
        "from": {
          "title": "from address",
          "type": "string",
          "pattern": "^0x[0-9,a-f,A-F]{40}$"
        },
        "hash": {
          "title": "transaction hash",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        },
        "transactionIndex": {
          "title": "transaction index",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
        }
      }
    }
  }
}
```


### eth_getTransactionReceipt.json echo -e 
```json
{
  "name": "eth_getTransactionReceipt",
  "summary": "Returns the receipt of a transaction by transaction hash.",
  "params": [
    {
      "name": "Transaction hash",
      "schema": {
        "title": "32 byte hex value",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      }
    }
  ],
  "result": {
    "name": "Receipt Information",
    "schema": {
      "type": "object",
      "title": "Receipt info",
      "required": [
        "blockHash",
        "blockNumber",
        "from",
        "cumulativeGasUsed",
        "gasUsed",
        "logs",
        "logsBloom",
        "transactionHash",
        "transactionIndex",
        "effectiveGasPrice"
      ],
      "properties": {
        "transactionHash": {
          "title": "transaction hash",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        },
        "transactionIndex": {
          "title": "transaction index",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
        },
        "blockHash": {
          "title": "block hash",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        },
        "blockNumber": {
          "title": "block number",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
        },
        "from": {
          "title": "from",
          "type": "string",
          "pattern": "^0x[0-9,a-f,A-F]{40}$"
        },
        "to": {
          "title": "to",
          "type": "string",
          "pattern": "^0x[0-9,a-f,A-F]{40}$",
          "description": "Address of the receiver or null in a contract creation transaction."
        },
        "cumulativeGasUsed": {
          "title": "cumulative gas used",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
          "description": "The sum of gas used by this transaction and all preceding transactions in the same block."
        },
        "gasUsed": {
          "title": "gas used",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
          "description": "The amount of gas used for this specific transaction alone."
        },
        "contractAddress": {
          "title": "contract address",
          "type": "string",
          "pattern": "^0x[0-9,a-f,A-F]{40}$",
          "description": "The contract address created, if the transaction was a contract creation, otherwise null."
        },
        "logs": {
          "title": "logs",
          "type": "array",
          "items": {
            "title": "log",
            "type": "object",
            "required": [
              "transactionHash"
            ],
            "properties": {
              "removed": {
                "title": "removed",
                "type": "boolean"
              },
              "logIndex": {
                "title": "log index",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "transactionIndex": {
                "title": "transaction index",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "transactionHash": {
                "title": "transaction hash",
                "type": "string",
                "pattern": "^0x[0-9a-f]{64}$"
              },
              "blockHash": {
                "title": "block hash",
                "type": "string",
                "pattern": "^0x[0-9a-f]{64}$"
              },
              "blockNumber": {
                "title": "block number",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "address": {
                "title": "address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "data": {
                "title": "data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "topics": {
                "title": "topics",
                "type": "array",
                "items": {
                  "title": "32 hex encoded bytes",
                  "type": "string",
                  "pattern": "^0x[0-9a-f]{64}$"
                }
              }
            }
          }
        },
        "logsBloom": {
          "title": "logs bloom",
          "type": "string",
          "pattern": "^0x[0-9a-f]{512}$"
        },
        "root": {
          "title": "state root",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$",
          "description": "The post-transaction state root. Only specified for transactions included before the Byzantium upgrade."
        },
        "status": {
          "title": "status",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
          "description": "Either 1 (success) or 0 (failure). Only specified for transactions included after the Byzantium upgrade."
        },
        "effectiveGasPrice": {
          "title": "effective gas price",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
          "description": "The actual value per gas deducted from the senders account. Before EIP-1559, this is equal to the transaction's gas price. After, it is equal to baseFeePerGas + min(maxFeePerGas - baseFeePerGas, maxPriorityFeePerGas)."
        }
      }
    }
  }
}
```


### eth_newFilter.json echo -e 
```json
{
  "name": "eth_newFilter",
  "summary": "Creates a filter object, based on filter options, to notify when the state changes (logs).",
  "params": [
    {
      "name": "Filter",
      "schema": {
        "title": "filter",
        "type": "object",
        "properties": {
          "fromBlock": {
            "title": "from block",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "toBlock": {
            "title": "to block",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "address": {
            "title": "Address(es)",
            "oneOf": [
              {
                "title": "Address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              {
                "title": "Addresses",
                "type": "array",
                "items": {
                  "title": "hex encoded address",
                  "type": "string",
                  "pattern": "^0x[0-9,a-f,A-F]{40}$"
                }
              }
            ]
          },
          "topics": {
            "title": "Topics",
            "type": "array",
            "items": {
              "title": "Filter Topic List Entry",
              "oneOf": [
                {
                  "title": "Any Topic Match",
                  "type": "null"
                },
                {
                  "title": "Single Topic Match",
                  "type": "string",
                  "pattern": "^0x[0-9a-f]{64}$"
                },
                {
                  "title": "Multiple Topic Match",
                  "type": "array",
                  "items": {
                    "title": "32 hex encoded bytes",
                    "type": "string",
                    "pattern": "^0x[0-9a-f]{64}$"
                  }
                }
              ]
            }
          }
        }
      }
    }
  ],
  "result": {
    "name": "Filter Identifier",
    "schema": {
      "title": "hex encoded unsigned integer",
      "type": "string",
      "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
    }
  }
}
```


### eth_newPendingTransactionFilter.json echo -e 
```json
{
  "name": "eth_newPendingTransactionFilter",
  "summary": "Creates a filter in the node, to notify when new pending transactions arrive.",
  "params": [],
  "result": {
    "name": "Filter Identifier",
    "schema": {
      "title": "hex encoded unsigned integer",
      "type": "string",
      "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
    }
  }
}
```


### eth_protocolVersion.json echo -e 
```json
{
  "name": "eth_protocolVersion",
  "summary": "Returns the current Ethereum protocol version.",
  "params": [],
  "result": {
    "name": "Protocol version",
    "schema": {
      "title": "version",
      "type": "string"
    }
  }
}
```


### eth_sendPrivateTransaction.json echo -e 
```json
{
  "name": "eth_sendTransaction",
  "summary": "Signs and submits a transaction.",
  "params": [
    {
      "name": "Transaction",
      "required": true,
      "schema": {
        "title": "Transaction object with sender",
        "type": "object",
        "required": [
          "from"
        ],
        "oneOf": [
          {
            "type": "object",
            "title": "EIP-1559 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "maxFeePerGas",
              "maxPriorityFeePerGas",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "maxPriorityFeePerGas": {
                "title": "max priority fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
              },
              "maxFeePerGas": {
                "title": "max fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "EIP-2930 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "Legacy transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          }
        ],
        "properties": {
          "from": {
            "title": "from",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          }
        }
      }
    }
  ],
  "result": {
    "name": "Transaction hash",
    "schema": {
      "title": "32 byte hex value",
      "type": "string",
      "pattern": "^0x[0-9a-f]{64}$"
    }
  }
}
```


### eth_sendRawTransaction.json echo -e 
```json
{
  "name": "eth_sendRawTransaction",
  "summary": "Submits a raw transaction.",
  "params": [
    {
      "name": "Transaction",
      "required": true,
      "schema": {
        "title": "hex encoded bytes",
        "type": "string",
        "pattern": "^0x[0-9a-f]*$"
      }
    }
  ],
  "result": {
    "name": "Transaction hash",
    "schema": {
      "title": "32 byte hex value",
      "type": "string",
      "pattern": "^0x[0-9a-f]{64}$"
    }
  }
}
```


### eth_sendTransaction.json echo -e 
```json
{
  "name": "eth_sendTransaction",
  "summary": "Signs and submits a transaction.",
  "params": [
    {
      "name": "Transaction",
      "required": true,
      "schema": {
        "title": "Transaction object with sender",
        "type": "object",
        "required": [
          "from"
        ],
        "oneOf": [
          {
            "type": "object",
            "title": "EIP-1559 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "maxFeePerGas",
              "maxPriorityFeePerGas",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "maxPriorityFeePerGas": {
                "title": "max priority fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
              },
              "maxFeePerGas": {
                "title": "max fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "EIP-2930 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "Legacy transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          }
        ],
        "properties": {
          "from": {
            "title": "from",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          }
        }
      }
    }
  ],
  "result": {
    "name": "Transaction hash",
    "schema": {
      "title": "32 byte hex value",
      "type": "string",
      "pattern": "^0x[0-9a-f]{64}$"
    }
  }
}
```


### eth_sign.json echo -e 
```json
{
  "name": "eth_sign",
  "summary": "Returns an EIP-191 signature over the provided data.",
  "params": [
    {
      "name": "Address",
      "required": true,
      "schema": {
        "title": "hex encoded address",
        "type": "string",
        "pattern": "^0x[0-9,a-f,A-F]{40}$"
      }
    },
    {
      "name": "Message",
      "required": true,
      "schema": {
        "title": "hex encoded bytes",
        "type": "string",
        "pattern": "^0x[0-9a-f]*$"
      }
    }
  ],
  "result": {
    "name": "Signature",
    "schema": {
      "title": "65 hex encoded bytes",
      "type": "string",
      "pattern": "^0x[0-9a-f]{512}$"
    }
  }
}
```


### eth_signTransaction.json echo -e 
```json
{
  "name": "eth_signTransaction",
  "summary": "Returns an RLP encoded transaction signed by the specified account.",
  "params": [
    {
      "name": "Transaction",
      "required": true,
      "schema": {
        "title": "Transaction object with sender",
        "type": "object",
        "required": [
          "from"
        ],
        "oneOf": [
          {
            "type": "object",
            "title": "EIP-1559 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "maxFeePerGas",
              "maxPriorityFeePerGas",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "maxPriorityFeePerGas": {
                "title": "max priority fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
              },
              "maxFeePerGas": {
                "title": "max fee per gas",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "EIP-2930 transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice",
              "chainId",
              "accessList"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "accessList": {
                "title": "accessList",
                "type": "array",
                "description": "EIP-2930 access list",
                "items": {
                  "title": "Access list entry",
                  "type": "object",
                  "properties": {
                    "address": {
                      "title": "hex encoded address",
                      "type": "string",
                      "pattern": "^0x[0-9,a-f,A-F]{40}$"
                    },
                    "storageKeys": {
                      "type": "array",
                      "items": {
                        "title": "32 byte hex value",
                        "type": "string",
                        "pattern": "^0x[0-9a-f]{64}$"
                      }
                    }
                  }
                }
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          },
          {
            "type": "object",
            "title": "Legacy transaction.",
            "required": [
              "type",
              "nonce",
              "gas",
              "value",
              "input",
              "gasPrice"
            ],
            "properties": {
              "type": {
                "title": "type",
                "type": "string",
                "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
              },
              "nonce": {
                "title": "nonce",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "to": {
                "title": "to address",
                "type": "string",
                "pattern": "^0x[0-9,a-f,A-F]{40}$"
              },
              "gas": {
                "title": "gas limit",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "value": {
                "title": "value",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
              },
              "input": {
                "title": "input data",
                "type": "string",
                "pattern": "^0x[0-9a-f]*$"
              },
              "gasPrice": {
                "title": "gas price",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "The gas price willing to be paid by the sender in wei"
              },
              "chainId": {
                "title": "chainId",
                "type": "string",
                "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
                "description": "Chain ID that this transaction is valid on."
              }
            }
          }
        ],
        "properties": {
          "from": {
            "title": "from",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          }
        }
      }
    }
  ],
  "result": {
    "name": "Encoded transaction",
    "schema": {
      "title": "hex encoded bytes",
      "type": "string",
      "pattern": "^0x[0-9a-f]*$"
    }
  }
}
```


### index.json echo -e 
```json
["eth_getBlockByHash"]
["eth_getBlockByNumber"]
["eth_getBlockTransactionCountByHash"]
["eth_getBlockTransactionCountByNumber"]
["eth_getUncleCountByBlockHash"]
["eth_getUncleCountByBlockNumber"]
["eth_protocolVersion"]
["eth_chainId"]
["eth_syncing"]
["eth_coinbase"]
["eth_accounts"]
["eth_blockNumber"]
["eth_call"]
["eth_estimateGas"]
["eth_gasPrice"]
["eth_feeHistory"]
["eth_newFilter"]
["eth_newBlockFilter"]
["eth_newPendingTransactionFilter"]
["eth_uninstallFilter"]
["eth_getFilterChanges"]
["eth_getFilterLogs"]
["eth_getLogs"]
["eth_mining"]
["eth_hashrate"]
["eth_getWork"]
["eth_submitWork"]
["eth_submitHashrate"]
["eth_sign"]
["eth_signTransaction"]
["eth_getBalance"]
["eth_getStorageAt"]
["eth_getTransactionCount"]
["eth_getCode"]
["eth_sendTransaction"]
["eth_sendRawTransaction"]
["eth_getTransactionByHash"]
["eth_getTransactionByBlockHashAndIndex"]
["eth_getTransactionByBlockNumberAndIndex"]
["eth_getTransactionReceipt"]
```


### result_ReturnData.json echo -e 
```json
{
  "name": "Return data",
  "schema": {
    "title": "hex encoded bytes",
    "type": "string",
    "pattern": "^0x[0-9a-f]*$"
  }
}
```


### result_eth_blockNumber.json echo -e 
```json
{
  "name": "Block number",
  "schema": {
    "title": "hex encoded unsigned integer",
    "type": "string",
    "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
  }
}
```


### result_eth_call.json echo -e 
```json
{
  "name": "Return data",
  "schema": {
    "title": "hex encoded bytes",
    "type": "string",
    "pattern": "^0x[0-9a-f]*$"
  }
}
```


### result_eth_chainId.json echo -e 
```json
{
  "name": "Chain ID",
  "schema": {
    "title": "hex encoded unsigned integer",
    "type": "string",
    "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
  }
}
```


### result_eth_getFilterLogs.json echo -e 
```json
{
  "name": "Log objects",
  "schema": {
    "title": "Filter results",
    "oneOf": [
      {
        "title": "new block hashes",
        "type": "array",
        "items": {
          "title": "32 byte hex value",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        }
      },
      {
        "title": "new transaction hashes",
        "type": "array",
        "items": {
          "title": "32 byte hex value",
          "type": "string",
          "pattern": "^0x[0-9a-f]{64}$"
        }
      },
      {
        "title": "new logs",
        "type": "array",
        "items": {
          "title": "log",
          "type": "object",
          "required": [
            "transactionHash"
          ],
          "properties": {
            "removed": {
              "title": "removed",
              "type": "boolean"
            },
            "logIndex": {
              "title": "log index",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "transactionIndex": {
              "title": "transaction index",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "transactionHash": {
              "title": "transaction hash",
              "type": "string",
              "pattern": "^0x[0-9a-f]{64}$"
            },
            "blockHash": {
              "title": "block hash",
              "type": "string",
              "pattern": "^0x[0-9a-f]{64}$"
            },
            "blockNumber": {
              "title": "block number",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "address": {
              "title": "address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "data": {
              "title": "data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "topics": {
              "title": "topics",
              "type": "array",
              "items": {
                "title": "32 hex encoded bytes",
                "type": "string",
                "pattern": "^0x[0-9a-f]{64}$"
              }
            }
          }
        }
      }
    ]
  }
}
```


### result_eth_getTransactionByBlockHashAndIndex.json echo -e 
```json
{
  "name": "Transaction information",
  "schema": {
    "type": "object",
    "title": "Transaction information",
    "required": [
      "blockHash",
      "blockNumber",
      "from",
      "hash",
      "transactionIndex"
    ],
    "oneOf": [
      {
        "title": "Signed 1559 Transaction",
        "type": "object",
        "required": [
          "accessList",
          "chainId",
          "gas",
          "input",
          "maxFeePerGas",
          "maxPriorityFeePerGas",
          "nonce",
          "r",
          "s",
          "type",
          "value",
          "yParity"
        ],
        "properties": {
          "type": {
            "title": "type",
            "type": "string",
            "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
          },
          "nonce": {
            "title": "nonce",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "to": {
            "title": "to address",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          },
          "gas": {
            "title": "gas limit",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "value": {
            "title": "value",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "input": {
            "title": "input data",
            "type": "string",
            "pattern": "^0x[0-9a-f]*$"
          },
          "maxPriorityFeePerGas": {
            "title": "max priority fee per gas",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
          },
          "maxFeePerGas": {
            "title": "max fee per gas",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
          },
          "accessList": {
            "title": "accessList",
            "type": "array",
            "description": "EIP-2930 access list",
            "items": {
              "title": "Access list entry",
              "type": "object",
              "properties": {
                "address": {
                  "title": "hex encoded address",
                  "type": "string",
                  "pattern": "^0x[0-9,a-f,A-F]{40}$"
                },
                "storageKeys": {
                  "type": "array",
                  "items": {
                    "title": "32 byte hex value",
                    "type": "string",
                    "pattern": "^0x[0-9a-f]{64}$"
                  }
                }
              }
            }
          },
          "chainId": {
            "title": "chainId",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Chain ID that this transaction is valid on."
          },
          "yParity": {
            "title": "yParity",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
          },
          "r": {
            "title": "r",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "s": {
            "title": "s",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          }
        }
      },
      {
        "title": "Signed 2930 Transaction",
        "type": "object",
        "required": [
          "accessList",
          "chainId",
          "gas",
          "gasPrice",
          "input",
          "nonce",
          "r",
          "s",
          "type",
          "value",
          "yParity"
        ],
        "properties": {
          "type": {
            "title": "type",
            "type": "string",
            "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
          },
          "nonce": {
            "title": "nonce",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "to": {
            "title": "to address",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          },
          "gas": {
            "title": "gas limit",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "value": {
            "title": "value",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "input": {
            "title": "input data",
            "type": "string",
            "pattern": "^0x[0-9a-f]*$"
          },
          "gasPrice": {
            "title": "gas price",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The gas price willing to be paid by the sender in wei"
          },
          "accessList": {
            "title": "accessList",
            "type": "array",
            "description": "EIP-2930 access list",
            "items": {
              "title": "Access list entry",
              "type": "object",
              "properties": {
                "address": {
                  "title": "hex encoded address",
                  "type": "string",
                  "pattern": "^0x[0-9,a-f,A-F]{40}$"
                },
                "storageKeys": {
                  "type": "array",
                  "items": {
                    "title": "32 byte hex value",
                    "type": "string",
                    "pattern": "^0x[0-9a-f]{64}$"
                  }
                }
              }
            }
          },
          "chainId": {
            "title": "chainId",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Chain ID that this transaction is valid on."
          },
          "yParity": {
            "title": "yParity",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
          },
          "r": {
            "title": "r",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "s": {
            "title": "s",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          }
        }
      },
      {
        "title": "Signed Legacy Transaction",
        "type": "object",
        "required": [
          "gas",
          "gasPrice",
          "input",
          "nonce",
          "r",
          "s",
          "type",
          "v",
          "value"
        ],
        "properties": {
          "type": {
            "title": "type",
            "type": "string",
            "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
          },
          "nonce": {
            "title": "nonce",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "to": {
            "title": "to address",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          },
          "gas": {
            "title": "gas limit",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "value": {
            "title": "value",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "input": {
            "title": "input data",
            "type": "string",
            "pattern": "^0x[0-9a-f]*$"
          },
          "gasPrice": {
            "title": "gas price",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The gas price willing to be paid by the sender in wei"
          },
          "chainId": {
            "title": "chainId",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Chain ID that this transaction is valid on."
          },
          "v": {
            "title": "v",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "r": {
            "title": "r",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "s": {
            "title": "s",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          }
        }
      }
    ],
    "properties": {
      "blockHash": {
        "title": "block hash",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      },
      "blockNumber": {
        "title": "block number",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      },
      "from": {
        "title": "from address",
        "type": "string",
        "pattern": "^0x[0-9,a-f,A-F]{40}$"
      },
      "hash": {
        "title": "transaction hash",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      },
      "transactionIndex": {
        "title": "transaction index",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      }
    }
  }
}
```


### result_eth_getTransactionByHash.json echo -e 
```json
{
  "name": "Transaction information",
  "schema": {
    "type": "object",
    "title": "Transaction information",
    "required": [
      "blockHash",
      "blockNumber",
      "from",
      "hash",
      "transactionIndex"
    ],
    "oneOf": [
      {
        "title": "Signed 1559 Transaction",
        "type": "object",
        "required": [
          "accessList",
          "chainId",
          "gas",
          "input",
          "maxFeePerGas",
          "maxPriorityFeePerGas",
          "nonce",
          "r",
          "s",
          "type",
          "value",
          "yParity"
        ],
        "properties": {
          "type": {
            "title": "type",
            "type": "string",
            "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
          },
          "nonce": {
            "title": "nonce",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "to": {
            "title": "to address",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          },
          "gas": {
            "title": "gas limit",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "value": {
            "title": "value",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "input": {
            "title": "input data",
            "type": "string",
            "pattern": "^0x[0-9a-f]*$"
          },
          "maxPriorityFeePerGas": {
            "title": "max priority fee per gas",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Maximum fee per gas the sender is willing to pay to miners in wei"
          },
          "maxFeePerGas": {
            "title": "max fee per gas",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei"
          },
          "accessList": {
            "title": "accessList",
            "type": "array",
            "description": "EIP-2930 access list",
            "items": {
              "title": "Access list entry",
              "type": "object",
              "properties": {
                "address": {
                  "title": "hex encoded address",
                  "type": "string",
                  "pattern": "^0x[0-9,a-f,A-F]{40}$"
                },
                "storageKeys": {
                  "type": "array",
                  "items": {
                    "title": "32 byte hex value",
                    "type": "string",
                    "pattern": "^0x[0-9a-f]{64}$"
                  }
                }
              }
            }
          },
          "chainId": {
            "title": "chainId",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Chain ID that this transaction is valid on."
          },
          "yParity": {
            "title": "yParity",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
          },
          "r": {
            "title": "r",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "s": {
            "title": "s",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          }
        }
      },
      {
        "title": "Signed 2930 Transaction",
        "type": "object",
        "required": [
          "accessList",
          "chainId",
          "gas",
          "gasPrice",
          "input",
          "nonce",
          "r",
          "s",
          "type",
          "value",
          "yParity"
        ],
        "properties": {
          "type": {
            "title": "type",
            "type": "string",
            "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
          },
          "nonce": {
            "title": "nonce",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "to": {
            "title": "to address",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          },
          "gas": {
            "title": "gas limit",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "value": {
            "title": "value",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "input": {
            "title": "input data",
            "type": "string",
            "pattern": "^0x[0-9a-f]*$"
          },
          "gasPrice": {
            "title": "gas price",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The gas price willing to be paid by the sender in wei"
          },
          "accessList": {
            "title": "accessList",
            "type": "array",
            "description": "EIP-2930 access list",
            "items": {
              "title": "Access list entry",
              "type": "object",
              "properties": {
                "address": {
                  "title": "hex encoded address",
                  "type": "string",
                  "pattern": "^0x[0-9,a-f,A-F]{40}$"
                },
                "storageKeys": {
                  "type": "array",
                  "items": {
                    "title": "32 byte hex value",
                    "type": "string",
                    "pattern": "^0x[0-9a-f]{64}$"
                  }
                }
              }
            }
          },
          "chainId": {
            "title": "chainId",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Chain ID that this transaction is valid on."
          },
          "yParity": {
            "title": "yParity",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature."
          },
          "r": {
            "title": "r",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "s": {
            "title": "s",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          }
        }
      },
      {
        "title": "Signed Legacy Transaction",
        "type": "object",
        "required": [
          "gas",
          "gasPrice",
          "input",
          "nonce",
          "r",
          "s",
          "type",
          "v",
          "value"
        ],
        "properties": {
          "type": {
            "title": "type",
            "type": "string",
            "pattern": "^0x([0-9,a-f,A-F]?){1,2}$"
          },
          "nonce": {
            "title": "nonce",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "to": {
            "title": "to address",
            "type": "string",
            "pattern": "^0x[0-9,a-f,A-F]{40}$"
          },
          "gas": {
            "title": "gas limit",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "value": {
            "title": "value",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "input": {
            "title": "input data",
            "type": "string",
            "pattern": "^0x[0-9a-f]*$"
          },
          "gasPrice": {
            "title": "gas price",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "The gas price willing to be paid by the sender in wei"
          },
          "chainId": {
            "title": "chainId",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "Chain ID that this transaction is valid on."
          },
          "v": {
            "title": "v",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "r": {
            "title": "r",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          },
          "s": {
            "title": "s",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
          }
        }
      }
    ],
    "properties": {
      "blockHash": {
        "title": "block hash",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      },
      "blockNumber": {
        "title": "block number",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      },
      "from": {
        "title": "from address",
        "type": "string",
        "pattern": "^0x[0-9,a-f,A-F]{40}$"
      },
      "hash": {
        "title": "transaction hash",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      },
      "transactionIndex": {
        "title": "transaction index",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      }
    }
  }
}
```


### result_eth_getTransactionReceipt.json echo -e 
```json
{
  "name": "Receipt Information",
  "schema": {
    "type": "object",
    "title": "Receipt info",
    "required": [
      "blockHash",
      "blockNumber",
      "from",
      "cumulativeGasUsed",
      "gasUsed",
      "logs",
      "logsBloom",
      "transactionHash",
      "transactionIndex",
      "effectiveGasPrice"
    ],
    "properties": {
      "transactionHash": {
        "title": "transaction hash",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      },
      "transactionIndex": {
        "title": "transaction index",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      },
      "blockHash": {
        "title": "block hash",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$"
      },
      "blockNumber": {
        "title": "block number",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
      },
      "from": {
        "title": "from",
        "type": "string",
        "pattern": "^0x[0-9,a-f,A-F]{40}$"
      },
      "to": {
        "title": "to",
        "type": "string",
        "pattern": "^0x[0-9,a-f,A-F]{40}$",
        "description": "Address of the receiver or null in a contract creation transaction."
      },
      "cumulativeGasUsed": {
        "title": "cumulative gas used",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
        "description": "The sum of gas used by this transaction and all preceding transactions in the same block."
      },
      "gasUsed": {
        "title": "gas used",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
        "description": "The amount of gas used for this specific transaction alone."
      },
      "contractAddress": {
        "title": "contract address",
        "type": "string",
        "pattern": "^0x[0-9,a-f,A-F]{40}$",
        "description": "The contract address created, if the transaction was a contract creation, otherwise null."
      },
      "logs": {
        "title": "logs",
        "type": "array",
        "items": {
          "title": "log",
          "type": "object",
          "required": [
            "transactionHash"
          ],
          "properties": {
            "removed": {
              "title": "removed",
              "type": "boolean"
            },
            "logIndex": {
              "title": "log index",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "transactionIndex": {
              "title": "transaction index",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "transactionHash": {
              "title": "transaction hash",
              "type": "string",
              "pattern": "^0x[0-9a-f]{64}$"
            },
            "blockHash": {
              "title": "block hash",
              "type": "string",
              "pattern": "^0x[0-9a-f]{64}$"
            },
            "blockNumber": {
              "title": "block number",
              "type": "string",
              "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
            },
            "address": {
              "title": "address",
              "type": "string",
              "pattern": "^0x[0-9,a-f,A-F]{40}$"
            },
            "data": {
              "title": "data",
              "type": "string",
              "pattern": "^0x[0-9a-f]*$"
            },
            "topics": {
              "title": "topics",
              "type": "array",
              "items": {
                "title": "32 hex encoded bytes",
                "type": "string",
                "pattern": "^0x[0-9a-f]{64}$"
              }
            }
          }
        }
      },
      "logsBloom": {
        "title": "logs bloom",
        "type": "string",
        "pattern": "^0x[0-9a-f]{512}$"
      },
      "root": {
        "title": "state root",
        "type": "string",
        "pattern": "^0x[0-9a-f]{64}$",
        "description": "The post-transaction state root. Only specified for transactions included before the Byzantium upgrade."
      },
      "status": {
        "title": "status",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
        "description": "Either 1 (success) or 0 (failure). Only specified for transactions included after the Byzantium upgrade."
      },
      "effectiveGasPrice": {
        "title": "effective gas price",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
        "description": "The actual value per gas deducted from the senders account. Before EIP-1559, this is equal to the transaction's gas price. After, it is equal to baseFeePerGas + min(maxFeePerGas - baseFeePerGas, maxPriorityFeePerGas)."
      }
    }
  }
}
```


### result_eth_newPendingTransactionFilter.json echo -e 
```json
{
  "name": "Filter Identifier",
  "schema": {
    "title": "hex encoded unsigned integer",
    "type": "string",
    "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
  }
}
```


### result_eth_protocolVersion.json echo -e 
```json
{
  "name": "Protocol version",
  "schema": {
    "title": "version",
    "type": "string"
  }
}
```


### result_eth_sendRawTransaction.json echo -e 
```json
{
  "name": "Transaction hash",
  "schema": {
    "title": "32 byte hex value",
    "type": "string",
    "pattern": "^0x[0-9a-f]{64}$"
  }
}
```


### result_eth_sendTransaction.json echo -e 
```json
{
  "name": "Transaction hash",
  "schema": {
    "title": "32 byte hex value",
    "type": "string",
    "pattern": "^0x[0-9a-f]{64}$"
  }
}
```


### return_eth_feeHistory.json echo -e 
```json
{
  "name": "feeHistoryResult",
  "description": "Fee history for the returned block range. This can be a subsection of the requested range if not all blocks are available.",
  "schema": {
    "title": "feeHistoryResults",
    "description": "Fee history results.",
    "type": "object",
    "required": [
      "firstBlock",
      "baseFeePerGas",
      "gasUsedRatio"
    ],
    "properties": {
      "oldestBlock": {
        "title": "oldestBlock",
        "type": "string",
        "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
        "description": "Lowest number block of returned range."
      },
      "baseFeePerGas": {
        "title": "baseFeePerGasArray",
        "description": "An array of block base fees per gas. This includes the next block after the newest of the returned range, because this value can be derived from the newest block. Zeroes are returned for pre-EIP-1559 blocks.",
        "type": "array",
        "items": {
          "title": "hex encoded unsigned integer",
          "type": "string",
          "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
        }
      },
      "reward": {
        "title": "rewardArray",
        "description": "A two-dimensional array of effective priority fees per gas at the requested block percentiles.",
        "type": "array",
        "items": {
          "title": "rewardPercentile",
          "description": "An array of effective priority fee per gas data points from a single block. All zeroes are returned if the block is empty.",
          "type": "array",
          "items": {
            "title": "rewardPercentile",
            "type": "string",
            "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$",
            "description": "A given percentile sample of effective priority fees per gas from a single block in ascending order, weighted by gas used. Zeroes are returned if the block is empty."
          }
        }
      }
    }
  }
}
```


### return_eth_gasPrice.json echo -e 
```json
{
  "name": "Gas price",
  "schema": {
    "title": "Gas price",
    "type": "string",
    "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
  }
}
```


### return_eth_newFilter.json echo -e 
```json
{
  "name": "Filter Identifier",
  "schema": {
    "title": "hex encoded unsigned integer",
    "type": "string",
    "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
  }
}
```


### reutnr_eth_estimateGas.json echo -e 
```json
{
  "name": "Gas used",
  "schema": {
    "title": "hex encoded unsigned integer",
    "type": "string",
    "pattern": "^0x([1-9a-f]+[0-9a-f]*|0)$"
  }
}
```
