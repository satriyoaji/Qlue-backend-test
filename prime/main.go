package prime

import "fmt"

func checkPrime(N int) (primes []int) {
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return
}

func main() {
	var n int
	fmt.Print("input n: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err.Error())
	}

	primesArray := []int{}
	primes := checkPrime(n)
	for _, p := range primes {
		primesArray = append(primesArray, p)
	}
	fmt.Println(primesArray)
}
