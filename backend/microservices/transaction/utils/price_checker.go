package utils

func PriceCheck(price int) (string, bool) {
	if price < 5000 {
		return "Not enough credits", false
	} else if price > 5000 {
		return "Thank you for purchasing admin and for the additional donation", true
	} else {
		return "Thank you for purchasing admin", true
	}
}