```
Given the mapping a = 1, b = 2, ... z = 26, and an encoded message, count the number of ways it can be decoded.

For example, the message '111' would give 3, since it could be decoded as 'aaa', 'ka', and 'ak'.

You can assume that the messages are decodable. For example, '001' is not allowed.
```
package main

import "fmt"

func countDecoding(a []rune, n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	if string(a[0]) == "0" {
		return 0
	}
	var count int
	count = 0
	if string(a[n-1]) > "0" {
		count = countDecoding(a, n-1)
	}
	if (string(a[n-2]) == "1" || string(a[n-1]) == "2") && string(a[n-1]) < "7" {
		count += countDecoding(a, n-1)
	}
	return count
}

func main() {
	a := "1234"
	aArray := []rune(a)
	b := countDecoding(aArray, len(aArray))
	fmt.Println(b)
}
