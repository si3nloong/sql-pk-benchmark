# SQL Primary Key Benchmarks

> This is an experiment to test the speed differences between different type of primary key.

## Setup

<p>Machine: Macbook Pro</p>
<p>CPU: M1</p>
<p>Memory: 16GB</p>
<p>OS: macOS Ventura 13.5.1</p>
<p>Record set: 1000000</p>

## Space Allocation

| Statement                         | Size (MB) in 1M records | Size (MB) in 5M records |
| --------------------------------- | ----------------------: | ----------------------: |
| Auto Increment BigInt Primary Key |                    32.6 |                   125.2 |
| UUID Primary Key                  |                    62.7 |                   506.8 |
| Ordered UUID Primary Key          |                    62.7 |                   350.0 |
| Binary UUID Primary Key           |                    41.6 |                   352.7 |
| Ordered Binary UUID Primary Key   |                    41.6 |                   205.3 |

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
BenchmarkGetListAutoIncrement-8         1000000000               0.004683 ns/op        0 B/op          0 allocs/op
BenchmarkGetListUUID-8                  1000000000               0.004217 ns/op        0 B/op          0 allocs/op
BenchmarkGetListOrderedUUID-8           1000000000               0.003579 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryUUID-8            1000000000               0.003749 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryOrderedUUID-8     1000000000               0.003376 ns/op        0 B/op          0 allocs/op
```

Get 100 records with order by asc.

```bash
BenchmarkGetListAutoIncrement-8         1000000000               0.004461 ns/op        0 B/op          0 allocs/op
BenchmarkGetListUUID-8                  1000000000               0.004077 ns/op        0 B/op          0 allocs/op
BenchmarkGetListOrderedUUID-8           1000000000               0.003914 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryUUID-8            1000000000               0.003745 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryOrderedUUID-8     1000000000               0.003627 ns/op        0 B/op          0 allocs/op
```

Get List using random primary key cursor and order by asc.

```bash
BenchmarkGetRandomlyAutoIncrement-8             1000000000               0.002429 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyUUID-8                      1000000000               0.002549 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyOrderedUUID-8               1000000000               0.002311 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyBinaryUUID-8                1000000000               0.002629 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyBinaryOrderedUUID-8         1000000000               0.002601 ns/op        0 B/op          0 allocs/op
```
