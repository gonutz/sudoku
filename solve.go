package sudoku

import (
	"errors"
	"strconv"
)

// Game is the 9x9 field of numbers, line by line.
// Values of 1-9 are treated as fixed.
// Values of 0 will be solved for.
type Game [9 * 9]int

func (g Game) String() string {
	s := ""
	for i, n := range g {
		if i%3 == 0 {
			s += " "
		}
		s += strconv.Itoa(n)
		if (i+1)%9 == 0 {
			s += "\n"
			if (i+1)%(3*9) == 0 {
				s += "\n"
			}
		}
	}
	return s
}

// Solve tries to find a valid combination of digits for all 0s in the given
// Game, considering all non-zero digits as fixed numbers. The first found valid
// solution is returned, even if multiple solutions would be possible.
// If a solution was found, the returned error is nil.
//
// Example usage:
//
//	fmt.Println(sudoku.Solve(sudoku.Game{
//		0, 0, 0, 0, 8, 4, 0, 0, 5,
//		0, 9, 0, 0, 0, 0, 0, 0, 3,
//		0, 0, 7, 0, 1, 3, 0, 0, 0,
//		0, 0, 5, 0, 0, 0, 1, 3, 0,
//		7, 0, 0, 0, 3, 0, 0, 0, 9,
//		0, 8, 3, 0, 0, 0, 2, 0, 0,
//		0, 0, 0, 6, 9, 0, 5, 0, 0,
//		2, 0, 0, 0, 0, 0, 0, 9, 0,
//		1, 0, 0, 5, 7, 0, 0, 0, 0,
//	}))
func Solve(g Game) (Game, error) {
	for _, n := range g {
		if !(0 <= n && n <= 9) {
			return g, errors.New("invalid digit " + strconv.Itoa(n) + " in game")
		}
	}

	if !validGame(g) {
		return g, errors.New("illegal input, conflicting fixed values")
	}

	if solve(&g) {
		return g, nil
	} else {
		return g, errors.New("unsolvable game")
	}
}

// validGame returns true if the non-zero digits so far are valid.
func validGame(g Game) bool {
	for i, n := range g {
		if n != 0 {
			g[i] = 0
			if !validAt(&g, n, i) {
				return false
			}
			g[i] = n
		}
	}
	return true
}

// solve returns true if a solution was found.
func solve(g *Game) bool {
	solvedAlready := true
	for i, n := range g {
		if n == 0 {
			solvedAlready = false
			for try := 1; try <= 9; try++ {
				if validAt(g, try, i) {
					g[i] = try
					if solve(g) {
						return true
					}
				}
			}
			g[i] = 0
			break
		}
	}
	return solvedAlready
}

// validAt returns true if the digit n can be placed at index at.
func validAt(g *Game, n, at int) bool {
	row := rowAt(at)
	col := colAt(at)
	block := blockAt(at)
	for i := range g {
		if g[i] == n && (row == rowAt(i) || col == colAt(i) || block == blockAt(i)) {
			return false
		}
	}
	return true
}

func rowAt(i int) int   { return i / 9 }
func colAt(i int) int   { return i % 9 }
func blockAt(i int) int { return (rowAt(i)/3)*3 + colAt(i)/3 }
