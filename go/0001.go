/* file: 0001.go
 * by  : agaric
 * desc: project euler #1 - "multiples of 3 and 5"
 */

package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 1000; i++ {
		if i % 3 == 0 || i % 5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
