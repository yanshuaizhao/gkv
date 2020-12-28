#### 简介
GKV 是Go语言基于map类型 Key/Value 内存缓存组件 (目前实现比较简单还需不断完善及功能增强)

#### 应用实例
```
package main

import (
	"fmt"
	"github.com/yanshuaizhao/gkv"
	"time"
)

func main() {
	c := gkv.New()
	if _, err := c.Set("key", "hi, golang", time.Second*100); err != nil {
		fmt.Println("set error:", err.Error())
	}
	value, _ := c.Get("key")
	fmt.Println(value)
}
```