package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Note struct {
	User     string `json:"user"`
	Note     int    `json:"Note"`
	Grader   string `json:"Grader"`
	Feedback int    `json:"Feedback"`
}

func main() {
	command := ""
	if len(os.Args) == 1 {
		command = os.Args[1]
	}
	var notes []Note
	reader := csv.NewReader(os.Stdin)
	reader.Read()
	for {
		line, err := reader.Read()
		if err == io.EOF || err != nil {
			break
		}
		grade := 0
		if line[1] != "" {
			grade, _ = strconv.Atoi(line[1])
		}
		feedback := 0
		if line[3] != "" {
			feedback, _ = strconv.Atoi(line[3])
		}
		notes = append(notes, Note{
			User:     line[0],
			Note:     grade,
			Grader:   line[2],
			Feedback: feedback,
		})
	}
	switch command {
	case "average":
		sum := 0
		for _, note := range notes {
			sum += note.Note
		}
		fmt.Println(sum / len(notes))
	case "average_user":
	case "moulinette_variance":
	}
}
