import getopt, sys
max = 0
marker = -1
elves = []
topN = 1

argumentList = sys.argv[1:]
options = "n:"
long_options = ["number="]
arguments, values = getopt.getopt(argumentList, options)

for currentArg, currentValue in arguments:
    if currentArg in ("-n", "--number"):
        topN = int(currentValue)

with open('input.txt') as f:
    calories = 0
    for i, line in enumerate(f):
        if line.strip():
            calories += int(line)
        else:
            elves.append(calories)
            if calories > max:
                max = calories
                marker = i+1
            calories = 0
    sortedElves = sorted(elves, reverse=True)
    print('Top', topN, 'elves are carrying:', sum(sortedElves[0:topN]), 'calories.')