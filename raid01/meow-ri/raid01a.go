package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {

	var topLeft rune = 'o'
	var topRight rune = 'o'
	var bottomLeft rune = 'o'
	var bottomRight rune = 'o'
	var vert rune = '|'
	var hori rune = '-'

	for j := 1; j <= y; j++ { //j is number of line
		for i := 1; i <= x; i++ { //i is number of column
			if j == 1 { //if it is top line
				if i == 1 { //if it is left column
					z01.PrintRune(topLeft)
				} else if i == x {
					z01.PrintRune(topRight)
				} else {
					z01.PrintRune(hori)
				}
			} else if j == y {
				if i == 1 { //if it is left column
					z01.PrintRune(bottomLeft)
				} else if i == x {
					z01.PrintRune(bottomRight)
				} else {
					z01.PrintRune(hori)
				}
			} else {
				if i == 1 { //if it is left column
					z01.PrintRune(vert)
				} else if i == x {
					z01.PrintRune(vert)
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}
