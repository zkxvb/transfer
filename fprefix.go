package main

import "fmt"

// prefixLen returns the length of the common prefix of a and b.
func prefixLen(a, b []byte) int {
	var i, length = 0, len(a)
	if len(b) < length {
		length = len(b)
	}
	for ; i < length; i++ {
		if a[i] != b[i] {
			break
		}
	}
	return i
}

func main() {
	i, n := 0, 0
	fmt.Scanf("%d\n", &n)
	if n == 1 {
		var s string
		fmt.Scanf("%s\n", &s)
		fmt.Printf("%s", s)
		return
	}
	var s1, s2 []byte
	if n&1 == 1 {
		fmt.Scanf("%s\n", &s2)
		i++
		fmt.Scanf("%s\n", &s1)
		i++
		l := prefixLen(s2, s1)
		fmt.Printf("%s", s2[:l+1])
		fmt.Scanf("%s\n", &s2)
		i++
	} else {
		fmt.Scanf("%s\n", &s1)
		i++
		fmt.Scanf("%s\n", &s2)
		i++
	}
	for {
		l := prefixLen(s1, s2)
		fmt.Printf("%s\n%s", s1[:l+1], s2[:l+1])
		if i <= n-2 {
			fmt.Scanf("%s\n", &s1)
			i++
			fmt.Scanf("%s\n", &s2)
			i++
		} else {
			break
		}
	}
}
