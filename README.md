# gedis

gedis = redis + fatih/pool

## hello world

```go
package main

import (
	"gedis"
	"fmt"
)

func main() {
	c, err := gedis.Dial("127.0.0.1", 6379)
	if err != nil {
		panic(err)
	}

	messages := [][]string{
		{"SET", "key", "1"},
		{"FUCK"},
		{"INCR", "key"},
		{"GET", "key"},
		{"RPUSH", "list", "foo"},
		{"RPUSH", "list", "bar"},
		{"RPUSH", "list", "Hello"},
		{"RPUSH", "list", "World"},
		{"LRANGE", "list", "0", "3"},
	}
	for _, message := range messages {
		cmd, err := c.Do(message...)
		if err != nil {
			panic(err)
		}
		fmt.Println(cmd)
	}

	c.Close()
}
```

## reference

- [golang中bufio包的用法](https://studygolang.com/articles/4367)
- [golang bufio包中的Write方法分析](https://blog.csdn.net/benben_2015/article/details/80614230)