package main

import (
	"fmt"
	useful "go-cmd-sample"
)

func main() {
	fmt.Println("go-cmd-sample - main running")

	usefulResult := useful.MyUsefulFunction()
	fmt.Printf("usefulResult: %s\n", usefulResult)

	fmt.Println("go-cmd-sample - main done")
}
