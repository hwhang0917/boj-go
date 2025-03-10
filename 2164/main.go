package main

import (
	"bufio"
	"fmt"
	"os"
)

type Node struct {
    data int
    next *Node
    prev *Node
}
type Deque struct {
    front *Node
    rear *Node
    size int
}
func CreateNode(data int) Node { return Node{ data: data } }
func CreateDeque() Deque { return Deque{} }
func (d *Deque) InsertFront(data int) {
    node := CreateNode(data)
    if d.front == nil {
        d.front = &node
        d.rear = d.front
    } else {
        node.next = d.front
        d.front.prev = &node
        d.front = &node
    }
    d.size++
}
func (d *Deque) InsertRear(data int) {
    node := CreateNode(data)
    if d.rear == nil {
        d.front = &node
        d.rear = d.front
    } else {
        node.prev = d.rear
        d.rear.next = &node
        d.rear = &node
    }
    d.size++
}
func (d *Deque) Shift() int {
    temp := d.front
    if temp == nil { return -1 }
    d.front = d.front.next
    if d.front == nil {
        d.rear = nil
    } else {
        d.front.prev = nil
    }
    d.size--
    return temp.data
}
func (d *Deque) Pop() int {
    temp := d.rear
    if temp == nil { return -1 }
    d.rear = d.rear.prev
    if d.rear == nil {
        d.front = nil
    } else {
        d.rear.next = nil
    }
    d.size--
    return temp.data
}
func (d *Deque) IsEmpty() bool {
    return d.front == nil
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    writer := bufio.NewWriter(os.Stdout)
    defer writer.Flush()

    var N int
    fmt.Fscanln(reader, &N)

    deque := CreateDeque()
    for i := 1; i <= N; i++ {
        deque.InsertRear(i)
    }

    for deque.size > 1 {
        deque.Shift()
        if deque.size > 1 {
            deque.InsertRear(deque.Shift())
        }
    }
    fmt.Fprintln(writer, deque.front.data)
}
