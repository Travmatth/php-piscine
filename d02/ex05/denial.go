package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/signal"
	"regexp"
	"strings"
	"syscall"
)

func registerSignals(sigs chan os.Signal) {
	<-sigs
	os.Exit(1)
}

func loadFile(file string) ([]map[string]string, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	var maps []map[string]string
	reader := csv.NewReader(f)
	keys, err := reader.Read()
	if err != nil {
		return nil, err
	}
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}
		m := make(map[string]string)
		for i := 0; i < len(line); i++ {
			m[keys[i]] = line[i]
		}
		maps = append(maps, m)
	}
	return maps, nil
}

var r = regexp.MustCompile("^\\((.*),\\s(.*)\\)\\s\"(.*)\"\\s\\[(.*)\\]$")

func parseInput(input string) (key, val, format string, vals []string) {
	matches := r.FindAllStringSubmatchIndex(input, -1)
	key = input[matches[0][2]:matches[0][3]]
	val = input[matches[0][4]:matches[0][5]]
	format = input[matches[0][6]:matches[0][7]]
	vals = strings.Fields(input[matches[0][8]:matches[0][9]])
	return
}

var helpString = "\n (key, value) \"format string %s %s %s\" [values to print]"

func startLoop(entries []map[string]string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your command (\"help\" for format):")
		in, _ := reader.ReadString('\n')
		switch in {
		case "help":
			fmt.Println(helpString)
		default:
			var entry map[string]string
			key, val, format, vals := parseInput(in)
			for _, currentEntry := range entries {
				if entry[key] == val {
					entry = currentEntry
					break
				}
			}
			args := make([]interface{}, len(vals))
			for _, val := range vals {
				args = append(args, entry[val])
			}
			fmt.Printf(format, args...)
		}
	}
}

func main() {
	if len(os.Args) != 3 {
		return
	}
	entries, err := loadFile(os.Args[1])
	if err != nil {
		return
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go registerSignals(sigs)
	startLoop(entries)
}
