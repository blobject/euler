/* file: 0003.go
 * by  : agaric
 * copy: public domain
 * desc: project euler #3 - "largest prime factor"
 * lang: go
 */

package main

import ("fmt"; "math")

func sieve(n int) []int {
	a := []int{}
	b := make([]bool, n)
	for i, _ := range b {
		b[i] = true
	}
	b[0] = false
	if n >= 1 { b[1] = false }
	for i, x := range b {
		if x {
			m := i * i;
			for m < n {
				b[m] = false
				m += i;
			}
		}
	}
	for i, x := range b {
		if x { a = append(a, i) }
	}
	return a
}

func main() {
	n := 600851475143
	p := sieve(int(math.Sqrt(float64(n))))
	max := 1
	for _, x := range p {
		if n % x == 0 && x > max {
			max = x
		}
	}
	fmt.Println(max)
}
