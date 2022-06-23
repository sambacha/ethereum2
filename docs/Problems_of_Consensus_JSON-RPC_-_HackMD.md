-   [Problems of Consensus JSON-RPC](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Problems-of-Consensus-JSON-RPC "Problems of Consensus JSON-RPC")
    -   [Consensus block validity](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Consensus-block-validity "Consensus block validity")
    -   [Independent message flow](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Independent-message-flow "Independent message flow")
    -   [HTTP overhead](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#HTTP-overhead "HTTP overhead")
    -   [Failure recovery](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Failure-recovery "Failure recovery")
    -   [Extensibility](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Extensibility "Extensibility")

Consensus API is an interface used by consensus-layer to communicate with execution-layer in Ethereum Mainnet client after the Merge.

There is an initial version of this interface called [Consensus JSON-RPC](https://github.com/ethereum/rayonism/blob/master/specs/merge.md#consensus-json-rpc) that was implemented by Ethereum Mainnet clients participated in [Rayonism](https://rayonism.io/) project. Though, this version worked well in Rayonism devnets it doesn’t seem to fit well all production use cases.

This document addresses a couple of corner cases not covered by the original version, introduces rich scenarious that might be desired but aren’t possible with the current protocol implementation. Thus, making points that should be taken in account in the design of the new version of Consensus API.

### [](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Consensus-block-validity "Consensus-block-validity")Consensus block validity

Obviously, if consensus block appears to be invalid then despite of validity of its execution payload the payload must be discarded by execution client. Otherwise, JSON-RPC requests may potentially serve data that will never be included onchain.

`SetHead` message could be used to signify the block validity but there are cases when recently imported beacon block doesn’t become the chain head eventually, thus, this approach is not 100% reliable.

The other way would be to propagate execution payload when corresponding consensus block has been validated. But it would require processing of beacon block and execution payload to be sequential which is undesirable.

New type of messages is introduced to propagate the information of whether consensus is validated to the execution client. Work name of this message is `ConsensusProcessed`.

### [](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Independent-message-flow "Independent-message-flow")Independent message flow

Messages of existing protocol are independent from each other. It means that every message is sufficient enough for execution client wrt to received information to act accordingly. For example, if the payload from `NewBlock` message has been validated then the response is sent back to the consensus client and the payload and resulting state are stored in the chain storage.

New `ConsensusProcessed` message does break this assumption and requires to keep an intermediate state which would accumulate meta data and take according action once there is enough information to make a decision. Also, intermediate state allows for not respecting the order of two causally dependent messages. For intance, `ConsensusProcessed` can be sent and/or processed before corresponding `NewBlock` message.

Key-value data structure (map) should be sufficiently enough to maintain the state of communication protocol. Entries of this map may be evicted after the cycle of payload processing is finished and the payload should either be discarded or persisted in the chain store.

There is one caveat to use `block_hash` of execution payload as a key for entries in the protocol state. There is possibility of more than one beacon block to refer to the execution payload with the same hash. And there is a possibility of one of these beacon blocks to be invalidate by corresponding `ConsensusProcessed` message. In this case one copy of execution payload should still stay in the protocol cache. This case can be handled by either using beacon block root as the key to the entry or by using refcounting. The former seems to be less complicated.

### [](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#HTTP-overhead "HTTP-overhead")HTTP overhead

There are two sources of overhead coming from JSON-RPC protocol

-   JSON encoding increases the size of messages
-   each HTTP request uses its own connection

Websockets is one of the options to alleviate the overhead and allow for asynchronous communication between consensus and execution clients.

**UPD** after [Merge Implementers’ Call 7](https://github.com/ethereum/pm/issues/345):

-   HTTP `Keep-Alive` header can be used to avoid creation of new connection with each request
-   JSON payloads can be gzipped reducing the size significantly given the specifics of data. Websockets support gzip too

### [](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Failure-recovery "Failure-recovery")Failure recovery

Suppose there is an execution or consensus client crash that lead to inconcistency between databases across the border. There are two potential cases. An easy one, when the payload has been persisted by execution side but not by consensus side. In this case consensus client will start from the most recent block it has and no consistency issue would arise.

The other case is when cosnensus client has persisted a block while execution client hasn’t. Execution client will have to fall back to state/block sync and get missing information from the network upon receiving `NewBlock` message from consensus. But in case of one or two, or `N` blocks this behaviour is an overhead.

Consensus and execution sides may exchange with status messages upon startup and get the knowledge of the chain head from each other. Then consensus side may decide what to do. It may replay recent blocks if the gap is short or initiate a sync process on execution side. Execution client would have enough information to make a decision on whether to pull a few blocks or better to fall back to state sync.

### [](https://hackmd.io/@n0ble/problems_of_consensus_json_rpc#Extensibility "Extensibility")Extensibility

Edge cases mentioned above requires an underlying communication channnel to be _bi-directional_. Messages like `ConsensusProcessed` doesn’t fit an RPC paradigm well. Also, we might want to make execution client more explicit about its state during the sync process.

The general goal is to design a protocol that will be easily extensible in the future. If some optimisation opportunities or any other edge cases will be found later on the design of the protocol should not lay additional restrictions on a solution for these cases.

Obviously, the cost of a new protocol comes with cycles spent on designing, implementating and testing it.