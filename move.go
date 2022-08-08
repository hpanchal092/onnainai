package main

import (
	"math"
)

func move(input string, b *[SIZE][SIZE]int) {
	// creates a slice representing a row/column on the board, then slides and
	// merges pieces on the row/column

	slice := []*int{}
	for i := 0; i < SIZE; i++ {
		switch input {
		case "u":
			slice = createVerticalSlice(b, i, true)
		case "d":
			slice = createVerticalSlice(b, i, false)
		case "l":
			slice = createHorizontalSlice(b, i, true)
		case "r":
			slice = createHorizontalSlice(b, i, false)
		}

		slideAndMerge(slice)
	}
}

func slideAndMerge(s []*int) {
	for i := 1; i < SIZE; i++ { // loop through the tiles
		if *s[i] != 0 { // if it has a proper value
			for pos := i; pos > 0; pos-- { // start sliding it towards the left
				currTile := s[pos]
				nextTile := s[pos-1]

				// merge if next tile is same as curr tile
				if *nextTile == *currTile {
					// makes it negative so other tiles can't merge with it in
					// the same move
					*nextTile = *currTile * -2
					*currTile = 0
					break
				}
				// stop sliding if next tile is not 0
				if *nextTile != 0 {
					break
				}
				// slide
				*nextTile = *currTile
				*currTile = 0
			}
		}
	}
	// remove the negative sign
	for i := 0; i < SIZE; i++ {
		*s[i] = int(math.Abs(float64(*s[i])))
	}
}

func createVerticalSlice(b *[SIZE][SIZE]int, n int, forwards bool) []*int {
	slice := make([]*int, 0, 4)

	if forwards {
		// create slice left, nth row
		for i := 0; i < SIZE; i++ {
			slice = append(slice, &b[i][n])
		}
	} else {
		// create slice right, nth row
		for i := SIZE - 1; i >= 0; i-- {
			slice = append(slice, &b[i][n])
		}
	}
	return slice
}

func createHorizontalSlice(b *[SIZE][SIZE]int, n int, forwards bool) []*int {
	slice := make([]*int, 0, 4)

	if forwards {
		// create slice upwards, nth column
		for i := 0; i < SIZE; i++ {
			slice = append(slice, &b[n][i])
		}
	} else {
		// create slice downwards, nth column
		for i := SIZE - 1; i >= 0; i-- {
			slice = append(slice, &b[n][i])
		}
	}
	return slice
}
