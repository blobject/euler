/* file: 0004.go
 * by  : agaric
 * desc: project euler #4 - "largest palindrome product"
 */

package main

import "fmt"

func reverse(a []int) []int {
	for i := len(a) / 2 - 1; i >= 0; i-- {
		j := len(a) - 1 - i
		a[i], a[j] = a[j], a[i]
	}
	return a
}

func isPalindrome(n int) bool {
	a := []int{}
	x := n
	for x >= 1 {
		a = append(a, x % 10)
		x /= 10
	}
	l, r := []int{}, []int{}
	for i, x := range a {
		if i < len(a) / 2 {
			l = append(l, x)
		} else {
			r = append(r, x)
		}
	}
	if len(l) < len(r) {
		r = append(r[:0], r[1:]...)
	}
	r = reverse(r)
	for i, x := range l {
		if x != r[i] {
			return false
		}
	}
	return true
}

func main() {
	a, b, c := []int{}, []int{}, []int{}
  // presuming the answer lies somewhere 800..999
	for i := 999; i > 800; i-- {
		for j := 999; j > 800; j-- {
			if isPalindrome(i * j) {
				a = append(a, i * j)
				b = append(b, i)
				c = append(c, j)
			}
		}
	}
	max := 0
	for i, x := range a {
		if x > a[max] {
			max = i
		}
	}
	fmt.Println(a[max], b[max], c[max])
}
