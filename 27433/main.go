package main

import (
	"bufio"
	"fmt"
	"os"
)

func factorial(n int) int {
    if (n == 0) {
        return 1
    }
    return n * factorial(n - 1)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var n int

    fmt.Fscanln(reader, &n)

    fmt.Fprintf(writer, "%d", factorial(n))
}
