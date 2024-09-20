package main

import (
	"errors"
	"fmt"
	"os"
)

const (
	rows, cols = 9, 9
	empty      = 0
)

type Cell struct {
	digit int8
	fixed bool
}

type Grid [rows][cols]Cell

var (
	ErrBounds     = errors.New("out of bounds")
	ErrDigit      = errors.New("invalid digit")
	ErrInRow      = errors.New("digit already present in the row")
	ErrInColumn   = errors.New("digit already present in this column")
	ErrInRegion   = errors.New("digit already present in this region")
	ErrFixedDigit = errors.New("initial digits cannot be overwritten")
)

// new sudoku grid
func NewSudoku(digits [rows][cols]int8) *Grid {
	var grid Grid

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			d := digits[r][c]
			if d != empty {
				grid[r][c].digit = d
				grid[r][c].fixed = true
			}
		}
	}

	return &grid
}

func (g *Grid) Set(row, col int, digit int8) error {
	switch {
	case !inBounds(row, col):
		return ErrBounds

	case !validDigit(digit):
		return ErrDigit

	case g.isFixed(row, col):
		return ErrFixedDigit

	case g.inRow(row, digit):
		return ErrInRow

	case g.inColumn(col, digit):
		return ErrInColumn

	case g.inRegion(row, col, digit):
		return ErrInRegion
	}

	g[row][col].digit = digit
	return nil
}

func (g *Grid) inRegion(row, col int, digit int8) bool {
	startRow, startCol := row/3*3, col/3*3

	for r := startRow; r < startRow+3; r++ {
		for c := startCol; c < startCol+3; c++ {
			if g[r][c].digit == digit {
				return true
			}
		}
	}
	return false
}

func (g *Grid) inRow(row int, digit int8) bool {
	for c := 0; c < cols; c++ {
		if g[row][c].digit == digit {
			return true
		}
	}

	return false
}

func (g *Grid) inColumn(col int, digit int8) bool {
	for r := 0; r < rows; r++ {
		if g[r][col].digit == digit {
			return true
		}
	}
	return false
}

func (g *Grid) isFixed(row, col int) bool {
	return g[row][col].fixed
}

func inBounds(row, col int) bool {
	if row < 0 || col < 0 || row >= rows || col >= cols {
		return false
	}
	return true
}

func validDigit(digit int8) bool {
	return digit >= 1 && digit <= 9
}

func main() {
	s := NewSudoku([rows][cols]int8{
		{5, 3, 0, 0, 7, 0, 0, 0, 0},
		{6, 0, 0, 1, 9, 5, 0, 0, 0},
		{0, 9, 8, 0, 0, 0, 0, 6, 0},
		{8, 0, 0, 0, 6, 0, 0, 0, 3},
		{4, 0, 0, 8, 0, 3, 0, 0, 1},
		{7, 0, 0, 0, 2, 0, 0, 0, 6},
		{0, 6, 0, 0, 0, 0, 2, 8, 0},
		{0, 0, 0, 4, 1, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 7, 9},
	})

	err := s.Set(1,1,4)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, row := range s{
		fmt.Println(row)
	}
}
