package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	strReverse := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(strReverse)
}
