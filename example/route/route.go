package main

import (
	"fmt"
	"strings"
)

func main() {
	var uri string = "column/article/419293"
	segments := strings.SplitN(uri, "/", 2)
	fmt.Println(segments)
}
