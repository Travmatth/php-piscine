package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"sort"
)

type Note struct {
	User     string `json:"user"`
	Note     float64    `json:"Note"`
	Grader   string `json:"Grader"`
	Feedback float64    `json:"Feedback"`
}

func main() {
	command := ""
	if len(os.Args) == 1 {
		command = os.Args[1]
	}
	var notes []Note
	totals := map[string]float64{
		"users": 0.0,
		"note": 0.0,
		"feedback": 0.0,
	}
	reader := csv.NewReader(os.Stdin)
	reader.Read()
	for {
		line, err := reader.Read()
		if err == io.EOF || err != nil {
			break
		}
		totals["users"] += 1
		grade := 0.0
		if line[1] != "" {
			grade, _ = strconv.ParseFloat(line[1], 64)
			totals["note"] += grade
		}
		feedback := 0.0
		if line[3] != "" {
			feedback, _ = strconv.ParseFloat(line[3], 64)
			totals["feedback"] += feedback
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
		fmt.Println(totals["note"] / totals["users"])
	case "average_user":
		sort.SliceStable(notes, func(i, j int) bool {
			return notes[i].User < notes[j].User
		})
	case "moulinette_variance":
	}
}
