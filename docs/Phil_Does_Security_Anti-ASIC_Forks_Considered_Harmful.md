A super hot topic in the [Ethereum](https://www.ccn.com/ethereum-users-support-hard-fork-to-brick-ethereum-asic-miners-poll/)/[Monero](https://www.trustnodes.com/2018/03/29/monero-emergency-pow-hardfork-asics-privacy-trial)/not-Bitcoin-cryptocurrency discussion space these days is the rise of ASICs for several coins: first Monero, then allegedly Ethereum, a coin that was designed to be and was sold as [ASIC resistant.](https://en.wikipedia.org/wiki/Ethash)  What are innocent coinholders to do in the event of such [evil companies like Bitmain](https://imgur.com/a/sFHnI) taking over decentralized mining ecosystems with their chips of doom?

In this post, I’m going to provide some counterpoints to the fervor that is gathering around repeated forks for ASIC resistance.  I’m going to argue that in the long term, these games won’t improve specialization or economies of scale, won’t be a productive use of developer time, and may pose the risk of _increasing_ centralization.  Finally, and perhaps most disturbingly, they make attacks on cryptocurrencies cheaper, primarily for sophisticated and well-resourced attackers like nation states.

## Rage Against Economies of Scale

The fundamental economic question at issue here is whether the cryptocurrency community can prevent anti-competitive [economies of scale](https://en.wikipedia.org/wiki/Economies_of_scale) from forming around mining, raising entry barriers to hobbyists and small operations while centralizing power and profit to large firms who will then hold [undue control over cryptocurrency evolution and governance.](https://www.reddit.com/r/Bitcoin/comments/6na8rf/analysis_of_jihad_wus_attack_on_bitcoin/)

In general, I do not believe such economies of scale are avoidable in any competitive mining ecosystem.  Miners have a variety of possible ways to gain “edges” over other miners, using even small differences in profitability at scale to dominate (and eventually purchase, if necessary) their competition, potentially centralizing production in an asymmetry of interest or resources.  Here are some example optimizations that are more accessible to large businesses than hobbyists:

1.  Electricity cost – The obvious one.  In some places, electricity is [just cheaper](https://www.theatlantic.com/technology/archive/2018/03/bitcoin-mining-arbitrages-cheap-electricity-into-money/555416/), and preferential agreements can be struck with governments that make mining anticompetitive in large swathes of the world with real, free-market electricity costs.
2.  Datacenter labor/cost – Running the huge datacenters associated with mining [isn’t cheap](https://qz.com/1055126/photos-china-has-one-of-worlds-largest-bitcoin-mines/); sophisticated actors can compete on labor price per GPU in operation, especially since these datacenters often require complex cooling systems or on-site repair shops for often crates of faulty equipment.
3.  Regulatory arbitrage – Running a large mining operation has [regulatory](https://www.ft.com/content/adfe7858-f4f9-11e7-88f7-5465a6ce1a00), accounting, and compliance costs, just to name a few; in many regions or for many businesses that are well connected, personal relationships or ability to operate in lax regulatory climates may naturally centralizing to these areas.
4.  Manufacturer relationships – Even if GPUs are all that’s mining, evil mining corporations and cartels [can still negotiate deals](https://cointelegraph.com/news/china-mining-co-bitmain-shows-higher-2017-profits-than-us-gpu-giant-nvidia-report-finds) that give them advantages in production from general-purpose manufacturers.  This includes the ability to rent GPUs at preferential rates, if renting is done at scale.
5.  Cybersecurity / trust / ability to “pool” funds and hashpower – The real bad boy centralization vectors in mining are not individual miners, they are pools, who ultimately have control over the content of blocks.  Some pools can compete simply by offering better security/block propagation/software/fraud detection/etc.  Removing all these barriers to entry will likely never happen.
6.  Frontrunning / ordering manipulation – Miners are in a privileged position to [frontrun decentralized exchanges](http://hackingdistributed.com/2017/08/13/cost-of-decent/), and many sophisticated miners are moving to profit off such games today (yes, I have proof; no, I can’t release it yet).  As cryptocurrency systems truly grow in value, sophisticated hedge funds and arbitrage actors may be able to partner with sophisticated mine operators to outcompete miners not using such strategies.
7.  In-protocol games, like [selfish mining](https://www.cs.cornell.edu/~ie53/publications/btcProcFC.pdf) – Generally, non-honest but rational/profitable strategies in mining rely on scale, introducing natural and substantial economies of scale.  Such economies have not been explored in practice, likely due to extraprotocol incentives, but it is very possible they _will_ be explored in the future as cryptocurrency systems become more entrenched and less fragile, reducing extraprotocol incentives to be honest for the “greater cryptocurrency good”.
8.  … and many more!

So, even if one vector supporting return to scale is removed (assuming that all ASICs can be defeated in the ideal case), the problem is not solved; manufacturers will still compete for and **realize** economies of scale, likely pushing mining further into the oligopolies observed even today on profitable GPU-mined coins:

[![Gee, that's a decentralized mining game you've built on GPUs... shame if something were to happen to it...](https://pdaian.com/blog/wp-content/uploads/2018/04/Screenshot-from-2018-04-02-00-40-58.png)](https://etherscan.io/stat/miner?range=7&blocktype=blocks)

_**Gee, that’s a decentralized mining game you’ve built on GPUs… shame if something were to happen to it…**_

## Stop Talking ASICs, Say Specialization

A lot of the current distraction seems to come in a distinction that is drawn between GPU mining and ASIC mining.  While it is true that a distinction currently exists between coins that are most profitable to mine on general purpose hardware, and coins that merit specialized hardware for their PoW-based security, this distinction is inevitably impermanent.  In the long term, specialization will arise in _any_ mining process and _any_ competitive market in which specialization is possible.

Imagine the plan of repeated hard forks works, and ASICs are never built for some coin (call it Ponero).  Ponero enjoys the “decentralization” associated with GPU mining, and happily ticks along with four pools dominating transaction sequencing and block production; pools that happen to both enjoy and lobby hard for the regular hard for process preventing the entry of more sophisticated competition whose “culture” they dislike.

Unfortunately, the manufacturer they are censoring, BitPrimary, is far more technologically sophisticated and well-connected than Ponero or any of its present mining pools.  BitPrimary is able to strike a deal with nVidia and AMD, with access to top-binned GPUs early, and production of specialized GPUs that don’t contain any of the “general purpose display equipment” and other niceties required by PC gamers anyway.  In the long term, if cryptocurrency mining contains to be an [O(billions)](https://medium.com/@jimmysong/just-how-profitable-is-bitmain-a9df82c761a) profitable business, GPU companies may even cut out the middleman and get in the games themselves (and have also said [they may explore cryptocurrency-specific lines](http://www.livebitcoinnews.com/amd-nvidia-envision-grim-future-gpu-mining/)).

At what point does such specialized hardware philosophically become an ASIC?  If it’s not a commodity GPU operation, nor an ASIC operation, what are we to call such specialized-GPU-at-scale operations?  Should we oppose those too, insisting on commodity hardware _only_?  Would the manufacturers of this commodity hardware doing preferential access deals or [binning agreements](http://www.overclock.net/forum/74-graphics-cards-general/767998-what-gpu-binning.html) violate a coin’s social contract against specialization and the centralization it brings?  The answers remain unclear and likely divisive.

No coin will be able to fight specialization through hard forks, and the only achievable result may make it **more difficult to specialize**, and maybe even make the **advantages of scaling up investment smaller.** Unfortunately, this actually _benefits_ sophisticated actors by allowing them to specialize on the available edges better than any competition, and to use even small competitive advantages to drive out opposition until they end up with the same control as sophisticated ASIC manufacturers with hardware edges.

The companies that are mining ASICs on Bitcoin today are [figuring out how to out-specialize hobbyists on GPU mining](https://www.bitsonline.com/antminer-g2-gpu-mining-rig-review/) _as we speak_, and a few rinky-dinky hard forks are not going to stop them.

## PoW Governance Considered Harmful

A key question if a community **does** decide to pursue a strategy of repeated ASIC forks is the [governance](https://medium.com/@Vlad_Zamfir/against-on-chain-governance-a4ceacd040ca) question; how are new PoW algorithms chosen, how are they evaluated for security, how are they tested for consistency by client teams, and how are they proposed to/”voted” on/adopted by a wide network of decentralized users with a [complex governance model](http://hackingdistributed.com/2016/01/03/time-for-bitcoin-user-voice/).

This process in and of itself provided a potent attack vector for adversaries with strong incentives.  It is quite obvious to any cryptographer that simply choosing an arbitrary hash function is not sufficient; even boring implementation-level engineering details of hash functions often lend themselves to [secret optimizations](https://arxiv.org/pdf/1604.00575.pdf) and “[miner edges](https://cointelegraph.com/news/bitcoin-mining-hardware-profit)” that can take **years to discover** and may be profitable even in the absence of ASICs.

So how to choose new hash functions?  How to make sure they are “fair”?  What does “fair” even mean and how can “fair” be evaluated for complex functions with few experts in the world truly understanding their internals?  One approach is to just let the nebulous “community” decide, absconding all personal responsibility for such a decision and thus its consequences:

[![Riccardo Spagni, Twitter:   It might entirely be less secure, but the community has made the hard call, and all we can do is see what happens. It’s the same as Monero’s tail emission - maybe that turns out to be unnecessary, or even net negative, but in the absence of strong evidence we must try our best.](https://pdaian.com/blog/wp-content/uploads/2018/04/Screenshot-from-2018-04-01-15-18-43.png)](https://twitter.com/fluffypony/status/975326552262107136)

_**The Community has ruled**_

Of course, this punts on all the hard and relevant governance questions.  Who are this nebulous “community”?  How was this actually decided?  Measured?  How were the minority rights of users who _did not_ want a fork respected?  Was this after all just a centralized decision, adopting the guise of “community driven” to avoid legal risk?

Any competent nation state has several skills: [backdooring cryptographic primitives](https://en.wikipedia.org/wiki/Dual_EC_DRBG), [human espionage / organizational infiltration](https://www.nytimes.com/2013/08/30/us/politics/leaked-document-outlines-us-spending-on-intelligence.html), and [manipulation of governance processes](https://www.channel4.com/news/factcheck/americas-long-history-of-meddling-in-other-countries-elections).  If this community decision isn’t seen as a backdoor for introducing potentially insecure primitives, a cryptocurrency’s threat model is likely not strong enough to warrant any mining rewards or economic security at all.

This is even before you consider the fact that many of the most vocal members of this “community” decision may in fact be miners, who are financially incentivized to lobby for hash functions they can either obtain edges on or have some expertise with.  In the long run, my prediction is that this will likely serve as a centralization vector _supporting_ economies of scale, prone to the wonderful world of [regulatory capture](https://en.wikipedia.org/wiki/Regulatory_capture), leading to many hours wasted on debates, distracting from real issues, and exhausting a coin’s capacity to make important and actually relevant governance choices.

## Emotional Censorship for the Short Term Only

So, I think I’ve made some case that in the long term, hard forking around specific actors to prevent specialization and economies of scale is a losing game.

What about in the short term?  In the short term, such hard forks can actually be remarkably effective in discriminating against certain actors.  It is undeniable that such hard forks can be a short-term annoyance for companies like Bitmain, that may actually be making capital investments into building ASICs.  And, perusing reddit and Twitter, this is what 99% of the debate comes down to.  Don’t believe me?

[![Reddit comment by "inb4_banned": yes, fuck bitmain.  even if there is only a 1% chances they have them we should still change the PoW](https://pdaian.com/blog/wp-content/uploads/2018/04/Screenshot-from-2018-04-01-17-23-16.png)](https://www.reddit.com/r/ethereum/comments/88jduy/are_you_in_favour_of_hard_fork_to_resist_asic/dwl0xz4/)

_**Pitchforks at the ready, my “community”**_

Unfortunately, I don’t think this is productive in the long run.  Bitmain is an operation that makes hundreds of millions of dollars mining cryptocurrencies.  [They will achieve economies of scale in any mining model, including general purpose hardware / GPU-dominated mining.](https://www.bitsonline.com/antminer-g2-gpu-mining-rig-review/)  Fighting this with a hard fork is at most a temporary band-aid to a very fundamental economic problem: the actors that end up dominating any competitive marketplace may not be your chosen favorite actors in the space.

## … and my attack cost!

My final argument against such forks is very simple: anti-ASIC forks reduce the cost to attack a cryptocurrency substantially.  There are a number of reasons for this.  Consider the ultimate cryptocurrency threat model: malicious nation states or highly resourced actors like malicious banks.  It was in this setting that cryptocurrency was first conceptualized as an anti-fragile system.

There are two problems with anti-ASIC hard forks.  The first is with ASIC resistance itself; in recent work actually analyzing the dollar cost to attack both Ethereum and Bitcoin, it’s quite apparent that ASICs are actually a good thing for the security of a cryptocurrency system ([extended abstract](http://materials.dagstuhl.de/files/17/17132/17132.JosephBonneau.ExtendedAbstract.pdf) / [full paper](https://fc18.ifca.ai/bitcoin/papers/bitcoin18-final17.pdf)).  General purpose computer-secured systems are quite cheaper to attack than specialized systems, as general purpose hardware can be rented and resold after an attack to substantially subsidize attack cost.  Worse still, an attacker can use general purpose hardware advantages to overwhelm any cryptocurrency seeking to avoid specialization by _**carrying attacks across PoW changing / defensive forks**_, a feature which is not possible with ASICs, where the community has a “nuclear option” available to disable miners at any time if objectively malicious behavior is detected.  These two factors combined severely cripple the economic security of networks secured by general purpose hardware.

The second problem with anti-ASIC hard forks is in the culture established itself; discouraging capital investment in the short term, and a governance process of repeated change create potent vectors for disruption, attack, and takeover, without accomplishing the stated goals of preserving decentralization in the long term.  In my opinion this will eventually lead to more centralization, as only the most sophisticated and dedicated companies are able to ride the waves of change and extract profits at scale.

## Block Rewards and PoW Instead

One alternative I’ve proposed for cryptocurrencies that really hate their miners is to [lower block rewards](https://github.com/ethereum/EIPs/issues/960) instead of trying to hard fork around wealthy, resourced actors; make it unprofitable to leverage economies of scale by making it mostly unprofitable to mine, and you will end up with a coin operated by hobbyists that are probably more decentralized than with ASIC-based PoW miners.  Of course your attack cost will go down, especially to nation states and resourced actors, but the trade-off may be worth it for decentralization; after all, how much dollar-valued security does a blockchain truly require to remain immutable?  This is a totally unanswered question in the research space, but my guess is you could get away much cheaper than, say, 12.5 BTC/10 mins.

## … or, PoS?

Of course, this is all a moot point for cryptocurrencies which [plan to move to full proof of stake](https://github.com/ethereum/wiki/wiki/Proof-of-Stake-FAQ), the ultimate anti-ASIC protection.

That is not to say that proof of stake will be any less centralized than proof of work, but to further argue against wasting time on pointless update games in the face of major protocol challenges ahead.

## So,

Keep the ASICs.  Learn to coexist.  Avoid the governance headaches and wasted developer hours, ending in the same economies of scale with likely more centralization.  Raise your coin’s security through capital investments and a diversity of stakeholders.

If you DO want to hard fork away from ASICs, be aware that it is a solution that can work in only the short term; the idea that you can somehow optimize for or enforce decentralization via regular, likely centralized-governance hard forks, is quite absurd.

**In the short term, such forks absolutely work to target particular actors**.  But if this is what is desired, it’s important to be politically honest about intentions.