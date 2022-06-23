---
created: 2022-02-06T17:45:39 (UTC -08:00)
tags: []
source: https://hackmd.io/@n0ble/merge-test-plan
author: 
---

# The Merge Test Plan - HackMD

> ## Excerpt
> # The Merge Test Plan  ## Spec documents * [consensus-specs](https://github.com/ethereum/consensus-s

---
## [](https://hackmd.io/@n0ble/merge-test-plan#Spec-documents "Spec-documents")Spec documents

-   [consensus-specs](https://github.com/ethereum/consensus-specs/tree/dev/specs/merge)
-   [EIP-3675](https://eips.ethereum.org/EIPS/eip-3675)
-   [EIP-4399](https://eips.ethereum.org/EIPS/eip-4399)
-   [Engine API](https://github.com/ethereum/execution-apis/blob/main/src/engine/specification.md)
-   Optimistic sync

## [](https://hackmd.io/@n0ble/merge-test-plan#Standalone-Unit "Standalone-Unit")Standalone (Unit)

### [](https://hackmd.io/@n0ble/merge-test-plan#consensus-specs "consensus-specs")consensus-specs

-   is\_merge\_transition\_complete
-   is\_merge\_transition\_block
-   is\_execution\_enabled
-   compute\_timestamp\_at\_slot
-   get\_inactivity\_penalty\_deltas
-   slash\_validator
-   process\_block
-   process\_execution\_payload
-   process\_slashings
-   initialize\_beacon\_state\_from\_eth1
-   is\_valid\_terminal\_pow\_block
-   validate\_merge\_block
-   upgrade\_to\_merge
-   get\_pow\_block\_at\_terminal\_total\_difficulty
-   get\_terminal\_pow\_block
-   prepare\_execution\_payload
-   get\_execution\_payload

### [](https://hackmd.io/@n0ble/merge-test-plan#EIP-3675--EIP-4399 "EIP-3675--EIP-4399")EIP-3675 + EIP-4399

-   Blockchain tests (client will have to switch to the PoS consensus engine to run these tests)
    -   PoW block processing
        -   decsendands of a terminal block are not processed
    -   Block structure
        -   blocks are discarded if the value of either of the following fields deviates from the corresponding constant: `ommersHash`, `difficulty`, `nonce`, `ommers`
        -   `mixHash` can be set to a random value
        -   `extraData` respects `MAX_EXTRA_DATA_BYTES`
        -   other fields remain unchanged
    -   Block validity
        -   `difficulty` isn’t verified against difficulty formula
        -   `nonce` and `mixHash` aren’t verified against Ethash rules
        -   list of ommers and each ommer isn’t verified (don’t know how to check though as the list must be empty)
        -   Transition block validity
            -   a parent of `TRANSITION_BLOCK` satisfies terminal PoW block conditions
        -   Other block validaty rules remain
    -   Block and ommer rewards
        -   `coinbase` account doesn’t receive block rewards
        -   `coinbase` account doesn’t receive ommer rewards (seems infeasible to check)
        -   transaction fees remain
    -   EVM
        -   `RANDOM` opcode returns expected value
-   Regression
    -   check that 3675 and 4399 aren’t taking effect before the `TRANSITION_BLOCK`
-   Fuzzing

## [](https://hackmd.io/@n0ble/merge-test-plan#Integration-Hive "Integration-Hive")Integration (Hive)

### [](https://hackmd.io/@n0ble/merge-test-plan#Engine-API "Engine-API")Engine API

-   -   `eth_` namespace is also exposed
    -   EL respects the order of `forkchoiceUpdated` calls
    -   `engine_executePayload`
    -   `engine_forkchoiceUpdated`
    -   `engine_getPayload`
-   CL + EL mock
    -   CL respects the order of method calls while assigning JSON-RPC request IDs
-   CL + EL
    -   CL implementation of Engine API methods
        -   `engine_executePayload`
        -   `engine_forkchoiceUpdated`
        -   `engine_getPayload`

### [](https://hackmd.io/@n0ble/merge-test-plan#EIP-3675--EIP-43991 "EIP-3675--EIP-43991")EIP-3675 + EIP-4399

-   CL mock + EL
    -   switches the fork choice rule upon the first `POS_FORKCHOICE_UPDATED` event
    -   `POS_FORKCHOICE_UPDATED` updates the head
    -   `POS_FORKCHOICE_UPDATED` updates finalized block
    -   head isn’t changed until `POS_FORKCHOICE_UPDATED` is received, even if the new payload is imported
    -   descendants of a terminal block are not propagated
    -   any block is not propagated after the first block gets finalized
    -   does not sync beyond terminal block
    -   block with transaction using `RANDOM` opcode is executed as expected starting with `TRANSITION_BLOCK`
    -   block with transaction using `DIFFICULTY` opcode is executed as expected before the `TRANSITION_BLOCK`
    -   Transition
        -   block structure, block validity, block rewards, and fork choice changes takes effect starting with `TRANSITION_BLOCK`
        -   payload is successfully created upon request _before_ the transition, i.e. the transition block
        -   switch between several transition PoW blocks happens successfully

### [](https://hackmd.io/@n0ble/merge-test-plan#consensus-specs--EIPs "consensus-specs--EIPs")consensus-specs + EIPs

-   CL + EL
    -   consensus-specs
        -   state transition happy case
        -   state transition invalid payload wrt execution
        -   fork choice (see the Transition below)
    -   Transition
        -   transition scenario with no overrides
            -   payload is successfully created upon request _before_ the transition, i.e. the transition block
            -   CL passes transition block while syncing optimistically and must go back to verify TTD or TBH
            -   list all other possible scenarios of the transition process
        -   case when `TERMINAL_BLOCK_HASH` is set
            -   `TRANSITION_BLOCK` must be a child of a block identified by `TERMINAL_BLOCK_HASH`
            -   list all other possible scenarios of the transition process with `TERMINAL_BLOCK_HASH` override
        -   block structure, block validity, block rewards, and fork choice changes takes effect starting with `TRANSITION_BLOCK`
            -   PoS chain takes precedence
        -   `TRANSITION_BLOCK` isn’t a child of a terminal PoW block
            -   in lock-step
            -   after syncing optimistically
            -   terminal PoW block is invalid wrt execution ruleset – `TRANSITION_BLOCK` can’t be verified at all
            -   `TRANSITION_BLOCK` is invalid wrt execution ruleset
        -   switch between several terminal PoW blocks happens successfully
        -   switch between several transition blocks happens successfully
        -   withheld transition PoW block is successfully processed – network partitioning
        -   transition block with the payload which parent is unavailable doesn’t brick a node and it can proceed with building the chain and attesting, and can successfully complete the transition process

## [](https://hackmd.io/@n0ble/merge-test-plan#System-Testnets "System-Testnets")System (Testnets)

-   general purpose
-   stress testnets
-   production sized

## [](https://hackmd.io/@n0ble/merge-test-plan#Useful-resources "Useful-resources")Useful resources

-   Mergemock
    -   [https://github.com/protolambda/mergemock](https://github.com/protolambda/mergemock)
-   Test cases for Merge
    -   [https://hackmd.io/z2h\_RAJoTHWSRka-9MDEVg](https://hackmd.io/z2h_RAJoTHWSRka-9MDEVg)
-   Engine API test vectors for The Merge™
    -   [https://notes.ethereum.org/rmVErCfCRPKGqGkUe89-Kg](https://notes.ethereum.org/rmVErCfCRPKGqGkUe89-Kg)
-   EIP 4399 test vectors (Kintsugi)
    -   [https://notes.ethereum.org/dfVXr2h2TqWDKSXZaMwGTw](https://notes.ethereum.org/dfVXr2h2TqWDKSXZaMwGTw)
-   Kintzugi testing
    -   [https://github.com/marioevz/kintsugi\_testing](https://github.com/marioevz/kintsugi_testing)
-   Consensus spec tests
    -   [https://github.com/ethereum/consensus-spec-tests](https://github.com/ethereum/consensus-spec-tests)
-   Retesteth
    -   [https://github.com/ethereum/retesteth](https://github.com/ethereum/retesteth)
-   EL tests
    -   [https://github.com/ethereum/tests](https://github.com/ethereum/tests)
-   Hive
    -   [https://github.com/ethereum/hive](https://github.com/ethereum/hive)
-   Eth2 Hive - take 2
    -   [https://github.com/ethereum/hive/pull/481](https://github.com/ethereum/hive/pull/481)
-   Prysmatic end-to-end testing tool
    -   [https://github.com/prysmaticlabs/prysm/tree/develop/testing/endtoend](https://github.com/prysmaticlabs/prysm/tree/develop/testing/endtoend)
-   EVM-fuzzing
    -   [https://github.com/MariusVanDerWijden/FuzzyVM](https://github.com/MariusVanDerWijden/FuzzyVM)
-   Engine-API fuzzing
    -   [https://github.com/MariusVanDerWijden/merge-fuzz](https://github.com/MariusVanDerWijden/merge-fuzz)
-   Transaction fuzzing
    -   [https://github.com/MariusVanDerWijden/tx-fuzz](https://github.com/MariusVanDerWijden/tx-fuzz)
-   EIP-3675 + EIP-4399 blockchain tests
    -   [https://docs.google.com/spreadsheets/d/1V8egI6VbtZR\_sqUNCdX2QeRVb9fCJ3QmzCu8VZFs5pA/edit#gid=0](https://docs.google.com/spreadsheets/d/1V8egI6VbtZR_sqUNCdX2QeRVb9fCJ3QmzCu8VZFs5pA/edit#gid=0)
-   FuzzyVM
    -   [https://github.com/MariusVanDerWijden/FuzzyVM](https://github.com/MariusVanDerWijden/FuzzyVM)
