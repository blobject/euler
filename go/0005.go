/* file: 0005.go
 * by  : agaric
 * copy: public domain
 * desc: project euler #5 - "smallest multiple"
 * lang: go
 */

package main

import "fmt"

// find all unique prime factors in n
func factors(res []int, n int) []int {
	i := 2
	for {
		if i > n { return res }
		if n % i == 0 {
			n /= i
			if len(res) == 0 || res[len(res) - 1] != i {
				res = append(res, i)
			}
		} else {
			i++
		}
	}
}

func main() {
	n := 20               // problem number
	fs := make([]bool, n) // prime factor occurrence
	cs := []int{}         // contributing multiples

	// 1. go through each number n, n-1, n-2, ..., 2
	// 2. find their unique prime factors
	// 3. if at least one of the prime factors have not yet occurred,
	//    this number is a contributing multiple
	for i := n; i > 1; i-- {
		f := factors([]int{}, i)
		contrib := false
		for _, j := range f {
			if !fs[j - 1] {
				fs[j - 1] = true
				contrib = true
			}
		}
		if contrib {
			cs = append(cs, i)
		}
	}

	mult := 1
	for _, m := range cs {
		mult *= m
	}

	// sweep through each number 2, 3, ..., n once more and make sure
	// each number divides the multiple
	for i := 2; i <= n; i++ {
		if mult % i != 0 {
			mult *= i
		}
	}

	fmt.Println(mult)
}
