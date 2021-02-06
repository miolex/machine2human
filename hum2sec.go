package machine2human

import (
	"math"
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
		return int(math.Abs(float64(x)))
	}

	var seconds int
	input = strings.TrimSpace(input)
	parsed := parser(input)

	for _, parsingCase := range parsed {

		t, _ := strconv.Atoi(parsingCase[0])
		seconds += t * stringToSec[parsingCase[1]]
	}

	return seconds
}

func parser(input string) [][]string {
	result := make([][]string, 0)
	tokens := make([][]byte, 0)
	tokenizer(&tokens, input)
	humanTokens := []rune("yMwdhmsглМндчмс")

	for index, value := range tokens {
		for _, v := range humanTokens {
			if string(value[0]) == string(v) && index > 0 || string(value[0:2]) == string(v) && index > 0 {
				if _, err := strconv.Atoi(string(tokens[index-1])); err == nil {

					result = append(result, []string{string(tokens[index-1]), string([]rune(string(tokens[index]))[0])})
				}
			}
		}
	}

	return result
}

func tokenizer(tokens *[][]byte, input string) {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))

	for token := s.Scan(); token != scanner.EOF; token = s.Scan() {
		*tokens = append(*tokens, []byte(s.TokenText()))
	}
}
