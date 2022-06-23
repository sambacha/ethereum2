---
created: 2022-04-03T11:26:18 (UTC -07:00)
tags: []
source: https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view
author: 
---

# Account Abstraction, Stateless Mining Eth1.x/Eth 2 Implementation, Rationale Document - HackMD

> ## Excerpt
> The following document discusses Quilt’s implementation plan and rationale for implementing account abstraction (AA) and early aspects of stateless mining into eth1.x. Feel free to follow the current tasks and work directly on trello. A longer discussion on how this work intersects with eth2, phase 2 can be found in the last section of the document.

---
The following document discusses Quilt’s implementation plan and rationale for implementing account abstraction (AA) and early aspects of stateless mining into eth1.x. Feel free to follow the current tasks and work directly [on trello](https://trello.com/b/pdY48Gez/qd-aa-rd-eip-sfork). A longer discussion on how this work intersects with eth2, phase 2 can be found in the last section of the document.

## [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Benefits-of-Account-Abstraction "Benefits-of-Account-Abstraction")Benefits of Account Abstraction

Vitalik released a [strong writeup/background](https://ethereum-magicians.org/t/implementing-account-abstraction-as-part-of-eth1-x/4020) on achieving a minimal implementation for account abstraction. This section does not define the features of account abstraction but rather states the benefits to the current eth1.x and eth2 ecosystem.

-   User Onboarding
    
    -   Skip onboarding steps for users that don’t already have ether and let the contract pay for initial gas fees
-   Privacy
    
    -   No need to rely on paid relayers for interactions that require privacy (ie mixers). In Tornado Cash for example, withdrawing your funds requires a relayer to preserve privacy. Otherwise, an account with eth needs to be in place which may defeat the purpose of mixing in the first place. With AA, the withdrawal funds can be used for gas payment
-   Smart Contract Wallets
    
    -   To use smart contract wallets today, an EOA account needs to pay for the necessary gas. With AA, smart contract wallets can use their balance to pay for gas/fees
    -   Social recovery can be simplified as part of multisig/multiparty interactions
-   dApp Usage via dApp Tokens
    
    -   A useful example may be around subscription style services such as gaming apps. In a gaming app, your users likely want to interact with the game without having to worry about low-level mechanics like ether - ideally, users should not even have to know what a blockchain is. In a system like this, users may be allocated (or win) a certain number of game tokens which are used as credits for underlying transactions (which the contract covers via its ether balance).
-   Multi Party Transactions
    
    -   Multisig wallets or transactions can use m-of-n cryptographic primitives to control account access and use funds from the account to pay transaction fees
-   Other Cryptographic Primitives
    
    -   BLS signatures, STARK verification, Schnorr signatures and more can be used instead of just ECDSA signatures
-   Cron Support
    
    -   Transactions can be re-submitted and used repeatedly until funds are drained or the AA account invalidates the transaction. A good example is a crowd funded alarm clock implmentation. The alarm clock contract allows calls to be registered to trigger a certain intervals. Triggering the alarm must happen externally and therefore is a common good. What we could do to improve this experience is construct an account with a single valid transaction that triggers the alarm clock at the defined intervals. The transaction is only valid at these intervals. Users who care about the alarm clock can then contribute to that account’s ether balance and anyone (including miners) can continue including that single transaction at its valid intervals. Keep in mind, this breaks the unique transaction hash rule, but is backwards compatible as it only applies to AA accounts
-   Efficiencies Around Failed Transactions
    
    -   During the ICO era, there were more people interested in buying tokens than the supply provided. As a result, blocks were filled/bogged down with invalid transactions which paid potentially high transaction fees. In AA, failed transactions can be guaranteed to not even make it on chain. This particular feature is also useful for potential DEX applications where multiple users try to claim the same arbitrage opportunity. This feature introduces a possible denial of service vector, which can be mitigated with new transaction pool limits.
-   Easier Rollup Transferrability
    
    -   For optimistic rollups, users technically have to wait a 2 week period to withdraw eth to the base chain (or to transfer the eth to another active rollup chain). Liquidity providers could be a part of the contract and could give users an immediate exit of their funds. The user can use funds from the rollup to pay for gas and transfer the funds to another application on chain. Account abstraction would simplify the mechanisms around such liquidity providers and reduce the overhead / number of blockchain interactions neccessary.
-   State Channel Improvements
    
    -   Griefing attacks can be reduced by already having the gas available in case an on chain dispute needs to be settled
    -   The transaction to close a state channel can use the funds already deposited into the state channel
-   Steps Towards Eth 2 Phase 1.5 and Phase 2
    
    -   AA takes a first step at iterating towards full account abstraction (Execution Environments) and expanded functionality for eth1.x on eth2
    -   There has been recent discussion on ethx64, where eth2 supports eth1 on 64 shards and eth1 gets upgrades (such as account abstraction and SSA). If ethx64 does not move forward, the work around account abstraction will move us forward on a new execution model we build

## [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#SSA-in-the-context-of-AA-and-Stateless-Mining "SSA-in-the-context-of-AA-and-Stateless-Mining")SSA in the context of AA and Stateless Mining

While we have come into contact with static state access (SSA) initially in the context of [research into full eth2 phase2 execution](https://ethresear.ch/t/state-provider-models-in-ethereum-2-0/6750), this concept is already relevant for current eth1.x work and its intersection with eth2:

-   Stateless Mining for Eth 2
    
    -   Even under eth1.x statelessness, miners are still required to hold the full eth1 chain state. One research topic within the eth1.x efforts is therefore the attempt to enable stateless mining. Moving from dynamic state access (DSA, the status quo on eth1) to SSA would be the most principled way of achieving this goal. Our plan is to investigate possible migration plans for eth1 to an SSA model.
-   Communication Between Accounts or Execution Environmnets
    
    -   If we have multiple execution frameworks or AA acccounts, we can achieve simpler communication between them (including possibly synchronous communication) under SSA. It moreover simplifies the requirements needed for the mempool and state network in addition to supporting proper dependency graphs when ordering/arranging incoming transactions to the mempool

## [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Implementation-Plan "Implementation-Plan")Implementation Plan

Our initial work is focused on the Geth client. Once the initial pieces are completed, we plan on also building or coordinating on Besu, parity and nethermind. We’ve broken it into 3 stages. You may follow our work [on trello](https://ethereum-magicians.org/t/implementing-account-abstraction-as-part-of-eth1-x/4020).

### [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Stage-1 "Stage-1")Stage 1

Stage 1 gets us a minimal viable implementation of account abstraction as [described by Vitalik](https://ethereum-magicians.org/t/implementing-account-abstraction-as-part-of-eth1-x/4020). Our preliminary plan includes support for multiple transactions with the same transaction hash (i.e. no enshrined AA nonce model), but we will revisit this decision at the end of the stage 1 implementation work.

#### [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Features "Features")Features

-   AA accounts are a subset of contract accounts (vs. adding a new account type)
-   Tx verification phase ends with a `PAYGAS` opcode. Transactions that do not call this opcode within a specified verification gas limit (e.g. 400,000 gas) are treated similar to normal transactions with invalid signatures and are dropped from the mempool.
-   Before calling `PAYGAS`, AA transactions cannot access external state or contracts (other than precompiles).
-   (0, 0) signature routes a transaction through `ENTRY_POINT` address `(= 2**160 - 1)`
-   AA contracts can only be directly called from the `ENTRY_POINT` address, cannot be called from other contracts.
-   Only one pending transaction at a time to each AA account

#### [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Tasks "Tasks")Tasks:

-   Write logic to check if tx is addressed to an AA - utilize (0, 0) signature and pass through the `ENTRY_POINT` `(= 2**160 - 1)`
-   Update mempool to validate/order AA transactions
    -   Generalize txList to support / enforce single transaction dependencies to AA
-   Restrict opcodes before `PAYGAS` runs during TX validation phase
-   Adding the PAYGAS opcode
-   Create the bytecode sequence for AA intro on contracts that falls in line with `require(msg.sender == ENTRY_POINT)`
-   Rough outline of tools and metrics that are important
-   Build simulation tools to benchmark performance of mempool with AA, stress test and generate transactions - [https://hackmd.io/X21CXYKbTbWT9DrcUAejHQ](https://hackmd.io/X21CXYKbTbWT9DrcUAejHQ)
-   Figure out how much gas is okay during validation via the simulating tools
-   Figure out how much gas is okay to support multiple transactions
-   Modify clients to support multiple transactions with the same hash
-   Build some hello world abstract accounts
-   Interface for calling a contract using an AA validation+gas pmt
-   Interface for calling a CREATE of AA validation+ gas pmt contract
-   Create illustrative abstract account examples with analysis

### [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Stage-2 "Stage-2")Stage 2

Stage 2 describes additional features to expand the functionality of account abstraction and build further tooling. This includes preliminary research on ways to support multiple transactions per account (e.g. via SSA). Tasks will be less detailed until we arrive closer to stage 2.

#### [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Research-amp-Tasks "Research-amp-Tasks")Research & Tasks

-   Research on expanding to multiple transactions per AA account.
-   Add support for AA contracts to solidity
-   Expand web tools to support AA
-   Begin discussion of migration to SSA
-   Research calls into AA contracts from EOA accounts and other contracts
-   Research and implement bundling transactions

### [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Stage-3 "Stage-3")Stage 3

Stage 3 may develop to either support a new execution framework for eth2 or be continued expansion on eth1.x.

#### [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Features1 "Features1")Features

-   Multiple transactions per AA account is supported
-   SSA migration
-   Dependency graph, parallel TX execution for mempools based on SSA access lists
-   Tooling around Async/Sync communication between abstract accounts
-   Support more execution environment style systems where AA accounts can form their own state network
-   Stateless mining roadmap in place

## [](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Further-Discussion-on-Intersection-with-Eth-2-Phase-2-Prior-Research "Further-Discussion-on-Intersection-with-Eth-2-Phase-2-Prior-Research")Further Discussion on Intersection with Eth 2 Phase 2 Prior Research

The eth1, eth2 merge timeline prioritizes a “phase 1.5 release”. Execution or “phase 2” for eth2 will be included after the merge and still is fairly greenfield and open (meaning there is no determinate time or concrete decisions on outcome). Since we are targeting a functional eth2 as soon as possible, account abstraction gives us a stepping stone or early iteration of what we would want as part of eth2, phase 2. Previous phase 2 goals targeted complete account abstraction (execution environments) and multiple executing engines. Execution environments and basic account abstraction share many of the same questions.

-   How does the BP know they will be paid without understanding/verifying the details of the execution environment or account?
-   How can we facilitate communication between different EEs or accounts (hint: asynchronous tooling)?
-   Complete or partial abstraction in a stateless world
-   How to relay/provide a state network for transactions to different accounts or execution environments?
-   How to bundle multiple transactions together into a single package
-   Rules on verification logic, witness refreshes, etc. across different AA accounts or execution environments and [its impact on the state network or mempool](https://ethresear.ch/t/state-provider-models-in-ethereum-2-0/6750)

To expand to a more “complete” form of account abstraction supporting multiple transactions per account (in both account abstraction and execution engines), we need significant expansions on the mempool. One possible implementation uses a dependency graph and witness refreshing. To make this as simple as possible, we suggest a Static State Access (SSA) model for accessing storage and have covered this topic in [past research](https://ethresear.ch/t/state-provider-models-in-ethereum-2-0/6750). This work towards SSA also gets us closer to stateless mining (which is otherwise quite difficult) within the eth2 system due to needing complex state provider models otherwise.

-   [Account Abstraction, Stateless Mining Eth1.x/Eth 2 Implementation, Rationale Document](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Account-Abstraction-Stateless-Mining-Eth1xEth-2-Implementation-Rationale-Document "Account Abstraction, Stateless Mining Eth1.x/Eth 2 Implementation, Rationale Document")
    -   [Benefits of Account Abstraction](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Benefits-of-Account-Abstraction "Benefits of Account Abstraction")
    -   [SSA in the context of AA and Stateless Mining](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#SSA-in-the-context-of-AA-and-Stateless-Mining "SSA in the context of AA and Stateless Mining")
    -   [Implementation Plan](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Implementation-Plan "Implementation Plan")
        -   [Stage 1](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Stage-1 "Stage 1")
        -   [Stage 2](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Stage-2 "Stage 2")
        -   [Stage 3](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Stage-3 "Stage 3")
    -   [Further Discussion on Intersection with Eth 2 Phase 2 Prior Research](https://hackmd.io/y7uhNbeuSziYn1bbSXt4ww?view#Further-Discussion-on-Intersection-with-Eth-2-Phase-2-Prior-Research "Further Discussion on Intersection with Eth 2 Phase 2 Prior Research")
