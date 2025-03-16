package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var N, K int
    fmt.Fscan(reader, &N)
    fmt.Fscan(reader, &K)

    var arr []int
    for i := 1; i <= N; i++ { arr = append(arr, i) }

    fmt.Fprintf(writer, "<")
    for len(arr) != 0 {
        i := (K - 1) % len(arr)
        if len(arr) == 1 {
            fmt.Fprintf(writer, "%d", arr[i])
        } else {
            fmt.Fprintf(writer, "%d, ", arr[i])
        }
        arr = append(arr[i + 1:], arr[:i]...)
    }
    fmt.Fprintf(writer, ">")
}
