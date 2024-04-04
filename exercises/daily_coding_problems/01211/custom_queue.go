package dailycodingproblems

type Queue struct {
	arr  []int
	size int
}

type QueueIntf interface {
	init() *Queue
	isEmpty() bool

	append(data int)
	popLeft() int
	popRight() int
}

func newQueue() *Queue {
	return new(Queue).initQueue()
}

func (queue *Queue) initQueue() *Queue {
	return &Queue{
		arr:  make([]int, 0),
		size: 0,
	}
}

func (queue *Queue) isEmpty() bool {
	return queue.size == 0
}

func (queue *Queue) append(data int) {
	queue.arr = append(queue.arr, data)
	queue.size++
}

func (queue *Queue) popLeft() int {
	top := queue.arr[0]
	queue.arr = queue.arr[1:]
	queue.size--

	return top
}

func (queue *Queue) popRight() int {
	top := queue.arr[queue.size-1]
	queue.arr = queue.arr[:queue.size-1]
	queue.size--

	return top
}
