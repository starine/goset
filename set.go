/*
* @Author: starine
* @Date:   2022/5/3 19:47
 */

package goset

//使用空结构体大小0，不占内存
type exist struct {}

//使用空接口interface{},支持任意类型元素
type set struct {
	m map[interface{}]exist
}

// NewSet 调用Add函数添加一堆元素到set中
func NewSet(items ...interface{}) *set {
	s := &set{}
	s.m = make(map[interface{}]exist)
	s.Add(items...)
	return s
}

// Add 将元素列表插入到set中
func (s *set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = exist{}
	}
}

// Remove 删除指定元素
func (s *set) Remove(item interface{}) {
	delete(s.m, item)
}

// Contains 查询指定元素是否存在
func (s *set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}

// Size 查询集合大小
func (s *set) Size() int {
	return len(s.m)
}