/* file: 0002.go
 * by  : agaric
 * desc: project euler #2 - "even Fibonacci numbers"
 */

package main

import "fmt"

func fib(n int) []int {
	a := make([]int, n)
	a[0], a[1] = 1, 2
	for i := 2; i < n; i++ {
		a[i] = a[i - 1] + a[i - 2]
	}
	return a
}

func main() {
	sum := 0
	for _, x := range fib(32) { // fib(33) > 4 mil
		if x % 2 == 0 {
			sum += x
		}
	}
	fmt.Println(sum)
}
