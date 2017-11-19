package main

import "golang.org/x/tour/pic"

var f int

// Pic draws a picture
func Pic(dx, dy int) [][]uint8 {
	p := make([][]uint8, dy)
	for i := range p {
		p[i] = make([]uint8, dx)
		for j := range p[i] {
			switch f {
			case 1:
				p[i][j] = (uint8)(i * j)
			case 2:
				p[i][j] = (uint8)((i + j) / 2)
			case 3:
				p[i][j] = (uint8)(i ^ j)
			}
		}
	}
	return p
}

func main() {
	f = 1
	pic.Show(Pic)
}
