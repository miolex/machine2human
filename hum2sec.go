package machine2human

import (
	"regexp"
	"strconv"
	"strings"
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
	var seconds int

	input = strings.TrimSpace(input)
	r := regexp.MustCompile(`(\d+)+\s*([yMwdhmsглМндчмс])`)
	matched := r.FindAll([]byte(input), -1)

	for _, m := range matched {
		temp := strings.Fields(string(m))
		t, _ := strconv.Atoi(temp[0])
		seconds += t * stringToSec[temp[1]]
	}

	return seconds
}
