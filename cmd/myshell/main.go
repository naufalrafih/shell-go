package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
	// Debug
	fmt.Fprint(os.Stdout, "$ ")

	// Wait for user input
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Printf("%s: command not found\n", strings.TrimRight(input, "\n"))
}
