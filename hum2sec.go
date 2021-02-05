package machine2human

import (
	"strconv"
	"strings"
	"text/scanner"
)

var stringToSec = map[string]int{
	"y": year,
	"г": year,
	"л": year,
	"M": month,
	"М": month,
	"w": week,
	"н": week,
	"d": day,
	"д": day,
	"h": hour,
	"ч": hour,
	"m": minute,
	"м": minute,
	"s": second,
	"с": second,
}

// Hum2Sec implement convertation readable string to seconds
func Hum2Sec(input string) int {
	// if string just is number for example "1332"
	x, err := strconv.Atoi(input)
	if err == nil {
		return x
	}

	var seconds int
	input = strings.TrimSpace(input)

	matched := matcher(input)
	for _, m := range matched {
		temp := m

		t, _ := strconv.Atoi(temp[0])
		seconds += t * stringToSec[temp[1]]
	}

	return seconds
}

func matcher(input string) [][]string {
	result := make([][]string, 0, len(input))

	var s scanner.Scanner
	s.Init(strings.NewReader(input))

	tokens := make([][]byte, 0, 16)

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, []byte(s.TokenText()))
	}

	t := []rune("yMwdhmsглМндчмс")

	for i, e := range tokens {
		for _, v := range t {
			if string(e[0]) == string(v) && i > 0 || string(e[0:2]) == string(v) && i > 0 {
				if _, err := strconv.Atoi(string(tokens[i-1])); err == nil {
					// fmt.Printf("%q looks like a number.\n", string(tokens[i-1])+" "+string([]rune(string(tokens[i]))[0]))
					result = append(result, strings.Split(string(tokens[i-1])+" "+string([]rune(string(tokens[i]))[0]), " "))
				}
			}
		}
	}

	return result
}
