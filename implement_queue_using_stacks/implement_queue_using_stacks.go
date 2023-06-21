// https://leetcode.com/problems/implement-queue-using-stacks

package implement_queue_using_stacks

type MyQueue struct {
	list []int
}

func Constructor() MyQueue {
	return MyQueue{
		list: []int{},
	}
}

func (q *MyQueue) Push(x int) {
	q.list = append(q.list, x)
}

func (q *MyQueue) Pop() (firstEl int) {
	firstEl = q.list[0]
	q.list = q.list[1:]
	return
}

func (q *MyQueue) Peek() int {
	return q.list[0]
}

func (q *MyQueue) Empty() bool {
	return len(q.list) == 0
}
