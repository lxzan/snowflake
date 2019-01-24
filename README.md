# snowflake
snowflake algorithm

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
