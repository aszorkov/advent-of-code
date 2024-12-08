## The problem

Given a series of inputs representing yours and your opponents moves in a game of rock paper scissors, determine your final score.

## Inputs:
Each line in the file is a round. Each line contains two characters.

The first letter is your opponents move:
- A for rock
- B for paper
- C for scissors

The second letter is your own move:
- X for rock
- Y for paper
- Z for scissors

## Scoring:
You are given a score  based on your selection:
- 1 for rock
- 2 for paper
- 3 for scissors

You are then given a score for the outcome of the round:
- 0 for a loss
- 3 for a draw
- 6 for a win

## Part 1
For the set of inputs provided, calculate the score if everything goes according to plan.

### Approach:

First we should consider what we need to calculate. While part 1 is only asking for our own score, we should also keep track of our opponent's score in the event that we may need to use it later.