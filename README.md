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

## Benchmarks

### 1M records

Insert 100 record per batch.

```bash
go test -bench=Insert -benchmem -benchtime=10s

BenchmarkInsertAutoIncrementID-8        1000000000               0.01174 ns/op         0 B/op          0 allocs/op
BenchmarkInsertUUID-8                   1000000000               0.02422 ns/op         0 B/op          0 allocs/op
BenchmarkInsertOrderedUUID-8            1000000000               0.01864 ns/op         0 B/op          0 allocs/op
BenchmarkInsertBinaryUUID-8             1000000000               0.03470 ns/op         0 B/op          0 allocs/op
BenchmarkInsertBinaryOrderedUUID-8      1000000000               0.01518 ns/op         0 B/op          0 allocs/op
```

Find record by id.

```bash
go test -bench=FindByID -benchmem -benchtime=10s

BenchmarkFindByIDAutoIncrementID-8      1000000000               0.004100 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDUUID-8                 1000000000               0.004609 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDOrderedUUID-8          1000000000               0.004292 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDBinaryUUID-8           1000000000               0.004229 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDBinaryOrderedUUID-8    1000000000               0.004241 ns/op        0 B/op          0 allocs/op
```

<!-- Get 100 records

```bash
BenchmarkGetListAutoIncrement-8         1000000000               0.004683 ns/op        0 B/op          0 allocs/op
BenchmarkGetListUUID-8                  1000000000               0.004217 ns/op        0 B/op          0 allocs/op
BenchmarkGetListOrderedUUID-8           1000000000               0.003579 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryUUID-8            1000000000               0.003749 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryOrderedUUID-8     1000000000               0.003376 ns/op        0 B/op          0 allocs/op
``` -->

Get 100 records with order by asc.

```bash
go test -bench=GetList -benchmem -benchtime=10s

BenchmarkGetListAutoIncrement-8         1000000000               0.004461 ns/op        0 B/op          0 allocs/op
BenchmarkGetListUUID-8                  1000000000               0.004077 ns/op        0 B/op          0 allocs/op
BenchmarkGetListOrderedUUID-8           1000000000               0.003914 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryUUID-8            1000000000               0.003745 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryOrderedUUID-8     1000000000               0.003627 ns/op        0 B/op          0 allocs/op
```

Get List using random primary key cursor and order by asc.

```bash
go test -bench=GetRandomly -benchmem -benchtime=10s

BenchmarkGetRandomlyAutoIncrement-8             1000000000               0.002429 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyUUID-8                      1000000000               0.002549 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyOrderedUUID-8               1000000000               0.002311 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyBinaryUUID-8                1000000000               0.002629 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyBinaryOrderedUUID-8         1000000000               0.002601 ns/op        0 B/op          0 allocs/op
```

### 5M records

Insert 100 record per batch.

```bash
go test -bench=Insert -benchmem -benchtime=10s

BenchmarkInsertAutoIncrementID-8        1000000000               0.002597 ns/op        0 B/op          0 allocs/op
BenchmarkInsertUUID-8                   1000000000               0.01232 ns/op         0 B/op          0 allocs/op
BenchmarkInsertOrderedUUID-8            1000000000               0.006094 ns/op        0 B/op          0 allocs/op
BenchmarkInsertBinaryUUID-8             1000000000               0.01839 ns/op         0 B/op          0 allocs/op
BenchmarkInsertBinaryOrderedUUID-8      1000000000               0.003050 ns/op        0 B/op          0 allocs/op
```

Find record by id.

```bash
go test -bench=FindByID -benchmem -benchtime=10s

BenchmarkFindByIDAutoIncrementID-8      1000000000               0.002176 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDUUID-8                 1000000000               0.002669 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDOrderedUUID-8          1000000000               0.003597 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDBinaryUUID-8           1000000000               0.004033 ns/op        0 B/op          0 allocs/op
BenchmarkFindByIDBinaryOrderedUUID-8    1000000000               0.003370 ns/op        0 B/op          0 allocs/op
```

Get 100 records with order by asc.

```bash
go test -bench=GetList -benchmem -benchtime=10s

BenchmarkGetListAutoIncrement-8         1000000000               0.003550 ns/op        0 B/op          0 allocs/op
BenchmarkGetListUUID-8                  1000000000               0.002468 ns/op        0 B/op          0 allocs/op
BenchmarkGetListOrderedUUID-8           1000000000               0.002353 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryUUID-8            1000000000               0.002485 ns/op        0 B/op          0 allocs/op
BenchmarkGetListBinaryOrderedUUID-8     1000000000               0.002294 ns/op        0 B/op          0 allocs/op
```

Get List using random primary key cursor and order by asc.

```bash
go test -bench=GetRandomly -benchmem -benchtime=10s

BenchmarkGetRandomlyAutoIncrement-8             1000000000               0.002429 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyUUID-8                      1000000000               0.002549 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyOrderedUUID-8               1000000000               0.002311 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyBinaryUUID-8                1000000000               0.002629 ns/op        0 B/op          0 allocs/op
BenchmarkGetRandomlyBinaryOrderedUUID-8         1000000000               0.002601 ns/op        0 B/op          0 allocs/op
```
