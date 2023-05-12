package main

import (
    "fmt"
    "os"
)

func main() {
	input := os.Args[1]
	res, err := ResolveStatusCode(input)
	if err != nil {
		fmt.Printf("unable to resolve status code: %s\n", err.Error())
		return
	}
	fmt.Println(res)
}