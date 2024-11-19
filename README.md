# cberrors
___
cberrors实现了golang内置error接口，强调易用性和可调试性

## Features
___
- [X] 错误信息的格式化创建
- [X] 自动附加堆栈信息
- [X] 使用cberrors封装系统panic

## Example
___

### 创建一个带错误信息和堆栈信息的error
```golang
package main

import (
	"fmt"
	"github.com/caibo86/cberrors"
)

func main() {
	s := "hello world"
	err := cberrors.New("这是的一个新的错误:%s", s)
	fmt.Println(err)
}

```

Output:
```shell
这是的一个新的错误:hello world
stack trace:
        file = /home/cb/go/src/cberrors/example/example.go, line = 17
        file = /usr/local/go/src/runtime/proc.go, line = 267
        file = /usr/local/go/src/runtime/asm_amd64.s, line = 1650
```

### 包装一个已有的系统错误
```golang
package main

import (
	"fmt"
	"github.com/caibo86/cberrors"
)

func main() {
	err1 := fmt.Errorf("原生golang的错误类型")
	err = cberrors.Wrap(err1)
	fmt.Println(err)
}

```

Output:
```shell
原生golang的错误类型
stack trace:
        file = /home/cb/go/src/cberrors/example/example.go, line = 26
        file = /home/cb/go/src/cberrors/example/example.go, line = 16
        file = /usr/local/go/src/runtime/proc.go, line = 267
        file = /usr/local/go/src/runtime/asm_amd64.s, line = 1650
```

### 输入格式化信息并panic
```golang
package main

import (
	"fmt"
	"github.com/caibo86/cberrors"
)

func main() {
	s := "hello world"
	cberrors.Panic("这是一个panic错误:%s", s)
}
    
```

Output:
```shell
panic: 这是一个panic错误:hello world
stack trace:
        file = /home/cb/go/src/cberrors/errors.go, line = 81
        file = /home/cb/go/src/cberrors/example/example.go, line = 33
        file = /home/cb/go/src/cberrors/example/example.go, line = 16
        file = /usr/local/go/src/runtime/proc.go, line = 267
        file = /usr/local/go/src/runtime/asm_amd64.s, line = 1650
```

### 包装一个已有的系统错误并panic
```golang
package main

import (
	"fmt"
	"github.com/caibo86/cberrors"
)

func main() {
	s := "hello world"
	err := fmt.Errorf("这是一个panic错误:%s", s)
	cberrors.PanicWrap(err)
}

```

Output:
```shell
panic: 这是一个panic错误:hello world
stack trace:
        file = /home/cb/go/src/cberrors/errors.go, line = 86
        file = /home/cb/go/src/cberrors/example/example.go, line = 40
        file = /home/cb/go/src/cberrors/example/example.go, line = 16
        file = /usr/local/go/src/runtime/proc.go, line = 267
        file = /usr/local/go/src/runtime/asm_amd64.s, line = 1650
```