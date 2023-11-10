# SnowFlake

[![Build Status][1]][2] [![codecov][3]][4]

[1]: https://github.com/lxzan/memorycache/workflows/Go%20Test/badge.svg?branch=main
[2]: https://github.com/lxzan/memorycache/actions?query=branch%3Amain
[3]: https://codecov.io/gh/lxzan/memorycache/graph/badge.svg?token=OHD6918OPT
[4]: https://codecov.io/gh/lxzan/memorycache

### Description

The reactivation code contrast uses atomic methods to implement the SnowFlake algorithm, which has extremely fast generation speed and avoids the use of `Mutex` lock method. The code implementation is simple and uses the factory pattern to create multiple objects. Can meet various ID generation needs, while using `CacheTimer` technology to reduce the waste of SnowFlake heap memory.

### Principle

-   200+ years long time support
-   Support 2048 processes(machines) at most
-   Generate 1,000,000 id/s per process

### Advantage

-   Simple and easy to use
-   No third-party dependencies
-   High performance
-   Low memory usage
-   Zero Allocs
-   Zero OP

### Methods

-   [x] `NextID()` uint64 // Generate a new ID

### Example

```go
import (
    "fmt"
    "github.com/lxzan/snowflake"
)

func main()  {
	s := snowflake.NewSnowFlake(1)
	fmt.Println(s.NextID())

	var o = Decode(1662468685074268160)
	fmt.Println(fmt.Sprintf("%v", o))
}

```

### Benchmark

```bash
goos: darwin
goarch: amd64
pkg: github.com/lxzan/snowflake
cpu: Intel(R) Xeon(R) CPU E5-2643 v2 @ 3.50GHz
BenchmarkNextID
BenchmarkNextID-12      195921657                6.105 ns/op           0 B/op          0 allocs/op
```
