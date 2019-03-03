package QueueArray

/*
 * Date: 2019/03/02
 * Email: xc8801@126.com
 *
 * Push Time Complexity:O(1), Pop Time Complexity:O(1)
 * Space Complexity:O(n)
 */

type QueueArray struct {
	data   []interface{}
	lenght int
}

const QUEUE_DEFAULT_LENGHT = 0

func NewQueueArray() *QueueArray {
	return &QueueArray{
		data: make([]interface{}, QUEUE_DEFAULT_LENGHT),
	}
}

func (queue *QueueArray) Push(data interface{}) {
	queue.data = append(queue.data, data)
	queue.lenght++
}

func (queue *QueueArray) Pop() (interface{}, bool) {
	if queue.Empty() {
		return nil, false
	}

	item := queue.data[queue.lenght-1]

	queue.data = queue.data[:queue.GetSize()-1]
	queue.lenght--

	return item, true
}

func (queue *QueueArray) Empty() bool {
	return queue.GetSize() == 0
}

func (queue *QueueArray) GetSize() int {
	return queue.lenght
}

func (queue *QueueArray) PopTop() (interface{}, bool) {
	if queue.Empty() {
		return nil, false
	}

	item := queue.data[0]

	queue.data = queue.data[1:]
	queue.lenght--

	return item, true
}
