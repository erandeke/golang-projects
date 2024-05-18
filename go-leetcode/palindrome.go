package main

func IsNumberPalindrome(a int) bool {

	if a < 0 || (a%10 == 0 && a != 0) {
		return false
	}

	reversed := 0

	for a > reversed {
		digit := a % 10
		reversed = reversed*10 + digit
		a = a / 10
	}

	if a == reversed || a == reversed/10 {
		return true
	}

	return false
}
