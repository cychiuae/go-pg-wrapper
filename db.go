/*
Copyright © 2020 Tomaz Lovrec <tomaz.lovrec@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package pgwrapper

import (
	"context"
	"io"
	"time"

	"github.com/go-pg/pg/v10"
	"gitlab.com/slax0rr/go-pg-wrapper/orm"
)

type DB interface {
	AddQueryHook(hook pg.QueryHook)
	Begin() (Tx, error)
	Close() error
	Conn() Conn
	Context() context.Context
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error)
	Exec(query interface{}, params ...interface{}) (res pg.Result, err error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	ExecOne(query interface{}, params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	Model(model ...interface{}) orm.Query
	ModelContext(c context.Context, model ...interface{}) orm.Query
	Options() *pg.Options
	Param(param string) interface{}
	PoolStats() *pg.PoolStats
	Prepare(q string) (Stmt, error)
	Query(model, query interface{}, params ...interface{}) (res pg.Result, err error)
	QueryContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (pg.Result, error)
	RunInTransaction(ctx context.Context, fn func(Tx) error) error
	String() string
	WithContext(ctx context.Context) DB
	WithParam(param string, value interface{}) DB
	WithTimeout(dur time.Duration) DB
}

type DBWrap struct {
	db *pg.DB
}

func NewDB(db *pg.DB) *DBWrap {
	return &DBWrap{db}
}

func (d *DBWrap) AddQueryHook(hook pg.QueryHook) {
	d.db.AddQueryHook(hook)
}

func (d *DBWrap) Begin() (Tx, error) {
	tx, err := d.db.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (d *DBWrap) Close() error {
	return d.db.Close()
}

func (d *DBWrap) Conn() Conn {
	return NewConn(d.db.Conn())
}

func (d *DBWrap) Context() context.Context {
	return d.db.Context()
}

func (d *DBWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.CopyFrom(r, query, params...)
}

func (d *DBWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.CopyTo(w, query, params...)
}

func (d *DBWrap) Exec(query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.Exec(query, params...)
}

func (d *DBWrap) ExecContext(ctx context.Context, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.ExecContext(ctx, query, params...)
}

func (d *DBWrap) ExecOne(query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.ExecOne(query, params...)
}

func (d *DBWrap) ExecOneContext(ctx context.Context, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.ExecOneContext(ctx, query, params...)
}

func (d *DBWrap) Model(model ...interface{}) orm.Query {
	query := d.db.Model(model...)
	return orm.NewQuery(query)
}

func (d *DBWrap) ModelContext(c context.Context, model ...interface{}) orm.Query {
	query := d.db.ModelContext(c, model...)
	return orm.NewQuery(query)
}

func (d *DBWrap) Options() *pg.Options {
	return d.db.Options()
}

func (d *DBWrap) Param(param string) interface{} {
	return d.db.Param(param)
}

func (d *DBWrap) PoolStats() *pg.PoolStats {
	return d.db.PoolStats()
}

func (d *DBWrap) Prepare(q string) (Stmt, error) {
	stmt, err := d.db.Prepare(q)
	if err != nil {
		return nil, err
	}
	return NewStmt(stmt), nil
}

func (d *DBWrap) Query(model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.Query(model, query, params...)
}

func (d *DBWrap) QueryContext(c context.Context, model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.QueryContext(c, model, query, params...)
}

func (d *DBWrap) QueryOne(model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.QueryOne(model, query, params...)
}

func (d *DBWrap) QueryOneContext(c context.Context, model, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return d.db.QueryOneContext(c, model, query, params...)
}

func (d *DBWrap) RunInTransaction(ctx context.Context, fn func(Tx) error) error {
	tx, err := d.db.Begin()
	if err != nil {
		return err
	}
	return NewTx(tx).RunInTransaction(ctx, fn)
}

func (d *DBWrap) String() string {
	return d.db.String()
}

func (d *DBWrap) WithContext(ctx context.Context) DB {
	db := d.db.WithContext(ctx)
	return NewDB(db)
}

func (d *DBWrap) WithParam(param string, value interface{}) DB {
	db := d.db.WithParam(param, value)
	return NewDB(db)
}

func (d *DBWrap) WithTimeout(dur time.Duration) DB {
	db := d.db.WithTimeout(dur)
	return NewDB(db)
}
