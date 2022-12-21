package utils

func isPrime_1(no int) bool {
	for i := 2; i < no; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}

func isPrime_2(no int) bool {
	max := (no / 2)
	for i := 2; i <= max; i++ {
		if no%i == 0 {
			return false
		}
	}
	return true
}
