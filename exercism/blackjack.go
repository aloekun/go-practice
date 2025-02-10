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
		fallthrough
	case "jack":
		fallthrough
	case "queen":
		fallthrough
	case "king":
		return 10
	case "other":
		return 0
	default:
		// joker, not expected string
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	value1 := ParseCard(card1)
	value2 := ParseCard(card2)
	valueDealer := ParseCard(dealerCard)
	sumValue := value1 + value2

	switch {
	case card1 == "ace" && card2 == "ace":
		// split
		return "P"
	case sumValue == 21:
		// switch play by dealer hand
		if valueDealer != 10 && valueDealer != 11 {
			// small number less 10
			return "W"
		}
		return "S"
	case 17 <= sumValue && sumValue <= 20:
		return "S"
	case 12 <= sumValue && sumValue <= 16:
		// switch play by dealer hand
		if valueDealer < 7 {
			return "S"
		}
		return "H"
	case sumValue <= 11:
		return "H"
	default:
		return "S"
	}
}
