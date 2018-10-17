package main

import "fmt"

func isLuckyNumber(i int) bool {
	strtI := fmt.Sprintf("%d", i)
	bsI := []byte(strtI)
	for i, j := 0, len(bsI)-1; i < j; i, j = i+1, j-1 {
		if bsI[i] == bsI[j] {
			return false
		}
	}
	return true
}

func main() {
	var a, b, cnt int
	fmt.Scanf("%d %d", &a, &b)
	cnt = 0
	for i := a; i < b+1; i++ {
		if isLuckyNumber(i) {
			cnt++
		}
	}
	fmt.Printf("%d", cnt)

}
