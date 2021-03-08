/* file: 0006.go
 * by  : agaric
 * desc: project euler #6 - "sum square difference"
 */

package main

import "fmt"

func sumsq(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func sqsum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum
}

func main() {
	n := 100
	fmt.Println(sqsum(n) - sumsq(n))
}
