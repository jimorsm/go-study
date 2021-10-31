# go语言特性
Go 语言是一个可以编译高效，支持高并发的，面向垃圾回收的全新语言.

- 秒级完成大型程序的单节点编译。
- 依赖管理清晰。
- 不支持继承，程序员无需花费精力定义不同类型之间的关系。
- 支持垃圾回收，支持并发执行，支持多线程通讯。
- 对多核计算机支持友好。

Go 语言特性衍生来源:
![](vx_images/1789631179197.png)

## go语言环境变量
- GOROOT：go的安装目录
- GOPATH
    - src：存放源代码
    - pkg：存放依赖包
    - bin：存放可执行文件
- 其他常用变量
    - GOOS，GOARCH，GOPROXY
    - 国内用户建议设置goproxy：`export GOPROXY=https://goproxy.cn`

## go 基本命令

- `bug`:     start a bug report
- `build`:    compile packages and dependencies
- `clean`:    remove object files and cached files
- `doc`:    show documentation for package or symbol
- `env`:    print Go environment information
- `fix`:    update packages to use new APIs
- `fmt`:    gofmt (reformat) package sources
- `generate`:     generate Go files by processing source
- `get`:    add dependencies to current module and install them
- `install`:    compile and install packages and dependencies
- `list`:    list packages or modules
- `mod`:    module maintenance
- `run`:    compile and run Go program
- `test`:    test packages
- `tool`:    run specified go tool
- `version`: print Go version
- `vet`: report likely mistakes in packages

### go build
- Go 语言不支持动态链接，因此编译时会将所有依赖编译进同一个二进制文件。
- 指定输出目录。
    - `go build –o bin/mybinary .`
- 常用环境变量设置编译操作系统和 CPU 架构。
    - `GOOS=linux GOARCH=amd64 go build`
- 全支持列表。
    - `$GOROOT/src/go/build/syslist.go`
    
### go test
Go 原生自带测试

```go
import "testing"

func TestIncrease(t *testing.T) {
    t.Log("Start testing")
    increase(1, 2)
}
```
`go test ./.... -v` 运行测试
go test命令扫描所有*_test.go为结尾的文件，惯例是将测试代码与正式代码放在同目录，
如 foo.go 的测试代码一般写在 foo_test.go

### go vet
代码静态检查，发现可能的bug或者可以的构造。

- Print-format 错误，检查类型不匹配的print
```go
str := “hello world!”
fmt.Printf("%d\n", str)
```

- Boolen 错误， 检查一直为 true、false 或者冗余的表达式
`fmt.Println(i != 0 || i != 1)`

- Range 循环，比如如下代码主协程会先退出，go routine无法被执行
```go
words := []string{"foo", "bar", "baz"}
for _, word := range words {
    go func() {
       fmt.Println(word).
    }()
}
```

- Unreachable的代码，如 return 之后的代码

- 其他错误，比如变量自赋值，error 检查滞后等
```go
res, err := http.Get("https://www.spreadsheetdb.io/")
defer res.Body.Close()
if err != nil {
    log.Fatal(err)
}
```

