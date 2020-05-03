package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	var s []string

	for i := 0; i < N; i++ {
		var tmp string
		fmt.Scan(&tmp)
		s = append(s, tmp)
	}

	m := make(map[string]struct{})

	delDup := make([]string, 0)

	for _, e := range s {
		_, ok := m[e]
		if !ok {
			m[e] = struct{}{}
			delDup = append(delDup, e)
		}
	}

	fmt.Println(len(delDup))
}
