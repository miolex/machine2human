// Package machine2human help you convert seconds to readble string and vice verca
package machine2human

import (
	"fmt"
	"strings"
)

var runeToRuString = map[rune][]string{
	's': {"секунд", "секунда", "секунды"},
	'm': {"минут", "минута", "минуты"},
	'h': {"часов", "час", "часа"},
	'd': {"дней", "день", "дня"},
	'w': {"недель", "неделя", "недели"},
	'M': {"Месяцев", "Месяц", "Месяца"},
	'y': {"лет", "год", "года"},
}

const (
	second = 1
	minute = 60
	hour   = 3600
	day    = 86400
	week   = 604800
	month  = 2592000
	year   = 31536000
)

// Sec2Hum implement convertation secons to readable string
func Sec2Hum(seconds int64) string {
	var readebleString string
	var temp int64

	if seconds == 0 {
		return fmt.Sprintf("0 %s", runeToRuString['s'][0])
	}

	// years
	temp = int64(seconds / year)
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['y'][yarusskiy(temp)])
		readebleString += " "
		seconds -= int64(temp * year)
		temp = 0
	}

	// months
	temp = int64(seconds / month)
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['M'][yarusskiy(temp)])
		readebleString += " "
		seconds -= int64(temp * month)
		temp = 0
	}

	// weeks
	temp = int64(seconds / week)
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['w'][yarusskiy(temp)])
		readebleString += " "
		seconds -= int64(temp * week)
		temp = 0
	}

	// days
	temp = int64(seconds / day)
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['d'][yarusskiy(temp)])
		readebleString += " "
		seconds -= int64(temp * day)
		temp = 0
	}

	// hours
	temp = int64(seconds / hour)
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['h'][yarusskiy(temp)])
		readebleString += " "
		seconds -= int64(temp * hour)
		temp = 0
	}

	// minutes
	temp = int64(seconds / minute)
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['m'][yarusskiy(temp)])
		readebleString += " "
		seconds -= int64(temp * minute)
		temp = 0
	}

	// seconds
	temp = int64(seconds / second)
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['s'][yarusskiy(temp)])
		readebleString += " "
		seconds -= int64(temp * second)
		temp = 0
	}

	readebleString = strings.Trim(readebleString, " ")
	return readebleString
}

func yarusskiy(digit int64) int {
	tmp := digit % 100
	if 11 <= tmp && tmp <= 19 {
		return 0
	}
	tmp = digit % 10
	if tmp == 1 {
		return 1
	}
	if 2 <= tmp && tmp <= 4 {
		return 2
	}
	if tmp == 0 || (5 <= tmp && tmp <= 9) {
		return 0
	}
	return 0
}
