package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)


func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var T int
    fmt.Fscanln(reader, &T)

    for i := 0; i < T; i++ {
        var rawPS string
        fmt.Fscanln(reader, &rawPS)
        PS := strings.Split(rawPS, "")

        validator := 0
        for j := 0; j < len(PS); j++ {
            ps := PS[j]
            switch ps {
                case "(":
                    validator++
                    break
                case ")":
                    validator--
                    break
                default:
                    panic("Unknown parenthesis")
            }
            if validator < 0 { break }
        }
        if validator == 0 {
            fmt.Fprintln(writer, "YES")
        } else {
            fmt.Fprintln(writer, "NO")
        }
    }
}
