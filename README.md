Work with Roman numerals in Go.

```go
package main

import (
	"fmt"
	"github.com/mndrix/roman.go"
)

func main() {
	seven, _ := roman.Encode(7)
	fmt.Printf("Final Fantasy %s\n" , seven)

	two, _ := roman.Decode("II")
	fmt.Printf("1+1=%d\n" , two)

	if roman.IsValid("IV") {
		fmt.Println("IV is a valid Roman numeral")
	}
}
```
