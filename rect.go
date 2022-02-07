package main

import "fmt"

type rect struct {
	t int
	b int
	l int
	r int
}

func collide(a rect, b rect) bool {
	if a.r > b.l {
		if a.l < b.r {
			if a.b > b.t {
				if a.t < b.b {
					fmt.Println("A: ", a, "B: ", b)
					return T
				}
			}
		}
	}

	return F
}
