package queue

type MtQueue struct {
	data   []int
	length int
	cap    int
	head   int
	rear   int
}

func (m *MtQueue) Init(data []int, cap int) {
	//
	if len(data) > 0 {
		m.data, m.length, m.head, m.rear = data, len(data), 0, len(data)-1
	}
}
