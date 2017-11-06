package main

import (
	"io/ioutil"
	"strings"

	"github.com/fatih/color"
)

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		color.Red(path)
	}
	check(err)
	return string(dat)
}

func writeFile(path string, newContent string) {
	err := ioutil.WriteFile(path, []byte(newContent), 0644)
	check(err)

	color.Green(path)
	//color.Blue(newContent)
}

func getLines(text string) []string {
	lines := strings.Split(text, "\n")
	return lines
}

func concatStringsWithJumps(lines []string) string {
	var r string
	for _, l := range lines {
		r = r + l + "\n"
	}
	return r
}
