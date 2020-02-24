package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

// Mardi 12 Novembre 2013 12:02:21
// Day_of_the_week Number_of_day Month Year Hours:Minutes:Seconds
var regex = "[a-zA-z][A-z]+ [0-9][0-9]? [a-zA-Z][a-z]* [0-9]{4} [0-9]{2}:[0-9]{2}:[0-9]{2}"

var days = map[string]string{
	"lundi":    "Monday",
	"mardi":    "Tuesday",
	"mercredi": "Wednesday",
	"jeudi":    "Thursday",
	"vendredi": "Friday",
	"samedi":   "Saturday",
	"dimanche": "Sunday",
}

var months = map[string]string{
	"janvier":   "Jan",
	"fevrier":   "Feb",
	"mars":      "Mar",
	"avril":     "Apr",
	"mai":       "May",
	"juin":      "Jun",
	"juillet":   "Jul",
	"aout":      "Aug",
	"septembre": "Sep",
	"octobre":   "Oct",
	"novembre":  "Nov",
	"décembre":  "Dec",
}

const (
	DayFr = iota
	DayNbr
	MonFr
	Year
	Timestamp
)

func fmtNbr(str string) string {
	if l := len(str); l == 2 {
		return str
	} else if l == 1 {
		return fmt.Sprintf("0%s", str)
	}
	return str[2:]
}

func makeDateString(str string) (string, error) {
	parts := strings.Fields(str)
	day, dayPresent := days[strings.ToLower(parts[DayFr])]
	month, monthPresent := months[strings.ToLower(parts[MonFr])]
	if !dayPresent || !monthPresent {
		return "", fmt.Errorf("Wrong Format")
	}
	date := fmt.Sprintf("%s-%s-%s", fmtNbr(parts[DayNbr]), month, fmtNbr(parts[Year]))
	value := fmt.Sprintf("%s, %s %s CET", day, date, parts[Timestamp])
	return value, nil
}

func main() {
	r := regexp.MustCompile(regex)
	if len(os.Args) <= 1 {
		os.Exit(0)
	} else if is := r.MatchString(os.Args[1]); is {
		if val, err := makeDateString(os.Args[1]); err == nil {
			t, err := time.Parse(time.RFC850, val)
			if err != nil {
				fmt.Println("Wrong Format")
			}
			fmt.Println(t.Unix() - 3600)
		} else {
			fmt.Println("Wrong Format")
		}
		return
	}
	fmt.Println("Wrong Format")
}
