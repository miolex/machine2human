# machine2human
Package machine2human help you convert seconds to readble string and vice versa.

## Installation
1. The first need Go installed.
```console
go get -u github.com/miolex/machine2human
```
2. Import it in your code
```golang
import "github.com/miolex/machine2human"
```

# Quick start
```golang
package main

import (
	"fmt"

	"github.com/miolex/machine2human"
)

func main() {
	fmt.Println(machine2human.Sec2Hum(228))
	fmt.Print(machine2human.Hum2Sec("12 минут"))
}
```
```console
$ go run example.go
3 минуты 48 секунд
720
```
