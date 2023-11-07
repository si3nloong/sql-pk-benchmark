# SQL Primary Key Benchmarks

> This is an experiment to test the speed differences between different type of primary key.

## Setup

<p>Machine: Macbook Pro</p>
<p>CPU: M1</p>
<p>Memory: 16GB</p>
<p>OS: macOS Ventura 13.5.1</p>

Record set: **1,000,000** (1M) and **5,000,000** (5M)

## Space Allocation

| Statement                         | Size (MB) in 1M records | Size (MB) in 5M records |
| --------------------------------- | ----------------------: | ----------------------: |
| Auto Increment BigInt Primary Key |                    32.6 |                   125.2 |
| UUID Primary Key                  |                    62.7 |                   506.8 |
| Ordered UUID Primary Key          |                    62.7 |                   350.0 |
| Binary UUID Primary Key           |                    41.6 |                   352.7 |
| Ordered Binary UUID Primary Key   |                    41.6 |                   205.3 |
| Snowflake ID                      |                   157.2 |                   202.3 |

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
                   │ autoincr.txt │                uuid.txt                │          ordered-uuid.txt           │              bin-uuid.txt              │        bin-ordered-uuid.txt         │            snowflake.txt             │
                  │    sec/op    │    sec/op      vs base                 │   sec/op     vs base                │    sec/op      vs base                 │   sec/op     vs base                │    sec/op     vs base                │
Insert-8             576.4µ ± 2%   2050.9µ ± 11%  +255.84% (p=0.000 n=10)   964.7µ ± 6%  +67.37% (p=0.000 n=10)   3175.2µ ± 44%  +450.90% (p=0.000 n=10)   861.1µ ± 2%  +49.40% (p=0.000 n=10)   1108.8µ ± 2%  +92.38% (p=0.000 n=10)
FindByID-8           63.90µ ± 5%    66.17µ ±  1%    +3.56% (p=0.000 n=10)   67.26µ ± 1%   +5.26% (p=0.000 n=10)    65.37µ ±  1%    +2.31% (p=0.000 n=10)   65.83µ ± 0%   +3.03% (p=0.000 n=10)   108.66µ ± 4%  +70.05% (p=0.000 n=10)
GetList-8            76.33µ ± 0%    83.68µ ±  2%    +9.63% (p=0.000 n=10)   84.36µ ± 3%  +10.52% (p=0.000 n=10)    74.73µ ±  0%    -2.10% (p=0.000 n=10)   74.97µ ± 1%   -1.78% (p=0.000 n=10)   123.87µ ± 5%  +62.28% (p=0.000 n=10)
GetListRandomly-8    111.6µ ± 1%    131.6µ ±  3%   +17.93% (p=0.000 n=10)   136.4µ ± 1%  +22.27% (p=0.000 n=10)    115.1µ ±  1%    +3.19% (p=0.000 n=10)   115.4µ ± 6%   +3.42% (p=0.000 n=10)    188.5µ ± 1%  +69.01% (p=0.000 n=10)
geomean              133.1µ         196.6µ         +47.74%                  165.3µ       +24.22%                   205.6µ         +54.47%                  148.8µ       +11.82%                   230.3µ       +73.07%
```

## Reference

- https://100go.co/89-benchmarks/
- https://blog.jiayu.co/2019/05/benchmarking-go-code/
- https://medium.com/@adamszpilewicz/benchmarking-and-performance-comparison-in-go-using-benchstat-3da48bfaeede
