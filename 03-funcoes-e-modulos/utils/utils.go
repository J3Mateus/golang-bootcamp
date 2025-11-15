package utils

// Function that receives a number and returns 1 if it's even and 0 if it's odd
func Even(value int) int {
	if value%2 == 0 {
		return 1
	}
	return 0
}

// Function that receives a number and returns true if it's prime, false otherwise
func IsPrime(value int) bool {
	if value <= 1 {
		return false
	}
	for i := 2; i*i <= value; i++ {
		if value%i == 0 {
			return false
		}
	}
	return true
}
