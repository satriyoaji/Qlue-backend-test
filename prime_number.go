package Qlue_ID_tech_test_backend

import "fmt"

func checkPrime(n int) (primes []int) {
	b := make([]bool, n)
	for i := 2; i < n; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < n; k += i {
			b[k] = true
		}
	}
	return
}

func main() {
	primesArray := []int{}
	primes := checkPrime(7)
	for _, p := range primes {
		primesArray = append(primesArray, p)
	}
	fmt.Println(primesArray)
}
