# Day 5

Wheeeeew the explanation for this one is a doozy.

Tl;dr, the input we are given is in two sections relating to instructions for a fictional printing process. The first section is a sequence of rules regarding page numbers, e.g.
```
47|53
97|13
97|61
97|47
```
Each line is a rule, stating that the page on the left hand side mnust be printed before the right hand side.

After a black line, the second part of the instructions begins, in the format:
```
75,47,61,53,29
97,61,53,29,13
75,29,13
```
Each line is an "update", in the form of a comma separated list of page numbers.

## Part One

For each update, we need to refer back to the rules from the first section and determine whether the update is valid (i.e. all pages are positioned correctly according to their specific ordering rules). If a page number was not represented in the rules section we can ignore it and move on.

Then, because reasons, we need to find the middle page in each update, and get the sum of all of the middle page numbers for all of the valid updates.

### Approach

Again this sounds like the perfect time for a map!

After taking in the first section of the input, we can iterate over the rules and create a map where the key is the page number and the value has a slice of page numbers that it must be before. For each rule, check if we've already added the page number as a key and if so we'll append the right hand side to the page number list. If we haven't, create a new map entry.

Then we can process each update in the second section of the input. For each of then, create a slice from the comma separated list of values and iterate over each of them. For each value, we'll need to iterate over the string from the beginning and ensure that for each of the values in it's rule list that none of them occur before it.

If an update passes these checks, we can increment the running total by it's middle value.

## Part Two

### Approach

## Learnings

-
