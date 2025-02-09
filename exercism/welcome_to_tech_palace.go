package techpalace

import "strings"

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	nameUpper := strings.ToUpper(customer)
	return "Welcome to the Tech Palace, " + nameUpper
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	borderStr := strings.Repeat("*", numStarsPerLine)
	return borderStr + "\n" + welcomeMsg + "\n" + borderStr
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	s := strings.Trim(oldMsg, "* ")
	s = strings.Replace(s, "*", "", -1)
	return strings.TrimSpace(s)
}
