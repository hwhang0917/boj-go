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
func (s *Stack) IsEmpty() int {
    if s.top == nil { return 1 }
    return 0
}
func (s *Stack) GetTop() int {
    if s.top == nil { return -1 }
    return s.top.data
}


func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var N int
    fmt.Fscanln(reader, &N)

    stack := CreateStack()
    for i := 0; i < N; i++ {
        var instruction, paramter int
        fmt.Fscanln(reader, &instruction, &paramter)

        switch instruction {
            case 1:
                stack.Push(paramter)
                break
            case 2:
                fmt.Fprintln(writer, stack.Pop())
                break;
            case 3:
                fmt.Fprintln(writer, stack.size)
                break;
            case 4:
                fmt.Fprintln(writer, stack.IsEmpty())
                break;
            case 5:
                fmt.Fprintln(writer, stack.GetTop())
                break;
            default:
                panic("Unknown command")
        }
    }
}
