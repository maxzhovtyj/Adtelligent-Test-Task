package main

import (
	"fmt"
	"github.com/maxzhovtyj/Adtelligent-Test-Task/task2/concat"
)

func main() {
	x := []string{"10", "11", "sl", "127"}

	str := concat.Concat(x)

	fmt.Println(str)
}
