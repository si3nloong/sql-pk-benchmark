# SQL Primary Key benchmark

> This is an experiment to test the speed differences between different type of primary key.

## Setup

<p>Machine: Macbook Pro</p>
<p>CPU: M1</p>
<p>Memory: 16GB</p>
<p>OS: macOS Ventura 13.5.1</p>
<p>Record set: 1000000</p>

## Space Allocation

| Statement                         | Size (MB) |
| --------------------------------- | --------- |
| Auto Increment BigInt Primary Key | 125.2     |
| UUID Primary Key                  | 506.8     |
| Ordered UUID Primary Key          | 350       |
| Binary UUID Primary Key           | 352.7     |
| Ordered Binary UUID Primary Key   | 205.3     |

## Benchmarks

Insert 100 record per batch

```bash
BenchmarkInsertAutoIncrementID-8        1000000000               0.01174 ns/op         0 B/op          0 allocs/op
BenchmarkInsertUUID-8                   1000000000               0.02422 ns/op         0 B/op          0 allocs/op
BenchmarkInsertOrderedUUID-8            1000000000               0.01864 ns/op         0 B/op          0 allocs/op
BenchmarkInsertBinaryUUID-8             1000000000               0.03470 ns/op         0 B/op          0 allocs/op
BenchmarkInsertBinaryOrderedUUID-8      1000000000               0.01518 ns/op         0 B/op          0 allocs/op
```

Find record by id

```bash
BenchmarkFindByIDAutoIncrementID-8      1000000000               0.004100 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDUUID-8                 1000000000               0.004609 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDOrderedUUID-8          1000000000               0.004292 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDBinaryUUID-8           1000000000               0.004229 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDBinaryOrderedUUID-8    1000000000               0.004241 ns/op        0 B/op          0 allocs/op
```

Get 100 records

```bash
BenchmarkGetListAutoIncrement-8         1000000000               0.003163 ns/op        0 B/op          0 allocs/op
BenchmarkGetListUUID-8                  1000000000               0.002804 ns/op        0 B/op          0 allocs/op
BenchmarkGetListOrderedUUID-8           1000000000               0.002517 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryUUID-8            1000000000               0.002311 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryOrderedUUID-8     1000000000               0.002190 ns/op        0 B/op          0 allocs/op
```
