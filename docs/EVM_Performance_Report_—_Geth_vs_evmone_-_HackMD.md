---
created: 2022-01-04T23:29:27 (UTC -08:00)
tags: []
source: https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone
author: 
---

# EVM Performance Report — Geth vs evmone - HackMD

> ## Excerpt
> # EVM Performance Report — Geth vs evmone  This report compares performance of the latests versions

---
This report compares performance of the latests versions of [Geth](https://geth.ethereum.org/) and [evmone](https://github.com/ethereum/evmone) EVMs.

-   [EVM Performance Report — Geth vs evmone](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#EVM-Performance-Report-%E2%80%94-Geth-vs-evmone "EVM Performance Report — Geth vs evmone")
    -   [EVM implementations](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#EVM-implementations "EVM implementations")
    -   [Benchmark Suite](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Benchmark-Suite "Benchmark Suite")
    -   [Results](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Results "Results")
        -   [Geth to evmone/Baseline (times)](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Geth-to-evmoneBaseline-times "Geth to evmone/Baseline (times)")
        -   [Geth to evmone/Baseline (gas rate)](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Geth-to-evmoneBaseline-gas-rate "Geth to evmone/Baseline (gas rate)")
        -   [evmone/Advanced to evmone/Baseline (times)](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#evmoneAdvanced-to-evmoneBaseline-times "evmone/Advanced to evmone/Baseline (times)")
        -   [evmone/Advanced to evmone/Baseline (gas rate)](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#evmoneAdvanced-to-evmoneBaseline-gas-rate "evmone/Advanced to evmone/Baseline (gas rate)")

## [](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#EVM-implementations "EVM-implementations")EVM implementations

In total we have compared 3 EVM implementations: Geth and 2 evmone interpreters (Advanced & Baseline).

The original evmone/Advanced has relatively sophisticated design: during code analysis it also collects information about gas cost and stack requirements of groups of instructions (basic blocks). This makes the later execution more efficient.  
In contrast, the evmone/Baseline has “classic” EVM design with the code analysis limited only to the mandatory collecting of valid `JUMPDEST` locations.

The latest stable releases of the projects have been used:

-   [geth](https://github.com/ethereum/go-ethereum) [1.10.12](https://github.com/ethereum/go-ethereum/releases/tag/v1.10.12) released 2021-11-08, built with go 1.16.2
-   [evmone](https://github.com/ethereum/evmone) [0.8.2](https://github.com/ethereum/evmone/releases/tag/v0.8.2) released 2021-08-26, built with clang 13.0.1

## [](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Benchmark-Suite "Benchmark-Suite")Benchmark Suite

The performance of EVM implementations is compared with the [evmone’s benchmark suite](https://github.com/ethereum/evmone/tree/master/test/benchmarks). The suite is a small set of computation heavy EVM bytecodes (e.g. cryptographic hash functions) and some micro-benchmarks (e.g. stress testing single `MSTORE` instruction). The suite **does not necessarily represent Mainnet workloads**.

To benchmark evmone interpreters, the internal tool `evmone-bench` based on the [benchmark library](https://github.com/google/benchmark) has been used. It outputs results in a JSON format. To benchmark geth, the `evm run --bench` tool has been used. The loading of the benchmark suite and execution runs have been conducted via the [`evm-bench.py`](https://github.com/ethereum/evmone/pull/396) Python script. Comparison and statistical analysis have been in the end processed with [`benchstat`](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat).

Each benchmark sample contains code analysis and execution times combined. These timings may be separated with some additional tuning of the tooling in the future.

## [](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Results "Results")Results

Hardware: Intel Haswell CPU [i7-4790K](https://ark.intel.com/content/www/us/en/ark/products/80807/intel-core-i74790k-processor-8m-cache-up-to-4-40-ghz.html) with turbo frequency 4.4 GHz, Linux x86-64 (Ubuntu 20.04)

The evmone/Baseline has been used as the _baseline_ performance.

### [](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Geth-to-evmoneBaseline-times "Geth-to-evmoneBaseline-times")Geth to evmone/Baseline (times)

```
name                              Baseline time        Geth time        delta
main/blake2b_huff/empty             22.1µs ± 0%     111.2µs ± 0%     +402.22%  (p=0.000 n=9+10)
main/blake2b_huff/2805nulls          360µs ± 0%      2145µs ± 1%     +495.70%  (p=0.000 n=9+10)
main/blake2b_huff/5610nulls          697µs ± 0%      4163µs ± 0%     +496.90%  (p=0.000 n=9+10)
main/blake2b_huff/8415nulls         1.02ms ± 1%      6.07ms ± 0%     +495.14%  (p=0.000 n=9+10)
main/blake2b_shifts/2805nulls       3.23ms ± 0%     16.77ms ± 0%     +419.30%  (p=0.000 n=10+8)
main/blake2b_shifts/5610nulls       6.63ms ± 0%     33.43ms ± 1%     +403.85%  (p=0.000 n=9+9)
main/blake2b_shifts/8415nulls       10.0ms ± 0%      49.7ms ± 0%     +399.47%  (p=0.000 n=9+8)
main/sha1_divs/empty                55.1µs ± 0%     222.1µs ± 0%     +302.91%  (p=0.000 n=10+9)
main/sha1_divs/1351                 1.12ms ± 0%      4.49ms ± 0%     +301.83%  (p=0.000 n=9+9)
main/sha1_divs/2737                 2.19ms ± 1%      8.74ms ± 0%     +298.66%  (p=0.000 n=10+9)
main/sha1_divs/5311                 4.25ms ± 0%     17.10ms ± 0%     +302.58%  (p=0.000 n=10+9)
main/sha1_shifts/empty              33.4µs ± 0%     176.7µs ± 1%     +429.77%  (p=0.000 n=9+10)
main/sha1_shifts/1351                691µs ± 3%      3640µs ± 1%     +427.11%  (p=0.000 n=10+10)
main/sha1_shifts/2737               1.34ms ± 0%      7.07ms ± 1%     +429.74%  (p=0.000 n=9+10)
main/sha1_shifts/5311               2.61ms ± 1%     13.85ms ± 1%     +430.30%  (p=0.000 n=10+10)
main/weierstrudel/0                  171µs ± 0%       494µs ± 0%     +188.37%  (p=0.000 n=10+10)
main/weierstrudel/1                  364µs ± 1%      1053µs ± 0%     +189.20%  (p=0.000 n=10+10)
main/weierstrudel/3                  565µs ± 0%      1653µs ± 0%     +192.78%  (p=0.000 n=8+10)
main/weierstrudel/9                 1.16ms ± 0%      3.47ms ± 0%     +197.77%  (p=0.000 n=9+10)
main/weierstrudel/14                1.67ms ± 0%      4.95ms ± 0%     +196.94%  (p=0.000 n=10+10)
micro/memory_grow_mload/nogrow      71.9µs ± 0%     295.6µs ± 1%     +310.99%  (p=0.000 n=9+10)
micro/memory_grow_mload/by1         73.7µs ± 0%     302.8µs ± 0%     +310.73%  (p=0.000 n=10+10)
micro/memory_grow_mload/by16        83.7µs ± 1%     369.1µs ± 0%     +341.01%  (p=0.000 n=10+10)
micro/memory_grow_mload/by32        90.7µs ± 0%     449.3µs ± 1%     +395.55%  (p=0.000 n=8+10)
micro/memory_grow_mstore/nogrow     75.6µs ± 0%     530.1µs ± 0%     +601.47%  (p=0.000 n=9+10)
micro/memory_grow_mstore/by1        77.4µs ± 0%     542.2µs ± 0%     +600.57%  (p=0.000 n=10+10)
micro/memory_grow_mstore/by16       87.1µs ± 0%     627.0µs ± 1%     +619.89%  (p=0.000 n=8+10)
micro/memory_grow_mstore/by32       95.1µs ± 0%     695.9µs ± 1%     +631.54%  (p=0.000 n=8+10)
[Geo mean]                           447µs           1874µs          +319.26%
```

### [](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#Geth-to-evmoneBaseline-gas-rate "Geth-to-evmoneBaseline-gas-rate")Geth to evmone/Baseline (gas rate)

```
name                             Baseline gas/s       Geth gas/s        delta
main/blake2b_huff/empty               632M ± 0%        126M ± 0%      -80.09%  (p=0.000 n=9+10)
main/blake2b_huff/2805nulls           853M ± 0%        143M ± 1%      -83.21%  (p=0.000 n=9+10)
main/blake2b_huff/5610nulls           860M ± 0%        144M ± 0%      -83.25%  (p=0.000 n=9+10)
main/blake2b_huff/8415nulls           862M ± 1%        145M ± 1%      -83.20%  (p=0.000 n=9+10)
main/blake2b_shifts/2805nulls        1.06G ± 0%       0.20G ± 0%      -80.74%  (p=0.000 n=10+8)
main/blake2b_shifts/5610nulls        1.12G ± 0%       0.22G ± 1%      -80.15%  (p=0.000 n=9+9)
main/blake2b_shifts/8415nulls        1.21G ± 0%       0.24G ± 0%      -79.98%  (p=0.000 n=9+8)
main/sha1_divs/empty                  661M ± 0%        164M ± 0%      -75.18%  (p=0.000 n=10+9)
main/sha1_divs/1351                   704M ± 0%        175M ± 0%      -75.11%  (p=0.000 n=9+9)
main/sha1_divs/2737                   701M ± 1%        176M ± 0%      -74.92%  (p=0.000 n=10+9)
main/sha1_divs/5311                   706M ± 0%        175M ± 0%      -75.16%  (p=0.000 n=10+9)
main/sha1_shifts/empty               1.06G ± 0%       0.20G ± 1%      -81.12%  (p=0.000 n=9+10)
main/sha1_shifts/1351                1.10G ± 3%       0.21G ± 1%      -81.03%  (p=0.000 n=10+10)
main/sha1_shifts/2737                1.11G ± 0%       0.21G ± 1%      -81.12%  (p=0.000 n=9+10)
main/sha1_shifts/5311                1.11G ± 1%       0.21G ± 1%      -81.14%  (p=0.000 n=10+10)
main/weierstrudel/0                   277M ± 0%         96M ± 0%      -65.32%  (p=0.000 n=10+10)
main/weierstrudel/1                   358M ± 1%        124M ± 0%      -65.42%  (p=0.000 n=10+10)
main/weierstrudel/3                   358M ± 0%        122M ± 0%      -65.84%  (p=0.000 n=8+10)
main/weierstrudel/9                   359M ± 0%        120M ± 0%      -66.42%  (p=0.000 n=9+10)
main/weierstrudel/14                  359M ± 0%        121M ± 0%      -66.32%  (p=0.000 n=10+10)
micro/memory_grow_mload/nogrow        797M ± 0%        194M ± 1%      -75.67%  (p=0.000 n=9+10)
micro/memory_grow_mload/by1           784M ± 0%        191M ± 0%      -75.65%  (p=0.000 n=10+10)
micro/memory_grow_mload/by16          857M ± 1%        194M ± 0%      -77.32%  (p=0.000 n=10+10)
micro/memory_grow_mload/by32         1.13G ± 0%       0.23G ± 1%      -79.82%  (p=0.000 n=8+10)
micro/memory_grow_mstore/nogrow       813M ± 0%        116M ± 0%      -85.74%  (p=0.000 n=9+10)
micro/memory_grow_mstore/by1          799M ± 0%        114M ± 0%      -85.73%  (p=0.000 n=10+10)
micro/memory_grow_mstore/by16         870M ± 0%        121M ± 1%      -86.11%  (p=0.000 n=8+10)
micro/memory_grow_mstore/by32        1.12G ± 0%       0.15G ± 1%      -86.33%  (p=0.000 n=8+10)
[Geo mean]                            784M             161M           -79.53%
```

### [](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#evmoneAdvanced-to-evmoneBaseline-times "evmoneAdvanced-to-evmoneBaseline-times")evmone/Advanced to evmone/Baseline (times)

```
name                              Baseline time    Advanced time        delta
main/blake2b_huff/empty             22.1µs ± 0%      67.0µs ± 0%     +202.72%  (p=0.000 n=9+9)
main/blake2b_huff/2805nulls          360µs ± 0%       349µs ± 0%       -3.19%  (p=0.000 n=9+10)
main/blake2b_huff/5610nulls          697µs ± 0%       629µs ± 0%       -9.76%  (p=0.000 n=9+9)
main/blake2b_huff/8415nulls         1.02ms ± 1%      0.90ms ± 0%      -12.00%  (p=0.000 n=9+10)
main/blake2b_shifts/2805nulls       3.23ms ± 0%      3.54ms ± 0%       +9.47%  (p=0.000 n=10+10)
main/blake2b_shifts/5610nulls       6.63ms ± 0%      7.05ms ± 0%       +6.23%  (p=0.000 n=9+9)
main/blake2b_shifts/8415nulls       10.0ms ± 0%      10.6ms ± 0%       +5.93%  (p=0.000 n=9+9)
main/sha1_divs/empty                55.1µs ± 0%      57.6µs ± 0%       +4.40%  (p=0.000 n=10+9)
main/sha1_divs/1351                 1.12ms ± 0%      1.06ms ± 0%       -5.23%  (p=0.000 n=9+9)
main/sha1_divs/2737                 2.19ms ± 1%      2.06ms ± 0%       -6.08%  (p=0.000 n=10+7)
main/sha1_divs/5311                 4.25ms ± 0%      4.01ms ± 0%       -5.65%  (p=0.000 n=10+10)
main/sha1_shifts/empty              33.4µs ± 0%      35.5µs ± 0%       +6.43%  (p=0.000 n=9+9)
main/sha1_shifts/1351                691µs ± 3%       616µs ± 0%      -10.78%  (p=0.000 n=10+9)
main/sha1_shifts/2737               1.34ms ± 0%      1.20ms ± 0%      -10.42%  (p=0.000 n=9+9)
main/sha1_shifts/5311               2.61ms ± 1%      2.33ms ± 0%      -10.93%  (p=0.000 n=10+9)
main/weierstrudel/0                  171µs ± 0%       231µs ± 0%      +34.69%  (p=0.000 n=10+9)
main/weierstrudel/1                  364µs ± 1%       429µs ± 0%      +17.71%  (p=0.000 n=10+10)
main/weierstrudel/3                  565µs ± 0%       631µs ± 0%      +11.82%  (p=0.000 n=8+9)
main/weierstrudel/9                 1.16ms ± 0%      1.24ms ± 0%       +6.11%  (p=0.000 n=9+7)
main/weierstrudel/14                1.67ms ± 0%      1.74ms ± 0%       +4.39%  (p=0.000 n=10+9)
micro/memory_grow_mload/nogrow      71.9µs ± 0%     148.2µs ± 1%     +106.10%  (p=0.000 n=9+10)
micro/memory_grow_mload/by1         73.7µs ± 0%     149.6µs ± 0%     +102.95%  (p=0.000 n=10+9)
micro/memory_grow_mload/by16        83.7µs ± 1%     162.4µs ± 0%      +94.07%  (p=0.000 n=10+9)
micro/memory_grow_mload/by32        90.7µs ± 0%     172.6µs ± 1%      +90.33%  (p=0.000 n=8+10)
micro/memory_grow_mstore/nogrow     75.6µs ± 0%     156.0µs ± 0%     +106.37%  (p=0.000 n=9+9)
micro/memory_grow_mstore/by1        77.4µs ± 0%     157.9µs ± 0%     +103.98%  (p=0.000 n=10+9)
micro/memory_grow_mstore/by16       87.1µs ± 0%     170.9µs ± 0%      +96.21%  (p=0.000 n=8+9)
micro/memory_grow_mstore/by32       95.1µs ± 0%     181.5µs ± 1%      +90.75%  (p=0.000 n=8+10)
[Geo mean]                           447µs            511µs           +14.27%
```

### [](https://notes.ethereum.org/@ipsilon/evm-performance-report-geth-vs-evmone#evmoneAdvanced-to-evmoneBaseline-gas-rate "evmoneAdvanced-to-evmoneBaseline-gas-rate")evmone/Advanced to evmone/Baseline (gas rate)

```
name                             Baseline gas/s   Advanced gas/s        delta
main/blake2b_huff/empty               632M ± 0%        209M ± 0%      -66.97%  (p=0.000 n=9+9)
main/blake2b_huff/2805nulls           853M ± 0%        881M ± 0%       +3.29%  (p=0.000 n=9+10)
main/blake2b_huff/5610nulls           860M ± 0%        954M ± 0%      +10.82%  (p=0.000 n=9+9)
main/blake2b_huff/8415nulls           862M ± 1%        980M ± 0%      +13.63%  (p=0.000 n=9+10)
main/blake2b_shifts/2805nulls        1.06G ± 0%       0.97G ± 0%       -8.65%  (p=0.000 n=10+10)
main/blake2b_shifts/5610nulls        1.12G ± 0%       1.05G ± 0%       -5.87%  (p=0.000 n=9+9)
main/blake2b_shifts/8415nulls        1.21G ± 0%       1.14G ± 0%       -5.60%  (p=0.000 n=9+9)
main/sha1_divs/empty                  661M ± 0%        633M ± 0%       -4.22%  (p=0.000 n=10+9)
main/sha1_divs/1351                   704M ± 0%        743M ± 0%       +5.52%  (p=0.000 n=9+9)
main/sha1_divs/2737                   701M ± 1%        746M ± 0%       +6.47%  (p=0.000 n=10+7)
main/sha1_divs/5311                   706M ± 0%        749M ± 0%       +5.99%  (p=0.000 n=10+10)
main/sha1_shifts/empty               1.06G ± 0%       0.99G ± 0%       -6.04%  (p=0.000 n=9+9)
main/sha1_shifts/1351                1.10G ± 3%       1.23G ± 0%      +12.05%  (p=0.000 n=10+9)
main/sha1_shifts/2737                1.11G ± 0%       1.24G ± 0%      +11.64%  (p=0.000 n=9+9)
main/sha1_shifts/5311                1.11G ± 1%       1.25G ± 0%      +12.27%  (p=0.000 n=10+9)
main/weierstrudel/0                   277M ± 0%        206M ± 0%      -25.76%  (p=0.000 n=10+9)
main/weierstrudel/1                   358M ± 1%        305M ± 0%      -15.05%  (p=0.000 n=10+10)
main/weierstrudel/3                   358M ± 0%        320M ± 0%      -10.57%  (p=0.000 n=8+9)
main/weierstrudel/9                   359M ± 0%        338M ± 0%       -5.75%  (p=0.000 n=9+7)
main/weierstrudel/14                  359M ± 0%        344M ± 0%       -4.20%  (p=0.000 n=10+9)
micro/memory_grow_mload/nogrow        797M ± 0%        387M ± 1%      -51.48%  (p=0.000 n=9+10)
micro/memory_grow_mload/by1           784M ± 0%        386M ± 0%      -50.73%  (p=0.000 n=10+9)
micro/memory_grow_mload/by16          857M ± 1%        441M ± 0%      -48.47%  (p=0.000 n=10+9)
micro/memory_grow_mload/by32         1.13G ± 0%       0.59G ± 1%      -47.46%  (p=0.000 n=8+10)
micro/memory_grow_mstore/nogrow       813M ± 0%        394M ± 0%      -51.54%  (p=0.000 n=9+9)
micro/memory_grow_mstore/by1          799M ± 0%        392M ± 0%      -50.98%  (p=0.000 n=10+9)
micro/memory_grow_mstore/by16         870M ± 0%        444M ± 0%      -49.03%  (p=0.000 n=8+9)
micro/memory_grow_mstore/by32        1.12G ± 0%       0.59G ± 1%      -47.57%  (p=0.000 n=8+10)
[Geo mean]                            784M             589M           -24.89%
```
