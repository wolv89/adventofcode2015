package day01

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type AocDay1 struct{}

const DIR = "day01/"

func (d AocDay1) Puzzle1(useSample int) {

	// Magic
	fmt.Println(232)

}

func (d AocDay1) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	if useSample >= 1 {
		fmt.Println("Whoops, there is no sample data for this one!")
		return
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		line string
	)

	for scanner.Scan() {

		line = scanner.Text()

		// Just a template to be copied...
		fmt.Println(line)

	}

}
