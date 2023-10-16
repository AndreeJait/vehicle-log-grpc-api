package db

import "github.com/uptrace/bun"

type DBI interface {
	NewValues(model interface{}) *bun.ValuesQuery
	NewSelect() *bun.SelectQuery
	NewInsert() *bun.InsertQuery
	NewUpdate() *bun.UpdateQuery
	NewDelete() *bun.DeleteQuery
	NewCreateTable() *bun.CreateTableQuery
	NewDropTable() *bun.DropTableQuery
	NewCreateIndex() *bun.CreateIndexQuery
	NewDropIndex() *bun.DropIndexQuery
	NewTruncateTable() *bun.TruncateTableQuery
	NewAddColumn() *bun.AddColumnQuery
	NewDropColumn() *bun.DropColumnQuery
	NewRaw(query string, args ...interface{}) *bun.RawQuery
}
