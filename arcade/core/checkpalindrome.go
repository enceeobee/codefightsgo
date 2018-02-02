package codefightsgo

// CheckPalindrome checks if inputString is a palindrome.
func CheckPalindrome(inputString string) bool {
	// We don't need to use Floor, because integer division does that automatically (I guess)
	// terminus := int(math.Floor(float64(len(inputString)) / float64(2)))

	for i := 0; i < len(inputString)/2; i++ {
		if string(inputString[i]) != string(inputString[len(inputString)-1-i]) {
			return false
		}
	}

	return true
}
