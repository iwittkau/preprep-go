//go:generate preprep-go -i main.go
package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello from main")
	//#ifdef FEATURE
	feature()
	//#endif
}
