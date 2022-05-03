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

