package main

import (
	"context"
	"database/sql/driver"
	"math/rand"
	"strings"
	"testing"

	"github.com/si3nloong/sqlgen/sequel"
	"github.com/si3nloong/sqlgen/sequel/db"
)

func setup[T interface {
	sequel.Migrator
	sequel.Tabler
}]() (context.Context, func()) {
	ctx := context.Background()
	if err := db.Migrate[T](ctx, dbConn); err != nil {
		panic(err)
	}
	return ctx, func() {
		// db.DropTable[T](ctx, dbConn)
	}
}

func benchmarkInsert[T interface {
	sequel.Migrator
	sequel.Columner
	sequel.Tabler
	sequel.Valuer
}](b *testing.B, callback func() []T) {
	var (
		ctx, cleanUp = setup[T]()
		data         []T
		err          error
	)
	defer cleanUp()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data = callback()
		b.StartTimer()
		if _, err = db.InsertInto(ctx, dbConn, data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkAutoIncrementIDInsert(b *testing.B) {
	benchmarkInsert(b, AutoIncrIDBatch)
}
func BenchmarkUUIDInsert(b *testing.B) {
	benchmarkInsert(b, NormalUUIDBatch)
}
func BenchmarkOrderedUUIDInsert(b *testing.B) {
	benchmarkInsert(b, NormalOrderedUUIDBatch)
}
func BenchmarkBinaryUUIDInsert(b *testing.B) {
	benchmarkInsert(b, BinaryUUIDBatch)
}
func BenchmarkBinaryOrderedUUIDInsert(b *testing.B) {
	benchmarkInsert(b, BinaryOrderedUUIDBatch)
}
func BenchmarkSnowflakeIDInsert(b *testing.B) {
	benchmarkInsert(b, SnowflakeIDBatch)
}

func findRandomly[T interface {
	sequel.Tabler
	sequel.Columner
}, Ptr sequel.Scanner[T]](ctx context.Context) ([]Ptr, error) {
	var (
		v     T
		query = `SELECT ` +
			strings.Join(v.Columns(), ",") +
			" FROM " + v.TableName() + " ORDER BY RAND() LIMIT 50"
	)
	result, err := db.QueryScan[T, Ptr](ctx, dbConn, query)
	if err != nil {
		return nil, err
	}
	ptrs := make([]Ptr, len(result))
	for i := 0; i < len(result); i++ {
		ptrs[i] = &result[i]
	}
	return ptrs, nil
}

func benchmarkFindByID[T interface {
	sequel.Migrator
	sequel.Tabler
	sequel.Columner
	sequel.Keyer
	sequel.Valuer
}, Ptr sequel.KeyValueScanner[T]](b *testing.B) {
	var (
		ctx, cleanUp = setup[T]()
		data         []Ptr
		idx          int
		err          error
	)
	defer cleanUp()

	data, err = findRandomly[T, Ptr](ctx)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		idx = rand.Intn(len(data))
		if err = db.FindOne(ctx, dbConn, data[idx]); err != nil {
			panic(err)
		}
	}
}

func BenchmarkAutoIncrementIDFindByID(b *testing.B) {
	benchmarkFindByID[AutoIncrID](b)
}
func BenchmarkUUIDFindByID(b *testing.B) {
	benchmarkFindByID[NormalUUID](b)
}
func BenchmarkOrderedUUIDFindByID(b *testing.B) {
	benchmarkFindByID[NormalOrderedUUID](b)
}
func BenchmarkBinaryUUIDFindByID(b *testing.B) {
	benchmarkFindByID[BinaryUUID](b)
}
func BenchmarkBinaryOrderedUUIDFindByID(b *testing.B) {
	benchmarkFindByID[BinaryOrderedUUID](b)
}
func BenchmarkSnowflakeIDFindByID(b *testing.B) {
	benchmarkFindByID[SnowflakeID](b)
}

// func getList[T interface {
// 	sequel.Keyer
// 	sequel.Tabler
// 	sequel.Columner
// }, Ptr sequel.Scanner[T]](ctx context.Context) ([]T, error) {
// 	var v T
// 	pkName, _, _ := v.PK()
// 	result, err := db.QueryScan[T, Ptr](ctx, dbConn,
// 		`SELECT `+strings.Join(v.Columns(), ",")+` FROM `+v.TableName()+` ORDER BY `+pkName+` LIMIT 100;`)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return result, nil
// }

func benchmarkGetList[T interface {
	sequel.Tabler
	sequel.Keyer
	sequel.Columner
	sequel.Migrator
}, Ptr sequel.Scanner[T]](b *testing.B) {
	var (
		ctx, cleanUp = setup[T]()
		v            T
		pkName       string
		query        string
		result       []T
		err          error
	)
	defer cleanUp()

	pkName, _, _ = v.PK()
	query = `SELECT ` + strings.Join(v.Columns(), ",") + ` FROM ` + v.TableName() + ` ORDER BY ` + pkName + ` LIMIT 100;`

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result, err = db.QueryScan[T, Ptr](ctx, dbConn, query)
		if err != nil {
			panic(err)
		} else if len(result) == 0 {
			panic("no record")
		}
	}
}

func BenchmarkAutoIncrementIDGetList(b *testing.B) {
	benchmarkGetList[AutoIncrID](b)
}
func BenchmarkUUIDGetList(b *testing.B) {
	benchmarkGetList[NormalUUID](b)
}
func BenchmarkOrderedUUIDGetList(b *testing.B) {
	benchmarkGetList[NormalOrderedUUID](b)
}
func BenchmarkBinaryUUIDGetList(b *testing.B) {
	benchmarkGetList[BinaryUUID](b)
}
func BenchmarkBinaryOrderedUUIDGetList(b *testing.B) {
	benchmarkGetList[BinaryOrderedUUID](b)
}
func BenchmarkSnowflakeIDGetList(b *testing.B) {
	benchmarkGetList[SnowflakeID](b)
}

func benchmarkGetRandomly[T interface {
	sequel.Tabler
	sequel.Columner
	sequel.Migrator
}, Ptr interface {
	sequel.Keyer
	sequel.Scanner[T]
}](b *testing.B) {
	var (
		ctx, cleanUp = setup[T]()
		data         []Ptr
		err          error
	)
	defer cleanUp()

	var (
		v      T
		pkName string
		idx    int
		pk     driver.Value
		query  string
		result []T
	)

	data, err = findRandomly[T, Ptr](ctx)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		idx = rand.Intn(len(data))
		pkName, _, pk = data[idx].PK()
		query = `SELECT ` + strings.Join(v.Columns(), ",") + ` FROM ` + v.TableName() + ` WHERE ` + pkName + ` >= ? ORDER BY ` + pkName + ` LIMIT 100;`
		b.StartTimer()
		result, err = db.QueryScan[T, Ptr](ctx, dbConn, query, pk)
		if err != nil {
			panic(err)
		} else if len(result) == 0 {
			panic("no record")
		}
	}
}

func BenchmarkAutoIncrementIDGetRandomly(b *testing.B) {
	benchmarkGetRandomly[AutoIncrID](b)
}
func BenchmarkUUIDGetRandomly(b *testing.B) {
	benchmarkGetRandomly[NormalUUID](b)
}
func BenchmarkOrderedUUIDGetRandomly(b *testing.B) {
	benchmarkGetRandomly[NormalOrderedUUID](b)
}
func BenchmarkBinaryUUIDGetRandomly(b *testing.B) {
	benchmarkGetRandomly[BinaryUUID](b)
}
func BenchmarkBinaryOrderedUUIDGetRandomly(b *testing.B) {
	benchmarkGetRandomly[BinaryOrderedUUID](b)
}
func BenchmarkSnowflakeIDGetRandomly(b *testing.B) {
	benchmarkGetRandomly[SnowflakeID](b)
}
