package main

func main(){

}

type MyQueue struct {
    queue []int
}

func Constructor() MyQueue {
    return MyQueue{
        queue: []int{},
    }
}

func (m *MyQueue) Push(x int)  {
    m.queue = append(m.queue, x)
}

func (m *MyQueue) Pop() int {
    tmp := m.queue[0]
    m.queue = m.queue[1:]
    return tmp
}

func (m *MyQueue) Peek() int {
    return m.queue[0]
}

func (m *MyQueue) Empty() bool {
    return len(m.queue) == 0
}
