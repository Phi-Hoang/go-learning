package main

import (
	"fmt"

	"github.com/phihdn/go-learning/lodash"
)

func main() {
	stringS := []interface{}{"abc", "123", "2", "4356", "pasd"}
	aString := "1234567890"
	fmt.Println(stringS)
	lodash.Reverse(stringS)
	fmt.Println(stringS)
	lodash.Reverse(aString)
	fmt.Println(aString)
}
