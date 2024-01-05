```
Given a sorted array, find the smallest positive integer that is not the sum of a subset of the array.

For example, for the input [1, 2, 3, 10], you should return 7.

```
package main

import "fmt"

func SmallestInt(arr []int, n int) int {
	var res int
	res = 1
	for i := 0; i < n && arr[i] <= res; i++ {
		res = res + arr[i]
	}
	return res
}

func main() {
	arrNew := []int{1, 2, 3, 10}
	answer := SmallestInt(arrNew, len(arrNew))
	fmt.Println(answer)
}
