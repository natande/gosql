package gosql

import (
	"database/sql"
	"reflect"
)

var _tableNamingType = reflect.TypeOf((*tableNaming)(nil)).Elem()

type DB struct {
	db         *sql.DB
	driverName string
}

// Open opens database
// dataSourceName's format: username:password@tcp(host:port)/dbName
func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{
		db:         db,
		driverName: driverName,
	}, nil
}

func (d *DB) SQLDB() *sql.DB {
	return d.db
}

func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.db.Exec(query, args...)
}

func (d *DB) MustExec(query string, args ...interface{}) {
	_, err := d.db.Exec(query, args...)
	if err != nil {
		panic(err)
	}
}

func (d *DB) Begin() (*Tx, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}

	return &Tx{
		tx:         tx,
		driverName: d.driverName,
	}, nil
}

func (d *DB) Close() error {
	return d.db.Close()
}

func (d *DB) Table(name string) *Table {
	return &Table{
		exe:        d.db,
		driverName: d.driverName,
		name:       name,
	}
}

func (d *DB) Insert(record interface{}) error {
	return d.Table(getTableName(record)).Insert(record)
}

func (d *DB) Update(record interface{}) error {
	return d.Table(getTableName(record)).Update(record)
}

func (d *DB) Save(record interface{}) error {
	return d.Table(getTableName(record)).Save(record)
}

func (d *DB) Select(records interface{}, where string, args ...interface{}) error {
	return d.Table(getTableNameBySlice(records)).Select(records, where, args...)
}

func (d *DB) SelectOne(record interface{}, where string, args ...interface{}) error {
	return d.Table(getTableName(record)).SelectOne(record, where, args...)
}
