package main

import (
	"flag"
	"fmt"
)

func main() {
	flags, err := newFlags()
	if err != nil {
		flag.Usage()
		fmt.Println()
		panic(err)
	}

	fmt.Println(flags.gaussianBlurDeviation)
	fmt.Println(flags.strength)
}
