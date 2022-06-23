---
created: 2022-03-18T18:16:19 (UTC -07:00)
tags: []
source: https://blog.0xproject.com/a-comprehensive-analysis-on-dex-liquidity-aggregators-performance-dfb9654b0723
author: Fulvia
---

# A comprehensive analysis on DEX liquidity aggregators’ performance | by Fulvia | 0x Blog

> ## Excerpt
> In the past few weeks, there has been a surge of new liquidity sources, with over $4Bn capital spread across more than 15 sources and new ones coming up every week. As a result, DEX liquidity has…

---
## _After accounting for gas costs, 7 out of 10 times 0x API offers better prices than 1inch, Dex Ag, Paraswap and Uniswap._

![](https://miro.medium.com/max/1400/1*-pWsPrFhA6UnsH-dAJP6VQ.png)

## TL;DR

-   **7 out of 10 times 0x API offers better adjusted prices than 1inch, Dex Ag, Paraswap and Uniswap.** That is, the price after transaction fees are paid.
-   **0x API v1 delivers market-leading gas efficiency.** With a custom, highly-optimized architecture, accessing **Uniswap’s liquidity through 0x API is now as cheap and often cheaper than going to Uniswap directly**.
-   Aggregators exist to remove the mental hurdle of finding the right platform to trade by guaranteeing that you get the best the ecosystem has to offer. However, **not all aggregators are created equal: they have different levels of performance in terms of pricing, quoted price accuracy, gas usage, revert rates and response times.**
-   **The best quoted price doesn’t always result in the best adjusted price.** Just like buying an airline ticket, you need to be careful about hidden fees/costs. While the base ticket might be cheaper, when you price in all of the add-ons (baggage fees, preferred boarding, seat assignments, etc.), you might end up paying more.
-   **0x API and** [**Matcha**](https://matcha.xyz/) — which uses 0x API under the hood — are two products launched by 0x Labs to **offer DeFi developers and traders a simple way to buy and sell assets at the best price possible across all DEX liquidity sources.**

## DEX liquidity wars and the rise of aggregation

In the past few weeks, there has been a surge of new liquidity sources, with over $4Bn capital spread across more than 15 sources and new ones coming up every week. As a result, DEX liquidity has become more fragmented than ever, making it more complex for DeFi users to find the best price for a given trade.

This in turn has led to the rise of aggregation, which has grown significantly over the last 12 months — [representing ~9.6% of total DEX trading volume in Q3](https://explore.duneanalytics.com/dashboard/dex-aggregators). However, aggregation is still somewhat misunderstood. For example, while **_quoted price_** is often used as the only metric to compare aggregators’ performance, it is not uncommon for some aggregators to quote $100 and consistently settle at $90. Quoted price does not reflect what the user ends up with, so it poses the question if this is the right metric to compare aggregators’ performance.

We aim to bring more transparency to the industry and open up the conversation on how to objectively measure and compare the performance of aggregators.

## Adjusted price — the golden metric for aggregators

We believe the following principles must guide a team building a high-performing, professional-grade aggregator:

1.  **Accuracy:** An aggregator should be accurate — that is, the price at settlement should equal the priced quoted. Slight variance is acceptable as there will always be market movements between the time a trade is submitted and the time that trade is confirmed on-chain. However, all else equal, an aggregator that quotes more accurate prices is a better product. What you see should be what you get.
2.  **Low cost:** More work performed on-chain results in more gas used and that in turn leads to higher gas cost. An aggregator should aim to lower the overhead as much as possible to ensure the user is getting the best deal. An aggregator should also adapt to the current gas price market and adjust routes accordingly. Every dollar spent on gas should earn the user more.
3.  **Low revert rates**: an aggregator protects its user by maintaining low revert rates across all trade sizes. It is a frustrating experience when a transaction fails as the user still pays for gas.

Ultimately, we combine these principles into a single golden metric: **_adjusted price_**.

In this study we analyze the above mentioned metrics individually, but we focus primarily on adjusted price. **Adjusted price is what a user receives after a transaction is completed on-chain and transaction costs (gas and fees) are paid.** The adjusted price can be calculated using the following formula:

_adjusted\_price = sold\_amount\_usd / ( bought\_amount\_usd — transaction\_costs\_usd )_

We believe this is a better metric to use for comparing aggregators as it is what the user ultimately gets in their wallet. Let’s use an example to illustrate this:

A user wants to trade 100 DAI for USDC. **Aggregator #1 quotes 100 USDC with a transaction cost of $5**. Using the formula above, we can determine that **adjusted price = 100 / ( 100–5 ) = $1.05**.

The same user checks out a second aggregator. **Aggregator #2 quotes 104 USDC for the same 100 DAI, but with a transaction cost of $10**. In this case, **adjusted price = 100 / (104–10) = $1.06**.

**In this example, Aggregator #1 wins as it ultimately finds a cheaper rate for the user to convert 100 DAI to 100 USDC.** While Aggregator #2 initially quoted a better rate and bought more tokens, you can observe how the transaction cost impacted the adjusted price. In conclusion, spending an additional $5 to only receive an additional 4 USDC resulted in a more expensive deal for the user.

We simulated 59,975 trades and calculated adjusted prices across different exchanges. Here’s what we found:

![](https://miro.medium.com/max/1400/1*NsOfD-0OZn_ZhMBKDtoK6g.png)

Adjusted prices winning rates for 0x API, 1inch, Dex Ag, Paraswap and Uniswap. 0x API outperforms other exchanges ~72% of the time by offering better prices after accounting for gas.

0x API outperforms other exchanges, on average, 72% of the time. This means that, **a trader using a 0x API integrator like Matcha will receive the best price available across all DEX 7 out of 10 times.**

When it comes to adjusted prices, we observe a difference in winning rates as the value size increases. For low value trades, Uniswap performs well since transaction costs are cheap relative to most aggregators. As trade sizes increase, aggregators perform their function and begin to split trades across multiple different sources. 0x API performs well on all trade sizes and other aggregators start to become competitive as the trade sizes increases. An aggregator like 1inch, for example, becomes increasingly competitive with 0x API in larger trade sizes due to their large number of supported sources.

## The methodology

The data shared in this section comes from a simulation tool we built at 0x Labs that allows us to continually quote and execute trades for various token pairs, trade sizes and confirmation delays. For the more technical crowd, the simulation is done by performing an `eth_call` on a Geth node and using state overrides to fund and track swaps, which includes metrics such as tokens bought and sold as well as gas used.

The exchanges under observation are 0x API, 1inch, Dex Ag, Paraswap and Uniswap. For this study we used the [20 most traded pairs across DEXs](https://explore.duneanalytics.com/queries/10293#20454) and, using their publicly available APIs, we sampled an average of 11,995 trades per exchange. Trade sizes ranged from $100 to $1M and settlement delays varied between 0 and 180 seconds (simulating the time spent waiting in the mempool). The simulation ran for a period of ~78 hours during which gas price fluctuated between 47 and 194 gwei, this allowed us to ensure the study reflected market conditions.

In the visuals presented throughout this report, 0x API is represented by the green colored bars and Matcha (0x API). A winner is considered a winner if it outperforms the other exchanges by more than 0.5 basis points. Prices within 0.5 basis points of the best price are considered a tie. A “tie with Matcha” means that 0x API and one or more exchanges tied. “Tie without Matcha” means that two or more exchanges, excluding 0x API, tied. Colors are consistent throughout the graphs.

## Accuracy, gas usage, revert rates and response times

In this study we analyze three different prices, so let’s define them:

-   **Quoted price** is the price initially advertised for a given trade — the “sticker” price users see in the UI.
-   **Realized price** is the price that is settled after a delay. The time it takes between the trade submission and confirmation on-chain is typically between 0 to 180 seconds.
-   **Adjusted price** is what a user gets once a transaction is realized and transaction costs (gas and fees) are paid. In other words, the realized price minus costs.

We will start by looking at quoted and realized prices.

![](https://miro.medium.com/max/1400/1*5CND9rsHRM9EViXz0h1bNg.png)

Quoted and adjusted prices winning rates for 0x API, 1inch, Dex Ag, Paraswap and Uniswap. Many aggregators seem to optimize for quoted price, but the best quoted price doesn’t always result in the best adjusted price.

In the graphs above, you can notice that Uniswap is no longer present in any of the trade size buckets. It is in stark contrast to the previous graphs showing the adjusted price winners. In other words, while it does not display the best price, sometimes it offers better adjusted prices for trades under $5K due to low fees. As mentioned before, this winning percentage shrinks significantly in the higher value trades as aggregators perform better by splitting trades across sources to find the best price.

You can also observe in these graphs that 0x API’s winning rate in quoted and realized prices is significantly lower than its winning rate for adjusted price — averaging 19% for quoted and realized prices and up to 57% and 63%, respectively, when we include ties. 0x API’s average winning rate for adjusted prices is ~72%. The reason why the winning rate is lower for quoted and realized prices is that we made an explicit trade-off in the design of our smart order routing logic to account for gas costs when calculating the trading route for an order. **By accounting for gas costs in 0x API’s smart order routing, we guarantee that the benefit in price outweighs the cost of sourcing liquidity from various sources.**

Unfortunately, many aggregators appear to optimize to win quoted price, that is, the “sticker” price, as it enables them to attract more users. However, as exemplified at the beginning of this report with Aggregator #2, aggregators who do not account for the cost of the transaction perform better in quoted price as they are willing to spend more, but it actually results in less value for the user. As a result, **the best quoted price doesn’t always result in the best adjusted price.**

A good analogy for this is airline tickets. Have you ever booked a flight through a certain airline because it seemed cheaper than the others, but then you had to pay extra for choosing your seat and then a little more for checking in your luggage? While the base ticket might be cheaper, when you price in all of the add-ons you might end up paying more. A similar thing happens with some aggregators: the price seems to be better than others when you get the quote, but once the transaction is completed you end up with a lot less money than you should have.

Wait, but why?

Let’s take 1inch for example, since they have the highest winning rate for quoted and realized prices. Their winning rate shrinks significantly from quoted to realized to adjusted price, declining from 22% to 19% to 14%. This means that **while 2 out of 10 times it _might appear to be_ a better deal, once settled and minus costs, it is not often the case**.

The variance between quoted and realized prices can generally be explained by either inaccuracy or changes in the market that occur during the time delay between quote and on-chain settlement — which can make the quoted price no longer available. The variance between realized and adjusted is simply the additional cost incurred by the aggregator.

Let’s take a look at accuracy across all of the exchanges. When an exchange quotes a certain price, how often can they deliver?

![](https://miro.medium.com/max/1400/1*7QO7-eFQrM7pBqOCxAOLVQ.png)

0x API and Uniswap consistently have positive variance (i.e above 0), which indicates that there was a price movement between the quote and the on-chain settlement of the trade in the user’s favor. This is known as positive slippage. When this occurs, 0x API always transfers the positive difference to the user.

Accuracy is important as it builds user trust and it demonstrates the price quoted by an aggregator is the price the end user will receive. Given enough time and samples, variance should approach 0.

As seen in the above graph, 0x API and Uniswap consistently have positive variance (i.e above 0), which indicates that there was a price movement between the quote and the on-chain settlement of the trade in the user’s favor. This is known as positive slippage. When this occurs, **0x API always transfers the positive difference to the user.**

Interestingly, 1inch’s variance is always below 0, which indicates that users never receive a better rate than the quoted price, regardless of how markets move. A variance below 0 could also indicate that an aggregator is using stale information to provide quotes or their response times are not fast enough. These issues may lead to a higher chance of transactions reverting. In a reverted transaction no tokens are bought or sold, but the user still has to pay the transaction fee. We will come back to this in detail later.

Now, let’s turn our attention to costs, which in Ethereum is measured as gas used. In the graph shown below we observe that there is a large amount of variance in the amount of gas used across all of the exchanges compared in this study.

When it comes to aggregators, the old adage “less is more” applies. It is good for an aggregator to spend more if it extracts more value for the user, but as we presented earlier in the adjusted price winners graph, most aggregators are unable to do so. Optimizing the code to lower the gas usage increases the likelihood that an aggregator can split the trade over more sources to extract more value.

![](https://miro.medium.com/max/1400/1*4MEkKoTPqCanzjSehIA8gg.png)

Despite increased gas usage, 0x API is able to extract more value for the users 7 out of 10 times. All other aggregators are at least double the cost of 0x API — which means it is double the cost for the user, with little additional value extracted.

Uniswap is one of the cheapest exchanges in terms of gas usage, so it is a goal all aggregators should aspire for, especially for the lower value buckets. It is extremely unlikely that a $100 trade will be split over multiple exchanges, even when gas prices are 1 gwei.

On average, 0x API has an increased cost over Uniswap. However, as shown in the adjusted prices winning rate graph, **despite the increased gas usage 0x API is able to extract more value for the users 7 out of 10 times**. **All other aggregators are at least double the cost of 0x API — which means it is double the cost for the user, with little additional value extracted.** The highest gas guzzlers are 1inch and Paraswap, which is reflected in their low winning rate for adjusted prices.

To put things in perspective, below is the average cost in dollar value of all simulated trades between $500 and $1K:

-   0x API — $6.43
-   1inch — $13.58
-   Dex Ag — $13.41
-   Paraswap — $18.59
-   Uniswap — $3.93 (we will come back to this later)

Reducing gas usage has been a focus for our team for the past eight months. Our initial target was to cut down transaction costs by 50%, however, in 0x API v1 we achieved to **cut down transaction costs, on average, by ~76%** vs v0. Gas usage is currently around 460K gas for a $50K trade, down from 1.5M gas for the same trade in v0. These optimizations have turned 0x API into the most efficient aggregator across all trade sizes.

In fact, gas prices have been quite higher than 100 gwei in the last month. We wanted to show you how 0x API’s adjusted price winning rate looks like in this scenario. We ran a second simulation artificially bumping gas prices and the results show that, **0x API outperform competitors in a high gas price environment too — offering better adjusted prices 8 out of 10 times when gas prices are above 400 gwei.**

![](https://miro.medium.com/max/1400/1*4D-NqY4GAjplRNRkwJHNnA.png)

0x API outperform competitors in a high gas price environment too — offering better adjusted prices 8 out of 10 times when gas prices are above 400 gwei.

Now, back to 0x API’s gas usage in comparison to Uniswap. On average, 0x API uses more gas than Uniswap and yet, **0x API offers better adjusted prices than Uniswap 7 out of 10 times**. How is this possible? At 0x Labs **we built a super optimized contract that allows 0x API to source liquidity that is 100% routed through Uniswap in a more gas efficient way than the default Uniswap contracts**. This is particularly beneficial for smaller trades (<$5K), which are typically routed to a single source.

To exemplify this, let’s take a look at the following two transactions:

![](https://miro.medium.com/max/1400/1*KMsE-bQcdnhKZunVsm0thQ.png)

0x API v1 delivers market-leading gas efficiency. With a custom, highly-optimized architecture, accessing Uniswap’s liquidity through 0x API is now as cheap and often cheaper than going to Uniswap directly.

The [first transaction](https://etherscan.io/tx/0xb567dd4262461e998cf98b61f439bfe6112feafc36e5c24488c6a9262b990365) (left) shows an ETH<>UNI trade made by going directly to Uniswap. The [second transaction](https://etherscan.io/tx/0xa76223b08efc5c293437b8f5673ed2b2757ba7f9c405dab5c98620c463ed9b57) (right) shows an equivalent ETH<>UNI trade being 100% routed to Uniswap via 0x API. As shown in the “gas used by transaction” field, the Uniswap transaction cost ~114K gas and the 0x API transaction cost ~110K gas, that is 3.7% cheaper. Certain on-chain state and conditions can change the results (for example, if the user needs to set new allowances), but in general, **using 0x API to trade on Uniswap is now as cheap and often cheaper than going to Uniswap directly**.

So far we’ve covered accuracy and gas usage, so it’s time to talk reverts. When a user makes a trade through one of these exchanges, what is the likelihood of the transaction failing?

Aggregators should aim to keep revert rates as low as possible as it is a frustrating experience when a transaction fails and the user still pays for gas. Common causes for reverts are: changes in the market between the time of quote and on-chain settlement, and inaccuracy of quoted prices (especially when using stale information to provide quotes).

Our simulation tool always executes swaps under perfect conditions, that is: fully funded, unlimited gas and no front-running (i.e. no information is leaked to the _dark forest_). We exclude reverts caused by these reasons as well as those that in real life can occur due to user behavior. As a result, simulated revert rates are lower than real world revert rates for all the exchanges, but nevertheless, they provide valuable insights into the exchanges’ performance.

![](https://miro.medium.com/max/1400/1*5hNjjhZkxMc5muWQfwkmSA.png)

A delay below 10s essentially means that the trade executed immediately or in the very next block after the transaction was submitted. A high revert rate in this delay bucket reflects inaccurate quoting as the window for market movements is very small.

In the graph above the 0–10 seconds delay bucket is particularly interesting to look at. A delay below 10s essentially means that the trade executed immediately or in the very next block after the transaction was submitted. A high revert rate in this delay bucket reflects inaccurate quoting.

In general, we observe that Paraswap’s revert rates are the highest in comparison to the other exchanges, followed by Dex Ag and 1inch. One possible explanation for reverts in a simulated environment where human behavior is not a cause for reverts, is that the price they quoted is gone (as seen in higher delay buckets) or never existed (as seen in lower delay buckets). If you recall, Dex Ag and Paraswap had the most inconsistent accuracy ratio.

Let’s explore the impact of revert rates for an end user. If a user makes 100 trades between $500 and $1K on each exchange, they would spend on failed transactions the following average amounts:

-   0x API — $6.89
-   1inch — $62.52
-   Dex Ag — $3.47
-   Paraswap — $119.22
-   Uniswap — $0.21

In the real world we have seen Uniswap’s daily revert rate up to [30%](https://explore.duneanalytics.com/queries/10443/source#20754), whereas in our simulation the highest revert rate we observed for it was ~1.3% for the longer delay buckets. The difference between realized and simulated revert rates speaks to a huge opportunity to improve the UX in all DEXs, but that’s a subject for a different discussion.

In the case of 0x API, the simulation shows we have, on average, the lowest simulated revert rates among aggregators across all the different delays. 0x API simulated reverts range between 0.6% for the shorter delays and 1.8% for the longer ones. In the real world, we have observed that the average weekly revert rate is typically between ~3% and 4% for a 0x API integrator like Matcha.

![](https://miro.medium.com/max/1400/1*gsA6j0GaAM3QFMR89fnJlA.png)

Matcha’s average weekly revert rate is typically between ~3% and 4%, 10X lower than other exchanges.

Now, there’s one important metric we haven’t covered yet: response times. This metric might not be immediately obvious to end users of an application like Matcha. However, **when you are building a product on top of third-party infrastructure, you need the service to be as reliable and fast possible**. Aggregators that are faster allow a user to get their trades in quicker, resulting in lower slippage and reduced chance of reverts. Not to mention the improved user experience.

There’s a technical challenge to maintaining low response times when we have to sample from more than 13 sources, including [RFQ](https://blog.0xproject.com/growing-defi-with-the-help-of-professional-market-makers-0xs-new-request-for-quote-system-8d5fb08594bf), a system where professional market makers provide liquidity. However, our goal is for 0x API to become the best professional grade liquidity aggregation API and offer response times around 1.5s is something we are proud of.

![](https://miro.medium.com/max/1400/1*qj6Mzch-7c6pCKGdqcF2pg.png)

When you are building a product on top of third-party infrastructure, you need the service to be as reliable and fast as possible. 0x API response times are consistently below 1.5s, up to 2.7X faster than other aggregators.

1inch‘s response times are the highest from all the exchanges analyzed, 2.7X times higher than 0x API’s response times. At 4s, 1inch’s response times are about one third the time it takes to mine a new block, which can make or break a time-sensitive swap — for example, in high volatility days when market conditions are in flux.

Given all of the information above, what is a simple summary of the best aggregator to choose? We can formulate the expected value of gains from switching from one exchange to another by calculating the average gain or loss from the switch across a wide range of trade scenarios. This will give an indication of the benefit a user can expect.

For users choosing to trade with a 0x API integrator like Matcha instead of one of the other exchanges, the average expected gains for a trade under $10K are:

-   1inch — $7.17
-   Dex Ag — $10.36
-   Paraswap — $14.89
-   Uniswap —$17.14

**Overall, a user trading on 1inch that switches over to a 0x API integrator like Matcha will observe, on average, an expected gain of $7.17 per trade. If they switch over from Uniswap, the average expected gain increases to $17.14 per trade.**

## Closing thoughts

First of all, we are glad you found this interesting and made it all the way here!

Our goal is to share what we have learned building 0x API and shine a light on key differences between aggregators that might not be obvious at first sight. As we mentioned earlier, we want to bring more transparency to the industry and open up the conversation about how to objectively measure and compare the performance of different aggregators so that DeFi developers and traders can make informed decisions.

By opening up [our data](https://metabase.spaceship.0x.org/public/dashboard/1b0e6e83-254c-4863-a327-16d515688d0c?run_id=buy-disney-timeshare&win_tolerance_(bps)=0.5), we hope to move the DEX aggregation space forward and continue building what is best for DeFi developers and users.

We know this is a lot of information to take in and that there will likely be a lot of questions, so we are hosting a [Reddit AMA](https://www.reddit.com/r/0xProject/comments/j3hyai/ama_0x_api_v1_october_7th_2020_at_12_pm_pt/) on Wednesday October 7th at 12pm PT.

If you are a trader, go to [Matcha](https://matcha.xyz/) and start saving on your trades today!

If you are a DeFi developer looking to source liquidity across DEXs, you can find our v1 announcement [here](https://blog.0xproject.com/announcing-0x-api-v1-34493d9bad3b) and the 0x API docs [here](https://0x.org/docs/api#swap). If you are an existing integrator using v0, the migration process to v1 is quite simple. We wrote [this guide](https://0x.org/docs/guides/migrate-0x-api-from-v0-to-v1) just in case.

For media inquiries, please email [clay@0xproject.com](mailto:clay@0xproject.com).

_Massive thanks to Ben, Jacob, Kroeger, Lawrence and Michael for their invaluable contributions and feedback on this piece._

## 🤝 Join the 0x community on…

[Twitter](https://twitter.com/0xproject?source=post_page---------------------------)⠀[Discord](https://discord.gg/d3FTX3M?source=post_page---------------------------)⠀[Telegram](https://t.me/official_0x)⠀[Reddit](https://www.reddit.com/r/0xproject?source=post_page---------------------------)⠀[YouTube](https://www.youtube.com/c/0xproject?source=post_page---------------------------) ⠀[LinkedIn](https://www.linkedin.com/company/0x?source=post_page---------------------------)

## 🛠️ Build on 0x

[Get started with our dev docs](https://0x.org/docs?source=post_page---------------------------)⠀[Contribute on GitHub](https://github.com/0xProject?source=post_page---------------------------)

## 🙋♀🙋♂️️ Work at 0x

[Join us in our mission](https://0x.org/about/jobs?source=post_page---------------------------)⠀[Check out our open positions](https://boards.greenhouse.io/0x?source=post_page---------------------------)
