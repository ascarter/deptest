package main

import (
	"fmt"

	"github.com/ascarter/deptest/lower"
	"github.com/ascarter/deptest/uniqueid"
	"github.com/ascarter/deptest/upper"
)

func main() {
	title := upper.Upper("app1")
	description := lower.Lower("A TEST APP1")
	id := uniqueid.UUID()

	fmt.Println(title)
	fmt.Println(description)
	fmt.Println(id)
}
