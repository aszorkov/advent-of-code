# Day 3

For day 3, we are given an input which is meant to simulate instructions for a program, where the only instruction we need to support is a multiplication function, invoked by **exactly** `mul(x,y)`.

## Part One

The input is "corrupted" so there are lots of instances of incorrect syntax, e.g. unclosed parenthesis, spaces in odd places, and so on. We need to parse the input, recognise only correct instances of the `mul` function, and sum up the result of all of these.

### Approach

Seems like this is a clear case of using regex to find all matches in the entire input for a single capture group. Let's hope Go has good native support for regex.

...

Good news, it [does](https://gobyexample.com/regular-expressions)!

So we should be able to just load the entire text file into a big string then create an array of all instances of a capture group which matches on `(mul\(\d,\d\))`. Once we have a slice of all valid instances, iterate over it and calculate a sum of all of the multiplication results.

## Part Two

In part two we're informed that there are actually two more functions we can respect from the input, `do()` and `don't()`. We're told that each time we encounter a `do()` we should respect all following `mul()` calls, until we encounter a `don't()`. All `mul()` should then be ignored until we encounter another `do()`. The input starts as if we have already encountered a `do()`.

### Approach

We should be able to solve this by using another regex to find all blocks in the input which start with a `do()` and end with a `don't()`, then repeat the solution from part one on these "valid" blocks.

`(?sm)(?:\A|do\(\)).*?(?:don't\(\)|\z)`

## Learnings

- Golang's regex library doesn't support lookahead/lookbehind
- Line terminators are evil. I had to refine my expression a few times as there was a newline in the middle of a block in my input, causing only half of the block to actually be matched
