package queue

// 用两个队列mock一个栈
// 比如，初始化一个栈为[1,2,3,4,5]
// 第一步
// tmp[1]
// tmp和data交换 tmp[] data[1]
// 第二步
// tmp[2], data[1]弹出入队tmp => tmp[2,1], data[]
// tmp和data交换 data[2,1] tmp[]
// 第三步
// tmp[3], data[2,1]弹出入队tmp => tmp[3,2,1], data[]
// tmp和data交换 data[3,2,1] tmp[]
// ...

type MockShed struct {
	tmpQueue  Queue
	dataQueue Queue
}

func (m *MockShed) Init(data []int) {
	m.tmpQueue, m.dataQueue = Queue{}, Queue{}
	for _, v := range data {
		m.Push(v)
	}
}

func (m *MockShed) Push(element int) {
	m.tmpQueue.Push(element)
	if m.dataQueue.Length > 0 {
		for i, l := 0, m.dataQueue.Length; i <= l-1; i++ {
			ele, _ := m.dataQueue.Pop()
			m.tmpQueue.Push(ele)
		}
	}
	m.tmpQueue, m.dataQueue = m.dataQueue, m.tmpQueue
}

func (m *MockShed) Pop() (int, error) {
	return m.dataQueue.Pop()
}
