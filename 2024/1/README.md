# Day 1

Okay cool, day 1. So we are given an input of two equal length "lists" of unique ids in the form of a text file where each line is an entry and the entries for each list are separated by a space. The lists are meant to match but they don't, oh no!

## Part One

Given these two lists, the first task is to see if the lists only differ by a small error by first sorting each lists and then sequentially comparing the entries of each list to determine the "distance" between each pair and returning the total of all of these distances.

### Approach

First I'll need to learn how to read from the filesystem in Go. Then I'll need to work out how to iterate over each line in a given file and take in each value.

Then I'll generate two arrays from this input, sort the arrays and iterate over the lists, calculating the distance between each pair and adding it to a running total.

## Part Two

So after solving part one, we realise that the lists are very different so we are given a different approach to take. We should look at each list and work out how often a value in the first list occurs in the second list. For each row, we calculate a "similarity score" which is the left value multiplied by how many instances of it are in the right list. From the sample input, it looks like values in each list are not necessarily unique, so the same number could occur multiple times on each side.

My first thought would be to generate a map, where each unique value acts as the key. For each entry in the map, we'll store a tuple with two values, these being the number of times the key appears in each list.

After building the map, we can iterate over the keys in the map and generate the similarity score by multiplying the key by the number of occurrences in each list, e.g.

```
--- The list ---
...
1234: {2, 3} // Similarity score = 1234 * 2 * 3 = 7404
5678: {1, 4} // Similarity score = 5678 * 1 * 4 = 22712
...
```

## Learnings

- `os` package is used for calling operating system functions, e.g. opening a file
- We also use `os.args` to grab command line arguments (in this case so I can provide an input file path from the cmd line)
- `bufio` provides buffered I/O, which provides the helpers we need to iterate over files
- `strings` provides methods for... strings... Such as splitting a line into distinct values delimited by spaced
- `slices` has helpers for sorting
- `for` is the only type of loop in Go. We can iterate over a range using `for index, val := range range_name`
- If we have to specify a variable but don't intend to use it, we can provide an underscore, e.g. `for index, _ := range range_name`
- When reading in a file, we'll get strings so we need to convert these to our needed data type, in this case `strconv.Atoi`
