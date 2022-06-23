---
created: 2022-02-25T13:31:28 (UTC -08:00)
tags: []
source: https://notes.ethereum.org/@vbuterin/SkeyEI3xv
author: 
---

# Vitalik’s Annotated Ethereum 2.0 Spec - HackMD

> ## Excerpt
> # Vitalik's Annotated Ethereum 2.0 Spec  **Notice**: This document was written in July-Aug 2020.  ##

---
**Notice**: This document was written in July-Aug 2020.

## [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Introduction "Introduction")Introduction

Ethereum 2.0 (aka eth2, aka Serenity) is the next major version of the Ethereum protocol, and is the culmination into [years](https://blog.ethereum.org/2014/10/21/scalability-part-2-hypercubes/) [of](https://github.com/vbuterin/scalability_paper/blob/master/scalability.pdf) [research](https://cdn.hackaday.io/files/10879465447136/Mauve%20Paper%20Vitalik.pdf) into [proof](https://blog.ethereum.org/2014/01/15/slasher-a-punitive-proof-of-stake-algorithm/) [of](https://blog.ethereum.org/2014/11/25/proof-stake-learned-love-weak-subjectivity/) [stake](https://medium.com/@VitalikButerin/minimal-slashing-conditions-20f0b500fc6c) and [sharding](https://ethresear.ch/t/a-proposal-for-structuring-committees-cross-links-etc/2118). The eth2 protocol is a full redesign of the consensus-critical parts of the Ethereum system, with a change of the consensus from proof of work to [proof of stake](https://eth.wiki/en/concepts/proof-of-stake-faqs) and the introduction of [sharding](https://eth.wiki/sharding/Sharding-FAQs) being the two most critical changes. As of the time of this writing, eth2 leaves the application-layer parts maximally untouched; that is, transactions and smart contracts would continue to work the same way that they did before, so applications would not have to change (except to compensate for a few gas cost changes) to be eth2 compatible. However, the engine that ensures that the network comes to consensus on the transactions is radically changed.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#What-are-proof-of-stake-and-sharding-and-why-do-they-matter "What-are-proof-of-stake-and-sharding-and-why-do-they-matter")What are proof of stake and sharding and why do they matter?

For long answers, see:

-   [Proof of stake FAQ](https://eth.wiki/en/concepts/proof-of-stake-faqs)
-   [Sharding FAQ](https://eth.wiki/sharding/Sharding-FAQs)

But in short:

-   **Proof of stake** is a more efficient consensus mechanism that avoids the need to consume extremely large amounts of electricity and hardware costs to maintain the network by using coins, rather than computing hardware, as the economic basis for the consensus
-   **Sharding** is a scalability technique that involves splitting up verification so that each node in the network only needs to verify a small portion of the transactions in the network, instead of verifying every transaction

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#How-does-eth2-sharding-work "How-does-eth2-sharding-work")How does eth2 sharding work?

In the current (eth1) architecture, the blockchain is simply a chain of blocks, with each block pointing to the previous block. Each block also contains transactions, and the history is just all of these transactions, in the order in which they appear in the chain.

![](https://i.stack.imgur.com/ukuq0.png)

In eth2, we instead have a **beacon chain** and N (currently 64) **shard chains**. The beacon chain is a central chain which everyone stores, downloads and verifies, and which functions roughly like an eth1 chain (except using proof of stake instead of proof of work). The beacon chain contains two kinds of information:

-   Consensus-critical information (eg. who the current proof of stake validators are, what rewards and penalties have been assigned to them, deposits and withdrawals…)
-   Pointers (ie. hashes) to shard chain blocks and shard state

Each block in each shard chain is only fully downloaded and verified by a small portion of nodes, and it’s the shard chains that contain all of the user-level transactions (except proof of stake-related functions). Users can choose which shard they publish transactions or contracts to. It is possible to move coins or contracts between shards, but only asynchronously (meaning, with a delay of 1 slot). The shard chains and the beacon chain are tightly coupled with each other, connected through hash-linking and crosslinks:

![](https://storage.googleapis.com/ethereum-hackmd/upload_3f4bade5829953a6bc0cfabbc7347673.png)

For a shard block to be considered part of the “canonical history” it must be (i) valid and (ii) referenced in the beacon chain via a **crosslink**. A crosslink is a set of signatures of a shard block, signed by a randomly selected **committee** of ~128 **validators** (consensus-participating PoS nodes) which attests to the shard block’s validity.

**NB: if the beacon chain ends up linking to an invalid shard block (this can only realistically happen in a 51% attack or similar extreme scenario), that beacon chain is invalid. Invalid transactions (or state transitions) by definition can never be part of the canonical history.**

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#How-does-eth2-proof-of-stake-work "How-does-eth2-proof-of-stake-work")How does eth2 proof of stake work?

The participants in eth2 proof of stake consensus are called **validators**. To become a validator, you need to deposit 32 ETH, either from the eth1 chain, or from a shard chain (when shard chains become enabled). Once you deposit 32 ETH, you are put into an **activation queue**, and some time later you become an **active validator**.

Time on the beacon chain is divided into **epochs** and **slots**. Each slot is 12 seconds long (eg. if the chain started today at 14:00:00, and the current time is 14:01:06, then you’re in the middle of slot 5; slot numbers start counting from 0). An epoch is a period of 32 slots (or 6.4 minutes). There are also longer units of time; by convention, a period of 2048 epochs (~9.1 days) is called an **eek** (“ethereum week”); some operations on the beacon chain that take a long time can be measured in eeks.

During each epoch, each validator makes an **attestation**. An attestation contains:

-   A hash of what the validator thinks is the “**head**” of the beacon chain (the same concept as the “head” of the chain in PoW)
-   A hash of what the validator thinks is the correct shard block to include
-   Some other hashes (the “**source**” and “**target**” in [Casper FFG](https://arxiv.org/abs/1710.09437))
-   A signature from the validator, proving that the validator endorses all of the above

The chain comes to consensus as a result of these attestations. Roughly speaking, if 2/3 of active validators sign an attestation supporting a block, that block becomes **finalized** (in reality it’s more complex, taking two rounds of signing; see [the Casper FFG paper](https://arxiv.org/abs/1710.09437) for details). A finalized block can never be reverted, unlike in PoW where any block can be reverted if someone creates an even longer chain.

If a validator correctly makes attestations, they get rewarded. If a validator misses their slot, or makes an incorrect attestation, they get penalized. If a validator unambiguously contradicts themselves (eg. voting for two conflicting blocks in the same epoch), they get **slashed**. A slashed validator (i) suffers a penalty of 3-100% of their deposit, (ii) is forcibly ejected from the validator set, and (iii) has their coins forcibly locked for an additional 4 eeks before they can withdraw. In general, as long as you are running correct validator software, this will not happen to you, and as long as you stay online more than ~55-70% of the time, validating will be profitable.

A validator can voluntarily initiate an **exit** at any time, though there is a limit on how many can exit per epoch, and if too many validators try to exit at the same time they will be put into a queue, and will remain active until they get to the front of the queue. After a validator successfully exits, after 1/8 of an eek they will be able to withdraw (though this functionality will only be turned off after [“the merge”](https://ethresear.ch/t/the-eth1-eth2-transition/6265) of eth1 and eth2).

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Phases "Phases")Phases

To reduce risks, the eth2 deployment process is split into phases:

-   **Phase 0**: a proof of stake beacon chain is launched, which can process deposits from the eth1 chain. No shard chains yet.
-   **Phase 1**: shard chains are activated, though they only process data, not transactions. Hence, the eth2 chain becomes useful as a data availability verification layer for rollups, but you cannot yet directly transact on it.
-   **Phase 1.5 (aka The Merge)**: the eth1 chain is shut down, and the state (ie. all balances, contracts, code, storage…) from the eth1 chain is moved into one shard of the eth2 chain. Users can transact on that shard.
-   **Phase 2**: all shards support full transaction capabilities.
-   **Phase 3+ (aka Ethereum 2.x)**: ongoing improvements to safety, efficiency and scalability, though the “basic structure” of the eth2 chain may well be never again changed.

See also [the roadmap](https://media.consensys.net/an-annotated-version-of-vitalik-buterins-ethereum-roadmap-5876498d4f3a):

![](https://storage.googleapis.com/ethereum-hackmd/upload_fae4498237fede550e6022528fd0f00e.jpeg)

## [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Organization-and-type-system "Organization-and-type-system")Organization and type system

The spec describes the data types used in eth2, as well as the state transition function (along with helper functions used in defining the state transition function). The functions are written in python to strike a balance between formality (the spec is fully executable) and ease of reading. All functions are strongly typed, with many distinct types to clarify what the inputs and outputs represent.

The type system is based on [SimpleSerialize](https://github.com/ethereum/eth2.0-specs/blob/dev/ssz/simple-serialize.md) (SSZ), a system of types including both simple types and compound types (vectors, structs…) along with algorithms for (i) serializing and (ii) Merkle-hashing objects of these types (the SSZ Merkle-hash of an object is often called the **root**). SSZ is intended to largely take over from RLP in eth1.

The most common basic types in SSZ are integers (usually `uint64`) and hashes (aka `Bytes32`); other rarer types like `bool`, variable-length `Bytes`, `Bytes4` and others also exist. There are four compound types in SSZ: (i) fixed-length lists (called **vectors**), (ii) variable-length lists though with a fixed maximum length (called **lists**), (iii) structs (called **containers**) and (iv) union types (“either an X or a Y”), not used in phase 0.

Now, let’s get to the actual spec… (phase 0)

## [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Custom-types "Custom-types")Custom types

We define the following Python custom types for type hinting and readability:

| Name | SSZ equivalent | Description |
| --- | --- | --- |
| `Slot` | `uint64` | a slot number |
| `Epoch` | `uint64` | an epoch number (generally, epoch `i` consists of slots `EPOCH_LENGTH*i ... EPOCH_LENGTH*(i+1)-1`) |
| `CommitteeIndex` | `uint64` | During every epoch, the validator set is randomly split up into `EPOCH_LENGTH` parts, one part for each slot in that epoch, but then within each slot that slot’s validators are further divided into committees. In phase 0 this division does nothing, but in phase 1 these different committees get assigned to different shard. `CommitteeIndex` is just the type of an integer when that integer refers to the index of a committee within a slot (is it the first committee, the second, the third?) |
| `ValidatorIndex` | `uint64` | Every validator is assigned a validator index upon depositing |
| `Gwei` | `uint64` | An amount in Gwei |
| `Root` | `Bytes32` | A Merkle root (typically of an SSZ object) |
| `Version` | `Bytes4` | A fork version number (this is used to prevent messages on one eth2 network from accidentally being valid on another eth2 network, eg. mainnet vs testnet or mainnet vs ETC-like fork chain) |
| `DomainType` | `Bytes4` | A domain type (different signed messages are given different domain tags to prevent messages signed for one function from accidentally being valid in another function) |
| `ForkDigest` | `Bytes4` | A digest of the current fork data (used for replay protection) |
| `Domain` | `Bytes32` | A signature domain (combines info from a domain type and a fork version, so we get replay protection along both dimensions) |
| `BLSPubkey` | `Bytes48` | a BLS12-381 public key (see [here](https://ethresear.ch/t/pragmatic-signature-aggregation-with-bls/2105) for an explanation of the BLS signature scheme and its benefits) |
| `BLSSignature` | `Bytes96` | a BLS12-381 signature |

When you see a function like `def get_block_root_at_slot(state: BeaconState, slot: Slot) -> Root:` (a real example in the spec), interpret it as “this function takes as input the beacon chain state and an integer representing a slot number, and outputs a Bytes32 which is a Merkle root”. In this case the Merkle root it outputs is the root hash of the block at the given slot (as you can tell from the name); but in general, paying attention to types will help make it easier for you to understand what’s going on. In addition to being a debugging aid, the strong type system also functions as a type of comment.

## [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Constants "Constants")Constants

The following values are (non-configurable) constants used throughout the specification. These constants are fairly boring; they’re just added to the spec for readability.

| Name | Value |
| --- | --- |
| `GENESIS_SLOT` | `Slot(0)` |
| `GENESIS_EPOCH` | `Epoch(0)` |
| `FAR_FUTURE_EPOCH` | `Epoch(2**64 - 1)` |
| `BASE_REWARDS_PER_EPOCH` | `uint64(4)` |
| `DEPOSIT_CONTRACT_TREE_DEPTH` | `uint64(2**5)` (= 32) |
| `JUSTIFICATION_BITS_LENGTH` | `uint64(4)` |
| `ENDIANNESS` | `'little'` |

## [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Configuration "Configuration")Configuration

Here, we have constants that are configurable, ie. if you adjust one of these up or down by 2x or even more, the network is likely not going to break. That said, a lot of thought went into setting these constants the way they are now, so it’s better to learn the reasoning for why each of these values are set as they are.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Misc "Misc")Misc

| `ETH1_FOLLOW_DISTANCE` | `uint64(2**10)` (= 1,024) |
| --- | --- |

To process eth1 deposits, the eth2 chain tracks block hashes of the eth1 chain. To simplify things, the eth2 chain only pays attention to eth1 blocks after a delay (`ETH1_FOLLOW_DISTANCE = 1,024` blocks). Assuming that the eth1 chain does not revert that far, this lets us rely on an assumption that if the eth2 sees an eth1 block it won’t “un-see” it (if eth1 does revert that far, emergency action will be required on the eth2 side). 1024 blocks corresponds to a delay of ~3.7 hours (and note that getting an eth1 block _accepted_ into eth2 would take another ~1.7 hours). Historically, all problems on the eth1 net have been responded to within this period of time. Pushing this time to be even longer would (i) increase deposit delays and (ii) make eth2 less convenient as a light client of eth1.

| `MAX_COMMITTEES_PER_SLOT` | `uint64(2**6)` (= 64) |
| --- | --- |

In phase 0, the whole idea of having multiple committees per slot serves no function; rather, this is preparatory work for phase 1, where each committee will be assigned to a different shard. We plan to have 64 shards at the start. Having fewer shards would lead to insufficient scalability; having more would lead to two undesirable consequences:

1.  Overhead of processing beacon chain blocks being too high
2.  The minimum amount of ETH needed to reach a full-sized committee for every shard in every slot (now 32 ETH \* 128 committee size \* 64 shards per slot \* 32 slots per epoch = 8,388,608 ETH) to be too high; we’re reasonably confident we can get 8.3m ETH staking, but getting 16.7m ETH staking would be harder, and if we can’t get that much, the system would be forced compromise by making cross-shard transactions take longer.

| `TARGET_COMMITTEE_SIZE` | `uint64(2**7)` (= 128) |
| --- | --- |

For a committee to be secure, the chance that 2/3 of it get corrupted in any given epoch (assuming <1/3 of the global validator set is made up of attackers) must be astronomically tiny. We can estimate this chance of corruption via binomial formulas:

```
>>> # Factorial
>>> def fac(n): return 1 if n==0 else n*fac(n-1)
>>> # How many distinct k-item combinations you can take from n items
>>> def choose(n, k): return fac(n) // fac(k) // fac(n-k)
>>> # If an event has chance p of occurring during each "trial", returns
>>> # the probability that in n trials the event will occur *exactly* k times
>>> def prob(n, k, p): return p**k * (1-p)**(n-k) * choose(n, k)
>>> # If an event has chance p of occurring during each "trial", returns
>>> # the probability that in n trials the event will occur *at least* k times
>>> def probge(n, k, p): return sum([prob(n, i, p) for i in range(k, n+1)])
```

Calling `probge(128, 86, 1/3)` (86 is the smallest integer above 128 \* 2/3) returns `5.55 * 10**-15` (ie. 5.55 in a quadrillion). This is an extremely low probability, with comfortable bounds to take into account the possibility an attacker will “grind” many random seeds to try to get a favorable committee (though this is extremely difficult with RANDAO and especially VDFs). If the committee size were 64, it would no longer sufficiently secure. Increasing it to 256 would add needless inefficiency.

| `MAX_VALIDATORS_PER_COMMITTEE` | `uint64(2**11)` (= 2,048) |
| --- | --- |

The maximum supported validator count is `2**22` (=4,194,304), or ~134 million ETH staking. Assuming 32 slots per epoch and 64 committees per slot, this gets us to a max 2048 validators in a committee.

| `MIN_PER_EPOCH_CHURN_LIMIT` | `uint64(2**2)` (= 4) |
| --- | --- |
| `CHURN_LIMIT_QUOTIENT` | `uint64(2**16)` (= 65,536) |

These two parameters set the rate at which validators can enter and leave the validator set. The minimum rate is 4 entering + 4 leaving per epoch, but if there are enough validators this rate increases: if there are more than 262,144 validators (8,388,608 ETH) then an amount of validators equal to 1/65536 of the validator set size can enter, and the same amount can leave, per epoch.

The goal of rate-limiting entry and exit is to prevent a large portion of malicious validators from performing some malicious action and then immediately leaving to escape being slashed. The main malicious action we are worried about is finalizing two incompatible blocks. The Casper FFG protocol (see paper [here](https://arxiv.org/abs/1710.09437)) ensures that this can only happen if at least 1/3 of validators commit a provably malicious action, which they can be slashed for; however, if they withdraw first they could conceivably dodge this penalty.

With the above numbers, if there is more than 8,388,608 ETH staking, it will take at least 65536/3 epochs, or 10.67 eeks, for 1/3 of validators to withdraw (however, if there is no attack, then the withdrawal queue will ordinarily be short).

The reason to have a long withdrawal delay is to ensure that even a fork that gets hidden from users for a long time and then published to clients that have been offline for some time would lead to slashing. Conceivably, an attacker could escape slashing by hiding a fork for longer than 10.67 eeks; for this reason, we have a rule that clients must go online at least once in that long (in reality a bit less than that) to retain their full security guarantees (this is called **weak subjectivity**).

Research:

-   Original post on weak subjectivity (2014): [https://blog.ethereum.org/2014/11/25/proof-stake-learned-love-weak-subjectivity/](https://blog.ethereum.org/2014/11/25/proof-stake-learned-love-weak-subjectivity/)
-   Why withdrawal queues are better than fixed withdrawal delays: [https://ethresear.ch/t/suggested-average-case-improvements-to-reduce-capital-costs-of-being-a-casper-validator/3844](https://ethresear.ch/t/suggested-average-case-improvements-to-reduce-capital-costs-of-being-a-casper-validator/3844)
-   Rate-limiting entry/exits, not withdrawals: [https://ethresear.ch/t/rate-limiting-entry-exits-not-withdrawals/4942](https://ethresear.ch/t/rate-limiting-entry-exits-not-withdrawals/4942)
-   Analyzing how long the de-facto weak subjectivity period is assuming the queue processes at a particular rate: [https://ethresear.ch/t/weak-subjectivity-under-the-exit-queue-model/5187](https://ethresear.ch/t/weak-subjectivity-under-the-exit-queue-model/5187)
-   Weak subjectivity in eth2 (by Aditya): [https://notes.ethereum.org/@adiasg/weak-subjectvity-eth2](https://notes.ethereum.org/@adiasg/weak-subjectvity-eth2)

| `SHUFFLE_ROUND_COUNT` | `uint64(90)` |
| --- | --- |

Number of rounds in the swap-or-not shuffle; for more into see the [`compute_shuffled_index` function description](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#compute_shuffled_index) below. Expert cryptographer advice told us `~4*log2(n)` is sufficient for safety; in our case, `n <= 2**22`, hence ~90 rounds.

| `MIN_GENESIS_ACTIVE_VALIDATOR_COUNT` | `uint64(2**14)` (= 16,384) |
| --- | --- |

Number of validators deposited needed to start the eth2 chain. This gives 524,288 ETH, high enough to put attacking out of the reach of all but a few very wealthy actors.

| `MIN_GENESIS_TIME` | `uint64(1578009600)` (Jan 3, 2020) |
| --- | --- |

Genesis will not start before this time, even if there are enough validators deposited.

| `HYSTERESIS_QUOTIENT` | `uint64(4)` |
| --- | --- |
| `HYSTERESIS_DOWNWARD_MULTIPLIER` | `uint64(1)` |
| `HYSTERESIS_UPWARD_MULTIPLIER` | `uint64(5)` |
| `EFFECTIVE_BALANCE_INCREMENT` | `Gwei(2**0 * 10**9)` (= 1,000,000,000) |

We store validator balances in two places: (i) the “effective balance” in the validator record, and (ii) the “exact balance” in a separate record. This is done for efficiency reasons: the exact balances get changed due to rewards and penalties in every epoch, so we store them in a compact array that requires rehashing only <32 MB to update. Effective balances (which are used for all other computations that require validator balances) get updated using a **[hysteresis](https://en.wikipedia.org/wiki/Hysteresis)** formula: if the effective balance is `n` ETH, then if the exact balance goes below `n-0.25` ETH, then the effective balance is set to `n-1` ETH, and if the exact balance goes above `n+1.25` ETH the effective balance is set to `n+1` ETH.

This ensures that an attacker can’t make effective balances update every epoch and thus cause processing the chain to become very slow by repeatedly nudging the exact balances above and then below some threshold; instead, the exact balance must change by at least a full 0.5 ETH to trigger an effective balance update.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Gwei-values "Gwei-values")Gwei values

| `MIN_DEPOSIT_AMOUNT` | `Gwei(2**0 * 10**9)` (= 1,000,000,000) |
| --- | --- |

A minimum deposit amount prevents DoS attacks that involve spamming the chain with very tiny deposits (note that 1 ETH just gets you a validator slot; it does not _activate_ unless you deposit the full 32 ETH).

| `MAX_EFFECTIVE_BALANCE` | `Gwei(2**5 * 10**9)` (= 32,000,000,000) |
| --- | --- |

There are two choices here that need justification. First, why force validator slots to be a fixed amount of ETH at all, instead of allowing them to be any size? Second, why a fixed size of 32 ETH, and not 1 ETH or 1000 ETH?

The problem with allowing variable balances is that algorithms for random selection (eg. of block proposers) and shuffling (for committees) become much more complicated. You would need an algorithm to select a block proposer such that the probability that the algorithm selects a particular proposer is proportional to the proposer’s balance, in a context where balances are changing and validators are always entering and exiting. This could be done with [fancy binary tree structures](https://algorithmist.com/wiki/Fenwick_tree) but would be complicated. In the case of committee selection, a wealthy validator cannot be assigned to one committee (as they would then dominate and be able to attack it); their weight would need to be split up among many committees. It’s much easier to solve both problems by simply formally representing wealthy validators as being many separate validators of the same size.

The 32 ETH choice is based on this logic: [https://medium.com/@VitalikButerin/parametrizing-casper-the-decentralization-finality-time-overhead-tradeoff-3f2011672735](https://medium.com/@VitalikButerin/parametrizing-casper-the-decentralization-finality-time-overhead-tradeoff-3f2011672735). If the deposit size is higher, then fewer people can participate, risking centralization, but if the deposit size is lower, then the chain suffers higher cost of verification (“overhead” in the post), risking sacrificing decentralization in a different way.

[Economic review](https://medium.com/@thomasborgers/ethereum-2-0-economic-review-1fc4a9b8c2d9) suggests that at the current 32 ETH level the hardware costs of staking are enough to make a significant, though not fatal, dent on validator returns. This implies that if the deposit size were reduced to 16 ETH, then the overhead of the chain would double, and the rewards to each validator would halve, so staking with a single validator slot would be four-times more difficult, already a potentially unsafe level. Hence, 32 ETH is the most inclusive deposit size that does not become self-defeating due to increasing overhead.

| `EJECTION_BALANCE` | `Gwei(2**4 * 10**9)` (= 16,000,000,000) |
| --- | --- |

Validators that go below 16 ETH get ejected (ie. forcibly exited). This minimum ensures that all active validators’ balances are (almost always) within a 2x “band” (maximum effective balance is 32 ETH; anything above is just saved rewards and does not count for staking purposes). This narrow range ensures that committees are stable; if higher disparities were permitted, there would be higher risk that a few wealthy malicious validators could randomly enter the same committee and take it over with their larger balances.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Initial-values "Initial-values")Initial values

| `GENESIS_FORK_VERSION` | `Version('0x00000000')` |
| --- | --- |

See explanation of the `Version` type [above](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Custom-types). The version starting at 0 is self-explanatory (why not 1? Because we’re computer scientists, not normal people, doge-dammit!).

| `BLS_WITHDRAWAL_PREFIX` | `Bytes1('0x00')` |
| --- | --- |

When a validator deposits, they provide two keys: a **signing key** and a **withdrawal key**. The withdrawal key is needed to access the funds when they are withdrawn; this dual-key structure reduces risks for validators, as they can keep their withdrawal key in cold storage.

The BLS withdrawal prefix is effectively a “version number” for the withdrawal key; the first version is just a hash of a public key; a withdrawal would reveal the public key as well as a signature signed with that public key specifying a further destination. Future versions will allow directly specifying a smart contract address on any shard of eth2.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Time-parameters "Time-parameters")Time parameters

| `GENESIS_DELAY` | `uint64(172800)` | seconds | 2 days |
| --- | --- | --- | --- |

When the deposit count becomes sufficient for the eth2 chain to start, the start is delayed by 2 days to give everyone time to prepare.

| `SECONDS_PER_SLOT` | `uint64(12)` | seconds | 12 seconds |
| --- | --- | --- | --- |

A tradeoff between blockchain speed and risk. Note that in future phases, multiple steps will have to happen within a slot: beacon block -> shard block -> beacon block, as well as eventually a round of data availability sampling, so it is good to be conservative.

![](https://storage.googleapis.com/ethereum-hackmd/upload_a2c119d0ba9d384203bcb079e33bfcd6.png)

Eth1 latency is generally ~1 second; 12 seconds gives a healthy safety margin on top of this.

| `SECONDS_PER_ETH1_BLOCK` | `uint64(14)` | seconds | 14 seconds |
| --- | --- | --- | --- |

An estimate of how often eth1 blocks appear on average.

| `MIN_ATTESTATION_INCLUSION_DELAY` | `uint64(2**0)` (= 1) | slots | 12 seconds |
| --- | --- | --- | --- |

Attestations made in slot N can be included in slot N+1.

| `SLOTS_PER_EPOCH` | `uint64(2**5)` (= 32) | slots | 6.4 minutes |
| --- | --- | --- | --- |

There are two reasons not to go lower than 32 slots per epoch:

1.  Either the slot duration would have to become longer (which would increase block times and hence reduce user experience) or the epoch duration would decrease, increasing overhead of processing the chain
2.  We want to have a guarantee that there will almost certainly be at least one honest proposer per epoch

Going higher than 32 would needlessly make it take longer for a block to reach finality (this takes 2 epochs). Hence, 32 slots per epoch appears optimal.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Aside-RANDAO-seeds-and-committee-generation "Aside-RANDAO-seeds-and-committee-generation")`[Aside: RANDAO, seeds and committee generation]`

In any proof of stake system, we need to have some mechanism for determining who is the proposer of a block (as well as other roles that don’t require all active validators to participate in simultaneously). In PoW, this happens automatically: everyone is trying to create a block, but on average only one person succeeds every (13 seconds in Ethereum | 600 seconds in Bitcoin), and you can’t predict who will succeed ahead of time. In PoS, however, this random selection must be done explicitly.

Clearly, there is no such thing as true “randomness” in a blockchain, because all nodes must come to consensus on the result and different computers calling `random()` will get different outputs. Instead, we generate **[pseudo-randomness](https://en.wikipedia.org/wiki/Pseudorandomness)** from a **seed** that is computed and updated as part of the blockchain.

The challenge is how to make the seed unpredictable. If the seed was predictable (say, we put `hash(42)` in as the seed at genesis), then validators could strategically decide when to deposit and withdraw or what their public keys is to target being a block proposer (or being part of a particular committee), which would open the door to attacks (this is one type of **stake grinding**).

To fully prevent this type of manipulation, we use a mechanism where the validator set is fixed 4 epochs in advance (ie. operations in epoch N can only affect the validator set from epoch N+5 onwards), and there is a procedure that constantly updates the seed. Hence, validator set manipulation is ineffective because the seed valid in some epoch can be guaranteed to update unpredictably for some time after the validator set for that epoch is fixed. That said, we still need a procedure for how the seed actually gets updated.

We update the seed (or rather, the **randao mix**, which is used to generate the seed) every block using a mechanism inspired by [RANDAO](https://github.com/randao/randao): the proposer of a block provides a hash that gets mixed into (ie. XOR’d into) the seed; this hash is unknown to the public ahead of time, but it is pre-committed, in the sense that there is only one valid hash that the proposer could submit. This is done with by BLS-signing the current epoch; the BLS signature scheme has the property that for any given key there is exactly one valid signature for any given message (as opposed to eg. ECDSA where there are many possible valid signatures that can be made with the same key for the same message).

![](https://storage.googleapis.com/ethereum-hackmd/upload_86d4ccefc8be1d92e2f0fb7be5efac80.png)

The randao mix at the start of epoch N is used to compute the seed for epoch N+1; this ensures that the proposer and committee roles are known one epoch ahead of time, giving validators a chance to prepare.

This mechanism ensures that the proposer has only one “bit of manipulation power”: they can either set the next seed to some hash R1 by publishing their block normally, or to some other hash R2 by not publishing (and sacrificing their block reward). Note also that only the last proposer truly has any manipulation power, because any other proposer knows that the seed will be altered in unpredictable ways by future proposers, so they have know way to know the effects of any manipulation they attempt. These two factors together make “stake grinding” _by manipulating the seed_ very difficult and almost always not worth it.

To see the proposer and committee selection algorithms, which take as input (i) the active validator set and (ii) the seed, and output the current block proposer and the committees, see [here](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#compute_shuffled_index).

| `MIN_SEED_LOOKAHEAD` | `uint64(2**0)` (= 1) | epochs | 6.4 minutes |
| --- | --- | --- | --- |
| `MAX_SEED_LOOKAHEAD` | `uint64(2**2)` (= 4) | epochs | 25.6 minutes |

See the diagram above. The `MIN_SEED_LOOKAHEAD` means that the seed used to compute proposers and committees is based on the randao mix from more than 1 epoch ago (specifically, the epoch N seed is based on the randao mix from the end of epoch N-2); this allows validators to determine their committee and proposal responsibilities >1 epoch ahead of time.

The `MAX_SEED_LOOKAHEAD` is actually the minimum delay on validator actiations and exits; it basically means that validators strategically activating and exiting can only affect the seed 4 epochs into the future, leaving a space of 3 epochs within which proposers can mix-in unknown info to scramble the seed and hence make stake grinding via activating or exiting validators non-viable.

| `MIN_EPOCHS_TO_INACTIVITY_PENALTY` | `uint64(2**2)` (= 4) | epochs | 25.6 minutes |
| --- | --- | --- | --- |

See [here](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#inactivity-quotient) for what the inactivity leak is; this constant simply says that the leaking starts after 4 epochs of non-finality.

| `EPOCHS_PER_ETH1_VOTING_PERIOD` | `uint64(2**5)` (= 32) | epochs | ~3.4 hours |
| --- | --- | --- | --- |

The eth2 chain learns about eth1 blocks (so that it can verify Merkle proofs of deposits) via a voting mechanism where block proposers vote on an eth1 block; the [honest validator guide](https://github.com/ethereum/eth2.0-specs/blob/dev/specs/phase0/validator.md#eth1-data) details which block validators choose. The voting period is set to 1024 slots to ensure a sufficient committee size and to give time to respond to potential failures; additionally, there is little value in the voting period being much shorter than the `ETH1_FOLLOW_DISTANCE` (also ~3.7 hours).

| `SLOTS_PER_HISTORICAL_ROOT` | `uint64(2**13)` (= 8,192) | slots | ~27 hours |
| --- | --- | --- | --- |

The eth2 chain contains a Merkle tree of its own historical blocks. This is done using two data structures: (i) a rotating “recent history log” (`state.block_roots` and `state.state_roots`) and (ii) an ongoing accumulator (`state.historical_roots`) which stores Merkle roots of recent history logs. The total state size is optimized if the two have roughly similar length, both `~sqrt(length of chain)`; setting this to 8192 slots ensures that this condition is satisfied at 67,108,864 slots (= 1,024 eeks, ~= 20 years). If desired, the lengths can be rebalanced after a century to improve efficiency, though the gains will be negligible.

| `MIN_VALIDATOR_WITHDRAWABILITY_DELAY` | `uint64(2**8)` (= 256) | epochs | ~27 hours |
| --- | --- | --- | --- |

Provides a reasonable amount of time to ensure that individual validators that misbehave can be slashed.

| `SHARD_COMMITTEE_PERIOD` | `uint64(2**8)` (= 256) | epochs | ~27 hours |
| --- | --- | --- | --- |

In phase 1, this is how often the proposer committees on a shard get reshuffled. In phase 0, validators are only allowed to exit if they have been active for at least this long; this prevents spamming the exit queue and repeatedly exiting and re-entering to force oneself into specific shards.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#State-list-lengths "State-list-lengths")State list lengths

| `EPOCHS_PER_HISTORICAL_VECTOR` | `uint64(2**16)` (= 65,536) | epochs | ~0.8 years |
| --- | --- | --- | --- |

How far back randomness seeds are visible; this is de-facto the maximum amount of time after which a validator can be slashed.

| `EPOCHS_PER_SLASHINGS_VECTOR` | `uint64(2**13)` (= 8,192) | epochs | ~36 days |
| --- | --- | --- | --- |

This is the minimum amount of time a validator must wait before they can withdraw if they are slashed; during that time, they get penalized an amount proportional to how many other validators get slashed in the same time period.

See the [section on slashings](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#slashings) for more details, and [here](https://notes.ethereum.org/@vbuterin/rkhCgQteN?type=view#Slashing-and-anti-correlation-penalties) for why this is done.

| `HISTORICAL_ROOTS_LIMIT` | `uint64(2**24)` (= 16,777,216) | historical roots | ~52,262 years |
| --- | --- | --- | --- |
| `VALIDATOR_REGISTRY_LIMIT` | `uint64(2**40)` (= 1,099,511,627,776) | validators |  |

All lists in SSZ have to have _some_ limit; 52,262 years is reasonably close to “forever” for practical purposes and ensures that Merkle branches do not get needlessly long. 1.1 trillion validators will also only be reached after a very long duration (assuming all ETH is staking, a maximum of 64 validators can be activated per epoch, so the list would take ~16 billion epochs ~= 209052 years to fill; this could be accelerated assuming clever use of 1 ETH validator slots, rewards, etc, but it would still take millennia).

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Rewards-and-penalties "Rewards-and-penalties")Rewards and penalties

| `BASE_REWARD_FACTOR` | `uint64(2**6)` (= 64) |
| --- | --- |

See `get_base_reward` in the section on [helpers](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#helpers) for details.

| `WHISTLEBLOWER_REWARD_QUOTIENT` | `uint64(2**9)` (= 512) |
| --- | --- |

If you submit evidence leading to a validator getting slashed, you get 1/512 of their balance as a reward.

| `PROPOSER_REWARD_QUOTIENT` | `uint64(2**3)` (= 8) |
| --- | --- |

As a general rule of thumb, the proposer of a block gets 1/8 of the rewards given to other validators in the block that they include. This ensures sufficient incentive to include attestations and other objects as well as incentive to produce blocks.

| `INACTIVITY_PENALTY_QUOTIENT` | `uint64(2**24)` (= 16,777,216) |
| --- | --- |

See the [Casper FFG paper](https://arxiv.org/abs/1710.09437) for a description of the inactivity leak, the mechanism by which if a chain fails to finalize, inactive validators start to suffer very high penalties until the validators that _are_ active get back up to above 2/3 of the total validator set (weighted by balance) and finalization can restart.

The size of the penalty during an epoch is proportional to the number of epochs that have passed since the most recent time the chain finalized; this leads to the total amount leaked growing _quadratically_ with time (note that the leak starts after 4 epochs of non-finality):

| Epochs since finality | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
| Leaked this epoch | 0 | 0 | 0 | 0 | 5 | 6 | 7 | 8 | 9 | 10 | 11 | 12 |
| Total leaked since finality | 0 | 0 | 0 | 0 | 5 | 11 | 18 | 26 | 35 | 45 | 56 | 68 |

`INACTIVITY_PENALTY_QUOTIENT` is the “unit” here, eg. if the total leaked is listed as 68, that means that you’ve lost `68/INACTIVITY_PENALTY_QUOTIENT ~= 1/246,723` of your balance.

Note that when the total leaked becomes a substantial fraction of deposit size, the amount leaked starts to decrease because it’s calculated as a percentage of current balance; hence, in those cases, the total leaked is best approximated as an exponential function: the portion of your original balance remaining is not `1 - 1/2 * epochs**2 / INACTIVITY_PENALTY_QUOTIENT`, but rather `(1 - 1/INACTIVITY_PENALTY_QUOTIENT) ** (epochs**2/2)`.

An alternative nearly-equivalent approximation of the remaining balance is `e ** -(epochs**2/(2*INACTIVITY_PENALTY_QUOTIENT))`, where `e ~= 2.71828`. This implies that after `2**12` epochs (2 eeks), the portion of your original balance remaining is `e**(-1/2)`, or roughly 60.6% of your original balance.

![](https://storage.googleapis.com/ethereum-hackmd/upload_54c8f64d00ade395640ad667ae1ccf40.png)

| `MIN_SLASHING_PENALTY_QUOTIENT` | `uint64(2**5)` (= 32) |
| --- | --- |

You lose at least 1/32 of your deposit if you get slashed (getting slashed has to hurt _some_ amount even if you’re the only one that got slashed…)

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Max-operations-per-block "Max-operations-per-block")Max operations per block

| Name | Value |
| --- | --- |
| `MAX_PROPOSER_SLASHINGS` | `2**4` (= 16) |
| `MAX_ATTESTER_SLASHINGS` | `2**1` (= 2) |
| `MAX_ATTESTATIONS` | `2**7` (= 128) |
| `MAX_DEPOSITS` | `2**4` (= 16) |
| `MAX_VOLUNTARY_EXITS` | `2**4` (= 16) |

These operations are set based on calculations of how many can be safely processed, though there is the additional constraint that `MAX_ATTESTATIONS` must equal the max number of committees (64) plus a safety margin to account for missed proposals or delayed or disagreeing attestations.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Domain-types "Domain-types")Domain types

| Name | Value |
| --- | --- |
| `DOMAIN_BEACON_PROPOSER` | `DomainType('0x00000000')` |
| `DOMAIN_BEACON_ATTESTER` | `DomainType('0x01000000')` |
| `DOMAIN_RANDAO` | `DomainType('0x02000000')` |
| `DOMAIN_DEPOSIT` | `DomainType('0x03000000')` |
| `DOMAIN_VOLUNTARY_EXIT` | `DomainType('0x04000000')` |
| `DOMAIN_SELECTION_PROOF` | `DomainType('0x05000000')` |
| `DOMAIN_AGGREGATE_AND_PROOF` | `DomainType('0x06000000')` |

These values are mixed into the messages of each type when those messages are being signed; this prevents messages signed for one purpose from being accidentally valid in another context.

## [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Containers "Containers")Containers

The following types are [SimpleSerialize (SSZ)](https://notes.ethereum.org/ssz/simple-serialize.md) containers.

_Note_: The definitions are ordered topologically to facilitate execution of the spec.

_Note_: If an object of some container type is initialized with some of its fields not set (eg. `x = Fork(epoch=6)`, with `previous_version` and `current_version` missing), those fields are set to zero (a zero container is of course recursively defined as a container with all fields set to zero).

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Misc-dependencies "Misc-dependencies")Misc dependencies

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Fork "Fork")`Fork`

```
class Fork(Container):
    previous_version: Version
    current_version: Version
    epoch: Epoch  # Epoch of latest fork
```

This struct exists in the state to store the current version of the eth2 protocol. The version number is changed when there is a hard fork: if some new hard fork rules are supposed to take place at epoch `N`, then as part of the epoch `N` state transition, `state.fork` is modified so that:

-   `state.fork.previous_version` equals the old `state.fork.current_version`
-   `state.fork.current_version` equals some newly chosen version number
-   `state.fork.epoch` equals `N`

The intention is that the “current fork version” equals `state.fork.previous_version` if the current epoch is less than `state.fork.epoch`, and `state.fork.current_version` if the current epoch is equal to or greater than `state.fork.epoch`. The current fork version is mixed into the signed data in all BLS-signed messages (see [`get_domain`](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#get_domain)).

Note that all messages (blocks, attestations, VoluntaryExits…) have some associated epoch number. Blocks are processed at the slot that is their declared slot, but attestations and other structures do have one edge case: an attestation could be crhe eated with some self-declared epoch `E1`, but only included on chain at some later epoch `E2 > E1`. The edge case is, what if `E1` is before the fork but `E2` is after it? Then, even though the message is processed during the new fork era, the message is validated assuming the signed data mixes in the old fork version. This is why we maintain `state.fork.previous_version` in the state.

If someone wants to continue the old chain, they can simply not implement the changes, including not changing `state.fork`. In this case, starting from the fork epoch, blocks from one fork will be invalid in the other fork. Attestations and other objects made before the fork could be included into both forks, but attestations and other objects made after the fork would only be valid either on one side or the other.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#ForkData "ForkData")`ForkData`

```
class ForkData(Container):
    current_version: Version
    genesis_validators_root: Root
```

This is a dummy struct that’s used to mix in not just the fork version, but also the genesis, in [`get_domain`](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#get_domain).

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Checkpoint "Checkpoint")`Checkpoint`

```
class Checkpoint(Container):
    epoch: Epoch
    root: Root
```

Ethereum’s Casper FFG implementation comes to consensus by dealing with epoch boundary hashes (that is, the most recent block hash in the chain before the start of an epoch). A Casper FFG vote normally consists of (i) source epoch, (ii) source block hash, (iii) target epoch, (iv) target block hash, and in the state we need to store the latest justified epoch (and hash) to know what source to accept, as well as the latest finalized epoch.

We simplify this by creating a container wrapper to represent the epoch and the hash, so we reduce to a Casper FFG vote containing two checkpoints (source and target), and the state storing the latest justified checkpoint of the current and previous epochs and the latest finalized checkpoint.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Validator "Validator")`Validator`

```
class Validator(Container):
    pubkey: BLSPubkey
    withdrawal_credentials: Bytes32  # Commitment to pubkey for withdrawals
    effective_balance: Gwei  # Balance at stake
    slashed: boolean
    # Status epochs
    activation_eligibility_epoch: Epoch  # When criteria for activation were met
    activation_epoch: Epoch
    exit_epoch: Epoch
    withdrawable_epoch: Epoch  # When validator can withdraw funds
```

This is the struct that contains all of the important information relating to a particular validator (except for its exact balance; see the [section on hysteresis above](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#hysteresis) for info on that). From top to bottom:

-   `pubkey`: public key for signing (ie. the “online staking key”)
-   `withdrawal_credentials`: a hash of the pubkey that will be used to withdraw (the private key can be kept in cold storage).
-   `effective_balance`: the balance of the validator that is used for all calculations (when counting total support for some attestation, when computing rewards and penalties, etc…)
-   `slashed`: has the validator been slashed?
-   `activation_eligibility_epoch`: when the validator became eligible for activation (this is used to process the activation queue: validators get activated in order of when they became eligible)
-   `activation_epoch`: when the validator was activated
-   `exit_epoch`: when the validator exited (whether voluntarily or due to low balance or slashing)
-   `withdrawable_epoch`: when the validator became eligible to withdraw their balance

The precise epochs are kept around in the state because we need to be able to not just calculate the current active validator set, but also the historical active validator set, so that we can compute historical committees and hence verify historical attestations and slashings.

Additionally, storing epochs for each phase transition simplifies the protocol. The alternative would be to store a variable `current_state` with flags (eg. 0 = not yet eligible for activation, 1 = eligible for activation, 2 = active, 3 = in exit queue, 4 = exited, 5 = withdrawn) as well as the epoch for the next transition, but that turns out to increase protocol complexity, as eg. currently (3), (4) and (5) are all handled by two lines of code (bottom of [`initiate_validator_exit`](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#initiate_validator_exit)), which would be difficult if they were separate.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#AttestationData "AttestationData")`AttestationData`

```
class AttestationData(Container):
    slot: Slot
    index: CommitteeIndex
    # LMD GHOST vote
    beacon_block_root: Root
    # FFG vote
    source: Checkpoint
    target: Checkpoint
```

For efficiency reasons, we require each validator to only sign a single attestation in each epoch. However, this attestation has a triple purpose: (i) Casper FFG voting, (ii) stabilizing the short-term block-by-block fork choice by voting on the current head, and (iii) shard block voting (added in phase 1). In each epoch, each validator is assigned to a single committee within a single slot, and they attest (ie. sign) this data structure that combines (i) (`source` and `target`) and (ii) (`beacon_block_root`).

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#IndexedAttestation "IndexedAttestation")`IndexedAttestation`

```
class IndexedAttestation(Container):
    attesting_indices: List[ValidatorIndex, MAX_VALIDATORS_PER_COMMITTEE]
    data: AttestationData
    signature: BLSSignature
```

An `AttestationData`, a signature and a list of indices that participated. This is the format of an attestation when it is included in `AttesterSlashing` objects to slash validators for misbehavior. The idea is that attestations imported from foreign chains may not share the same committees as the current chain, so we need to explicitly provide a list of which validators participated so that the attestation can be verified, and if needed the participants slashed.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#PendingAttestation "PendingAttestation")`PendingAttestation`

```
class PendingAttestation(Container):
    aggregation_bits: Bitlist[MAX_VALIDATORS_PER_COMMITTEE]
    data: AttestationData
    inclusion_delay: Slot
    proposer_index: ValidatorIndex
```

For efficiency reasons we do not process the full effects of attestations included in a block immediately; that would require editing 1/32 of all validators’ exact balances in the Merkle tree, which costs almost as much as simply re-hashing the entire vector. Additionally, it would require storing an extra data structure for “who already attested” to prevent double-counting. Instead, we simply store all attestations that we received in the state, minus their signatures (as those are no longer necessary), plus information about who included them and with what delay (to calculate rewards). These pending attestations are then processed at the end of an epoch.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Eth1Data "Eth1Data")`Eth1Data`

```
class Eth1Data(Container):
    deposit_root: Root
    deposit_count: uint64
    block_hash: Bytes32
```

Every eth2 block contains a vote for an eth1 block. This vote contains the hash of the eth1 block, and to allow more convenient verification of deposits, it also contains the root of the deposit tree and the number of deposits already made. Technically, the deposit tree root and size would be Merkle-proven from the eth1 block hash, but this would involve verifying a hexary RLP Patricia tree Merkle branch, which is needlessly complex.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#HistoricalBatch "HistoricalBatch")`HistoricalBatch`

```
class HistoricalBatch(Container):
    block_roots: Vector[Root, SLOTS_PER_HISTORICAL_ROOT]
    state_roots: Vector[Root, SLOTS_PER_HISTORICAL_ROOT]
```

See the discussion in the [`SLOTS_PER_HISTORICAL_ROOT`](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#slots_per_historical_root) section above.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Aside-note-on-the-deposit-process "Aside-note-on-the-deposit-process")`[Aside: note on the deposit process]`

The process for someone on eth1 to deposit to become a validator on eth2 is as follows:

1.  Generate an online signing key (the `pubkey`) and an offline withdrawal key (the hash of this is called the `withdrawal_credentials`).
2.  Send 32 ETH to the deposit contract by calling its `deposit` function, and in that call provide as arguments the `pubkey`, the `withdrawal_credentials` and a `signature`, signed with the pubkey, of both keys (and the amount deposited, which could under special circumstances be something other than 32 ETH).
3.  The deposit contract maintains a **deposit tree**; when it processes your function call, it adds your deposit (a `DepositData` record) to the deposit tree
4.  Once the eth2 chain becomes aware of an eth1 block after you deposited (through the Eth1Data voting procedure), the eth2 chain “knows” the root hash of a deposit tree that your `DepositData` record is part of.
5.  Once this happens, eventually an eth2 proposer will include a Merkle branch proving your deposit, and at that point you will get scheduled for activation. Note that proposers are _required_ to include as many deposits as possible, and to process deposits sequentially, so there is no possibility that your deposit will never be processed because proposers are mean or lazy.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#DepositMessage "DepositMessage")`DepositMessage`

```
class DepositMessage(Container):
    pubkey: BLSPubkey
    withdrawal_credentials: Bytes32
    amount: Gwei
```

This is the data that the `signature` in the `DepositData` is signing over. The signature is necessary, and must sign over all three of these fields, for two reasons:

1.  There must be a signature signing over the pubkey itself as a proof of possession to guard against rogue key attacks (see [here](https://crypto.stanford.edu/~dabo/pubs/papers/BLSmultisig.html#mjx-eqn-eqagg) for a description on what rogue key attacks on BLS aggregation are and how proofs of possession, or “prov\[ing\] knowledge of the secret key (POSK)” solve this problem).
2.  There is a rule that there can only be one validator with each pubkey (this rule is there to prevent deposits from being replayed; while this is a non-issue in phase 0 as deposits must be processed sequentially, in phase 2+ depositors from shards are responsible for getting their own deposits included, so they may be included out-of-order). However, this opens up an attack: if someone is already depositing, you can front-run their deposit with your own deposit that has the same pubkey but a different withdrawal credentials (or a lower amount), effectively causing their funds to be destroyed. Requiring the withdrawal credentials and the amount to be signed by the public key prevents this.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#DepositData "DepositData")`DepositData`

```
class DepositData(Container):
    pubkey: BLSPubkey
    withdrawal_credentials: Bytes32
    amount: Gwei
    signature: BLSSignature  # Signing over DepositMessage
```

The struct that goes into the deposit tree saved by the deposit contract on the eth1 side. Note that this data does NOT get verified on the eth1 side (as we do not yet have BLS-12-381 on eth1); invalid signatures could make it into the deposit tree, and it’s the eth2 chain’s responsibility to ignore them.

```
class BeaconBlockHeader(Container):
    slot: Slot
    proposer_index: ValidatorIndex
    parent_root: Root
    state_root: Root
    body_root: Root
```

The block header of a beacon chain block: contains the slot, the proposer index, the state root, the root hash of the parent block, and the root hash of everything else in the beacon chain block. Fundamentally similar to block headers in other blockchains including bitcoin and eth1.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Aside-domain-separation "Aside-domain-separation")\[Aside: domain separation\]

Domain separation in eth2 is done to prevent situations where a signature of an object of one type and context is accidentally a valid signature of an object in a different type or in a different context. This could happen because the same data happens to be valid as multiple data types; such situations could conceivably be triggered by attackers to cause slashings or other problems. Domain separation unambiguously makes this impossible.

There are two main types of domain separation in eth2:

-   **Cross-chain domain separation**: a message signed for one eth2 chain should not be valid in any other eth2 chain (eg. mainnet vs testnet, testnet vs another testnet, a hypothetical eth2 vs “eth2 classic” fork)
-   **Cross-context domain separation**: a signature of an attestation should not be valid as a signature of a beacon block, etc etc.

We achieve domain separation by mixing in a domain hash whenever we sign a message; that is, when we sign some `object`, we’re really signing `hash(root_hash(object), domain_hash)`. The domain hash itself mixes together the `domain_type` and the `fork_version` (representing the chain), see [`get_domain`](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#get_domain) for the logic of how this works.

The `domain_type` is a 4-byte value; see [the list of domain types](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Domain-types). The `fork_version` (think of this as like a chain ID, except it changes every hard fork to facilitate replay protection during deliberate ETH/ETC-like splits) is computed based on the logic described [above in the section on forks](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Fork).

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#SigningData "SigningData")`SigningData`

```
class SigningData(Container):
    object_root: Root
    domain: Domain
```

Eth2 heavily makes use of **signed containers**: structures where there is some inner container `C1`, and an outer container `C2(message: C1, signature: BLSSignature)`. To implement domain separation (see [the section above](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#domain_separation)), instead of directly signing the root hash of the message being signed, the signature signs over the root hash of a struct containing that message and the domain (computed via [`get_domain`](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#get_domain)).

`SigningData` is a dummy struct used to compute `hash(root_hash(object), domain_hash)`; in general the spec has made the aesthetic decision to broadly eschew explicit inline bit concatenations (`hash(x + y)`) in favor of using structs, with SSZ Merkle-hashing internally performing these bit concatenations.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Beacon-operations "Beacon-operations")Beacon operations

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#ProposerSlashing "ProposerSlashing")`ProposerSlashing`

```
class ProposerSlashing(Container):
    signed_header_1: SignedBeaconBlockHeader
    signed_header_2: SignedBeaconBlockHeader
```

A proposer can get slashed for signing two distinct headers at the same slot. This object can be included on-chain to perform that slashing.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#AttesterSlashing "AttesterSlashing")`AttesterSlashing`

```
class AttesterSlashing(Container):
    attestation_1: IndexedAttestation
    attestation_2: IndexedAttestation
```

An attester can get slashed for signing two attestations that together violate the Casper FFG slashing conditions. This object can be included on-chain to perform that slashing.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Attestation "Attestation")`Attestation`

```
class Attestation(Container):
    aggregation_bits: Bitlist[MAX_VALIDATORS_PER_COMMITTEE]
    data: AttestationData
    signature: BLSSignature
```

A record specifying that part of some committee (using a bitfield to identify which part) signed some `AttestationData`. Eth2 [uses BLS signature aggregation](https://ethresear.ch/t/pragmatic-signature-aggregation-with-bls/2105) for efficiency, so instead of every validator’s attestation being included separately, attestations first get broadcasted to an aggregation layer in the network, and then the block proposer can merge all attestations that sign the exact same `AttestationData` (in the normal case this is most of them) into a single `Attestation` that gets included on chain.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Deposit "Deposit")`Deposit`

```
class Deposit(Container):
    proof: Vector[Bytes32, DEPOSIT_CONTRACT_TREE_DEPTH + 1]  # Merkle path to deposit root
    data: DepositData
```

Proof that a validator deposited. These get processed sequentially in order of index; each proof is a Merkle branch proving that the deposit actually is in the correct position in the deposit tree that was created by the eth1 deposit contract.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#VoluntaryExit "VoluntaryExit")`VoluntaryExit`

```
class VoluntaryExit(Container):
    epoch: Epoch  # Earliest epoch when voluntary exit can be processed
    validator_index: ValidatorIndex
```

When a validator wishes to exit voluntarily, they may create and sign and broadcast a message of this type.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Beacon-blocks "Beacon-blocks")Beacon blocks

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#BeaconBlockBody "BeaconBlockBody")`BeaconBlockBody`

```
class BeaconBlockBody(Container):
    randao_reveal: BLSSignature
    eth1_data: Eth1Data  # Eth1 data vote
    graffiti: Bytes32  # Arbitrary data
    # Operations
    proposer_slashings: List[ProposerSlashing, MAX_PROPOSER_SLASHINGS]
    attester_slashings: List[AttesterSlashing, MAX_ATTESTER_SLASHINGS]
    attestations: List[Attestation, MAX_ATTESTATIONS]
    deposits: List[Deposit, MAX_DEPOSITS]
    voluntary_exits: List[SignedVoluntaryExit, MAX_VOLUNTARY_EXITS]
```

The “main” part of a beacon block. The most important thing in here is the attestations, as these need to get included so the chain can track its own finality status and apply rewards and penalties, but this also includes slashings, deposits, voluntary exits, the revealed value to adjust the block’s randomness seed, an eth1 vote and an open “graffiti” field.

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#BeaconBlock "BeaconBlock")`BeaconBlock`

```
class BeaconBlock(Container):
    slot: Slot
    proposer_index: ValidatorIndex
    parent_root: Root
    state_root: Root
    body: BeaconBlockBody
```

A full beacon block; basically a beacon block header but with the body root replaced by the full body.

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Beacon-state "Beacon-state")Beacon state

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#BeaconState "BeaconState")`BeaconState`

```
class BeaconState(Container):
    # Versioning
    genesis_time: uint64
    genesis_validators_root: Root
    slot: Slot
    fork: Fork
    # History
    latest_block_header: BeaconBlockHeader
    block_roots: Vector[Root, SLOTS_PER_HISTORICAL_ROOT]
    state_roots: Vector[Root, SLOTS_PER_HISTORICAL_ROOT]
    historical_roots: List[Root, HISTORICAL_ROOTS_LIMIT]
    # Eth1
    eth1_data: Eth1Data
    eth1_data_votes: List[Eth1Data, EPOCHS_PER_ETH1_VOTING_PERIOD * SLOTS_PER_EPOCH]
    eth1_deposit_index: uint64
    # Registry
    validators: List[Validator, VALIDATOR_REGISTRY_LIMIT]
    balances: List[Gwei, VALIDATOR_REGISTRY_LIMIT]
    # Randomness
    randao_mixes: Vector[Bytes32, EPOCHS_PER_HISTORICAL_VECTOR]
    # Slashings
    slashings: Vector[Gwei, EPOCHS_PER_SLASHINGS_VECTOR]  # Per-epoch sums of slashed effective balances
    # Attestations
    previous_epoch_attestations: List[PendingAttestation, MAX_ATTESTATIONS * SLOTS_PER_EPOCH]
    current_epoch_attestations: List[PendingAttestation, MAX_ATTESTATIONS * SLOTS_PER_EPOCH]
    # Finality
    justification_bits: Bitvector[JUSTIFICATION_BITS_LENGTH]  # Bit set for every recent justified epoch
    previous_justified_checkpoint: Checkpoint  # Previous epoch snapshot
    current_justified_checkpoint: Checkpoint
    finalized_checkpoint: Checkpoint
```

This is the most important data structure here; it is the thing that the state transition function that is defined in this spec is modifying. This contains all of the information that is needed to process the next beacon block, and can be split into a few categories:

-   Miscellanea and versioning (slot, fork version, genesis time…)
-   History that needs to be accessed: historical block hashes, state roots, randomness seeds…
-   Data needed to maintain the eth1 voting system
-   The validator registry (plus the separate array of exact balances)
-   Per-epoch sums of slashings (used to keep track of how many got slashed so that proportional penalties can be calculated)
-   Pending attestations
-   Casper FFG-related data

### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#Signed-envelopes "Signed-envelopes")Signed envelopes

These are just signed versions of many of the containers above:

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#SignedVoluntaryExit "SignedVoluntaryExit")`SignedVoluntaryExit`

```
class SignedVoluntaryExit(Container):
    message: VoluntaryExit
    signature: BLSSignature
```

#### [](https://notes.ethereum.org/@vbuterin/SkeyEI3xv#SignedBeaconBlock "SignedBeaconBlock")`SignedBeaconBlock`

```
class SignedBeaconBlock(Container):
    message: BeaconBlock
    signature: BLSSignature
```

```
class SignedBeaconBlockHeader(Container):
    message: BeaconBlockHeader
    signature: BLSSignature
```
