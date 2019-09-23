package main

import "testing"

func TestFindEmptyCase(t *testing.T) {
	game, _ := loadGame("864371259325849761971265843436192587198657432257483916689734125713528694542916370")
	i, j, err := findEmptyCase(game)
	if !(i == 8 && j == 8) || err != nil {
		t.Errorf("Should find empty case at position (8,8), returned (%d,%d) %s", i, j, err)
	}
}

func TestFindEmptyCaseNoMoreEmpty(t *testing.T) {
	game, _ := loadGame("864371259325849761971265843436192587198657432257483916689734125713528694542916378")
	i, j, err := findEmptyCase(game)
	if err == nil {
		t.Errorf("Should raise an error as no more space, but returned (%d,%d) %s", i, j, err)
	}
}

