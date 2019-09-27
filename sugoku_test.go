package main

import (
	"errors"
	"fmt"
	"testing"
	"github.com/stretchr/testify/assert"
)

type testPairArrayBool struct {
	array    [9]int
	expected bool
}

type testPairDataBool struct {
	data     string
	expected bool
}

type testPairDataError struct {
	data string
	err  error
}

type testTupleDataPosition struct {
	data string
	row  int
	col  int
	err  error
}

func TestLoadGame(t *testing.T) {
	var tests = [] testPairDataError{
		{"", errors.New("incorrect size for game raw data")},
		{"3821731298", errors.New("incorrect size for game raw data")},
		{"86437125932584976197126584343619258719865743225748391668973412571352869454291637812312312", errors.New("incorrect size for game raw data")},
		{"864371259325849761971265843436192587198657432257483916689734125713528694542916378", nil},
		{"abc371259325849761971265843436192587198657432257483916689734125713528694542916378", errors.New("unable to parse valid int")},
		{"884371259325849761971265843436192587198657432257483916689734125713528694542916378", errors.New("invalid game")},
	}

	for _, test := range tests {
		_, err := loadGame(test.data)
		assert.IsType(t, test.err, err, test.data)
	}
}

func TestIsValidGame(t *testing.T) {
	var tests = [] testPairDataBool{
		{"864371259325849761971265843436192587198657432257483916689734125713528694542916378", true},
		{"004300209005009001070060043006002087190007400050083000600000105003508690042910300", true},
		{"884300209005009001070060043006002087190007400050083000600000105003508690042910300", false},
		{"804300209805009001070060043006002087190007400050083000600000105003508690042910300", false},
	}

	for _, test := range tests {
		game, _ := loadGame(test.data)
		assert.Equal(t, test.expected, isValidGame(game), test.data)
	}
}

func TestHasDuplicates(t *testing.T) {
	var tests = [] testPairArrayBool{
		{[9] int{4, 0, 0, 0, 0, 0, 0, 0, 0}, false},
		{[9] int{4, 4, 0, 0, 0, 0, 0, 0, 0}, true},
		{[9] int{}, false},
		{[9] int{0, 2, 5, 2, 3, 0, 8, 7, 0}, true},
		{[9] int{0, 2, 5, 1, 3, 0, 8, 7, 0}, false},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, hasDuplicates(test.array), test.array)
	}

}

func TestIsGameOver(t *testing.T) {
	var tests = [] testPairDataBool{
		{"864371259325849761971265843436192587198657432257483916689734125713528694542916378", true},
		{"004300209005009001070060043006002087190007400050083000600000105003508690042910300", false},
	}

	for _, test := range tests {
		game, _ := loadGame(test.data)
		assert.Equal(t, test.expected, isGameOver(game), test.data)
	}
}

func TestFindEmptyCase(t *testing.T) {
	var tests = [] testTupleDataPosition{
		{"864371259325849761971265843436192587198657432257483916689734125713528694542916370", 8, 8, nil},
		{"864371259325849761971265843436192587198657432257483916689734125713528694542916378", -1, -1, errors.New("no empty case found")},
	}

	for _, test := range tests {
		game, _ := loadGame(test.data)
		i, j, err := findEmptyCase(game)
		assert.Equal(t, test.row, i, test.data)
		assert.Equal(t, test.col, j, test.data)
		assert.IsType(t, test.err, err, test.data)
	}
}

func TestSolve(t *testing.T) {
	init, _ := loadGame("004300209005009001070060043006002087190007400050083000600000105003508690042910300")
	solution, _ := loadGame("864371259325849761971265843436192587198657432257483916689734125713528694542916378")

	result, err := solve(init)
	assert.Equal(t, solution, result, init)
	assert.Nil(t, err, init)
}

func TestLoadAndSolveGame(t *testing.T) {
	originalLogFatal := logFatal

	// After this test, replace the original fatal function
	defer func() { logFatal = originalLogFatal }()

	var errors []string
	logFatal = func(args ...interface{}) {
		if len(args) > 0 {
			errors = append(errors, fmt.Sprint(args))
		} else {
			errors = append(errors)
		}
	}

	errors = nil
	loadAndSolveGame("")
	assert.Equal(t, len(errors), 1)

	errors = nil
	loadAndSolveGame("004300209005009001070060043006002087190007400050083000600000105003508690042910300")
	assert.Equal(t, len(errors), 0)
}
