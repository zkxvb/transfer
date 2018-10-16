package main

import (
	"fmt"
)

// 递归思想：
// 一个一个字母安排，安排完则计数 ++，安排不了则返回

// 安排的了？
//func canJoin(m, n, k int, c byte) bool {
//	switch c {
//	case 'r':
//		if n+k == 0 {
//			return false
//		}
//	case 'b':
//		if m+k == 0 {
//			return false
//		}
//	case 'g':
//		if m+n == 0 {
//			return false
//		}
//	}
//	return true
//}

// 递归安排
//func join(m, n, k int, c byte, cnt *int) {
//	// 递归出口
//	if m+n+k == 0 {
//		//l.Lock()
//		*cnt++
//		//l.Unlock()
//		return
//	}
//	if !canJoin(m, n, k, c) {
//		return
//	}
//	if c != 'r' && m > 0 {
//		join(m-1, n, k, 'r', cnt)
//	}
//	if c != 'b' && n > 0 {
//		join(m, n-1, k, 'b', cnt)
//	}
//	if c != 'g' && k > 0 {
//		join(m, n, k-1, 'g', cnt)
//	}
//	return
//}

//var l sync.Mutex
var db [20][20][20][4]int

func init() {
	for i := 0; i < 20; i++ {
		for j := 0; j < 20; j++ {
			for k := 0; k < 20; k++ {
				for q := 0; q < 4; q++ {
					db[i][j][k][q] = -1
				}
			}
		}
	}
	for i := 0; i < 4; i++ {
		if i == 0 {
			db[0][0][0][i] = 0
		} else {
			db[0][0][0][i] = 1
		}
	}
	for i := 0; i < 4; i++ {
		if i == 1 {
			db[1][0][0][i] = 0
		} else {
			db[1][0][0][i] = 1
		}
	}
	for i := 0; i < 4; i++ {
		if i == 2 {
			db[0][1][0][i] = 0
		} else {
			db[0][1][0][i] = 1
		}
	}
	for i := 0; i < 4; i++ {
		if i == 3 {
			db[0][0][1][i] = 0
		} else {
			db[0][0][1][i] = 1
		}
	}
}

func stat(m, n, k int) int {
	var res int = 0
	if m > 0 {
		res = 1 << 2
	}
	if n > 0 {
		res += 1 << 1
	}
	if k > 0 {
		res += 1
	}
	return res
}

func getUAdj(m, n, k int, c byte) int {

	t := stat(m, n, k)

	switch c {
	case ' ':
		if db[m][n][k][0] == -1 {
			switch t {
			case 0:
				db[m][n][k][0] = 0
			case 1:
				db[m][n][k][0] = getUAdj(m, n, k-1, 'g')
			case 2:
				db[m][n][k][0] = getUAdj(m, n-1, k, 'b')
			case 3:
				db[m][n][k][0] = getUAdj(m, n-1, k, 'b') + getUAdj(m, n, k-1, 'g')
			case 4:
				db[m][n][k][0] = getUAdj(m-1, n, k, 'r')
			case 5:
				db[m][n][k][0] = getUAdj(m-1, n, k, 'r') + getUAdj(m, n, k-1, 'g')
			case 6:
				db[m][n][k][0] = getUAdj(m-1, n, k, 'r') + getUAdj(m, n-1, k, 'b')
			case 7:
				db[m][n][k][0] = getUAdj(m-1, n, k, 'r') + getUAdj(m, n-1, k, 'b') + getUAdj(m, n, k-1, 'g')
			}
		}
		return db[m][n][k][0]
	case 'r':
		if db[m][n][k][1] == -1 {
			switch t {
			case 1, 5:
				db[m][n][k][1] = getUAdj(m, n, k-1, 'g')
			case 2, 6:
				db[m][n][k][1] = getUAdj(m, n-1, k, 'b')
			case 3, 7:
				db[m][n][k][1] = getUAdj(m, n-1, k, 'b') + getUAdj(m, n, k-1, 'g')
			case 4:
				db[m][n][k][1] = 0
			}
		}
		return db[m][n][k][1]
	case 'b':
		if db[m][n][k][2] == -1 {
			switch t {
			case 1, 3:
				db[m][n][k][2] = getUAdj(m, n, k-1, 'g')
			case 4, 6:
				db[m][n][k][2] = getUAdj(m-1, n, k, 'r')
			case 5, 7:
				db[m][n][k][2] = getUAdj(m-1, n, k, 'r') + getUAdj(m, n, k-1, 'g')
			case 2:
				db[m][n][k][2] = 0
			}
		}
		return db[m][n][k][2]
	case 'g':
		if db[m][n][k][3] == -1 {
			switch t {
			case 4, 5:
				db[m][n][k][3] = getUAdj(m-1, n, k, 'r')
			case 2, 3:
				db[m][n][k][3] = getUAdj(m, n-1, k, 'b')
			case 6, 7:
				db[m][n][k][3] = getUAdj(m-1, n, k, 'r') + getUAdj(m, n-1, k, 'b')
			case 1:
				db[m][n][k][3] = 0
			}
		}
		return db[m][n][k][3]
	default:
		return -1
	}
}

func main() {

	var m, n, k int
	for {
		fmt.Scanf("%d%d%d", &m, &n, &k)
		var cnt int = 0
		cnt = getUAdj(m, n, k, ' ')
		fmt.Printf("%d\n", cnt)
	}
	//fmt.Scanf("%d%d%d", &m, &n, &k)
	//var cnt int = 0
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
	//join(m-1, n, k, 'r', &cnt) //顺序运行
	//join(m, n-1, k, 'b', &cnt)
	//join(m, n, k-1, 'g', &cnt)
	//cnt = getUAdj(m, n, k, ' ')
	//fmt.Printf("%d\n", cnt)
}
