package main

import "fmt"

func app(a []int, n int) []int {
	for j := 0; j < n; j++ {
		for i := 0; i < n-1; i++ {
			if a[i+1] < a[i] {
				m := a[i+1]
				a[i+1] = a[i]
				a[i] = m
			}
		}
	}
	return a
}

func main() {
	var a []int
	var l, r, n int
	fmt.Scanf("%d%d%d", &n, &l, &r)
	for i := 0; i < n; i++ {
		var x int
		fmt.Scanf("%d", &x)
		a = append(a, x)
	}
	b := a[l : r+1]
	c := app(b, r-l+1)
	for i := 0; i < len(c); i++ {
		a = append(a, c[i])
	}
	for i := 0; i < len(a)-(r-l+1); i++ {
		fmt.Printf("%d", a[i])
	}
}
