package machine2human_test

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/miolex/machine2human"
)

var (
	testString = genString(1000000)
	testInt    = randInt(100000, 999999999)
)

func randInt(min, max int) int {
	rand.Seed(time.Now().Unix())
	rand := rand.Intn(max-min) + min
	return rand
}

func genString(n int) string {
	rand.Seed(time.Now().Unix())
	x := []rune("yлMМwнdдhчmмsс")
	b := make([]string, 2*n)
	for i := range b {
		if i%2 == 0 {
			ii := rand.Int31n(100-1) + 1

			b[i] = strconv.Itoa(int(ii))
			continue
		}
		b[i] = string(x[rand.Int31n(int32(len(x)))])
	}
	return strings.Join(b, "")
}
func BenchmarkSec2Hum(b *testing.B) {

	for i := 0; i < b.N; i++ {
		machine2human.Sec2Hum(testInt)
	}
}

func BenchmarkHum2Sec(b *testing.B) {

	for i := 0; i < b.N; i++ {
		machine2human.Hum2Sec(testString)
	}
}
