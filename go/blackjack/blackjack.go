package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	total := ParseCard(card1) + ParseCard(card2)
	// Check for blackjack (Ace + Ten/Face Card)
	isBlackjack := (card1 == "ace" && (card2 == "ten" || card2 == "king" || card2 == "queen" || card2 == "jack")) ||
		(card2 == "ace" && (card1 == "ten" || card1 == "king" || card1 == "queen" || card1 == "jack"))

	if isBlackjack {
		// Stand if dealer has a chance of blackjack (Ace or Ten/Face Card)
		if dealerCard == "ace" || dealerCard == "ten" || dealerCard == "king" || dealerCard == "queen" || dealerCard == "jack" {
			return "S"
		}
		// Otherwise, Win
		return "W"
	}

	switch {
	case total == 22: // Two aces
		return "P"
	case total == 21 || total >= 17 && total <= 20: // Stand on 21 or 17–20
		return "S"
	case total >= 12 && total <= 16:
		if ParseCard(dealerCard) >= 7 {
			return "H" // Hit if dealer shows 7 or higher
		}
		return "S"
	default:
		return "H" // Hit in all other cases
	}
}
