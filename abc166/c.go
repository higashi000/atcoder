package main

import (
	"fmt"
)

type Observatory struct {
	Id   int
	Conn []int
	Flg  bool
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var h []int
	var a []int
	var b []int

	for i := 0; i < n; i++ {
		var tmp int
		fmt.Scan(&tmp)
		h = append(h, tmp)
	}

	for i := 0; i < m; i++ {
		var tmp int
		fmt.Scan(&tmp)
		a = append(a, tmp)
		fmt.Scan(&tmp)
		b = append(b, tmp)
	}

	res := make([]bool, n)

	for i := 0; i < n; i++ {
		res[i] = true
	}

	for i := 0; i < m; i++ {
		if h[a[i]-1] == h[b[i]-1] {
			res[a[i]-1] = false
			res[b[i]-1] = false
		} else if h[a[i]-1] < h[b[i]-1] {
			res[a[i]-1] = false
		} else if h[a[i]-1] > h[b[i]-1] {
			res[b[i]-1] = false
		}

	}

	cnt := 0
	for i := 0; i < n; i++ {
		if res[i] {
			cnt++
		}
	}

	fmt.Println(cnt)
}
