# snowflake
snowflake algorithm

### Feature
- 200+ years long time support
- support 1024 processes(machines) at most
- generate 1,000,000 id/s per process

### Example
```go
import "github.com/lxzan/snowflake"

func main()  {
	snowflake.Initialize(1)
	Println(snowflake.NextID())
	
	var o = Decode(1662468685074268160)
	println(fmt.Sprintf("%v",o))
}

// output 1662468685074268160

// &{1548294616 1 0}
```

### Benchmark
- goos: darwin
- goarch: amd64
- pkg: snowflake
- BenchmarkKey-4   	100000000	        11.0 ns/op
- PASS
