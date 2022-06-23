---
created: 2022-04-05T05:36:57 (UTC -07:00)
tags: []
source: https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ
author: 
---

# Sharding P2P POC Performance Test Plan - HackMD

> ## Excerpt
> https://github.com/ethresearch/sharding-p2p-poc
https://github.com/ethresearch/sharding-p2p-poc/tree/testing

---
###### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#tags-p2p-testing-spec "tags-p2p-testing-spec")tags: `p2p` `testing` `spec`

## [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Test-Phase-v10 "Test-Phase-v10")Test Phase v1.0

[https://github.com/ethresearch/sharding-p2p-poc](https://github.com/ethresearch/sharding-p2p-poc)  
[https://github.com/ethresearch/sharding-p2p-poc/tree/testing](https://github.com/ethresearch/sharding-p2p-poc/tree/testing)

**Project Manager:** Zak Cole  
**Test Team:** Daniel Choi, Nate Blakely

Whiteblock  
[www.whiteblock.io](http://www.whiteblock.io/)

## [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Overview "Overview")Overview

The network is segmented into X number of shards. Every ~10 minutes, validators are randomly assigned to a shard, so the stress point is observing and testing the ability of validators to subscribe to new topics and send/receive messages pertaining to this new topic in an adequate amount of time.

## [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Test-Utilities "Test-Utilities")Test Utilities

Tests will be conducted using the [Whiteblock](https://www.whiteblock.io/) platform in accordance with the proposed scope of work outlined within this document.

## [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Test-Procedure "Test-Procedure")Test Procedure

Per test case:

1.  Build network
2.  Provision nodes
3.  Configure network conditions between nodes according to specified test case
4.  Configure actions and behavior between nodes according to specified test case
5.  Output performance data in CSV format
6.  Aggregate data, parse, & present data
7.  Push data to appropriate repo
8.  Reset environment

Additional details pertaining to the test setup and procedure can be found within the [source code test repo](https://github.com/ethresearch/sharding-p2p-poc/tree/master/test) or referenced in [this document](https://notes.ethereum.org/xCbNDb3qQjmvHGnMEMt_HA?view).

## [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Performance-Metrics "Performance-Metrics")Performance Metrics

| Value | Description |
| --- | --- |
| Subscription Time | The length of time it takes for a node to subscribe to a topic, or in otherwords, join a shard, and begin receiving and broadcasting messages pertaining to that topic. |
| Discovery Time | The length of time it takes for a node to become aware of its peers within their subscribed shard. |
| Message Propagation Time | (Broadcast time) The length of time it takes for a message, once broadcast, to be received by a majority (99%) of peers within the shard. |

## [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Performance-Tests "Performance-Tests")Performance Tests

The following tables define each test series within this test phase. A test series focuses on observing and documenting the effects of certain conditions on performance. Each test series is comprised of three separate test cases which define the variable to be tested.

It is important to note that each test series may yield unexpected results which may influence the configuration of subsequential test cases or series. Accounting for this notion, this test plan should be considered a living document subject to change. Based on the results of this test phase, a consecutive test phase may be developed.

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-1-Number-of-Validators "Series-1-Number-of-Validators")Series 1: Number of Validators

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 300 | 600 | 1000 |
| Static Nodes | 600 | 600 | 600 |
| Shards | 5 | 5 | 5 |
| Message Size | 100KB | 100KB | 100KB |
| Bandwidth | 5Mb | 5Mb | 5Mb |
| Network Latency | 0ms | 0ms | 0ms |
| Packet Loss | 0% | 0% | 0% |

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-2-Number-of-Static-Nodes "Series-2-Number-of-Static-Nodes")Series 2: Number of Static Nodes

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 600 | 600 | 600 |
| Static Nodes | 300 | 600 | 1000 |
| Shards | 5 | 5 | 5 |
| Message Size | 100KB | 100KB | 100KB |
| Bandwidth | 5Mb | 5Mb | 5Mb |
| Network Latency | 0ms | 0ms | 0ms |
| Packet Loss | 0% | 0% | 0% |

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-3-Number-of-Shards "Series-3-Number-of-Shards")Series 3: Number of Shards

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 1000 | 1000 | 1000 |
| Static Nodes | 1000 | 1000 | 1000 |
| Shards | 5 | 7 | 10 |
| Message Size | 100KB | 100KB | 100KB |
| Bandwidth | 5Mb | 5Mb | 5Mb |
| Network Latency | 0ms | 0ms | 0ms |
| Packet Loss | 0% | 0% | 0% |

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-4-Message-Size "Series-4-Message-Size")Series 4: Message Size

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 1000 | 1000 | 1000 |
| Static Nodes | 1000 | 1000 | 1000 |
| Shards | 10 | 10 | 10 |
| Message Size | 1KB | 100KB | 1MB |
| Bandwidth | 5Mb | 5Mb | 5Mb |
| Network Latency | 0ms | 0ms | 0ms |
| Packet Loss | 0.01% | 0.5% | 1.0% |

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-4-Bandwidth "Series-4-Bandwidth")Series 4: Bandwidth

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 1000 | 1000 | 1000 |
| Static Nodes | 1000 | 1000 | 1000 |
| Shards | 5 | 5 | 5 |
| Message Size | 100KB | 100KB | 100KB |
| Bandwidth | 5Mb | 500Mb | 1G |
| Network Latency | 0ms | 0ms | 0ms |
| Packet Loss | 0% | 0% | 0% |

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-5-Network-Latency "Series-5-Network-Latency")Series 5: Network Latency

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 1000 | 1000 | 1000 |
| Static Nodes | 1000 | 1000 | 1000 |
| Shards | 10 | 10 | 10 |
| Message Size | 100KB | 100KB | 100KB |
| Bandwidth | 5Mb | 5Mb | 5Mb |
| Network Latency | 25ms | 50ms | 100ms |
| Packet Loss | 0% | 0% | 0% |

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-6-Packet-Loss "Series-6-Packet-Loss")Series 6: Packet Loss

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 1000 | 1000 | 1000 |
| Static Nodes | 1000 | 1000 | 1000 |
| Shards | 10 | 10 | 10 |
| Message Size | 100KB | 100KB | 100KB |
| Bandwidth | 5Mb | 5Mb | 5Mb |
| Network Latency | 0ms | 0ms | 0ms |
| Packet Loss | 0.01% | 0.5% | 1.0% |

### [](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-7-Stress-Test "Series-7-Stress-Test")Series 7: Stress Test

| Variable | Test Case A | Test Case B | Test Case C |
| --- | --- | --- | --- |
| Validators | 1000 | 500 | 300 |
| Static Nodes | 2000 | 2000 | 2000 |
| Shards | 10 | 20 | 30 |
| Message Size | 1MB | 1MB | 1MB |
| Bandwidth | 5Mb | 5Mb | 5Mb |
| Network Latency | 50ms | 50ms | 50ms |
| Packet Loss | 0.01% | 0.01% | 0.01% |

-   [Sharding P2P POC Performance Test Plan](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Sharding-P2P-POC-Performance-Test-Plan "Sharding P2P POC Performance Test Plan")
    -   [Test Phase v1.0](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Test-Phase-v10 "Test Phase v1.0")
    -   [Overview](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Overview "Overview")
    -   [Test Utilities](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Test-Utilities "Test Utilities")
    -   [Test Procedure](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Test-Procedure "Test Procedure")
    -   [Performance Metrics](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Performance-Metrics "Performance Metrics")
    -   [Performance Tests](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Performance-Tests "Performance Tests")
        -   [Series 1: Number of Validators](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-1-Number-of-Validators "Series 1: Number of Validators")
        -   [Series 2: Number of Static Nodes](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-2-Number-of-Static-Nodes "Series 2: Number of Static Nodes")
        -   [Series 3: Number of Shards](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-3-Number-of-Shards "Series 3: Number of Shards")
        -   [Series 4: Message Size](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-4-Message-Size "Series 4: Message Size")
        -   [Series 4: Bandwidth](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-4-Bandwidth "Series 4: Bandwidth")
        -   [Series 5: Network Latency](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-5-Network-Latency "Series 5: Network Latency")
        -   [Series 6: Packet Loss](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-6-Packet-Loss "Series 6: Packet Loss")
        -   [Series 7: Stress Test](https://notes.ethereum.org/Q_kQKXZUQD29YCshej1qPQ#Series-7-Stress-Test "Series 7: Stress Test")
