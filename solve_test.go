package sudoku

import "testing"

func TestRowAt(t *testing.T) {
	checkInt(t, rowAt(0), 0)
	checkInt(t, rowAt(8), 0)
	checkInt(t, rowAt(9), 1)
	checkInt(t, rowAt(80), 8)
}

func TestColAt(t *testing.T) {
	checkInt(t, colAt(0), 0)
	checkInt(t, colAt(1), 1)
	checkInt(t, colAt(8), 8)
	checkInt(t, colAt(9), 0)
	checkInt(t, colAt(80), 8)
}

func TestBlockAt(t *testing.T) {
	checkInt(t, blockAt(0), 0)
	checkInt(t, blockAt(1), 0)
	checkInt(t, blockAt(2), 0)
	checkInt(t, blockAt(3), 1)
	checkInt(t, blockAt(8), 2)
	checkInt(t, blockAt(9), 0)
	checkInt(t, blockAt(26), 2)
	checkInt(t, blockAt(27), 3)
	checkInt(t, blockAt(80), 8)
}

func TestInvalidDigit(t *testing.T) {
	g := Game{12}
	unchanged, err := Solve(g)
	checkError(t, err, "invalid digit 12 in game")
	checkGame(t, unchanged, g)
}

func TestImpossibleGame(t *testing.T) {
	g := Game{
		// the top-left block cannot have a 1
		0, 0, 0, 1, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 1, 0, 0,
		4, 5, 6, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	unchanged, err := Solve(g)
	checkError(t, err, "unsolvable game")
	checkGame(t, unchanged, g)
}

func TestInvalidFixedNumbers(t *testing.T) {
	g := Game{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 7, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 7, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}
	unchanged, err := Solve(g)
	checkError(t, err, "illegal input, conflicting fixed values")
	checkGame(t, unchanged, g)
}

func TestSolveValidPuzzle(t *testing.T) {
	solved, err := Solve(Game{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		7, 8, 9, 1, 2, 3, 4, 5, 6,
		4, 5, 6, 7, 8, 9, 1, 2, 3,
		9, 1, 2, 3, 4, 5, 6, 7, 8,
		6, 7, 8, 9, 0, 2, 3, 4, 5, // <- missing 1 in the center
		3, 4, 5, 6, 7, 8, 9, 1, 2,
		8, 9, 1, 2, 3, 4, 5, 6, 7,
		5, 6, 7, 8, 9, 1, 2, 3, 4,
		2, 3, 4, 5, 6, 7, 8, 9, 1,
	})
	if err != nil {
		t.Fatal("unexpected error:", err)
	}
	checkGame(t, solved, Game{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		7, 8, 9, 1, 2, 3, 4, 5, 6,
		4, 5, 6, 7, 8, 9, 1, 2, 3,
		9, 1, 2, 3, 4, 5, 6, 7, 8,
		6, 7, 8, 9, 1, 2, 3, 4, 5,
		3, 4, 5, 6, 7, 8, 9, 1, 2,
		8, 9, 1, 2, 3, 4, 5, 6, 7,
		5, 6, 7, 8, 9, 1, 2, 3, 4,
		2, 3, 4, 5, 6, 7, 8, 9, 1,
	})
}

func TestSolutionCountCanBeCalculated(t *testing.T) {
	if !HasUniqueSolution(Game{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		7, 8, 9, 1, 2, 3, 4, 5, 6,
		4, 5, 6, 7, 8, 9, 1, 2, 3,
		2, 3, 4, 5, 6, 7, 8, 9, 1,
		8, 9, 1, 2, 3, 4, 5, 6, 7,
		5, 6, 7, 8, 9, 1, 2, 3, 4,
		3, 4, 5, 6, 7, 8, 9, 1, 2,
		9, 1, 2, 3, 4, 5, 6, 7, 8,
		6, 7, 8, 9, 1, 2, 3, 4, 5,
	}) {
		t.Error("finished game has 1 solution")
	}

	if !HasUniqueSolution(Game{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		7, 8, 9, 1, 2, 3, 4, 5, 6,
		4, 5, 6, 7, 8, 9, 1, 2, 3,
		2, 3, 4, 5, 6, 7, 0, 9, 1,
		8, 9, 1, 2, 3, 4, 5, 6, 7,
		5, 6, 7, 8, 9, 1, 2, 3, 4,
		3, 4, 5, 6, 7, 8, 9, 1, 2,
		9, 1, 2, 3, 4, 5, 6, 7, 8,
		6, 7, 8, 9, 1, 2, 3, 4, 5,
	}) {
		t.Error("game with only one missing number has unique solution")
	}

	if HasUniqueSolution(Game{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}) {
		t.Error("empty game has multiple solutions")
	}

	if HasUniqueSolution(Game{
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 10, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}) {
		t.Error("invalid numbers in game means no solution")
	}

	if HasUniqueSolution(Game{
		1, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 1, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 1, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0,
	}) {
		t.Error("invalid game has no solution")
	}
}

func TestStringRepresentation(t *testing.T) {
	have := Game{
		1, 2, 3, 4, 5, 6, 7, 8, 9,
		4, 5, 6, 7, 8, 9, 1, 2, 3,
		7, 8, 9, 1, 2, 3, 4, 5, 6,
		2, 3, 4, 5, 6, 7, 8, 9, 1,
		5, 6, 7, 8, 9, 1, 2, 3, 4,
		8, 9, 1, 2, 3, 4, 5, 6, 7,
		3, 4, 5, 6, 7, 8, 9, 1, 2,
		6, 7, 8, 9, 1, 2, 3, 4, 5,
		9, 1, 2, 3, 4, 5, 6, 7, 8,
	}.String()
	want := `123 456 789
456 789 123
789 123 456

234 567 891
567 891 234
891 234 567

345 678 912
678 912 345
912 345 678`
	if have != want {
		t.Errorf("want:\n%s\nbut have:\n%s", want, have)
	}
}

func checkInt(t *testing.T, have, want int) {
	if have != want {
		t.Errorf("want %v but have %v", want, have)
	}
}

func checkError(t *testing.T, have error, wantMsg string) {
	if have == nil {
		t.Fatal("error expected but was nil")
	}
	msg := have.Error()
	if msg != wantMsg {
		t.Errorf("unexpected error message, want '%s' but have '%s'", wantMsg, msg)
	}
}

func checkGame(t *testing.T, have, want Game) {
	if have != want {
		t.Errorf("games differ, want %v but have %v", want, have)
	}
}
