# expr

Package expr implements parsing of basic infix arithmetic expressions.


Installation
------------

The recommended way to install expr

```
go get github.com/stianwa/expr
```

Examples
--------

```go

package main
 
import (
       "fmt"
       "github.com/stianwa/expr"
       "log"
)

func main() {
       f, err := expr.Calc("4+3*5/(9-6)")
       if err != nil {
               log.Fatalf("failed to parse expression: %v\n", err)             
       }
       fmt.Printf("result: %f\n", f)
}
```

State
-------
The expr package is currently under development. Do not use for production.


License
-------

MIT, see [LICENSE.md](LICENSE.md)
