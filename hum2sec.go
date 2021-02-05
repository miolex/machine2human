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
	// if string just is number for example "1332"
	x, err := strconv.Atoi(input)
	if err == nil {
		return x
	}

	var seconds int
	input = strings.TrimSpace(input)
	r := regexp.MustCompile(`(\d+)+\s*([yMwdhmsглМндчмс])`)

	matched := r.FindAll([]byte(input), -1)
	for _, m := range matched {
		temp := strings.Fields(string(m))
		if len(temp) < 2 {
			for i := range stringToSec {
				if strings.Contains(temp[0], i) {
					tmp := string(temp[0])
					tmp = strings.ReplaceAll(tmp, i, " "+i)
					temp = strings.Fields(tmp)
				}
			}
		}

		t, _ := strconv.Atoi(temp[0])
		seconds += t * stringToSec[temp[1]]
	}

	return seconds
}
