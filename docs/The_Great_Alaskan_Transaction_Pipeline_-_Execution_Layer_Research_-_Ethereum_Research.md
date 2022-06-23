---
created: 2022-04-03T11:16:42 (UTC -07:00)
tags: []
source: https://ethresear.ch/t/the-great-alaskan-transaction-pipeline/8472/9
author: 
---

# The Great Alaskan Transaction Pipeline - Execution Layer Research - Ethereum Research

> ## Excerpt
> That’s a step in the right direction, but there’s something critical missing: how is the validation cost of blocks bounded?   If it’s purely in bytes, then a malicious miner can add 1,000,000,000,000,000 gas worth of valid txs to the tx queue in a single block. This also breaks EIP-1559, as you can’t make the miner pay for this spike in gas since they’re not supposed to know how much gas txs use. If it’s in gas (which includes bytes metered in gas), then a malicious user can send the miner a ti...

---
The Alaskan Pipeline (Inspired by Nervos’ inspirational [NC-Max 17](https://eprint.iacr.org/2020/1101.pdf)) is a proposal for decoupling block propagation from witness propagation. Once these two things have been broken apart witness sizes have a drastically lessened impact on chain security. Paired with a mechanism for incentivizing witness propagation this proposal aims to solve two large outstanding difficulties with Stateless Ethereum.

Currently, transactions are immediately processed. By the time a block has been sealed all the transactions inside that block have been processed and the results of their execution are reflected in the state root. Alaska breaks transaction processing into a two-step process: inclusion and execution:

-   **Inclusion**: Each block commits to a list of transactions which are propagated along with the block. However, those transactions are not yet processed: their execution is not reflected in the state root and no receipts for them are included.\[1\]
    -   There is a limit on how many transactions each block can include, but how that limit is determined is an open question (more on this below)
-   **Execution**: Each block commits to a list of transactions which are executed in that block.
    -   It must execute transactions in the order they were included without skipping any transactions, starting with the oldest included transaction which has not yet been executed.
    -   After execution a transaction has been fully applied: that execution is reflected in the block’s state root and list of receipts
    -   There is a cooldown window: transactions may only be executed if they were included at least $\eta$ blocks ago, where $\eta$ is a configurable parameter.
    -   There are limits on how few/many transactions blocks may execute, but those limits are also an open question.

This is a fair amount of additional complication, why is it worth doing? **Separating inclusion from execution introduces a delay which affords us the time to build and propagate witnesses for each transaction before it is executed.** Because included transactions have a total ordering and are never skipped, executed transactions have completely deterministic witnesses and network participants have $\eta$ blocks to disseminate those witnesses. If everybody is honest then block propagation will not need to include witness data because by the time a block has been sealed everybody already has the witness.

This decoupling does not add any delay to transaction latency. The result of transaction execution is completely determined once the transaction has been included, and participants such as miners will know exactly how the transaction will resolve. Receipts and state roots are delayed, which means there is a delay before light clients and stateless clients can have transaction execution proven to them, but Alaska will not slow down transactions for network participants who are willing to fully process blocks.

For Alaska to work we will need an additional network primitive peers can use to gossip block witnesses to each other, honest miners will gossip the witness for a block as soon as they know what the witness will be. Honest nodes will encourage miners share block witnesses by introducing a delay to block propagation: _any node which receives a block for which they have not yet received a witness will wait for one second before gossipping that block to their peers_.

This rule ensures that any miner which does not pre-propagate witnesses risks having their blocks delayed which puts them at significant risk of having their blocks uncled which probabilistically loses them a significant amount of revenue. If a miner does pre-propagate witnesses they pay an additional cost in terms of bandwidth but for any reasonable witness size this is outweighed by the revenue they earn from having their blocks become canonical blocks.

This proposal does not change how much bandwidth Stateless Ethereum consumes, witnesses must still be propagated. However, Alaska propagates those witnesses during the time we are not bandwidth constrained: between block propagation. **As a result, larger witnesses do not slow block propagation**.

That’s the Great Alaskan Transaction Pipeline!

If we suddenly care much less about witness sizes then the only remaining blocker to a workable Stateless Ethereum is a solution to gas cost accounting. We don’t even need to come to agreement on a witness format: Alaska keeps the specifics of witness outside of consensus meaning we can easily evolve the format over time.

##### Open questions:

-   The obvious: What should $\eta$ be? How long should the propagation delay be?
-   What should the rules be around the queue of unexecuted transactions? How can we prevent the queue from growing unboundedly? Any rule we use should ensure the next few block witnesses are determinstic.
-   Miners are providing two separate services: including transactions and executing transactions. How should the fee be split between those two operations?
-   How does Alaska interact with ReGenesis?
-   How does Alaska interact with 1559?
-   For Alaska to work witnesses need to be deterministic. This means opcodes like BLOCKHASH must act as if the current block is the block in which the transaction was included, not the current block during transaction execution. Is there anything else which might cause witnesses to no longer be deterministic?
-   What does the transition strategy look like, how do we switch over to these new rules?

##### A proposed rule for bounding the size of the transaction queue:

This solution isn’t very clean but it seems like a decent starting point: the transaction queue is allowed to grow until the total number of transactions is $(\eta * blockGasLimit) / 21000$ . Blocks must execute as many transactions from the queue as they can without going over the block gas limit (and without breaking the rule that transactions must have been included at least $\eta$ blocks ago. This rule means that witnesses for future blocks are completely deterministic: miners do not have control over how many transactions they process. It also means some transactions will be executed/proven twice: once to show that completely processing this transaction would cause the block to go over the limit; once more to execute the transaction in the following block.

##### Footnote:

\[1\]: This description is a simplification: Each block commits not to the set of transactions it includes but to the full queue of unexecuted transactions. Each block appends some transactions to the end of the queue of unprocessed transactions and commits to the new queue. This allows network participants to statelessly process blocks.

-   #### created
    
    Jan '21
    
-   [
    
    #### last reply
    
    ](https://ethresear.ch/t/the-great-alaskan-transaction-pipeline/8472/17)
    
    [](https://ethresear.ch/t/the-great-alaskan-transaction-pipeline/8472/17)Jan '21
    
-   2.1k
    
    #### views
    
-   7
    
    #### users
    
-   6
    
    #### likes
    
-   5
    
    #### links
    

Deferring transaction execution from ordering is a DoS vector. If you try to avoid the DoS vector, it reduces to immediate execution.

Can you expand on this. It’s not immediately clear to me what the DOS vector is.

The potentially valuable thing that this concept gets us is a mechanism that side-steps the need for doing gas metering for witnesses at execution time by side-stepping the DSA problem, allowing transactions to be coupled with a witnesses prior to execution.

At present, we can’t include witnesses in transactions because of the DSA problem and not being able to predict what data a transaction will touch.

At present, we haven’t figured out a good solution to have miners generate witnesses because this would require us to do gas metering for the total witness data, and complex changes to the gas schedule are backwards incompatible and thus complex.

So, by decoupling inclusion and execution, we can side step these problems, allowing witnesses to be generated once a transaction has been included, but before it is executed.

I think the block proposer will need to pre-execute the transactions anyway since the blockproposer needs to know how much money exactly it will make from fees (and optimize).

In addition John is correct regarding DoS. By the way, DoS can be addressed by requiring each transaction to include Proof of Work, but it is likely to be lots of PoW.

One way to definitely optimize things without delaying execution is for the block to include hashes of transactions instead of transactions. Nodes already have most of transactions in pending queue, so there is no reason to include transactions in the block.

I expect this would be part of the standard workflow of miners, but I don’t think it would be strictly necessary. We assert that once a transaction is in the queue it’s behavior is fully deterministic (more on this below). Under this model, a miner would likely maintain a version of the state that represents the resulting state after executing the full transaction queue. This would allow them to to apply the standard transaction validity checks to new transactions they are adding to the queue in the same manner as is done today when building a block.

As for transactions in the queue being fully deterministic… opcodes like `BLOCKHASH` and `TIMESTAMP` and maybe `BLOCKNUMBER` are problematic and we’d need some solution to deal with this.

I think we have determinism if those opcodes act as if the current block is the block in which the transaction was included. As far as the state transition is concerned the transaction is executed as soon as it is included. “All” we’re doing is deferring the generation of the state root and receipts.

> Deferring transaction execution from ordering is a DoS vector. If you try to avoid the DoS vector, it reduces to immediate execution.

Couldn’t you immediately charge for the cost of storing the transaction bytes on chain, then later charge for execution gas?

That’s a step in the right direction, but there’s something critical missing: how is the validation cost of blocks bounded?

1.  If it’s purely in bytes, then a malicious miner can add 1,000,000,000,000,000 gas worth of valid txs to the tx queue in a single block. This also breaks EIP-1559, as you can’t make the miner pay for this spike in gas since they’re not supposed to know how much gas txs use.
2.  If it’s in gas (which includes bytes metered in gas), then a malicious user can send the miner a tiny tx that fills up the block with a high gas price, and have that tx be invalid (or valid, but revert immediately). The user gets their data made available and pay for this, but the rest of the block is empty and the miner got screwed.

Interesting. I guess one way to make this work is to always charge transactions their GAS\_LIMIT, no matter the actual gas used?  
This seems harsh on reverting transactions, but doable?

That would work, and be terrible UX! It’s in fact exactly what Cosmos does currently (deferred execution by 1 block and no gas refunds).

A better idea might be to execute transactions immediately, but include the witness _in the next block_.

Good idea!

For Eth1x historians, and for context, below is a summary related discussions. I copy/paste quotes here since some channels may not be publicly linkable.

In 2019 and early 2020, we discussed Alexey and Igor’s breakthorugh experiments ([1 1](https://medium.com/@akhounov/data-from-the-ethereum-stateless-prototype-8c69479c8abc), [2](https://medium.com/@akhounov/the-shades-of-statefulness-in-ethereum-nodes-697b0f88cd04), [3](https://medium.com/@mandrigin/stateless-ethereum-binary-tries-experiment-b2c035497768), [4](https://medium.com/@mandrigin/semi-stateless-initial-sync-experiment-897cc9c330cb)) which made statelessness feasible by using a cache of recent witness data.

On February 15th, I suggested something which resembles the Alaskan Pipeline:

> How about a future-cache in the form of a consensus tx pool?  
> The closest thing that I can find is Tendermint, where txs are propagated before they are committed.  
> Blocks would include some txs without witnesses, and these txs would go into the consensus tx pool after a delay period during which their witnesses propagate. To accommodate Turing-complete (modulo system limits) contract execution, some witnesses may need to be sent with the block.

On March 3rd M H Swende wisely said that caches bring complexity:

> I found the write-up: [https://gist.github.com/holiman/2fae5769b0334b857443b53a5aa746ec 1](https://gist.github.com/holiman/2fae5769b0334b857443b53a5aa746ec) (still a bit unpolished).  
> Example: if a client has a consensus-mandated cache of N blocks, but a new block comes along, where the proof is missing a piece of state that is present in N-1. So unless the client correctly cleaned out every last remnant of N-1, it will incorrectly deem the block+witness as valid.  
> And if there’s a reorg, so it needs to evaluate a sibling block, then it needs to reconstruct the previous state of the cache (or even several generations back). And in order to have such fine-grained control, we need journalling. And in the end, we’ll likely back this thing with an _actual_ cache that is suited to the actual memory available in the machine we’re running on.

This ended our discussions about caching, and we focused on pure statelessness.

Then, in June 2020, ReGenesis was [proposed](https://ethresear.ch/t/regenesis-resetting-ethereum-to-reduce-the-burden-of-large-blockchain-and-state/7582). ReGenesis uses a form of caching, but is simpler than the others. ReGenesis had potential to get consensus, so we all supported it.

Also around June 2020, there was another idea to decouple witnesses from blocks entirely, see Alexey’s [pre-consensus 1](https://ledgerwatch.github.io/regenesis_plan.html#Pre-consensus-network), an avalanche-like consensus for peers to agree on which witnesses are important before blocks are mined.

In January 2021, The Great Alaskan Transaction Pipeline is proposed.

[@adlerjohn](https://ethresear.ch/u/adlerjohn) I appreciate the difficulties you are pointing out and I’m very game to discuss them. I am aware that there are unsolved issues with this concept that would need to be figured out and it’s good to get them identified. I would also like to avoid getting too lost down any of those paths before focusing on the part which I believe is most valuable.

[@poemm](https://ethresear.ch/u/poemm) 's history here is right on target. Over the course of the stateless ethereum effort we’ve struggled to find a viable solution for in-protocol witnesses. This proposal is by no means a polished solution, but it hints at a direction that I believe is worth exploring. I believe that Brian and I will plan on presenting this concept in the next stateless call for discussion and to answer questions people may have.

The “core” idea here is side-stepping the DSA problem by decoupling transaction ordering from execution. Transaction ordering gets defined, and then execution happens “later”. With some assumptions about ensuring execution is deterministic, this gives a window of time during which a witness can be generated prior to execution. If we pull on this thread, maybe it leads us closer to in-protocol witnesses.

As far as I can tell there is no DOS. As soon as a transaction has been included the end result is completely deterministic and miners are 100% capable of knowing what the result of execution is going to be. This “reduces to immediate execution”, as John puts it, but that’s okay, under Stateless Ethereum miners have the entire state and are more than capable of operating with it.

I didn’t write this in the original post but after some thought I think the entire gas fee must go to the miner which includes the transaction. If we do this there is no need to resort to ideas such as charging for the entire `GAS_LIMIT`. If we use the proposed rule where miners must execute as many transactions as will fit under the $BlockGasLimit$ then executing miners have no discretion and will execute transactions, with the block reward for mining a valid block as their incentive.

This is a neat idea. What if we kept everything as it is in Istanbul and added just the block propagation rule: When a network participant receives a block for which it has not received a witness for the parent block it waits 1 second before forwarding. Miners which do not propagate witnesses have a high risk of having their blocks orphaned. I need to think about it more but my first impression is that this Alaska-lite gives you a lot of the benefit of Alaska without the complication.

Also, thank you for adding the context here [@poemm](https://ethresear.ch/u/poemm), this is great!

An interesting idea to address DOS is to charge the PROPOSER penalty for each transaction that ends up not having minimum gas.

Then it will be up to the proposer to execute or may be find a lighter way of ensuring the block is not filled with no-gas transactions.

> An interesting idea to address DOS

Please spell out exactly how the DOS works. I am willing to believe there is one but do not currently see any.

In this proposal it is possible for miners to include invalid transactions (the state root is not yet updated, so stateless clients would not be able to validate blocks if we were asserting on transaction validation). However, if including miners are given the full transaction fee then this is equivalent to how things work today.

Miners today are capable of mining empty blocks and sometimes do. This is not a DOS. Any miner which mines an empty block has reduced their own revenue and left the revenue of everybody else unchanged.

Miners under the current Alaska proposal are allowed to include invalid transactions. Any miner who does so has lost revenue but not reduced the revenue of any other miner. It is possible to pre-process the transaction queue and predict the results of executing any candidate transaction, so miners will be 100% capable of knowing ahead of time whether the transactions they are including are valid.

May be I did not understand the proposal correctly …

If you are pre-processing transactions, isn’t it equivalent to block transaction execution? If you are executing all transactions in a block, then I guess you will have state root.

Or you are saying that updating the state root is much more computationally intense than executing the block but not updating the root?)
