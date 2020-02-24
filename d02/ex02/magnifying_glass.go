package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

var titleExpr = "<a.*?title=\"(.*?)\">"
var anchorExpr = "<a.*?>(.*?)<"

func replaceRegex(file string, regex string) string {
	r := regexp.MustCompile(regex)
	matches := r.FindAllStringSubmatchIndex(file, -1)
	for i := len(matches) - 1; i >= 0; i-- {
		start := matches[i][2]
		end := matches[i][3]
		pre := file[:start]
		match := strings.ToUpper(file[start:end])
		post := file[end:]
		file = pre + match + post
	}
	return file
}

func main() {
	if len(os.Args) > 1 {
		f := os.Args[1]
		if file, err := ioutil.ReadFile(f); err == nil {
			str := replaceRegex(replaceRegex(string(file), titleExpr), anchorExpr)
			fmt.Println(str)
		}
	}
}
