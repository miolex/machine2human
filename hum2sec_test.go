package machine2human

import "testing"

func TestHum2Sec(t *testing.T) {
	var testTable = []struct {
		input    string
		expected int
	}{
		{"", 0},
		{"0", 0},
		{"100", 100},
		{"0 секунд", 0},
		{"1 секунда", 1},
		{"1 sec", 1},
		{"1с", 1},
		{"1s", 1},
		{"1 минута", 60},
		{"1 mic", 60},
		{"1м", 60},
		{"1m", 60},
		{"1 час", 3600},
		{"1 hour", 3600},
		{"1ч", 3600},
		{"1h", 3600},
		{"1 день", 86400},
		{"1 day", 86400},
		{"1д", 86400},
		{"1d", 86400},
		{"1 неделя", 604800},
		{"1 week", 604800},
		{"1н", 604800},
		{"1w", 604800},
		{"1 неделя 1 день 1 час 1 минута 1 секунда", 694861},
		{"2 неделя 2 день 2 час 2 минута 2 секунды", 1389722},
		{"5 неделя 5 день 5 час 5 минута 5 секунда", 3474305},
	}

	for _, testCase := range testTable {
		actual := Hum2Sec(testCase.input)

		if actual != testCase.expected {
			t.Errorf("Sec2Hum(%s): expected `%d` actual `%d`\n", testCase.input, testCase.expected, actual)
		}
	}
}
