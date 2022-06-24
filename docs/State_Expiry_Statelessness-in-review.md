---
created: 2022-06-24T07:19:39 (UTC -07:00)
tags: []
source: https://www.notion.so/State-Expiry-Statelessness-in-Review-8d531abcc2984babb9bf76a44459e611
author: 
---

# State Expiry & Statelessness in Review

> ## Excerpt
> A new tool that blends your everyday work apps into one. It's the all-in-one workspace for you and your team

---
The goal of this dev update is to clarify and contextualize the efforts to moves towards state expiry in Ethereum. There is a lot written on the topic, including multiple summaries and roadmaps, and this post attempts to put everything together as of the end of June 2021. Links to the primary sources will be provided, including a master list at the end of this post.

The Ethereum Virtual Machine (EVM) maintains a memory store. In addition to account balances, nonces and contract code, any contract can request the allocation of a piece of storage. Transaction must pay for allocated storage (which we'll often refer to as "state"), as well as subsequent reads and writes. As long as the state is accessed frequently, the cost of reads and writes should be enough to pay for the cost of maintaining the state. However, even state that is never accessed must be maintained. Not only does this state live rent-free on every full node, but as dead state accretes, it makes bringing a full node online (a process known as "sync") significantly more difficult than it would be otherwise.

To give you an idea, a full node (which maintains the state, but does only keeps the N (usually N=128) latest blocks) [must currently store > 400 GB](https://etherscan.io/chartsync/chaindefault) (or more than double that if using Geth). Of this, the state only comprises [only about 30-60GB](https://ethereum-magicians.org/t/state-management-review/6602/) (however it is the only part who could grow unbounded). This is still better than archive nodes (who keep the whole block history), which must store [a whopping 7.5 TB](https://etherscan.io/chartsync/chainarchive).

This is a good point to mention a related issue. In any scheme we can come up with, someone will have to know about the state (although maybe only a small fraction of all nodes). How can we make sure that other people can learn about the state when they need to? If full state storage is required for some node, can we devise a more efficient solution to acquire this state than to download the full blockchain (7.5TB!) and replay every single transaction to find the current state?

Before looking at solutions, one precision that will become hugely relevant later. State in Ethereum is represented by key-value stores, stored in [Merkle trees](https://en.wikipedia.org/wiki/Merkle_tree). A Merkle tree is a [radix tree](https://en.wikipedia.org/wiki/Radix_tree) (a compressed [trie](https://en.wikipedia.org/wiki/Trie)) where each node is labelled with the hash of its children (or the hash of the data for leaf nodes). The keys are always [keccak](https://en.wikipedia.org/wiki/SHA-3) hashes of the "real" key — this ensures that all keys have the same size (256 bits) and makes it difficult to find key that share a prefix, which could by an attacker to forcibly reduce the amount of compression in the tree. There is a Merkle tree mapping (hashes of) accounts to account data (the state trie). The account data itself includes the hash of (the root of) another Merkle tree storing the memory for the account. These storage tries (one per account) are separate from the state tree, but make up for the bulk of what the call "the state". See [this illustration](https://raw.githubusercontent.com/4c656554/BlockchainIllustrations/master/Ethereum/EthBlockchain5.png).

These trees are often called "modified patricia Merkle trees", which is just another way to say these are radix trees of branching factor 16 annotated whose nodes are annonated with a Merkle hash.

How are those trees actually stored in practice? We can't just build the Merkle trees in- memory — hundreds of GB of RAM is too much. In practice, most nodes use the [LevelDB](https://github.com/google/leveldb) (or its fork [RocksDB](https://github.com/facebook/rocksdb)) key-value store. Because these stores are not implemented using Merkle trees, we need to store the hashes for the internal nodes of the Merkle trees as separate key-value pairs!

There are a few ways to attack the problem, but there are two major directions we can go in:

State Expiry: Make sure the state is "paid for", otherwise it expires and can't be used anymore (potentially, the state can be resurrected if payments resume).

Statelessness: Do not force nodes (even full nodes) to store any state. Instead all the state accessed by a transaction must be provided with the transaction, alongside proofs that the provided state is indeed part of the canonical EVM state.

A first idea, dating back to 2015, was called ["state rent"](https://github.com/ethereum/EIPs/issues/35), a form of state expiry. Under this proposal, rent would be extracted from the balance of the contracts making use of the state. Its major inconvenient is that the proposal is not backward compatible with existing contracts, as fees are paid on an ongoing basis, and not only when processing a transaction — instantly bankrupting existing contracts with an empty balance (and disrupting the correct operation of others). The mechanics of keeping a contract funded so that it may pay its rent are also a bit awkward. Ideally, the costs would be passed transparently to users of the contract, just like other gas fees.

A subsequent idea is that of statelessness. This is outlined in ["The Stateless Client Concept"](https://ethresear.ch/t/the-stateless-client-concept/172) post. The idea relies on the ability to build proofs (also called "witnesses") that particular pieces of state have a certain value. Only the hash of the root of the account's data trie is required to validate this proof. Under this model, miners could bundle witnesses along with the transactions, which would enable other full nodes to validate them without storing any state. This is known as "weak statelessness".

The post takes the idea further: could we make miners stateless and achieve "full statelessness"? We could, if we push the responsibility of supplying the witnesses to the transaction originators. This runs into two issues. First, the value of a piece of storage might change between transaction origination and transaction execution. This is easily fixed by letting miners store all recent (~24h) state changes and letting them alter the witnesses in case the state changed. The other issue is that originators might only be light nodes (i.e. nodes that do not validate blocks, but instead trust some other well-known nodes). There must be an easy way for such nodes to get hold of a witness for the data it is trying to access. We'll talk about some of what is available today and planned for the future in the [Distributing State](https://www.notion.so/8d531abcc2984babb9bf76a44459e611#621e2a542e8c4bc683089e7955ce0d36) section.

Another issue with full statelessness is that currently, a contract might not know in advance which state it is going to access (this is called [dynamic state access](https://ethresear.ch/t/state-provider-models-in-ethereum-2-0/6750)) — which makes it impossible to provide proof for! This is potentially a big issue, because it makes statelessness backward incompatible with old contracts.

In passing, there is currently an optional way to specify which accounts you'll interact with, [address lists](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2930.md), though it's meant as a way to make transaction processing easier. In particular it enables parallel transaction processing in some cases.

A final hurdle is the large size of proofs when using Merkle trees. As we'll see in a bit, this problem finds a solution in the Verkle tree data structure, which enables much smaller proofs.

Finally, the direction currently being pursued by the Ethereum core developers is that of state expiry, but without rent. Under [this proposal](https://notes.ethereum.org/@vbuterin/state_expiry_eip), nodes must only maintain state that has been accessed recently (e.g. in the last year). If a transaction wants to access older state, it must supply the state's value, as well as a witness (proof) for that value.

Compared to statelessness, this require all full nodes to keep track of some state — although the size of this state becomes effectively bounded (as the number of blocks in a period is limited, and the gas limit for each block is finite). However, it comes with the advantage that blocks do not need to bundle witnesses, nor do these witnesses need to be transmitted on the network when propagating transaction.

For another overview of the design space, I encourage you to read Vitalik Buterin's [A Theory of Ethereum State Size](https://hackmd.io/@vbuterin/state_size_management), up to and including the "When to expire" section.

One of the problems we started with was that state currently lives rent-free on all full nodes. Statelessness and rent-free state expiry help with this by putting an effective bound on the amount of state that nodes need to keep. However this begs the question of who will keep the historical state. The schemes do not directly incentivize this. Under state expiry, transaction senders will pay a gas fee for resurrecting old state, but they are also responsible to supply the old values and their witnesses.

There are two reasons one could use to dismiss this concern, at least temporarily. First, contract stakeholders have good incentives to track and publicize data related to their contract. Second, even if the total state grows to tens of terabytes, this is still small enough to store of fairly cheap mechanical hard drives. Mechanical hard drives are too slow to use to store the state used while mining, but it's perfectly fine to store expired values — miners will not have to look them up, as it is the transaction senders' responsibility to provide them. So it is not unrealistic to expect many nodes to altruistically provide these values in the future, especially if leading node implementations make it easy to do so.

This is all conjecture, however. Please read [The Data Availability Problem under Stateless Ethereum](https://ethresear.ch/t/the-data-availability-problem-under-stateless-ethereum/6973), which makes a good case as to why this might become an issue if stateless nodes become common enough. I believe this post was written with full statelessness in mind however, and the problem wouldn't be that bad under state expiry. It also talks about how to improve state propagation, which is the topic of the next section.

Another problem we outlined at the start was how to distribute state to new nodes.

Full statelessness solves this problem trivially, although the cost of bringing a "state distribution node" online may eventually become a concern. For state expiry, things could work as they currently do for non-expired state, while expired state would need to be acquired on demand.

Let's review how Ethereum client implementation let you retrieve the current state. There are a couple of possible ways.

full — Simply download the full blockchain and replay every single transaction. This requires terabytes of bandwidth, and can take a few days — the bottleneck being disk bandwidth while processing transactions. It also relies on the altruism of archive nodes that store the whole blockchain.

fast — [Fast sync](https://github.com/ethereum/go-ethereum/pull/1889) which might only take ~10 hours at present. With fast sync you still download the full blockchain and validate the [proof of work](https://norswap.com/blockchain-how/) of each block. You do not however process all transactions. Instead, you download the state at the most recent checkpoint (or "pivot") from other nodes, then start processing transactions normally from there on. This is secure, because blocks include the Merkle root hash of the state tree, so you can verify that the state tree that you download matches the pivot block, whose authenticity you ascertained by verifying the PoW chain. Fast sync works but is (ironically) still pretty slow, and hard on your bandwidth and disks. You can find more details on fast sync [here](https://github.com/ethereum/go-ethereum/pull/1889). One of the reason for it still being so slow is the way that state is retrieved: you ask for other nodes to give you the root node (identified by its hash). Once you have the answer, you can ask for the children of the root, and so on and so forth. A lot of time is l[ost in these network roundtrips](https://blog.ethereum.org/2021/03/03/geth-v1-10-0/#snap-sync).

snap/warp — Snap sync (OpenEthereum name) or warp sync (Go-Ethereum name) is a more efficient way to sync. Instead of downloading the state node-by-node, it relies on clients generating "state snapshots" that contain the whole state and can be requested by other clients. The requesting clients can then rebuild the state tree themselves, from the leaves up. This avoids the latency of network roundtrips, and lets the sending clients store the snapshot data in a continuous disk chunk (instead of spread all over as they would if they were values in a key-value store). There are some differences between snap & warp sync, which you can [read more about here](https://blog.ethereum.org/2021/03/03/geth-v1-10-0/#snap-sync). Note that snap & warp still download the whole blockchain like full & fast.

beam — In [beam sync](https://medium.com/@jason.carver/intro-to-beam-sync-a0fd168be14a), nodes do not need to download the whole state before they start validating. Instead, they request state from the network as they need it. In particular, they request a block witness for each block, which contains all the values accessed by the transaction as well as a Merkle proof for these values. A block proof is more compact than a sequence of individual value proofs, due o the fact the multiple values might share common ancestors (we'll review the structure of Merkle proofs in the [Reducing Witness Size with Verkle Trees](https://www.notion.so/8d531abcc2984babb9bf76a44459e611#8e4d71f24aaf468d9ee77ce16ad53f9d) section). Beam sync uses the [Ethereum Witness Protocol (wit)](https://github.com/ethereum/devp2p/blob/master/caps/wit.md) to transmit these witnesses. With beam sync, a node is supposed to save the state it gets from the network, filling up its state tries little by little. However, we need to consider that some people might create [fully stateless clients](https://github.com/ethereum/devp2p/blob/master/caps/wit.md) using this API (though it's not quite clear why they would chose to validate the blockchain, but refuse to store a few 100 GBs of recent state). Unfortunately, so far beam sync seems to be only implement by the (alpha-stage) [trinity client](https://github.com/ethereum/trinity) and the [Nethermind client](https://github.com/NethermindEth/nethermind). For implementing state expiry, we will need something like the wit protocol, in order to be able to retrieve expired state on-demand.

Let's briefly touch on something closely related with state distribution — decentralized light clients.

Currently, most "end-user clients" need to trust some well-known peer. In reality, these clients are not even real nodes — they do not participate in the Ethereum network protocols, nor keep even a superficial view of the blockchain state. The typical Ethereum wallet is a client that talks to a server (either a wallet provider, or a generic API provider like [Infura](https://infura.io/)) over the [JSON-RPC protocol](https://eth.wiki/json-rpc/). The server does not have custody of your coins, but it provides all the information that you need to use your wallets to make transactions, including balances and gas fee estimates.

Ideally, we would like to have decentralized lightweight clients, which are "real" Ethereum nodes talking to other Ethereum nodes directly. In "The Winding Road To Functional Light Client" ([Part 1](https://snakecharmers.ethereum.org/the-winding-road-to-functional-light-clients/), [Part 2](https://snakecharmers.ethereum.org/the-winding-road-to-functional-light-clients-part-2/), [Part 3](https://snakecharmers.ethereum.org/the-winding-road-to-functional-light-clients-part-3/)), Piper Merriam explains why (decentralized) lightweight clients are desirable, what makes them difficult to pull off today, and what we could do to make them a reality. I'll attempt to summarize the main points here, as relevant to the current conversation, but the original is well worth reading for more context.

Strictly speaking, we do not need decentralized light clients. However (and here I'm injecting my own opinion), immediately capitulating on that point moves the community and the consensus away from the platonic ideal of the decentralized zero-trust blockchain. This is the first step down a path we might not want to walk.

Light clients (I'll drop the "decentralized" from now on) are hard to pull off because of the way the state is disseminated in Ethereum currently. [The Ethereum Wire Protocol (ETH)](https://github.com/ethereum/devp2p/blob/master/caps/eth.md), which is used for full and fast sync, expects every node to have all the state. If you can never answer any requests, you risk "losing your peering" — peers will stop querying you, and also stop answering your requests.

An alternative protocol geared specifically towards light clients — the [Light Ethereum Subprotocol (LES)](https://github.com/ethereum/devp2p/blob/master/caps/les.md) — has been proposed and implemented. However, its design is asymmetric: lightweight clients query data without contributing anything back, while running the servers that provide this data is potentially expensive and accrues no benefits. As a result, there are too few of them and the protocol is too unreliable to make use of.

So, the ETH protocol asks too much of light clients (storing the whole state), while the LES protocol has them simply leeching off altruistic nodes. This suggests a middle-of-the-road solution, where every light client keeps track of some state. The idea would be to implement as a [Kademlia distributed hash table (DHT)](https://en.wikipedia.org/wiki/Kademlia) (which is notably used for trackerless BitTorrent).

This [new protocol](https://github.com/ethereum/stateless-ethereum-specs/blob/master/state-network.md) could also be made more lightweight. In the ETH protocol, if you want to grab a value, you need to grab every intermediate node in the trie (because [GetNodeData](https://github.com/ethereum/devp2p/blob/master/caps/eth.md#getnodedata-0x0d) accepts node hashes, not key hashes)! The new state protocol could simply store (key, value) pairs along with their proofs, and supply them with a single request.

Because of the decentralized trustless nature of the endeavour, we'll need proofs to be attached to the state. It's interesting to note the similar to the [wit](https://github.com/ethereum/devp2p/blob/master/caps/wit.md) protocol — but remember wit only covers block-level proofs.

You might think that statelessness or state expiry might solve issues of state distribution to light clients. However, that isn't the case. Under a "stateless" regime, light clients will need to acquire witness for values to be accessed in their transactions (statelessness) or to be revived (state expiry), so the problem persists.

All of the above was written in the context of the current version of Ethereum (1.x). Work is currently ongoing to transition Ethereum to [proof-of-stake](https://norswap.com/blockchain-how/#proof-of-stake), and then to [sharding](https://ethereum.org/en/eth2/shard-chains/). This has a certain number of consequences for state handling.

In [Why it's so important to go stateless](https://dankradfeist.de/ethereum/2021/02/14/why-stateless.html), Dankrad Feist argues for weak statelessness (as opposed to state expiry, which is the current direction). His main argument is that as miners get replaced with block proposers and validators in proof-of-stake, with weak statelessness only proposers would have to keep the state. This is important because it's mostly validators that ensure the chain security, so we'd like to have a large and diverse set of validators. My own understanding of this is that if validating on your own hardware (or even on a cloud provider) is difficult, it means that people are more likely to delegate validation to staking pools, increasing centralization (and giving custody of their keys to boot).

Interestingly, Dankrad dismisses the fact that nodes will have to acquire and attach witnesses to their transactions (even if proposers have the full state — nodes [validate account nonces and balances](https://github.com/ethereum/go-ethereum/blob/9e23610b0f49c5b9b173825d74507461452741cf/core/tx_pool.go#L596-L604) before propagating transactions [via gossip](https://github.com/ethereum/devp2p/blob/master/caps/eth.md#newpooledtransactionhashes-0x08)): "Creating these witnesses is so cheap that it’s unimaginable that there won’t be a plethora of services offering it." Clearly, this is quite different from the thinking around the portal network. Not everything is consensual!

Let's go back to state expiry proper. All proposals must contend with a thorny problem: how do you distinguish new state as opposed to stated that was written to before, but expired.

There are two main approaches to tackle the issue. First, we can propose some form of merge semantics so that if expired state is revived, and the state was written while expired, the resulting value is well-defined. This turns out to be rather tricky. The issue is that smart contracts are accounts too. The contract doesn't itself know if state is expired or not — it is the transaction's originator responsibility to revive expired pieces of state. The risk then is that a transaction originator would not supply witnesses to revive old state, which could lead to (expired or current) state being corrupted because of a "missing" value that was in fact expired. Another issue is that a user must not be allowed to revive an old version of the state that is not also the most recent version.

The second approach is to add "stubs" at the location of the expired storage. But this can at most reduce the amount of storage required by a constant factor: state consumption is still O(n) of total storage locations ever accessed, and the internal nodes in the Merkle tree have to be preserved.

The [proposed solution](https://ethresear.ch/t/resurrection-conflict-minimized-state-bounding-take-2/8739) (see also [the draft EIP](https://notes.ethereum.org/@vbuterin/state_expiry_eip)) is quite ingenious. It relies in dividing times into epochs of approximately one year, and on [extending account addresses](https://ethereum-magicians.org/t/increasing-address-size-from-20-to-32-bytes/5485) ([draft EIP](https://notes.ethereum.org/@ipsilon/address-space-extension-exploration)) with an epoch counter. Conceptually, there would be one version of the state tree per epoch (actually the state tree + the storage trees, but let's keep things brief). Nodes would have to keep the state tree for the current epoch (n) and the previous epoch (n-1).

If an account with epoch x tries to write a value in its storage, it can do so immediately if x = n. If x < n and the value was last modified in epoch y ≥ x, then the account must send a witness for the value at epoch y (or sometimes for its absence if x = y) and for its absence at all epochs y+1, ..., n-1. We could allow omitting the witness for n - 1, and if y = n, then no witnesses are required.

Reading works in the same way. The only significant difference is that the proven most recent value will be copied in the n tree, effectively reviving it.

In this scheme, reviving is implicit (it does not need specific EVM instructions), though not transparent as witnesses need to be provided. The "rent" is charged in the form of a per-witness gas cost. There is a [draft EIP](https://notes.ethereum.org/@vbuterin/witness_gas_cost_2) specifying these costs.

The scheme also has the advantage of being backward compatible, as all existing storage location will have epoch 0. However, one undesirable consequence is that accounts (including contracts) with epoch x will incur higher and higher costs to write an entirely new storage slot as time passes. One solution [discussed here](https://ethresear.ch/t/resurrection-conflict-minimized-state-bounding-take-2/8739/13) is to allow contracts to create storage space in the new epochs.

So, in the brave new world of expiring-state Ethereum, state expires and can be resurrected by clients. New full nodes can start validating directly by using the beam protocol, and light clients are made possible through the portal network.

All these things rely on the availability of state proofs — cryptographic proofs that show that a key-value pair is indeed part of the canonical state.

There's is one issue here, and it's that when using Merkle tree, state proofs get pretty large.

How large? To prove that a value belongs to the Merkle tree with a given root hash, we need to be able to recompute the root hash given the value. This requires not only the value, but also the hash of every sibling of every internal node that is an ancestor of the value. Ethereum uses hexary (16-branching) trees, which makes the max depth 64 (for a 256 bits key). This means that in the worst case, a proof will include 15\*64=960 hashes of 32 bytes (256 bits) each — 30 kB for each 256 bits value! (In general the proof is O(b\*log\_b(n))where b is the branching factor (16) and n is the size of tree.) This worst case is unrealistic (because the tree is sparse in practice), but the problem remains that the overhead is prohibitively large.

Let's do some back of the envelope calculations. Vitalik [asserts](https://vitalik.ca/general/2021/06/18/verkle.html) that 256 (2^8) random values will saturate the top level of a tree and fill about 63% of the second level. There are currently 83M accounts in the world state tree, so > 2^26. This would mean we could expect to saturate 1.63\*4~=7 levels and get on the order of 7\*15=105 hashes for a total of 3.2kB. And sure enough, [this states](https://notes.ethereum.org/@vbuterin/verkle_tree_eip): "A witness accessing an account in today’s hexary Patricia tree is, in the average case, close to 3 kB, and in the worst case it may be three times larger." Another way to get to the same number would be to take log16(80M) = 6.56, since a complete n-ary tree has height log\_n(size) .

![](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fc3d98732-d0da-4c6e-b68e-74619d182129%2Fverkle2.png?table=block&id=57f2c59e-861d-40b6-ba3f-c4a919cd5fcf&spaceId=69370dc8-c714-413c-9049-99172496c74d&width=1720&userId=35e2b552-8057-4dd9-a1dd-dbaff01abd7c&cache=v2)

In red: the nodes whose hash must be supplied as part of a Merkle proof for the (4ce, HORSE) pair. ([Source](https://vitalik.ca/general/2021/06/18/verkle.html)

One solution that was conceived was to use binary trees instead of hexary. This would make the worst-case proof size 255 hashes (8kB), almost a four-fold reduction. Using the same computation as above, the average size in the 80M account world state tree of today should be closer to 26 hashes (832B) — very reasonable!

But an even better solution was found in the form of Verkle trees (which [stands for](https://news.ycombinator.com/item?id=27554404) V(ector Commitment Mer)kle Tree). They were introduced in 2018, and [that paper](https://math.mit.edu/research/highschool/primes/materials/2018/Kuszmaul.pdf) is very readable.

Vitalik Buterin wrote [an excellent article](https://vitalik.ca/general/2021/06/18/verkle.html) explaining the problem and Verkle trees. I'll cover more or less the same grounds below, but I think you really ought that article as well (it's short enough).

Verkle trees rely on elliptic curve cryptography, and in particular on a construct called KZG or Kate ("kah-tay") commitments. I don't master this, but I will offer a brief summary of what Kate commitments offer. If you want to learn more about it yourself, here is a trail of articles that introduces the many layers necessary to grasp Kate commitments:

In brief, a Kate commitment is a "vector commitment" which means it allows to "commit" to a list of values a0, a1, ... aN. Committing means that you can send "a commitment" that does not reveal the values, but allows you to later show the values and prove that the commitment was created using those values, and none others.

Merkle trees can be used to implement a vector commitment: just stuff all the values in the leaves of a Merkle tree, then send the Merkle root (the hash of the root of the tree) as commitment. Because we're using [cryptographic hash functions](https://en.wikipedia.org/wiki/Cryptographic_hash_function), it's computationally infeasible to find a Merkle root that can be derived from two different value lists.

Once we have a commitment, we can make a proof for a value in that commitment. For Merkle trees, this is the Merkle proof we describe above.

Kate commitments work with elliptic curve cryptography, and in particular allow to commit to a polynomial function (i.e. commit to a function f, such that we then later send proofs that f(x) = y. We can build vector commitments on top of polynomial commitments. This involves building a polynomial of degree N to commit to a vector of size N.

Just like cryptographic hashes, it's also computationally infeasible to find another list of values that yields the same commitment. However, unlike cryptographic hashes, Kate commitments [are not always good at hiding the polynomial](https://dankradfeist.de/ethereum/2020/06/16/kate-polynomial-commitments.html) (and hence the list of values) used to construct the commitment. In Ethereum, we don't care, because these values are public state anyway.

The great thing is Kate proofs are much smaller than Merkle proofs, and notably their size is independent from the degree of the polynomial / size of the vector.

Even more impressive is the ability to combine multiple proofs into one, constant-size proofs. The math is out of scope here, but it involves building an "interpolation polynomial".

A Verkle tree is a trie just like a Merkle tree, but its inner nodes are labelled with a vector commitment for its children.

In reality, the vector commitment for a node commits to the hash of the vector commitment of the children. I'm assuming this makes the hash computation a bit more regular and maybe adds a bit of [collision resistance](https://en.wikipedia.org/wiki/Collision_resistance). However, you cannot verify a Kate proof using such a hash — you need the original commitment.

As such, a Verkle proof for a key-value pair consists of one Kate proof for each internal node on the path to the leaf that contains the value. Because these are constant size, the size of the value proof is proportional to the depth of the key in the tree (i.e. O(log\_b(n))). But because the branching factor does not impact the size of the proof, we can have a very large branching factor to minimize depth. In [Verkle trie for Eth1 state](https://dankradfeist.de/ethereum/2021/06/18/verkle-trie-for-eth1.html), Dankrad Feist proposes b=1024, which leads to an average depth of 3.67 for a tree with size n=2^30 inserted randomly.

I was a bit confused by the use of the term root to refer to the Kate commitment labelling each node. I think this is by analogy to Merkle root, where the hash of each internal node is "the Merkle root" of the subtree rooted at that node.

![](https://www.notion.so/image/https%3A%2F%2Fs3-us-west-2.amazonaws.com%2Fsecure.notion-static.com%2F164d8cc0-d5d7-4ab2-a596-f70c52938f65%2Fverkle3.png?table=block&id=974ae6d4-abdf-4c26-a5ce-a67eb4afe639&spaceId=69370dc8-c714-413c-9049-99172496c74d&width=1720&userId=35e2b552-8057-4dd9-a1dd-dbaff01abd7c&cache=v2)

In red: nodes for which a proof must be supplied as part of the Verkle proof for the (4ce, HORSE) pair. ([Source](https://vitalik.ca/general/2021/06/18/verkle.html))

But actually, we can even roll all these proofs into one. Using [this proposed scheme](https://dankradfeist.de/ethereum/2021/06/18/pcs-multiproofs.html), a proof for a single key-value pair takes 128 bytes, independently of its depth, the size of the tree, collisions etc. Remember this was ~3kB for hexary Merkle trees, ~832B for binary Merkley trees. We do however need to send the commitments (32B) for each node in the path (the verifier of the proof potentially only has the Verkle root: each Kate proof in the Verkle proof proves that a node is the child of another node on the path. This proof is done against the commitment for the node (which we supply). Because (a) the verifier has the commitment for the root, and this commitment commits to the hash of the commitment of its children and (b) it's computationally impossible to find a collision on the hash of the child commitment, then the child must be in the tree. We can apply the same reasoning recursively for every pair of nodes on the path until we reach the value.

So for a tree with 2^30 items and a branching factor of 256, we expect a single key-value proof in today's 80M accounts world-state tree to require on average 128 + log256(80M)\*32 = 230B. Remember that the hexary Merkle proof required on average 3kB while the binary Merkle proof required 832B in hashes. And both of these came with significantly higher worst-cases due to the much smaller branching factor.

And that's not all! We can even merge proofs for different key-value pairs together and save more space. The idea is that if multiple key-value pairs share some nodes on their path, there is no need to duplicate the proof for these nodes.

Because of locality effect, it would be particularly interesting to store all account related data in a single tree. Currently the account, nonce and code live in the "world state tree"; while the rest of the account data (that transactions can write to) lives in the "account storage trees" — see [this illustration](https://raw.githubusercontent.com/4c656554/BlockchainIllustrations/master/Ethereum/EthBlockchain5.png).

find about ongoing implementation efforts, e.g. verkle trees & portal network protocols

identify state of research / implementatin

Implement wit/0 and beam sync in other clients

find out how much state is active

list available Verkle tree implementations
