package utils

func MaskCard(cardNumber string) string {
	if len(cardNumber) < 4 {
		return "**** **** **** ****"
	}
	return "**** **** **** " + cardNumber[len(cardNumber)-4:]
}