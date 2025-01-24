package day03

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay3 struct{}

const DIR = "day03/"

// Magic number style
// Full data set has around 8,000 moves, assuming we won't ever head directly out that far
// (Famous last words??)
const SIZE = 4096

type Point struct {
	x, y int
}

func (d AocDay3) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	} else if useSample == 2 {
		datafile = DIR + "sample2.txt"
	} else if useSample == 3 {
		datafile = DIR + "sample3.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// One liner
	scanner.Scan()
	line := scanner.Text()

	grid := make([][]bool, SIZE)
	for g := range grid {
		grid[g] = make([]bool, SIZE)
	}

	houses := 1
	santa := Point{SIZE / 2, SIZE / 2}

	grid[santa.x][santa.y] = true

	for m := 0; m < len(line); m++ {

		switch line[m] {
		case '^':
			santa.x--
		case 'v':
			santa.x++
		case '<':
			santa.y--
		case '>':
			santa.y++
		}

		if !grid[santa.x][santa.y] {
			grid[santa.x][santa.y] = true
			houses++
		}

	}

	fmt.Println("")
	fmt.Println("Houses: ", houses)

}

func (d AocDay3) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	} else if useSample == 2 {
		datafile = DIR + "sample2.txt"
	} else if useSample == 3 {
		datafile = DIR + "sample3.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// One liner
	scanner.Scan()
	line := scanner.Text()

	grid := make([][]bool, SIZE)
	for g := range grid {
		grid[g] = make([]bool, SIZE)
	}

	houses := 1
	santa, robot := Point{SIZE / 2, SIZE / 2}, Point{SIZE / 2, SIZE / 2}
	isSantaMove := true

	grid[santa.x][santa.y] = true

	for m := 0; m < len(line); m++ {

		if isSantaMove {
			switch line[m] {
			case '^':
				santa.x--
			case 'v':
				santa.x++
			case '<':
				santa.y--
			case '>':
				santa.y++
			}

			if !grid[santa.x][santa.y] {
				grid[santa.x][santa.y] = true
				houses++
			}
		} else {
			switch line[m] {
			case '^':
				robot.x--
			case 'v':
				robot.x++
			case '<':
				robot.y--
			case '>':
				robot.y++
			}

			if !grid[robot.x][robot.y] {
				grid[robot.x][robot.y] = true
				houses++
			}
		}

		isSantaMove = !isSantaMove

	}

	fmt.Println("")
	fmt.Println("Houses: ", houses)

}
