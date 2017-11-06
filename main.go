package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

func main() {
	//savelog()

	log.Println("kunigu")

	scrapDirectory(".")
}

func scrapDirectory(path string) {
	filesList, _ := ioutil.ReadDir(path)
	for _, f := range filesList {
		fileNameSplitted := strings.Split(f.Name(), ".")
		if len(fileNameSplitted) > 1 {
			if fileNameSplitted[1] == "txt" || fileNameSplitted[1] == "html" {
				file := readFile(path + "/" + f.Name())
				if strings.Contains(file, "{{kunigu ") {
					fmt.Println(path + "/" + f.Name())
					r := kuniguFile(file)
					writeFile(path+"/"+fileNameSplitted[0]+"OUT."+fileNameSplitted[1], r)
				}
			}
		} else {
			//is a directory
			//fmt.Println(path + "/" + f.Name())
			scrapDirectory(path + "/" + f.Name())
		}
	}
}

func kuniguFile(file string) string {
	lines := getLines(file)
	var resultL []string
	for _, l := range lines {
		if strings.Contains(l, "{{kunigu ") {
			var htmlcontent string
			includefile := strings.Split(l, " @")[1]
			includefile = strings.Replace(includefile, "}}", "", -1)
			if strings.Contains(l, "--md-to-html") {
				mdcontent := readFile(includefile)
				htmlcontent = string(blackfriday.Run([]byte(mdcontent)))
			} else {
				htmlcontent = readFile(includefile)
			}
			resultL = append(resultL, htmlcontent)
		} else {
			resultL = append(resultL, l)
		}
	}
	result := concatStringsWithJumps(resultL)
	return result
}
