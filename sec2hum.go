// Package machine2human help you convert seconds to readble string and vice verca
package machine2human

import (
	"fmt"
	"math"
	"strings"
)

const (
	second = 1
	minute = 60
	hour   = 3600
	day    = 86400
	week   = 604800
	month  = 2592000
	year   = 31536000
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

// Sec2Hum implement convertation secons to readable string
func Sec2Hum(seconds int) string {
	seconds = int(math.Abs(float64(seconds)))
	var readebleString string
	var temp int

	if seconds == 0 {
		return fmt.Sprintf("0 %s", runeToRuString['s'][0])
	}

	// years
	temp = seconds / year
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['y'][yarusskiy(temp)])
		readebleString += " "
		seconds -= temp * year
		temp = 0
	}

	// months
	temp = seconds / month
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['M'][yarusskiy(temp)])
		readebleString += " "
		seconds -= temp * month
		temp = 0
	}

	// weeks
	temp = seconds / week
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['w'][yarusskiy(temp)])
		readebleString += " "
		seconds -= temp * week
		temp = 0
	}

	// days
	temp = seconds / day
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['d'][yarusskiy(temp)])
		readebleString += " "
		seconds -= temp * day
		temp = 0
	}

	// hours
	temp = seconds / hour
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['h'][yarusskiy(temp)])
		readebleString += " "
		seconds -= temp * hour
		temp = 0
	}

	// minutes
	temp = seconds / minute
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['m'][yarusskiy(temp)])
		readebleString += " "
		seconds -= temp * minute
		temp = 0
	}

	// seconds
	temp = seconds / second
	if temp > 0 {
		readebleString += fmt.Sprintf("%d %s", temp, runeToRuString['s'][yarusskiy(temp)])
		readebleString += " "
		seconds -= temp * second
		temp = 0
	}

	readebleString = strings.TrimSpace(readebleString)
	return readebleString
}

func yarusskiy(digit int) int {
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

	return 0
}
