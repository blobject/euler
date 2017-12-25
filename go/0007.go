/* file: 0007.go
 * by  : agaric
 * copy: public domain
 * desc: project euler #7 - "10001st prime"
 * lang: go
 */

package main

import "fmt"

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
	fmt.Println(sieve(110000)[10000])
}
