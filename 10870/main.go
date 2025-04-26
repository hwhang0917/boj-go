package main

import (
    "bufio"
    "os"
    "fmt"
)

func fib(n int) int {
    if (n <= 0) {
        return 0
    }
    if (n == 1) {
        return 1
    }
    return fib(n - 1) + fib(n - 2)
}

func main() {
    r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()

    var n int
    fmt.Fscanln(r, &n)
    fmt.Fprintf(w, "%d", fib(n))
}
