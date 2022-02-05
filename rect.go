package main

type rect struct {
	t int
	b int
	l int
	r int
}

func collide(a rect, b rect) bool {
	// if the bottom of B is higher than the top of A
	if b.b > a.t {
		return F
	}

	if b.r < a.l {
		return F
	}
	return T

}
