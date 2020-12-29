package main

import (
	"fmt"

	"./greeting"
)

func main() {
	var gstr string
	gstr = greeting.Do()
	fmt.Println(gstr)
}
