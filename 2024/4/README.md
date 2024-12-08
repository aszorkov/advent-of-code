# Day 4

Ooo a word search.

## Part One

Given an input in the form of a grid, find all instances of the word `XMAS`. Valid matches are horizontal, vertical and diagonal within the grid, and both forward and backwards.

### Approach

So we're going to need to read in the file into a two dimensional array and iterate over it to find matches for each space in the grid.

If we check in all directions from the first character in the string we want to match, "X" in this case, we should be able to look up, down, left, right, and diagonally in all directions to try and build the whole word square by square.

The first control structure would be two `for` loops for the x/y directions of the grid, and then inside of this we can call to a helper function which will search in all directions if the character at a certain co-ordinate is the start of our search term. 

## Part Two

Now we're told that this wasn't actually an "XMAS" word search, but instead was meant to be an "X-MAS" word search, meaning we're meant to find all instances of the word "MAS" which cross over each other... in an X. For example:
```
.M.M.
..A..
.S.S.
```

So we need to refine our search to find these. The "MAS" can appear either forward or backwards. What the hell.

### Approach

It looks like we can use a similar approach to before, but instead of only kicking off our search from the start of the word XMAS, we can look for all occurrences of the middle of the word, "A", and then search diagonally to make sure that we find either "MAS" or "SAM" in both diagonal directions.

## Learnings

- When assigning the character at an index in a string to a var, it will be assigned the unicode value of the character by default, meaning it will actually be an int. In this case it's fine because we just need to match values, but this was confusing when I was printing out values and wondering why I was getting numbers!
- Thinking in two dimensional arrays is always annoying! I was scratching my head for a while wondering why my output wouldn't match the sample input, then realised I had accidentally made my "down" directions all subtract instead of add to the y index
- Based on a very short time Googling, it doesn't look like there is a helper string for reversing a given string in Go, so you need to implement your own function for this