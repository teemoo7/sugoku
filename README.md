# sugoku

[![Build Status](https://travis-ci.org/teemoo7/sugoku.svg?branch=master)](https://travis-ci.org/teemoo7/sugoku) [![Go Report Card](https://goreportcard.com/badge/github.com/teemoo7/sugoku)](https://goreportcard.com/report/github.com/teemoo7/sugoku) [![Coverage](https://sonarcloud.io/api/project_badges/measure?project=ch.teemoo%3Asugoku&metric=coverage)](https://sonarcloud.io/dashboard?id=ch.teemoo%3Asugoku) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=ch.teemoo%3Asugoku&metric=alert_status)](https://sonarcloud.io/dashboard?id=ch.teemoo%3Asugoku)

sugoku is a basic Sudoku solver, written in Golang.

It takes a raw input string representing a sudoku grid (where `0` means empty) and tries to solve it using a simple backtracking algorithm. The idea was not to make it as efficient as possible, but just to play a bit with Golang features.

Such inputs can be found on many sites, among others [here](https://www.kaggle.com/bryanpark/sudoku).

## Example

Below is an example demonstrating how the solver works. 

### Raw input

    004300209005009001070060043006002087190007400050083000600000105003508690042910300

### Initial board

    _ _ 4 3 _ _ 2 _ 9
    _ _ 5 _ _ 9 _ _ 1
    _ 7 _ _ 6 _ _ 4 3
    _ _ 6 _ _ 2 _ 8 7
    1 9 _ _ _ 7 4 _ _
    _ 5 _ _ 8 3 _ _ _
    6 _ _ _ _ _ 1 _ 5
    _ _ 3 5 _ 8 6 9 _
    _ 4 2 9 1 _ 3 _ _


### Solution
    8 6 4 3 7 1 2 5 9
    3 2 5 8 4 9 7 6 1
    9 7 1 2 6 5 8 4 3
    4 3 6 1 9 2 5 8 7
    1 9 8 6 5 7 4 3 2
    2 5 7 4 8 3 9 1 6
    6 8 9 7 3 4 1 2 5
    7 1 3 5 2 8 6 9 4
    5 4 2 9 1 6 3 7 8
