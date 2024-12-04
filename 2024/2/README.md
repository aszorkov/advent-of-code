# Day 2

This time, we're given a list of "reports" as input, where each line in the input is a single report and the values separated by spaces are the sequential "levels" in a report.

## Part One

The task is to review each report in the input and determine if each report is "safe", then determine the total amount of "safe" reports from the given input.

The rules to determine if a report is "safe" are:

1. All levels in the input must follow the same direction (e.g. all increasing or all decreasing)
2. The distance between each adjacent level in a report can be between one and three (inclusive)

### Approach

This seems simple enough, we'll load all of the rows from the input file same as last time, then some validation logic which iterates over the levels in a report and bails out early when we find a failure case, or marks as safe if we make it to the end with no issues.

We can first work out the difference and the direction for each pair of adjacent levels by getting the difference of each value to the value preceding it, and the sign of the result tells us the direction (positive for increasing, negative for decreasing, who woulda thunk it)

Then we mark the report as unsafe if the difference is 0, less than -3 or greater than 3. If we make it to the end, the report is safe.

## Part Two

Now we're given an extra piece of information that we can actually tolerate a single bad level in a given report. Essentially, if removing one of the items in the report would create a safe report then we'll accept it. If we can't create a safe report by removing any one item, then it is considered unsafe.

While in part one we can build this out by simply iterating over each report one time, the first solution that jumps out to solve part two is recursive in nature.

We'll create a loop over the levels in the report, starting at index 0. If the report is safe, we can break out of the loop and move to the next report.

If the report was unsafe, we'll create a new report of all levels minus the current index and check this report. If the new report would have been safe, we can mark the original report as safe and specify which index we removed to make it safe.

If removing the current index doesn't make the report safe, try again after removing the next index, and the next, until we've tried removing all of the indices.

We could even extend this to take a value as a "tolerance" argument to the function to determine how many entries we might allow to be deleted; i.e. how many levels of recursion to allow. The answer to part two would take a "tolerance" of 1.

## Learnings

- We can define `struct` types to create objects with many different typed values
- Since the only type of loop in Go is a `for`, we have to emulate a `do while` by using a flag in the condition, e.g. `for i, break := 0, false; i < length && !break; i++`
- If we want to allow for optional arguments to a function, we can create a [variadic function](https://go.dev/play/p/T6-_T7KOIg), where we specify that we will accept zero, one or many arguments of the same type, e.g. `func buildReport(data []int, args ...int)`
