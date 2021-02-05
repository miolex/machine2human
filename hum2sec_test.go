package machine2human

import "testing"

func TestHum2Sec(t *testing.T) {
	var testTable = []struct {
		input    string
		expected int
	}{
		{"1 y", 31536000},
		{"1 г", 31536000},
		{"1 year", 31536000},
		{"1 год", 31536000},
		{"3 года", 31536000 * 3},
		{"5 лет", 31536000 * 5},
		{"2 дня 5 минут", 86400*2 + 60*5},
		{"2 недели", 604800 * 2},
		{"43 секунды", 43},
		{"28 часов", 3600 * 28},
		{"12 минут", 60 * 12},
		{"30 дней", 86400 * 30},
	}

	for _, testCase := range testTable {
		actual := Hum2Sec(testCase.input)

		if actual != testCase.expected {
			t.Errorf("Sec2Hum(%s): expected `%d` actual `%d`\n", testCase.input, testCase.expected, actual)
		}
	}
}
