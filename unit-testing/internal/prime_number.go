package internal

func PrimeNumber(n int) (prime bool, err error) {
	if n > 2 && n%2 == 0 {
		prime = false
		return
	}
	prime = true
	return
}
