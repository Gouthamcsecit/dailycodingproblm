```
Problem

Given an array of integers, return a new array such that each element at index i of the new array is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. If our input was [3, 2, 1], the expected output would be [2, 3, 6].
```

package main

import "fmt"

func ProductSum(a []int) []int {
	arrNew := []int{}
	for i := 0; i < len(a); i = i + 1 {
		temp := 1
		for j := 0; j < len(a); j = j + 1 {
			if a[i] != a[j] {
				temp = temp * a[j]
			}
		}
		arrNew = append(arrNew, temp)
	}
	return arrNew
}

func main() {
	var a []int
	a = []int{1, 2, 3, 4, 5}
	arrNew := ProductSum(a)
	fmt.Println(arrNew)
}
