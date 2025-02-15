# [Advent of Code](/)

- [[About]](/2021/about)
- [[Events]](/2021/events)
- [[Shop]](https://teespring.com/stores/advent-of-code)
- [[Settings]](/2021/settings)
- [[Log Out]](/2021/auth/logout)

TecHunter 13\*

# 0.0.0.0: [2021](/2021)

- [[Calendar]](/2021)
- [[AoC++]](/2021/support)
- [[Sponsors]](/2021/sponsors)
- [[Leaderboard]](/2021/leaderboard)
- [[Stats]](/2021/stats)

Our [sponsors](/2021/sponsors) help make Advent of Code possible:

[corpuls](https://corpuls.world/karriere/stellenangebote/code-to-save-lives.php) \- <3<3<3 where your code might actually save lives .....\_\_...\_\_..... ..../...v...\\.... ....\|.......\|.... .....\\...../..... .......\ /.......

## \-\-\- Day 7: The Treachery of Whales ---

A giant [whale](https://en.wikipedia.org/wiki/Sperm_whale) has decided your submarine is its next meal, and it's much faster than you are. There's nowhere to run!

Suddenly, a swarm of crabs (each in its own tiny submarine - it's too deep for them otherwise) zooms in to rescue you! They seem to be preparing to blast a hole in the ocean floor; sensors indicate a _massive underground cave system_ just beyond where they're aiming!

The crab submarines all need to be aligned before they'll have enough power to blast a large enough hole for your submarine to get through. However, it doesn't look like they'll be aligned before the whale catches you! Maybe you can help?

There's one major catch - crab submarines can only move horizontally.

You quickly make a list of _the horizontal position of each crab_ (your puzzle input). Crab submarines have limited fuel, so you need to find a way to make all of their horizontal positions match while requiring them to spend as little fuel as possible.

For example, consider the following horizontal positions:

```
16,1,2,0,4,2,7,1,2,14
```

This means there's a crab with horizontal position `16`, a crab with horizontal position `1`, and so on.

Each change of 1 step in horizontal position of a single crab costs 1 fuel. You could choose any horizontal position to align them all on, but the one that costs the least fuel is horizontal position `2`:

- Move from`16` to `2`: `14` fuel
- Move from`1` to `2`: `1` fuel
- Move from`2` to `2`: `0` fuel
- Move from`0` to `2`: `2` fuel
- Move from`4` to `2`: `2` fuel
- Move from`2` to `2`: `0` fuel
- Move from`7` to `2`: `5` fuel
- Move from`1` to `2`: `1` fuel
- Move from`2` to `2`: `0` fuel
- Move from`14` to `2`: `12` fuel

This costs a total of `<em>37</em>` fuel. This is the cheapest possible outcome; more expensive outcomes include aligning at position `1` ( `41` fuel), position `3` ( `39` fuel), or position `10` ( `71` fuel).

Determine the horizontal position that the crabs can align to using the least fuel possible. _How much fuel must they spend to align to that position?_

Your puzzle answer was `345197`.

The first half of this puzzle is complete! It provides one gold star: \*

## \-\-\- Part Two ---

The crabs don't seem interested in your proposed solution. Perhaps you misunderstand crab engineering?

As it turns out, crab submarine engines don't burn fuel at a constant rate. Instead, each change of 1 step in horizontal position costs 1 more unit of fuel than the last: the first step costs `1`, the second step costs `2`, the third step costs `3`, and so on.

As each crab moves, moving further becomes more expensive. This changes the best horizontal position to align them all on; in the example above, this becomes `5`:

- Move from`16` to `5`: `66` fuel
- Move from`1` to `5`: `10` fuel
- Move from`2` to `5`: `6` fuel
- Move from`0` to `5`: `15` fuel
- Move from`4` to `5`: `1` fuel
- Move from`2` to `5`: `6` fuel
- Move from`7` to `5`: `3` fuel
- Move from`1` to `5`: `10` fuel
- Move from`2` to `5`: `6` fuel
- Move from`14` to `5`: `45` fuel

This costs a total of `<em>168</em>` fuel. This is the new cheapest possible outcome; the old alignment position ( `2`) now costs `206` fuel instead.

Determine the horizontal position that the crabs can align to using the least fuel possible so they can make you an escape route! _How much fuel must they spend to align to that position?_

Answer:

Although it hasn't changed, you can still [get your puzzle input](7/input).

You can also [Shareon
[Twitter](https://twitter.com/intent/tweet?text=I%27ve+completed+Part+One+of+%22The+Treachery+of+Whales%22+%2D+Day+7+%2D+Advent+of+Code+2021&url=https%3A%2F%2Fadventofcode%2Ecom%2F2021%2Fday%2F7&related=ericwastl&hashtags=AdventOfCode) [Mastodon](javascript:void(0);)] this puzzle.

