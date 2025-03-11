package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	defer writer.Flush()

	for {
		scanner.Scan()
		line := scanner.Text()
		if line == "." {
			break
		}

		chars := strings.Split(line, "")
        valid := true
		var stack []string

		for i := 0; i < len(chars); i++ {
			char := chars[i]
			if char != "(" && char != ")" && char != "[" && char != "]" {
				continue
			}
            stackSize := len(stack)
            switch char {
                case "(", "[":
                    stack = append(stack, char)
                case ")":
                    if stackSize == 0 || stack[stackSize - 1] == "[" {
                        valid = false
                    } else {
                        stack = stack[:len(stack) -1]
                    }
                case "]":
                    if stackSize == 0 || stack[stackSize - 1] == "(" {
                        valid = false
                    } else {
                        stack = stack[:len(stack) -1]
                    }
            }
		}
        if valid && len(stack) == 0 {
            fmt.Fprintln(writer, "yes")
        } else {
            fmt.Fprintln(writer, "no")
        }

		if err := scanner.Err(); err != nil {
			panic(err)
		}
	}
}
