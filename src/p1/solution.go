```
Problem

Given a list of numbers and a number k, return whether any two numbers from the list add up to k.

For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
```
package main

import "fmt"

func Sumup(arrayNew []int, k int) {
	for i := 0; i < len(arrayNew); i = i + 1 {
		for j := 0; j < len(arrayNew); j = j + 1 {
			if arrayNew[i]+arrayNew[j] == k {
				fmt.Println(arrayNew[i], arrayNew[j])
			}
		}
	}
}

func main() {
	arrayNew := []int{10, 15, 3, 7}
	var k int
	k = 17
	Sumup(arrayNew, k)
}
