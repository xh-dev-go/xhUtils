# xhUtils

## install
```shell
go get github.com/xh-dev-go/xhUtils

```


## Demo

Empty String
```go
import fmt
fmt.Println(common.StringEmpty) // avoiding defining "" string
```

Simplified flag variable
```go
strVar := flagString.New("string", "usage of string").BindCmd()
intVar := flagInt.New("int_variable", "usage of int").BindCmd()
flag.Parse()
```