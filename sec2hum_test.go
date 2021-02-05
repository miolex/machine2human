package machine2human

import "testing"

func TestSec2Hum(t *testing.T) {
	var testTable = []struct {
		seconds  int
		expected string
	}{
		{0, "0 секунд"},
		{31536000, "1 год"},
		{3 * 31536000, "3 года"},
		{8 * 31536000, "8 лет"},
		{2592000, "1 месяц"},
		{3 * 2592000, "3 месяца"},
		{8 * 2592000, "8 месяцев"},
		{604800, "1 неделя"},
		{3 * 604800, "3 недели"},
		{8 * 604800, "1 месяц 3 недели 5 дней"},
		{123456789, "3 года 11 месяцев 3 дня 21 час 33 минуты 9 секунд"},
		{86400, "1 день"},
		{3 * 86400, "3 дня"},
		{5 * 86400, "5 дней"},
		{3600, "1 час"},
		{3 * 3600, "3 часа"},
		{12 * 3600, "12 часов"},
		{60, "1 минута"},
		{3 * 60, "3 минуты"},
		{12 * 60, "12 минут"},
		{1, "1 секунда"},
		{3, "3 секунды"},
		{12, "12 секунд"},
	}

	for _, testCase := range testTable {
		actual := Sec2Hum(testCase.seconds)
		if actual != testCase.expected {
			t.Errorf("Sec2Hum(%d): expected `%s` actual `%s`\n", testCase.seconds, testCase.expected, actual)
		}
	}
}
