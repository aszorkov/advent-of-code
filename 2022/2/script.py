# Constants
win = 6
draw = 3
loss = 0
rock = 1
paper = 2
scissors = 3

# Expects inputs as numbers representing the value of each move:
# 1 for rock
# 2 for paper
# 3 for scissors
# Returns an array in the form [player_score, opponent_score]
def score_round(opponent, player):
    if (player == "X" and opponent == "C") or (player == "Y" and opponent == "A") or (player == "Z" and opponent == "B"):
        player_outcome = 6
        opponent_outcome = 0
    elif (player == "X" and opponent == "B") or (player == "Y" and opponent == "C") or (player == "Z" and opponent == "A"):
        player_outcome = 0
        opponent_outcome = 6
    else:
        player_outcome = opponent_outcome = 3
    return [convert_input(opponent) + opponent_outcome, convert_input(player) + player_outcome]

# Helper function to convert character inputs to their score values
def convert_input(input):
    if input in ['A', 'X']:
        return rock
    elif input in ['B', 'Y']:
        return paper
    elif input in ['C', 'Z']:
        return scissors

with open('input.txt') as f:
    opponent_score = player_score = 0
    for i, line in enumerate(f):
        if(line.strip()):
            inputs = line.split()
            round_scores = score_round(inputs[0], inputs[1])
            opponent_score += round_scores[0]
            player_score += round_scores[1]
            print(inputs, "player:", round_scores[1], ", opponent:", round_scores[0], "player_total:", player_score)
    if player_score < opponent_score:
        result = "Loss"
    elif player_score == opponent_score:
        result = "Draw"
    elif player_score > opponent_score:
        result = "Win"
    print('Player score: ', player_score)
    print('Opponent score: ', opponent_score)
    print('Result: ', result)