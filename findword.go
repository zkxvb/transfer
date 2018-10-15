package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type dictionary struct {
	col   int
	row   int
	chars [][]byte
}

func NewDict(col, row int, chars [][]byte) *dictionary {
	return &dictionary{
		col:   col,
		row:   row,
		chars: chars,
	}
}

type pos struct {
	x, y int
}

func (d dictionary) findChar(c byte) []*pos {
	var ps []*pos
	for i := 0; i < d.row; i++ {
		for j := 0; j < d.col; j++ {
			if d.chars[i][j] == c {
				p := &pos{x: i, y: j}
				ps = append(ps, p)
			}
		}
	}
	return ps
}

func (d dictionary) findCharNearPos(c byte, p *pos) (bool, *pos) {
	if p.x < 0 || p.x >= d.row || p.y < 0 || p.y >= d.col {
		return false, nil
	}

	// zuo you bian jie
	if p.y == 0 || p.y == d.col-1 {
		if p.y == 0 { // zuo bian jie
			// quan bu xiang you zhao
			if d.chars[p.x][p.y+1] == c {
				return true, &pos{p.x, p.y + 1}
			}
		} else { // you bian jie
			// quan bu xiang zuo zhao
			if d.chars[p.x][p.y-1] == c {
				return true, &pos{p.x, p.y + 1}
			}
		}
		if p.x == 0 { // shang jiao
			if d.chars[p.x+1][p.y] == c {
				return true, &pos{p.x + 1, p.y}
			}
		} else if p.x == d.row-1 { // xia jiao
			if d.chars[p.x-1][p.y] == c {
				return true, &pos{p.x - 1, p.y}
			}
		} else { // normal
			if d.chars[p.x+1][p.y] == c {
				return true, &pos{p.x + 1, p.y}
			}
			if d.chars[p.x-1][p.y] == c {
				return true, &pos{p.x - 1, p.y}
			}
		}
		return false, nil
	}

	// shang xia bian jie
	if p.x == 0 || p.x == d.row-1 {
		if p.x == 0 { // shang bian jie
			// quan bu xiang xia zhao
			if d.chars[p.x+1][p.y] == c {
				return true, &pos{p.x + 1, p.y}
			}
		} else { // xia bian jie
			// quan bu xiang shang zhao
			if d.chars[p.x-1][p.y] == c {
				return true, &pos{p.x - 1, p.y}
			}
		}
		if p.y == 0 { // zuo jiao
			if d.chars[p.x][p.y+1] == c {
				return true, &pos{p.x, p.y + 1}
			}
		} else if p.y == d.col-1 { // you jiao
			if d.chars[p.x][p.y-1] == c {
				return true, &pos{p.x, p.y - 1}
			}
		} else { // normal
			if d.chars[p.x][p.y+1] == c {
				return true, &pos{p.x, p.y + 1}
			}
			if d.chars[p.x][p.y-1] == c {
				return true, &pos{p.x, p.y - 1}
			}
		}
		return false, nil
	}

	// normal
	if d.chars[p.x][p.y+1] == c {
		return true, &pos{p.x, p.y + 1}
	}
	if d.chars[p.x][p.y-1] == c {
		return true, &pos{p.x, p.y - 1}
	}
	if d.chars[p.x+1][p.y] == c {
		return true, &pos{p.x + 1, p.y}
	}
	if d.chars[p.x-1][p.y-1] == c {
		return true, &pos{p.x - 1, p.y}
	}
	return false, nil
}

func (d dictionary) findWord(word string) bool {
	w := []byte(word)
	ps := d.findChar(w[0])
	if len(ps) == 0 {
		return false
	}
	if len(w) == 1 {
		return true
	}

	for _, p := range ps {
		var innerBreak = false
		tp := p
		for i := 0; i < len(w)-1; i++ {
			ct, np := d.findCharNearPos(w[i+1], tp)
			if !ct {
				innerBreak = true
				break
			}
			tp = np
		}
		if innerBreak {
			continue
		}
		return true
	}

	return false
}

func main() {
	var m, n, k int
	fmt.Scanf("%d %d %d\n", &m, &n, &k)

	reader := bufio.NewReader(os.Stdin)
	strWord, _, _ := reader.ReadLine()
	strWords := strings.Fields(string(strWord))

	var chars [][]byte
	for i := 0; i < n; i++ {
		var s []byte
		ch, _, _ := reader.ReadLine()
		chs := strings.Fields(string(ch))
		for j := 0; j < m; j++ {
			s = append(s, []byte(chs[j])[0])
		}
		chars = append(chars, s)
	}

	dict := NewDict(m, n, chars)
	var first = true
	for _, w := range strWords {
		if dict.findWord(w) {
			if !first {
				fmt.Println()
			}
			fmt.Printf("%s", w)
			first = false
		}
	}
}
