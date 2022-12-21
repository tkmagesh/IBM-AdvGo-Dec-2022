package utils

func IsPrime(no int) bool {
	for i := 2; i < (no - 1); i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
