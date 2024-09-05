with open("input.txt") as file:
    lines = file.read().split('\n')

cards = {
    'A': 14, 
    'K': 13, 
    'Q': 12, 
    'J': 11, 
    'T': 10, 
    '9': 9, 
    '8': 8, 
    '7': 7, 
    '6': 6, 
    '5': 5, 
    '4': 4, 
    '3': 3, 
    '2': 2
    }

hand_strength = {
    7: '5 of a kind',
    6: '4 of a kind',
    5: 'full house',
    4: '3 of a kind',
    3: 'two pair',
    2: 'one pair',
    1: 'high card'
}

hand_types = []

for line in lines:
    hand, value = line.split(' ')
    counts = set()
    for card in hand:
        count = hand.count(card)
        if count > 1:
            counts.add((card, count))
    
    counts = list(counts)

    if len(counts) == 1:
        if counts[0][1] == 5:
            strength = 7
        elif counts[0][1] == 4:
            strength = 6
        elif counts[0][1] == 3:
            strength = 4
        elif counts[0][1] == 2:
            strength = 2
    elif len(counts) == 2:
        if counts[0][1] == 2 and counts[1][1] == 2:
            strength = 3
        elif (counts[0][1] == 2 and counts[1][1] == 3) or (counts[0][1] == 3 and counts[1][1] == 2):
            strength = 5
    else:
        strength = 1

    
    hand_types.append((strength, hand_strength[strength], hand, value))

def get_card_value(card):
    return cards[card[0]]

def sorting_key(hand):
    strength, _, cards_str, _ = hand
    card_values = [get_card_value(card) for card in cards_str]
    return (strength, card_values)

sorted_hand_types = sorted(hand_types, key=sorting_key)

total = 0

for i, hand in enumerate(sorted_hand_types):
    total += ((i + 1) * int(hand[3]))

print('part 1 total:', total)