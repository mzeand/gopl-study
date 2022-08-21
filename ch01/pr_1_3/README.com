```
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^(BenchmarkAppend_naive|BenchmarkAppend_join)$ github.com/mzeand/gopl-study/ch01/pr_1_3 -v

goos: darwin
goarch: arm64
pkg: github.com/mzeand/gopl-study/ch01/pr_1_3
BenchmarkAppend_naive
BenchmarkAppend_naive-8         18362730                65.16 ns/op           96 B/op          2 allocs/op
BenchmarkAppend_join
BenchmarkAppend_join-8          33879562                34.85 ns/op           24 B/op          1 allocs/op
PASS
ok      github.com/mzeand/gopl-study/ch01/pr_1_3        3.539s

```
