package day02

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type AocDay2 struct{}

const DIR = "day02/"

type Gift struct {
	l, w, h int
}

func (g Gift) MeasureWrappingPaper() int {

	s1, s2, s3 := g.l*g.w, g.l*g.h, g.w*g.h

	return 2*s1 + 2*s2 + 2*s3 + min(s1, s2, s3)

}

func (g Gift) MeasureRibbon() int {

	sides := []int{g.l, g.w, g.h}

	slices.Sort(sides)

	return 2*sides[0] + 2*sides[1] + g.l*g.w*g.h

}

func (d AocDay2) Puzzle1(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		gift      Gift
		line      string
		wp, total int
	)

	for scanner.Scan() {

		line = scanner.Text()

		dimensions := strings.Split(line, "x")

		gift = Gift{}
		gift.l, _ = strconv.Atoi(dimensions[0])
		gift.w, _ = strconv.Atoi(dimensions[1])
		gift.h, _ = strconv.Atoi(dimensions[2])

		wp = gift.MeasureWrappingPaper()
		total += wp

		fmt.Println("[", gift.l, "x", gift.w, "x", gift.h, "] | ", wp, " | ", total)

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}

func (d AocDay2) Puzzle2(useSample int) {

	datafile := DIR + "data.txt"
	if useSample == 1 {
		datafile = DIR + "sample.txt"
	}

	f, err := os.Open(datafile)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var (
		gift      Gift
		line      string
		rb, total int
	)

	for scanner.Scan() {

		line = scanner.Text()

		dimensions := strings.Split(line, "x")

		gift = Gift{}
		gift.l, _ = strconv.Atoi(dimensions[0])
		gift.w, _ = strconv.Atoi(dimensions[1])
		gift.h, _ = strconv.Atoi(dimensions[2])

		rb = gift.MeasureRibbon()
		total += rb

		fmt.Println("[", gift.l, "x", gift.w, "x", gift.h, "] | ", rb, " | ", total)

	}

	fmt.Println("")
	fmt.Println("Total: ", total)

}
