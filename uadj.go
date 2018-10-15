package main

import (
	"fmt"
	"time"
)

// 递归思想：
// 一个一个字母安排，安排完则计数 ++，安排不了则返回

// 安排的了？
func canJoin(m, n, k int, c byte) bool {
	switch c {
	case 'r':
		if n+k == 0 {
			return false
		}
	case 'b':
		if m+k == 0 {
			return false
		}
	case 'g':
		if m+n == 0 {
			return false
		}
	}
	return true
}

// 递归安排
func join(m, n, k int, c byte, cnt *int) {
	// 递归出口
	if m+n+k == 0 {
		//l.Lock()
		*cnt++
		//l.Unlock()
		return
	}
	if !canJoin(m, n, k, c) {
		return
	}
	if c != 'r' && m > 0 {
		join(m-1, n, k, 'r', cnt)
	}
	if c != 'b' && n > 0 {
		join(m, n-1, k, 'b', cnt)
	}
	if c != 'g' && k > 0 {
		join(m, n, k-1, 'g', cnt)
	}
	return
}

//var l sync.Mutex

func main() {
	ts := time.Now()
	var m, n, k int
	fmt.Scanf("%d%d%d", &m, &n, &k)
	var cnt int = 0
	//ch := make(chan int, 3) // 多线程运行
	//go func(c chan int) {
	//	join(m-1, n, k, 'r', &cnt)
	//	c <- 1
	//}(ch)
	//go func(c chan int) {
	//	join(m, n-1, k, 'b', &cnt)
	//	c <- 1
	//}(ch)
	//go func(c chan int) {
	//	join(m, n, k-1, 'g', &cnt)
	//	c <- 1
	//}(ch)
	//<-ch
	//<-ch
	//<-ch
	join(m-1, n, k, 'r', &cnt) //顺序运行
	join(m, n-1, k, 'b', &cnt)
	join(m, n, k-1, 'g', &cnt)
	fmt.Printf("%d\n", cnt)
	fmt.Print(time.Since(ts))

}
