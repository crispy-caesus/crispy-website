package code_pages

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func BlogHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fileReader()))
}

func fileReader() string {
	f, err := os.Open("blog/source/hi.md")
	if err != nil {
		log.Fatal(err)
	}

	var html string
	var text string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		text += line + "\n"

		if len(line) > 1 {
			if line[0] == '#' {
				var currentCharIndex int
				for currentCharIndex = 1; currentCharIndex < 3; currentCharIndex++ {
					if line[currentCharIndex] == '#' {
						continue
					} else if line[currentCharIndex] == ' ' {
						break
					} else {
						log.Fatal("tag detected")
					}
				}
				html += fmt.Sprintf("<h%d>%s</h%d>\n", currentCharIndex, line[currentCharIndex:], currentCharIndex)
			}
		}
	}

	return html
}
