# sugoku

[![Build Status](https://travis-ci.org/teemoo7/sugoku.svg?branch=master)](https://travis-ci.org/teemoo7/sugoku) [![Go Report Card](https://goreportcard.com/badge/github.com/teemoo7/sugoku)](https://goreportcard.com/report/github.com/teemoo7/sugoku)

sugoku is a basic Sudoku solver, written in Golang.

It takes a raw input string representing a sudoku grid (where `0` means empty) and tries to solve it using a simple backtracking algorithm. The idea was not to make it as efficient as possible to see how it works.

Such inputs can be found on many sites, among others [here](https://www.kaggle.com/bryanpark/sudoku).
