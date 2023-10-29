package main

import (
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/si3nloong/sqlgen/sequel/types"
)

func (NormalOrderedUUID) CreateTableStmt() string {
	return "CREATE TABLE IF NOT EXISTS normal_ordered_uuid (id VARCHAR(36) NOT NULL,created DATETIME NOT NULL,PRIMARY KEY (id));"
}
func (NormalOrderedUUID) AlterTableStmt() string {
	return "ALTER TABLE normal_ordered_uuid MODIFY id VARCHAR(36) NOT NULL,MODIFY created DATETIME NOT NULL AFTER id;"
}
func (NormalOrderedUUID) TableName() string {
	return "normal_ordered_uuid"
}
func (NormalOrderedUUID) Columns() []string {
	return []string{"id", "created"}
}
func (v NormalOrderedUUID) IsAutoIncr() bool {
	return false
}
func (v NormalOrderedUUID) PK() (columnName string, pos int, value driver.Value) {
	return "id", 0, (driver.Valuer)(v.ID)
}
func (v NormalOrderedUUID) Values() []any {
	return []any{(driver.Valuer)(v.ID), time.Time(v.record.Created)}
}
func (v *NormalOrderedUUID) Addrs() []any {
	return []any{(sql.Scanner)(&v.ID), (*time.Time)(&v.record.Created)}
}

func (AutoIncrID) CreateTableStmt() string {
	return "CREATE TABLE IF NOT EXISTS auto_incr_id (created DATETIME NOT NULL,id BIGINT NOT NULL AUTO_INCREMENT,PRIMARY KEY (id));"
}
func (AutoIncrID) AlterTableStmt() string {
	return "ALTER TABLE auto_incr_id MODIFY created DATETIME NOT NULL,MODIFY id BIGINT NOT NULL AUTO_INCREMENT AFTER created;"
}
func (AutoIncrID) TableName() string {
	return "auto_incr_id"
}
func (AutoIncrID) Columns() []string {
	return []string{"created", "id"}
}
func (v AutoIncrID) IsAutoIncr() bool {
	return true
}
func (v AutoIncrID) PK() (columnName string, pos int, value driver.Value) {
	return "id", 1, int64(v.ID)
}
func (v AutoIncrID) Values() []any {
	return []any{time.Time(v.record.Created), int64(v.ID)}
}
func (v *AutoIncrID) Addrs() []any {
	return []any{(*time.Time)(&v.record.Created), types.Integer(&v.ID)}
}

func (TimestampID) CreateTableStmt() string {
	return "CREATE TABLE IF NOT EXISTS timestamp_id (id DATETIME(6) NOT NULL,created DATETIME NOT NULL,PRIMARY KEY (id));"
}
func (TimestampID) AlterTableStmt() string {
	return "ALTER TABLE timestamp_id MODIFY id DATETIME(6) NOT NULL,MODIFY created DATETIME NOT NULL AFTER id;"
}
func (TimestampID) TableName() string {
	return "timestamp_id"
}
func (TimestampID) Columns() []string {
	return []string{"id", "created"}
}
func (v TimestampID) IsAutoIncr() bool {
	return false
}
func (v TimestampID) PK() (columnName string, pos int, value driver.Value) {
	return "id", 0, time.Time(v.ID)
}
func (v TimestampID) Values() []any {
	return []any{time.Time(v.ID), time.Time(v.record.Created)}
}
func (v *TimestampID) Addrs() []any {
	return []any{(*time.Time)(&v.ID), (*time.Time)(&v.record.Created)}
}

func (BigIntID) CreateTableStmt() string {
	return "CREATE TABLE IF NOT EXISTS big_int_id (created DATETIME NOT NULL,id BIGINT NOT NULL,PRIMARY KEY (id));"
}
func (BigIntID) AlterTableStmt() string {
	return "ALTER TABLE big_int_id MODIFY created DATETIME NOT NULL,MODIFY id BIGINT NOT NULL AFTER created;"
}
func (BigIntID) TableName() string {
	return "big_int_id"
}
func (BigIntID) Columns() []string {
	return []string{"created", "id"}
}
func (v BigIntID) IsAutoIncr() bool {
	return false
}
func (v BigIntID) PK() (columnName string, pos int, value driver.Value) {
	return "id", 1, int64(v.ID)
}
func (v BigIntID) Values() []any {
	return []any{time.Time(v.record.Created), int64(v.ID)}
}
func (v *BigIntID) Addrs() []any {
	return []any{(*time.Time)(&v.record.Created), types.Integer(&v.ID)}
}

func (BinaryUUID) CreateTableStmt() string {
	return "CREATE TABLE IF NOT EXISTS binary_uuid (id BINARY(16) NOT NULL,created DATETIME NOT NULL,PRIMARY KEY (id));"
}
func (BinaryUUID) AlterTableStmt() string {
	return "ALTER TABLE binary_uuid MODIFY id BINARY(16) NOT NULL,MODIFY created DATETIME NOT NULL AFTER id;"
}
func (BinaryUUID) TableName() string {
	return "binary_uuid"
}
func (BinaryUUID) Columns() []string {
	return []string{"id", "created"}
}
func (v BinaryUUID) IsAutoIncr() bool {
	return false
}
func (v BinaryUUID) PK() (columnName string, pos int, value driver.Value) {
	return "id", 0, types.BinaryMarshaler(v.ID)
}
func (v BinaryUUID) Values() []any {
	return []any{types.BinaryMarshaler(v.ID), time.Time(v.record.Created)}
}
func (v *BinaryUUID) Addrs() []any {
	return []any{types.BinaryUnmarshaler(&v.ID), (*time.Time)(&v.record.Created)}
}

func (BinaryOrderedUUID) CreateTableStmt() string {
	return "CREATE TABLE IF NOT EXISTS binary_ordered_uuid (id BINARY(16) NOT NULL,created DATETIME NOT NULL,PRIMARY KEY (id));"
}
func (BinaryOrderedUUID) AlterTableStmt() string {
	return "ALTER TABLE binary_ordered_uuid MODIFY id BINARY(16) NOT NULL,MODIFY created DATETIME NOT NULL AFTER id;"
}
func (BinaryOrderedUUID) TableName() string {
	return "binary_ordered_uuid"
}
func (BinaryOrderedUUID) Columns() []string {
	return []string{"id", "created"}
}
func (v BinaryOrderedUUID) IsAutoIncr() bool {
	return false
}
func (v BinaryOrderedUUID) PK() (columnName string, pos int, value driver.Value) {
	return "id", 0, types.BinaryMarshaler(v.ID)
}
func (v BinaryOrderedUUID) Values() []any {
	return []any{types.BinaryMarshaler(v.ID), time.Time(v.record.Created)}
}
func (v *BinaryOrderedUUID) Addrs() []any {
	return []any{types.BinaryUnmarshaler(&v.ID), (*time.Time)(&v.record.Created)}
}

func (NormalUUID) CreateTableStmt() string {
	return "CREATE TABLE IF NOT EXISTS normal_uuid (id VARCHAR(36) NOT NULL,created DATETIME NOT NULL,PRIMARY KEY (id));"
}
func (NormalUUID) AlterTableStmt() string {
	return "ALTER TABLE normal_uuid MODIFY id VARCHAR(36) NOT NULL,MODIFY created DATETIME NOT NULL AFTER id;"
}
func (NormalUUID) TableName() string {
	return "normal_uuid"
}
func (NormalUUID) Columns() []string {
	return []string{"id", "created"}
}
func (v NormalUUID) IsAutoIncr() bool {
	return false
}
func (v NormalUUID) PK() (columnName string, pos int, value driver.Value) {
	return "id", 0, (driver.Valuer)(v.ID)
}
func (v NormalUUID) Values() []any {
	return []any{(driver.Valuer)(v.ID), time.Time(v.record.Created)}
}
func (v *NormalUUID) Addrs() []any {
	return []any{(sql.Scanner)(&v.ID), (*time.Time)(&v.record.Created)}
}