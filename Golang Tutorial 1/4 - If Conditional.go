package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	iAge := 8

	if (iAge >= 1) && (iAge <= 18) {
		pl("Young Birthday")
	} else if (iAge == 21) || (iAge == 50) {
		pl("Important Birthday")
	} else if iAge >= 65 {
		pl("Old Birthday")
	} else {
		pl("Not Important")
	}
}
