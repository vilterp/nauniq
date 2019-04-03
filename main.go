package main

import (
	"fmt"
	"os"
	"bufio"
)

func main() {
	seen := map[string]struct{}{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := seen[line]; !ok {
			fmt.Print(line)
			seen[line] = struct{}{}
		}
	}
}
