Some of you may be familiar with the perspective that [MEV extraction is theft](https://www.coindesk.com/miners-front-running-service-theft). In this article, I’m going to go deep into my personal arguments for why [extracting MEV in cryptocurrencies](http://explore.flashbots.net/) isn’t like theft, why it is a critical metric for network security in **any** distributed system secured by economic incentives (yes, including [centralized ones](https://apnews.com/article/coinbase-stock-ipo-price-c3b802074ce4349b5bccf9ba43022800)), and what we should do about MEV in the next 3-5 years as a community. I will argue that MEV is fundamental, and that no known magic wand exists for remedying the problem.

![Image - 194749] | Wat Do? | Know Your Meme](https://i.kym-cdn.com/photos/images/original/000/194/749/pc131567047911_petichan.jpg)

my MEV journey

I will cover the relationship between [distributed fairness protocols](https://blog.chain.link/chainlink-fair-sequencing-services-enabling-a-provably-fair-defi-ecosystem/) and MEV, and the relationship between distributed fairness protocols and real fairness. I will argue that MEV is a sticky double-edged sword, broad enough that all attempts to reduce to “good” or “bad” are reductionist. Lastly, I will provide a roadmap for what we should do about MEV as a community, what I plan to do about it, and what **you** have to do about it, regardless of your role in the ecosystem.

**This piece represents my opinion alone. While I have numerous stakes in the MEV game, personal and professional, I encourage you to read the below AS AN OPINION PIECE and WITH A BUCKET OF SALT.**

Buckle up, Buckaroo!

## MEV is Fundamental

First, we will argue that MEV is fundamental to the cryptocurrency experiment. It is not going away. This logically follows from the cryptocurrency experiment for several reasons:

**Cryptographic transcriptability:** All distributed systems that can function as currencies must contain the key property of “[auditability](https://bitcoinist.com/zcash-inflation-bug-infinite-tokens/)“, or the ability to validate system state transitions and/or user actions. Auditability is the key property provided by the Bitcoin blockchain’s chain of Merkle-transaction-roots, as well as by many Ethereum smart contracts. We achieve auditability through the creation of cryptographic transcripts, or proof arguments, that can be validated by others. This is key to the way distributed economic systems are built, and without it, nobody would be able to validate anything about the system. Transcriptability and auditability, however, also open the door to MEV. There will always be users that prefer one version of events to another, and there will always be users to whom one transcript is more valuable than another. [These users will be able to condition payments to express their preferred outcome, and these payments (as well as the preferences themselves) furnish a form of MEV by altering the incentives of those selecting the transcript.](https://hackingdistributed.com/2018/07/02/on-chain-vote-buying/)

**Interoperability under heterogeneous trust:** There will never be a single blockchain architcture to rule them all (sorry, Maxis). In a world where multiple systems with multiple trust assumptions interoperate, the boundaries of these systems always provide value to those who are in privileged positions at those boundaries. For example, even if ETH had 0 MEV, there would likely be MEV available arbitraging Binance Smart Chain with ETH, especially available to those validating both. The same goes for MEV controlling Bitcoin transaction order and therefore [centralized exchange arbitrage in periods of high-market-activity](https://cointelegraph.com/news/xrp-price-spikes-to-0-90-crashes-in-seconds-as-coinbase-goes-down). For MEV to go away entirely, the world must operate within a single trust zone, which seems… unlikely.

[![](https://pdaian.com/blog/wp-content/uploads/2021/04/image-2.png)](https://docs.google.com/presentation/d/1oZcp2a1fQUdJSVAOrkXSZKr9f4h_Ll66ZUPtrd1e8cc/edit)

This image shamelessly modified from [Charlie Noyes’s thought-provoking presentation on interoperability x MEV.](https://docs.google.com/presentation/d/1oZcp2a1fQUdJSVAOrkXSZKr9f4h_Ll66ZUPtrd1e8cc/edit)  

**“Permissionlessness”:** One of the common design-patterns in cryptocurrency is to provide a payment that can be received by any user to provide an action useful to the network. This is the backbone of how many distributed applications operate. For example, in Uniswap, without the MEV provided arbitraging markets with external/centralized exchanges and fellow dexes, the price would not reflect the market and not provide a useful trading product for users. In MakerDAO, without the MEV that [pays liquidation bots](https://twitter.com/makerrekt), there would be no incentive to pay gas to update the state of loans in the system, leading to immediate systemic collapse. In Cryptokitties, there would be no birthing of bred kitties without the MEV that pays for the gas of birthing. The existence of a “bribe to anyone” (permissionless bribe) runs these systems, and indeed runs the blockchains they operate on themselves. MEV is therefore fundamental, and not always harmful. It is often an essential part of how decentralized/”permissionless” protocols operate and how they are secured, and _why_ they do not require permission to achieve [certain equilibria](https://arxiv.org/abs/1904.05234). Sometimes, we want MEV!

Note that other than permissionlessness, these properties can also exist in centralized systems. This should be obvious to anyone who thinks through MEV on a system like Binance Chain, and thinks about who controls who extracts what, but it also applies to distributed databases, non-financial systems, messengers, and any system in which a user may want to securely bribe a participant to perform a certain action. These bribes can be seen as generalizations of MEV.

## MEV must be extracted for economic security

Assuming you buy my arguments up until here, some amount of MEV will always exist in distributed systems. Let’s take this as an axiom, and reason only about worlds in which MEV exists and can be extracted ([such as, empirically, the world today](http://explore.flashbots.net/)).

Consider the economic security model of cryptocurrencies. Cryptocurrencies were [posited by Nakamoto](http://bitcoin.org/bitcoin.pdf) as “secure as long as honest nodes collectively control more CPU power than any cooperating group of attacker nodes.” When can we expect “honest” nodes (nodes that run the protocol without deviating) to outnumber attacking nodes? This is where the rubber of the Nakamoto-honesty model meets the road.

One reason why cryptocurrencies have been successful is [their security under weaker assumptions than honesty](https://medium.com/blockchain-at-berkeley/understanding-crypto-economic-security-through-game-theory-526e810c7736). If we were fine on relying on a majority-honesty assumption, say of public keys, we could simply use [permissioned consensus protocols](https://en.wikipedia.org/wiki/Byzantine_fault#Advanced_solutions) and establish a money abstraction. Instead, cryptocurrencies allow for existing in a world with more abstract economic assumptions than such “permission”.

One such economic assumption is that miners are incentivized away from engaging in dishonest mining behavior due to “[extraprotocol incentives](https://twitter.com/hasufl/status/1230523981247193090)” like the desire to convert mining profits into fiat in the long-term. Another such assumption is that CPU power costs money, so Nakamoto consensus imposes a non-trivial attack cost. No matter what security assumption makes you believe Bitcoin can work, it’s likely an **economic assumption**, not an **honesty assumption**. While the honesty may underlie the technical analysis, the economics underlies the _raison d’être_, and I would be willing to bet underlies the reason many of you are reading this post right now (directly or otherwise).

[Economics matters, and we don’t want to purely rely on identity or trust.](https://en.wikiversity.org/wiki/Cryptoeconomics) In this model, to incentivize people to behave the way that is desired, we as a community leverage the concept of economic rewards. Such incentives provide the security for any interesting cryptocurrency system that could not simply be replaced with a database. In a world of incentives where cryptocurrencies are secured by validators, the presence of MEV (extracted or not) inherently changes validator incentives. A validator who plays the economic game ignoring MEV is going to lose power systematically to [one who does not](https://www.coindesk.com/ethermine-adds-front-running-software-to-help-miners-offset-eip-1559-revenue-losses).

If we want any form of robust economic security, **every validator must extract available MEV at around the same rate**. Any validator that extracts MEV at a much higher rate than others is inherently concentrating economic rewards, and therefore influence in the system’s security. Any miner that extracts MEV at a higher rate is able to concentrate CPU control, and even more directly in proof-of-stake, stakers that are efficient concentrate staked capital.

A system where validators “leave MEV on the table” is one where an obvious attacker subsidy is readily available. This is self-defeating and degrades security in a pure-economic-rationality model, as it causes centralization/oligopoly. This degradation is even more severe in an [economically-expensive-to-attack-to-counter-nation-state model,](https://news.bitcoin.com/snowden-releases-nsa-documents-showing-bitcoin-1-priority/) because malicious nation states are likely totally happy to use MEV to lower attack costs for systems.

![North Korea's Kim ends freeze of nuclear and missile tests | Kim Jong Un  News | Al Jazeera](https://www.aljazeera.com/wp-content/uploads/2020/01/5725cee6113f42b4ad829dff628f13ce_18.jpeg?resize=770%2C513)

MEV extraction? Now you’re going too far…

To keep our economic assumptions strong, we must therefore keep MEV extraction **efficient** and **democratic**.

## MEV and Fairness

One natural question to those new to the MEV landscape is “[why not slap a fairness protocol on it (Layer 1) and be done?](https://blog.chain.link/chainlink-fair-sequencing-services-enabling-a-provably-fair-defi-ecosystem/)” Surely, there can be cryptographic protocols that create an ordering that is more fair and less manipulable than the current proof-of-work mining game?

Therein lies the rub, dear reader. Let us engage in this exercise of “adding a fairness protocol” in good faith. The first step is defining “a fairness protocol”. We need to specify a protocol to integrate into Layer 1. One natural way to choose the protocol is to read the [state-of-the-art academic literature on fair ordering](https://eprint.iacr.org/2020/269.pdf), which provides the following guide:

![](https://pdaian.com/blog/wp-content/uploads/2021/04/image-1.png)

Upon dissecting further, one sees that there’s no such thing as the “fair ordering protocol” the same way there’s no such thing as “the consensus protocol”. Instead, we have what appears to be a family of protocols, parameterized by network assumptions, fault thresholds, the sets of nodes which operate them, cryptographic assumptions, and more. Which one to choose, and which one to prefer?

The interesting conclusion here is that as long as two fair ordering protocols exist, so does MEV arbitraging these protocols. Note that the definition of MEV does not require honesty assumptions, and instead deals with profit-maximizing behavior only. If the validators in both protocols profit-maximize, they can corrupt the fair ordering protocols in a number of ways and profit.

This is similar to market design in many ways. [Many proposals exist for fairer market design.](https://faculty.chicagobooth.edu/eric.budish/research/HFT-FrequentBatchAuctions.pdf) They are awesome. I want to see them deployed, tested, experimented with. However, claims that any of these designs are silver-bullets for better fairness often stems from their analysis in an isolated model that does not adequately capture real-world complexity, and in a financially incentivized modality, [the differences become salient and profitable for those willing and able to exploit them.](https://hackingdistributed.com/2017/06/19/bancor-is-flawed/)

We need more experimentation, and more research. **Distributed fairness is a very promising research direction that could benefit all of humanity through market design.** But we should not exaggerate claims, either.

## Fairness and Fairness

When diving down this rabbit-hole, it helps to remember how we define fairness in the real world:

![](https://pdaian.com/blog/wp-content/uploads/2021/04/image.png)

Obviously, fairness must be impartial. But how can you have a single impartial protocol when multiple options exist, and people (rightfully) have different preferences?

I consider this equivalent to the [Libertarian pipe-dream of having a global property ledger](https://www.libertarian-labyrinth.org/proudhon-library/notes-on-what-is-property-2019/) from which all of the world’s justice hails. Sorry not sorry for the alienation, my opinion is that human-scale agreement on subjects this complex and impactful is impossible and absurd.

Back to fairness… we then have that some fairness protocols will work for some users. Others will work for others. Which protocol is more fair? The one [preferred by Group A or Group B](https://ourworldindata.org/war-and-peace)? Impossible to say a-priori.

Fair-ordering protocols may try to approximate a real-world notion of fairness, but they will never reach an idealized unparameterized definition, because removing the parameters requires humanity agreeing on what parameters are fair, which requires humanity agreeing on what is fair in the first place!

So, fair-ordering is at best a helpful approximation. Useful, but not a silver bullet to ordering woes, and certainly not one that can be universally deployed.

What’s worse is that [most fairness protocols rely on honesty assumptions](https://eprint.iacr.org/2020/269.pdf), and still do not carry robust economic arguments in favor of their operation that would be required to make them suitable for operations in settings such as cryptocurrencies. **A critical area of future work,** but not yet one on which we can afford to stake our financial future, and certainly not one that objectively satisfies the real-world definition of “fair” posted above.

## Fairness sucks and MEV is righteous!!!

Note that my blog post here exists to counter the prevailing obvious narratives that [a clear ethical divide, as well as a mutual exclusivity](https://www.coindesk.com/miners-front-running-service-theft), exists between MEV quantification, MEV extraction, and fair ordering.

**I am not trying to say fair ordering is uselesss**. Fair ordering under reasonable assumptions, especially those tailored to protocols, can likely massively reduce MEV and improve user outcomes. It is a critical area of study and research. Fair ordering, however, cannot completely eliminate MEV for the reasons stated above.

Our full solutions must go deeper, and must encompass more aspects of the problem in a far more robust way.

There is clearly a deep relationship between MEV and fairness. Some forms of MEV [seem patently unfair.](https://www.coindesk.com/everything-you-ever-wanted-to-know-about-the-defi-flash-loan-attack) But some may actually incentivize users or validators to run fair ordering protocols. Indeed, running a fair ordering protocol on Ethereum would likely require paying validator rewards that would furnish MEV directly to L1. **In a twist of irony, it may be impossible to achieve economically-secure proof-of-work fairness without MEV!** (this is just my conjecture, and nobody has proven it _yet_…)

From a research perspective, these relationships remain poorly understood. However, they are obviously intuitively there, and they warrant a lot of study, debate, and thought.

## Wat do?

![I Want YOU!” – The Story of James Montgomery Flagg's Iconic Poster -  Illustration Chronicles](https://payload.cargocollective.com/1/0/4579/12920296/prt_300x300_1500464488_2x.jpg)

frfr tho….

So now, finally, a prescription. And yes, it’s for you. And yes, you should take it (I’m not P🤡rdue Pharma, I promise):

**DAPP DEVELOPERS:** Limit MEV to what is _strictly required/inherent_. Think carefully about MEV in your protocol. Where does it exist? How can it harm users? How can it benefit users? Can it be redistributed or minimized? How does inter-operation with other systems change this over time? And how is this communicated to users? Figure your shit out, and do it properly, please, or your distributed system **will break**.

**USERS:** Demand the above. Don’t use dapps with predatory MEV. Understand what MEV means to you and your transactions. Learn about MEV, and dive deep into the systems you use yourself.

**L1/L2:** Redesign with MEV in mind. Think about the use cases you are designing for, and how best to manage MEV. Think about whether you have a plan for how changing MEV incentives as a result of dapps being deployed may change your economic incentives, and in what ways MEV may violate assumptions you are making around network participants.

**Miners:** [Extract, extract, extract, and don’t be ashamed.](https://github.com/flashbots/mev-geth/) Your continued extraction is required to the game-theoretic security of any network on which MEV exists. I would also recommend further aligning incentives and expectations by enforcing a code of conduct on fellow miners/network participants if necessary, as short-term MEV extraction can harm long-term incentives that have the ability to benefit everyone in the ecosystem.

**Everyone else**: Help us democratize the profits of MEV, while keeping overall system users from being harmed as much as possible. Stay involved in the conversation, and keep trying to [mitigate 👏 negative 👏 externalities.](https://medium.com/flashbots/frontrunning-the-mev-crisis-40629a613752)