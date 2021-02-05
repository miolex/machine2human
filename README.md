# machine2human
Package machine2human help you convert seconds to readble string and vice verca.

## Installation
1. The first need Go installed.
```golang
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
	fmt.Println(machine2human.Sec2Hum(13453))
}
```
```console
$ go run example.go
3 часа 44 минуты 13 секунд
```