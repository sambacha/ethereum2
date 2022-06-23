# OpenMEV Transaction and Pricing API

> This is a work in progress


> OpenMEV exposes a JSON-RPC protocol to send transactions/bundles, get transaction status updates, cancel transactions, and retrieve ETH network gas fees through its backbone platform providing connectivity to providers like Flashbots and Eden Network.
---


## `api.txprice.com`

### Usage

#### Get Ethereum Transaction Pricing

```bash
curl -L api.txprice.com
```

#### Get Pricing at 95% confidence interval

```bash
curl -s -L api.txprice.com | jq .blockPrices[0].estimatedPrices[1]
```

```json
{
  "confidence": 95,
  "price": 107,
  "maxPriorityFeePerGas": 2.52,
  "maxFeePerGas": 212.92
}
```

## Transaction Pricing API (api.txprice.com)

Returns a range of confidence intervals for gas prices needed to qualify a
transaction for inclusion in the next block. Order of confidence intervals is
subject to change.

##### Response

```jsonc
{
  "system": "ethereum",
  "network": "main",
  "unit": "gwei",
  "maxPrice": 123,
  "currentBlockNumber": 13005095,
  "msSinceLastBlock": 3793,
  "blockPrices": [
    {
      "blockNumber": 13005096,
      "baseFeePerGas": 94.647990462,
      "estimatedTransactionCount": 137,
      "estimatedPrices": [
        // json.blockPrices[0].estimatedPrices[0]
        {
          "confidence": 99, // json.blockPrices[0].estimatedPrices[0].confidence
          "price": 104, // json.blockPrices[0].estimatedPrices[0].price
          "maxPriorityFeePerGas": 9.86, // json.blockPrices[0].estimatedPrices[0].maxFeePerGas
          "maxFeePerGas": 199.16 // json.blockPrices[0].estimatedPrices[0].maxFeePerGas
        },
        {
          "confidence": 95,
          "price": 99,
          "maxPriorityFeePerGas": 5.06,
          "maxFeePerGas": 194.35
        },
        {
          "confidence": 90,
          "price": 98,
          "maxPriorityFeePerGas": 4.16,
          "maxFeePerGas": 193.45
        },
        {
          "confidence": 80,
          "price": 97,
          "maxPriorityFeePerGas": 2.97,
          "maxFeePerGas": 192.27
        },
        {
          "confidence": 70,
          "price": 96,
          "maxPriorityFeePerGas": 1.74,
          "maxFeePerGas": 191.04
        }
      ]
    }
  ]
}
```

## How it works


### Sending many transactions using eth\_sendRawTransactions

1.  send transactions using **eth\_sendRawTransactions**, the response will include a bundle ID
2.  use the bundle ID to regularly query the status of the bundle using **eth\_getBundleStatusById**
3.  the status can be either
    -   **PENDING\_BUNDLE**: the bundle is in progess, being sent to miners via flashbots for inclusion
    -   **FAILED\_BUNDLE**: the bundle has not been included. The reason will be shown in the response error message
    -   **SUCCESSFUL\_BUNDLE**: the bundle has been included by a miner

### Sending one transaction using eth\_sendRawTransaction


1.  send a single transaction using **eth\_sendRawTransaction**, the response will include the transaction hash
2.  use the transaction hash to regularly query the status using **eth\_getTransactionStatusByHash**
3.  the status can be either
    -   **PENDING\_BUNDLE**: the bundle is in progess, being sent to miners via flashbots for inclusion
    -   **FAILED\_BUNDLE**: the bundle has not been included. The reason will be shown in the response error message
    -   **SUCCESSFUL\_BUNDLE**: the bundle has been included by a miner

### Sending and canceling a transaction by bundle ID


1.  send a bundle of transactions using eth\_sendRawTransactions, the response will include a bundle ID
2.  use the bundle ID to regularly cancel the bundle using **eth\_cancelBundleById**
3.  use the bundle ID to regularly query the status of the bundle using **eth\_getBundleStatusById**
4.  the status can be either
    -   any of the above statuses
    -   **CANCEL\_BUNDLE\_SUCCESSFUL**: the bundle was canceled successfully

### Sending and canceling a transaction by hash

1.  send a transaction using eth\_sendRawTransaction, the response will include the transaction hash
2.  use the hash to regularly cancel the bundle using **eth\_cancelBundleByTransactionHash**
3.  use the transaction hash to regularly query the status of the bundle using **eth\_getTransactionStatusByHash**
4.  the status can be either
    -   any of the above statuses
    -   **CANCEL\_BUNDLE\_SUCCESSFUL**: the bundle was canceled successfully

### Recommended fees (optional)

To build a successful transaction, it is important to include enough fees to cover both the Ethereum **baseFee** as well as the **miner tip** (incentive for a miner to include your transaction in a block). It may be complex to estimate what the fees should be, especially when the network is heavily used and more transactions are competing with each other.

To help on that matter, the API provides a method that will send back **mistX recommended fees**. Those have been adjusted to ensure a good chance of inclusion while minimizing the miner payment to avoid overpayment. The method **eth\_gasFees** will send back fees organized in three tiers, so you can choose which one fits best your use case

-   **low**: lowest chance of inclusion, it will take more time to get included. Higher risk of not covering the baseFee in times of heavy traffic
-   **medium**: good chance of inclusion, it will be faster to get included. Medium risk of not covering the baseFee in times of heavy traffic
-   **high**: highest chance and fastest inclusion. Lowest risk of not covering the baseFee in times of heavy traffic

## Methods

### eth\_sendRawTransactions

Sends one or many raw transactions.

Parameters

Response

Example

Method name

Default:

eth\_sendRawTransactions

Array of serialized raw transactions

### eth\_sendRawTransaction

Sends a single raw transaction.

Parameters

Response

Example

The transaction hash

Match pattern:

```
^0x\[a-fA-F\\d\]{64}$
```

### eth\_getBundleStatusById

Get the status of a bundle

Parameters

Response

Example

Method name

Default:

eth\_getBundleStatusById

Array containing one bundle id

### eth\_getTransactionStatusByHash



Get the status of a transaction

Parameters

Response

Example

Method name

Default:

eth\_getTransactionStatusByHash

Array containing one transaction hash

### eth\_cancelBundleById



Cancel a bundle

Parameters

Response

Example

Method name

Default:

eth\_cancelBundleById

Array containing one bundle id

### eth\_cancelBundleByTransactionHash



Cancel a bundle containing a transaction using the transaction hash

Parameters

Response

Example

Method name

Default:

eth\_cancelBundleByTransactionHash

Array containing one bundle id

###  eth\_gasFees



Get current gas fees

Parameters

Response

Example

Empty array

\>= 0 items<= 0 items
