# SQL Primary Key Benchmarks

> This is an experiment to test the speed differences between different type of primary key.

## Setup

<p>Machine: Macbook Pro</p>
<p>CPU: M1</p>
<p>Memory: 16GB</p>
<p>OS: macOS Ventura 13.5.1</p>

Record set: **1,000,000** (1M) and **5,000,000** (5M)

## Space Allocation

| Statement                         | Size (MB) in 5M records |
| --------------------------------- | ----------------------: |
| Auto Increment BigInt Primary Key |                   125.2 |
| UUID Primary Key                  |                   506.8 |
| Ordered UUID Primary Key          |                   350.0 |
| Binary UUID Primary Key           |                   352.7 |
| Ordered Binary UUID Primary Key   |                   205.3 |
| Snowflake ID                      |                   202.3 |

## How to run tests?

```bash
go install golang.org/x/perf/cmd/benchstat@latest

# go test -benchmem -run=^$ -bench=^BenchmarkAutoIncrementID -benchtime=5s -count=10 | tee stat1.txt
go test -run=^$ -bench=^BenchmarkAutoIncrementID -benchtime=5s -count=10 | tee stat1.txt
go test -run=^$ -bench=^BenchmarkUUID -benchtime=5s -count=10 | tee stat2.txt
go test -run=^$ -bench=^BenchmarkOrderedUUID -benchtime=5s -count=10 | tee stat3.txt
go test -run=^$ -bench=^BenchmarkBinaryUUID -benchtime=5s -count=10 | tee stat4.txt
go test -run=^$ -bench=^BenchmarkBinaryOrderedUUID -benchtime=5s -count=10 | tee stat5.txt
go test -run=^$ -bench=^BenchmarkSnowflakeID -benchtime=5s -count=10 | tee stat6.txt

benchstat stat1.txt stat2.txt stat3.txt stat4.txt stat5.txt stat6.txt
```

## Benchmarks

```bash
goos: darwin
goarch: arm64
pkg: github.com/si3nloong/sql-pk-benchmark
              │ autoincr_id.txt │                uuid.txt                │           ordered-uuid.txt           │             bin-uuid.txt              │         bin-ordered-uuid.txt         │           snowflake_id.txt           │
              │     sec/op      │    sec/op      vs base                 │    sec/op     vs base                │    sec/op      vs base                │    sec/op     vs base                │    sec/op     vs base                │
Insert-8            913.2µ ± 9%   1934.3µ ± 23%  +111.82% (p=0.000 n=10)   1351.1µ ± 7%  +47.95% (p=0.000 n=10)   1529.1µ ± 20%  +67.45% (p=0.000 n=10)   1118.6µ ± 1%  +22.49% (p=0.000 n=10)   1050.7µ ± 1%  +15.06% (p=0.000 n=10)
FindByID-8          95.32µ ± 1%   101.81µ ±  2%    +6.80% (p=0.000 n=10)    98.90µ ± 3%   +3.75% (p=0.001 n=10)   101.31µ ±  4%   +6.28% (p=0.000 n=10)    99.90µ ± 4%   +4.80% (p=0.001 n=10)   107.83µ ± 4%  +13.12% (p=0.000 n=10)
GetList-8           117.9µ ± 3%    121.5µ ±  1%    +3.04% (p=0.000 n=10)    127.1µ ± 2%   +7.85% (p=0.000 n=10)    117.3µ ±  1%        ~ (p=0.912 n=10)    114.0µ ± 1%   -3.26% (p=0.001 n=10)    118.3µ ± 2%        ~ (p=1.000 n=10)
GetRandomly-8       176.2µ ± 0%    211.8µ ±  2%   +20.19% (p=0.000 n=10)    215.7µ ± 3%  +22.40% (p=0.000 n=10)    181.9µ ±  3%   +3.19% (p=0.002 n=10)    180.6µ ± 3%   +2.50% (p=0.000 n=10)    191.9µ ± 5%   +8.92% (p=0.000 n=10)
geomean             206.2µ         266.8µ         +29.38%                   246.0µ       +19.31%                   239.8µ        +16.27%                   219.0µ        +6.22%                   225.2µ        +9.21%
```

## Reference

- https://100go.co/89-benchmarks/
- https://blog.jiayu.co/2019/05/benchmarking-go-code/
- https://medium.com/@adamszpilewicz/benchmarking-and-performance-comparison-in-go-using-benchstat-3da48bfaeede
