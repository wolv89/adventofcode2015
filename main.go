package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/wolv89/adventofcode2015/day01"
)

var (
	flagday, flagpuzzle, flagsample int
)

type AocDay interface {
	Puzzle1(int)
	Puzzle2(int)
}

func init() {

	flag.IntVar(&flagday, "day", 0, "Advent of Code day to run (Between 1 and 25)")
	flag.IntVar(&flagday, "d", 0, "Advent of Code day to run (Between 1 and 25)")

	flag.IntVar(&flagpuzzle, "puzzle", 1, "Which puzzle to run on the given day (1 or 2, defaults to 1)")
	flag.IntVar(&flagpuzzle, "p", 1, "Which puzzle to run on the given day (1 or 2, defaults to 1)")

	flag.IntVar(&flagsample, "sample", 0, "Use sample data, instead of full data? (Defaults to false)")
	flag.IntVar(&flagsample, "s", 0, "Use sample data, instead of full data? (Defaults to false)")

}

func validateFlags(lim int) {

	if flagday < 1 || flagday > lim {
		log.Fatalf("Please enter a day between 1 and %d", lim)
	}

	if flagpuzzle < 1 || flagpuzzle > 2 {
		log.Fatalf("Puzzle step can only be 1 or 2")
	}

}

func main() {

	days := []AocDay{
		day01.AocDay1{},
	}

	flag.Parse()
	validateFlags(len(days))

	day := days[flagday-1]

	if flagpuzzle == 1 {
		day.Puzzle1(flagsample)
	} else {
		day.Puzzle2(flagsample)
	}

	fmt.Println("")

}
