package main

import (
    "bufio"
    "fmt"
    "os"
)

func recurse(s string, l int, r int, rCount *int) int {
    *rCount++
    if l >= r {
        return 1
    } else if s[l] != s[r] {
        return 0
    }
    return recurse(s, l + 1, r - 1, rCount)
}

func isPalindrome(s string, rCount *int) int {
    return recurse(s, 0, len(s) - 1, rCount)
}

func main() {
    r := bufio.NewReader(os.Stdin)
    w := bufio.NewWriter(os.Stdout)
    defer w.Flush()

    var T int
    fmt.Fscanln(r, &T)
    cases := make([]string, T)
    for i := 0; i < T; i++ {
        fmt.Fscanln(r, &cases[i])
    }

    for i := 0; i < T; i++ {
        rCount := 0
        result := isPalindrome(cases[i], &rCount)
        fmt.Fprintf(w, "%d %d\n", result, rCount)
    }
}
