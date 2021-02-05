package machine2human_test

import (
	"testing"

	"github.com/miolex/machine2human"
)

func TestSec2Hum(t *testing.T) {
	var testTable = []struct {
		seconds  int
		expected string
	}{
		{0, "0 секунд"},
		{1, "1 секунда"},
		{-1, "1 секунда"},
		{2, "2 секунды"},
		{5, "5 секунд"},
		{15, "15 секунд"},
		{60, "1 минута"},
		{120, "2 минуты"},
		{300, "5 минут"},
		{3600, "1 час"},
		{7200, "2 часа"},
		{18000, "5 часов"},
		{86400, "1 день"},
		{172800, "2 дня"},
		{432000, "5 дней"},
		{604800, "1 неделя"},
		{1209600, "2 недели"},
		{2592000, "1 Месяц"},
		{5184000, "2 Месяца"},
		{12960000, "5 Месяцев"},
		{31536000, "1 год"},
		{63072000, "2 года"},
		{157680000, "5 лет"},

		{34822861, "1 год 1 Месяц 1 неделя 1 день 1 час 1 минута 1 секунда"},
		{69645722, "2 года 2 Месяца 2 недели 2 дня 2 часа 2 минуты 2 секунды"},
		{174546305, "5 лет 6 Месяцев 2 недели 1 день 5 часов 5 минут 5 секунд"},
	}

	for _, testCase := range testTable {
		actual := machine2human.Sec2Hum(testCase.seconds)
		if actual != testCase.expected {
			t.Errorf("Sec2Hum(%d): expected `%s` actual `%s`\n", testCase.seconds, testCase.expected, actual)
		}
	}
}
