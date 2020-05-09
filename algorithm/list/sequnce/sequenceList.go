package sequnce

//线性表-顺序表的实现

//SequenceList 数据结构
type SequenceList struct {
	Len      int
	Capacity int
	Data     []interface{}
}

//NewSequenceList 初始化
func NewSequenceList(capacity int) *SequenceList {
	list := new(SequenceList)
	list.Len = 0
	list.Capacity = capacity

	list.Data = make([]interface{}, capacity)
	return list
}

//IsFull 判断是否已满
func (s *SequenceList) IsFull() bool {
	return s.Len == s.Capacity
}

//IsEmpty 判断是否为空
func (s *SequenceList) IsEmpty() bool {
	return s.Len == 0
}

//Clear 将线性表置为空表
func (s *SequenceList) Clear() {
	s.Len = 0
	s.Data = nil
}

//Insert 插入元素
func (s *SequenceList) Insert(index int, elem interface{}) bool {
	//判断是否满足条件
	if s.IsFull() || index < 0 || index >= s.Capacity {
		return false
	}

	//插入数据
	s.Data = append(s.Data[:index], append([]interface{}{elem}, s.Data[index:]...)...)
	s.Len++
	return true
}

//Delete 删除元素
func (s *SequenceList) Delete(index int) bool {
	//判断是否满足条件
	if s.IsEmpty() {
		return false
	}
	//删除元素
	s.Data = append(s.Data[:index], s.Data[index+1:]...)
	s.Len--
	return true
}

//Get 获取索引的值
func (s *SequenceList) Get(index int) (interface{}, bool) {
	//判断是否满足条件
	if s.IsEmpty() || index > s.Len {
		return nil, false
	}
	return s.Data[index], true
}

//Append 追加数据
func (s *SequenceList) Append(elem interface{}) bool {
	//判断是否满足条件
	if s.IsFull() || s.Len >= s.Capacity {
		return false
	}
	//插入数据
	s.Data = append(s.Data, elem)
	s.Len++
	return true
}
