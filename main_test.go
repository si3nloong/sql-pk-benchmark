package main

import (
	"context"
	"strings"
	"testing"

	"github.com/si3nloong/sqlgen/sequel"
	"github.com/si3nloong/sqlgen/sequel/db"
)

const (
	round     = 20
	findCount = 10
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
		db.DropTable[T](ctx, dbConn)
	}
}

func BenchmarkInsertAutoIncrementID(b *testing.B) {
	ctx, cleanUp := setup[AutoIncrID]()
	defer cleanUp()

	b.ResetTimer()
	for i := 0; i < round; i++ {
		b.StopTimer()
		data := AutoIncrIDBatch()
		b.StartTimer()
		if _, err := db.InsertInto(ctx, dbConn, data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertUUID(b *testing.B) {
	ctx, cleanUp := setup[NormalUUID]()
	defer cleanUp()

	b.ResetTimer()
	for i := 0; i < round; i++ {
		b.StopTimer()
		data := NormalUUIDBatch()
		b.StartTimer()
		if _, err := db.InsertInto(ctx, dbConn, data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertOrderedUUID(b *testing.B) {
	ctx, cleanUp := setup[NormalOrderedUUID]()
	defer cleanUp()

	b.ResetTimer()
	for i := 0; i < round; i++ {
		b.StopTimer()
		data := NormalOrderedUUIDBatch()
		b.StartTimer()
		if _, err := db.InsertInto(ctx, dbConn, data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertBinaryUUID(b *testing.B) {
	ctx, cleanUp := setup[BinaryUUID]()
	defer cleanUp()

	b.ResetTimer()
	for i := 0; i < round; i++ {
		b.StopTimer()
		data := BinaryUUIDBatch()
		b.StartTimer()
		if _, err := db.InsertInto(ctx, dbConn, data); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkInsertBinaryOrderedUUID(b *testing.B) {
	ctx, cleanUp := setup[BinaryOrderedUUID]()
	defer cleanUp()

	b.ResetTimer()
	for i := 0; i < round; i++ {
		b.StopTimer()
		data := BinaryOrderedUUIDBatch()
		b.StartTimer()
		if _, err := db.InsertInto(ctx, dbConn, data); err != nil {
			b.Fatal(err)
		}
	}
}

func findRandomly[T interface {
	sequel.Tabler
	sequel.Columner
}, Ptr sequel.Scanner[T]](ctx context.Context) (*T, error) {
	var v T
	result, err := db.QueryScan[T, Ptr](ctx, dbConn, `SELECT `+
		strings.Join(v.Columns(), ",")+
		" FROM "+v.TableName()+" ORDER BY RAND() LIMIT 1")
	if err != nil {
		return nil, err
	}
	return &result[0], nil
}

func BenchmarkFindByIDAutoIncrementID(b *testing.B) {
	ctx, cleanUp := setup[AutoIncrID]()
	defer cleanUp()

	randomRecord := make([]*AutoIncrID, findCount)
	for i := 0; i < findCount; i++ {
		data, err := findRandomly[AutoIncrID](ctx)
		if err != nil {
			panic(err)
		}
		randomRecord[i] = data
	}

	b.ResetTimer()
	b.StartTimer()

	for i := range randomRecord {
		if err := db.FindOne(ctx, dbConn, randomRecord[i]); err != nil {
			panic(err)
		}
	}

	b.StopTimer()
}

func BenchmarkFindByIDUUID(b *testing.B) {
	ctx, cleanUp := setup[NormalUUID]()
	defer cleanUp()

	randomRecord := make([]*NormalUUID, findCount)
	for i := 0; i < findCount; i++ {
		data, err := findRandomly[NormalUUID](ctx)
		if err != nil {
			panic(err)
		}
		randomRecord[i] = data
	}

	b.ResetTimer()
	b.StartTimer()

	for i := range randomRecord {
		if err := db.FindOne(ctx, dbConn, randomRecord[i]); err != nil {
			panic(err)
		}
	}

	b.StopTimer()
}

func BenchmarkFindByIDOrderedUUID(b *testing.B) {
	ctx, cleanUp := setup[NormalOrderedUUID]()
	defer cleanUp()
	randomRecord := make([]*NormalOrderedUUID, findCount)
	for i := 0; i < findCount; i++ {
		data, err := findRandomly[NormalOrderedUUID](ctx)
		if err != nil {
			panic(err)
		}
		randomRecord[i] = data
	}

	b.ResetTimer()
	b.StartTimer()
	for i := range randomRecord {
		if err := db.FindOne(ctx, dbConn, randomRecord[i]); err != nil {
			panic(err)
		}
	}
	b.StopTimer()
}

func BenchmarkFindByIDBinaryUUID(b *testing.B) {
	ctx, cleanUp := setup[BinaryUUID]()
	defer cleanUp()
	randomRecord := make([]*BinaryUUID, findCount)
	for i := 0; i < findCount; i++ {
		data, err := findRandomly[BinaryUUID](ctx)
		if err != nil {
			panic(err)
		}
		randomRecord[i] = data
	}

	b.ResetTimer()
	b.StartTimer()
	for i := range randomRecord {
		if err := db.FindOne(ctx, dbConn, randomRecord[i]); err != nil {
			panic(err)
		}
	}
	b.StopTimer()
}

func BenchmarkFindByIDBinaryOrderedUUID(b *testing.B) {
	ctx, cleanUp := setup[BinaryOrderedUUID]()
	defer cleanUp()
	randomRecord := make([]*BinaryOrderedUUID, findCount)
	for i := 0; i < findCount; i++ {
		data, err := findRandomly[BinaryOrderedUUID](ctx)
		if err != nil {
			panic(err)
		}
		randomRecord[i] = data
	}

	b.ResetTimer()
	b.StartTimer()
	for i := range randomRecord {
		if err := db.FindOne(ctx, dbConn, randomRecord[i]); err != nil {
			panic(err)
		}
	}
	b.StopTimer()
}

func getList[T interface {
	sequel.Keyer
	sequel.Tabler
	sequel.Columner
}, Ptr sequel.Scanner[T]](ctx context.Context) ([]T, error) {
	var v T
	pkName, _, _ := v.PK()
	result, err := db.QueryScan[T, Ptr](ctx, dbConn,
		`SELECT `+strings.Join(v.Columns(), ",")+` FROM `+v.TableName()+` ORDER BY `+pkName+` LIMIT 100;`)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func benchmarkGetList[T interface {
	sequel.Tabler
	sequel.Keyer
	sequel.Columner
	sequel.Migrator
}, Ptr sequel.Scanner[T]](b *testing.B) {
	ctx, cleanUp := setup[T]()
	defer cleanUp()

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < round; i++ {
		if _, err := getList[T, Ptr](ctx); err != nil {
			panic(err)
		}
	}
	b.StopTimer()
}

func BenchmarkGetListAutoIncrement(b *testing.B) {
	benchmarkGetList[AutoIncrID](b)
}
func BenchmarkGetListUUID(b *testing.B) {
	benchmarkGetList[NormalUUID](b)
}
func BenchmarkGetListOrderedUUID(b *testing.B) {
	benchmarkGetList[NormalOrderedUUID](b)
}
func BenchmarkGetListBinaryUUID(b *testing.B) {
	benchmarkGetList[BinaryUUID](b)
}
func BenchmarkGetListBinaryOrderedUUID(b *testing.B) {
	benchmarkGetList[BinaryOrderedUUID](b)
}

func getListByCursor[T interface {
	sequel.Tabler
	sequel.Columner
}, Ptr interface {
	sequel.Scanner[T]
	sequel.Keyer
}](ctx context.Context, vi Ptr) ([]T, error) {
	var v T
	pkName, _, pk := vi.PK()
	result, err := db.QueryScan[T, Ptr](ctx, dbConn,
		`SELECT `+strings.Join(v.Columns(), ",")+` FROM `+v.TableName()+` WHERE `+pkName+` = ? ORDER BY `+pkName+` LIMIT 100;`, pk)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func benchmarkGetRandomly[T interface {
	sequel.Tabler
	sequel.Columner
	sequel.Migrator
}, Ptr interface {
	sequel.Keyer
	sequel.Scanner[T]
}](b *testing.B) {
	ctx, cleanUp := setup[T]()
	defer cleanUp()

	randomRecord := make([]*T, findCount)
	for i := 0; i < findCount; i++ {
		data, err := findRandomly[T, Ptr](ctx)
		if err != nil {
			panic(err)
		}
		randomRecord[i] = data
	}

	b.ResetTimer()
	b.StartTimer()
	for i := 0; i < findCount; i++ {
		if result, err := getListByCursor[T, Ptr](ctx, randomRecord[i]); err != nil {
			panic(err)
		} else if len(result) == 0 {
			panic("no record")
		}
	}
	b.StopTimer()
}

func BenchmarkGetRandomlyAutoIncrement(b *testing.B) {
	benchmarkGetRandomly[AutoIncrID](b)
}
func BenchmarkGetRandomlyUUID(b *testing.B) {
	benchmarkGetRandomly[NormalUUID](b)
}
func BenchmarkGetRandomlyOrderedUUID(b *testing.B) {
	benchmarkGetRandomly[NormalOrderedUUID](b)
}
func BenchmarkGetRandomlyBinaryUUID(b *testing.B) {
	benchmarkGetRandomly[BinaryUUID](b)
}
func BenchmarkGetRandomlyBinaryOrderedUUID(b *testing.B) {
	benchmarkGetRandomly[BinaryOrderedUUID](b)
}
