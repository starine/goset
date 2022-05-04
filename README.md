## goset

golang没有内置的set，因此使用map[interface{}]struct{}自定义一个，方便使用。

## 使用方式

```go
package main

import (
  "fmt"
  mapset "github.com/starine/goset"
)

func main() {
   s := mapset.NewSet(8, "beijing")
   arr := [5]int {1,2,3,4,5}
   s.Add(arr)
   s.Add(false)
   fmt.Println(s.Contains(false)) //true
   fmt.Println(s.Contains(8)) //true
   fmt.Println(s.Contains(arr)) //true

}
```

## NewSet 调用Add函数添加一堆元素到set中
```go
func NewSet(items ...interface{}) *set {
	s := &set{}
	s.m = make(map[interface{}]exists)
	s.Add(items...)
	return s
}
```

## Add 将元素列表插入到set中
```go
func (s *set) Add(items ...interface{}) {
	for _, item := range items {
		s.m[item] = exists{}
	}
}
```

## Remove 删除指定元素
```go
func (s *set) Remove(item interface{}) {
	delete(s.m, item)
}
```

## Contains 查询指定元素是否存在
```go
func (s *set) Contains(item interface{}) bool {
	_, ok := s.m[item]
	return ok
}
```

## Size 查询集合大小
```go
func (s *set) Size() int {
	return len(s.m)
}
```
