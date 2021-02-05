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
		{2592000, "1 Месяц"},
		{3 * 2592000, "3 Месяца"},
		{8 * 2592000, "8 Месяцев"},
		{604800, "1 неделя"},
		{3 * 604800, "3 недели"},
		{8 * 604800, "1 Месяц 3 недели 5 дней"},
		{123456789, "3 года 11 Месяцев 3 дня 21 час 33 минуты 9 секунд"},
	}

	for _, testCase := range testTable {
		actual := Sec2Hum(int64(testCase.seconds))
		if actual != testCase.expected {
			t.Errorf("Sec2Hum(%d): expected `%s` actual `%s`\n", testCase.seconds, testCase.expected, actual)
		}
	}
}
