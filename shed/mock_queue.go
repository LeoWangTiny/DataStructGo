package shed

// 思路: 无论是两个栈实现一个队列还是两个队列实现一个栈,这个问题的本质其实就是汉诺塔问题的变种
// 核心思想就是通过一个空栈或是空队列进行数据顺序的切换，然后在另一个栈或是队列进行数据的合并
// 两个栈实现一个队列的思路
// 演示
// input: [1,2,3,4,5]
//
// line1(Push 1)----------------------------
// 将第一个元素数据栈中
// shedData:[1] shedNil:[]
//
// line2(Push 2)----------------------------
// 将数据栈中的元素依序弹出，压入空栈中进行数据反转
// shedData:[] shedNil: [1]
// 将新增数据压到数据栈中
// shedData:[2] shedNil: [1]
// 将临时栈中的数据弹出，压入数据栈中
// shedData:[2,1] shedNil: []
//
// line3(Push 3)----------------------------
// 将数据栈中的元素依序弹出，压入空栈中进行数据反转
// shedData:[] shedNil[1,2]
// 将新增数据压到数据栈中
// shedData:[3] shedNil[1,2]
// 将临时栈中的数据弹出，压入数据栈中
// shedData:[3,2,1] shedNil: []
// ..........

type MockQueue struct {
	emptyShed Shed
	dataShed  Shed
}

func (m *MockQueue) Init(data []int) {
	m.emptyShed = Shed{}
	m.dataShed = Shed{}
	for _, v := range data {
		m.Push(v)
	}
}

func (m *MockQueue) Push(element int) {
	if m.dataShed.Length != 0 {
		// 先将数据栈的数据反压到临时栈中
		for i, l := 0, m.dataShed.Length; i <= l-1; i++ {
			e, _ := m.dataShed.Pop()
			m.emptyShed.Push(e)
		}
	}
	m.dataShed.Push(element)
	if m.emptyShed.Length != 0 {
		for i, l := 0, m.emptyShed.Length; i <= l-1; i++ {
			e, _ := m.emptyShed.Pop()
			m.dataShed.Push(e)
		}
	}
}

func (m *MockQueue) Pop() (int, error) {
	return m.dataShed.Pop()
}
