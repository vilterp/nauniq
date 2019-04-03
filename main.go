package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	seen := map[string]struct{}{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if _, ok := seen[line]; !ok {
			fmt.Println(line)
			seen[line] = struct{}{}
		}
	}
}
