package main

import (
	"database/sql"
	"time"
	t "time"

	"github.com/godruoyi/go-snowflake"

	"github.com/gofrs/uuid/v5"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/si3nloong/sqlgen/sequel/dialect/mysql"
)

const (
	recordCount = 100
)

var (
	dbConn *sql.DB
)

func init() {
	var err error
	dbConn, err = sql.Open("mysql", "root:abcd1234@/sqlbenchmark?parseTime=true")
	if err != nil {
		panic(err)
	}
}

type record struct {
	Created t.Time
}

func newRecord() record {
	return record{Created: t.Now()}
}

type AutoIncrID struct {
	record
	ID int64 `sql:",pk,auto_increment"`
}

type TimestampID struct {
	ID time.Time `sql:",pk,size:6"`
	record
}

func TimestampIDBatch() []TimestampID {
	data := make([]TimestampID, recordCount)
	for i := range data {
		v := TimestampID{}
		v.ID = t.Now()
		v.Created = t.Now()
		data[i] = v
	}
	return data
}

type BigIntID struct {
	record
	ID int64 `sql:",pk"`
}

func AutoIncrIDBatch() []AutoIncrID {
	data := make([]AutoIncrID, recordCount)
	for i := range data {
		v := AutoIncrID{}
		v.Created = t.Now()
		data[i] = v
	}
	return data
}

type binaryUUID struct {
	ID uuid.UUID `sql:",uuid,binary,pk"`
	record
}

type BinaryUUID struct {
	ID uuid.UUID `sql:",uuid,binary,pk"`
	record
}

type BinaryOrderedUUID struct {
	ID uuid.UUID `sql:",uuid,binary,pk"`
	record
}

func BinaryUUIDBatch() []BinaryUUID {
	data := make([]BinaryUUID, recordCount)
	for i := range data {
		v := BinaryUUID{}
		v.ID, _ = uuid.NewV4()
		v.Created = t.Now()
		data[i] = v
	}
	return data
}

func BinaryOrderedUUIDBatch() []BinaryOrderedUUID {
	data := make([]BinaryOrderedUUID, recordCount)
	for i := range data {
		v := BinaryOrderedUUID{}
		v.ID, _ = uuid.NewV7()
		v.Created = t.Now()
		data[i] = v
	}
	return data
}

type NormalUUID struct {
	ID uuid.UUID `sql:",pk"`
	record
}

func NormalUUIDBatch() []NormalUUID {
	data := make([]NormalUUID, recordCount)
	for i := range data {
		v := NormalUUID{}
		v.ID, _ = uuid.NewV4()
		v.Created = t.Now()
		data[i] = v
	}
	return data
}

type NormalOrderedUUID struct {
	ID uuid.UUID `sql:",pk"`
	record
}

func NormalOrderedUUIDBatch() []NormalOrderedUUID {
	data := make([]NormalOrderedUUID, recordCount)
	for i := range data {
		v := NormalOrderedUUID{}
		v.ID, _ = uuid.NewV7()
		v.Created = t.Now()
		data[i] = v
	}
	return data
}

type SnowflakeID struct {
	ID uint64 `sql:",pk"`
	record
}

func SnowflakeIDBatch() []SnowflakeID {
	data := make([]SnowflakeID, recordCount)
	for i := range data {
		v := SnowflakeID{}
		v.ID = snowflake.ID()
		v.Created = t.Now()
		data[i] = v
	}
	return data
}
