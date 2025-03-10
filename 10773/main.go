package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
    data int
    next *Node
}
type Stack struct {
    top *Node
    size int
}
func CreateNode(data int) Node { return Node{ data:data } }
func CreateStack() Stack { return Stack{} }
func (s *Stack) Pop() int {
    if s.top == nil { return -1 }
    temp := s.top
    s.top = s.top.next
    s.size--
    return temp.data
}
func (s *Stack) Push(data int) {
    node := CreateNode(data)
    if s.top != nil { node.next = s.top }
    s.top = &node
    s.size++
}
func (s *Stack) Sum() int {
    sum := 0
    node := s.top
    for node != nil {
        sum += node.data
        node = node.next
    }
    return sum
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var K int
    fmt.Fscanln(reader, &K)

    stack := CreateStack()
    for i := 0; i < K; i++ {
        var num int
        fmt.Fscanln(reader, &num)
        if num == 0 {
            stack.Pop()
        } else {
            stack.Push(num)
        }
    }

    fmt.Fprintln(writer, stack.Sum())
}
