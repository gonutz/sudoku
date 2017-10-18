Sudoku Solver
=============

This package provides a Go function to solve sudokus.

Example:

```Go
package main

import (
	"fmt"
	"github.com/gonutz/sudoku"
)

func main() {
	fmt.Println(sudoku.Solve(sudoku.Game{
		0, 0, 0, 0, 8, 4, 0, 0, 5,
		0, 9, 0, 0, 0, 0, 0, 0, 3,
		0, 0, 7, 0, 1, 3, 0, 0, 0,
		0, 0, 5, 0, 0, 0, 1, 3, 0,
		7, 0, 0, 0, 3, 0, 0, 0, 9,
		0, 8, 3, 0, 0, 0, 2, 0, 0,
		0, 0, 0, 6, 9, 0, 5, 0, 0,
		2, 0, 0, 0, 0, 0, 0, 9, 0,
		1, 0, 0, 5, 7, 0, 0, 0, 0,
	}))
}
```

Output:

```
 631 784 925
 894 265 713
 527 913 684

 965 427 138
 712 836 459
 483 159 276

 378 692 541
 256 341 897
 149 578 362

 <nil>
```