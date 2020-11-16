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

	"github.com/cychiuae/go-pg-wrapper/v10/orm"
	"github.com/go-pg/pg/v10"
	pgorm "github.com/go-pg/pg/v10/orm"
)

type Tx interface {
	Begin() (Tx, error)
	Close() error
	Commit() error
	Context() context.Context
	CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error)
	CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error)
	Exec(query interface{}, params ...interface{}) (pg.Result, error)
	ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	ExecOne(query interface{}, params ...interface{}) (pg.Result, error)
	ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error)
	Formatter() pgorm.QueryFormatter
	Model(model ...interface{}) orm.Query
	ModelContext(c context.Context, model ...interface{}) orm.Query
	Prepare(q string) (*pg.Stmt, error)
	Query(model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOne(model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	QueryOneContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error)
	Rollback() error
	RunInTransaction(context.Context, func(Tx) error) error
	Stmt(stmt Stmt) Stmt
}

type TxWrap struct {
	tx *pg.Tx
}

func NewTx(tx *pg.Tx) *TxWrap {
	return &TxWrap{tx}
}

func (t *TxWrap) Begin() (Tx, error) {
	tx, err := t.tx.Begin()
	if err != nil {
		return nil, err
	}
	return NewTx(tx), nil
}

func (t *TxWrap) Close() error {
	return t.tx.Close()
}

func (t *TxWrap) Commit() error {
	return t.tx.Commit()
}

func (t *TxWrap) Context() context.Context {
	return t.tx.Context()
}

func (t *TxWrap) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return t.tx.CopyFrom(r, query, params...)
}

func (t *TxWrap) CopyTo(w io.Writer, query interface{}, params ...interface{}) (res pg.Result, err error) {
	return t.tx.CopyTo(w, query, params...)
}

func (t *TxWrap) Exec(query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.Exec(query, params...)
}

func (t *TxWrap) ExecContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.ExecContext(c, query, params...)
}

func (t *TxWrap) ExecOne(query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.ExecOne(query, params)
}

func (t *TxWrap) ExecOneContext(c context.Context, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.ExecOneContext(c, query, params...)
}

func (t *TxWrap) Formatter() pgorm.QueryFormatter {
	return t.tx.Formatter()
}

func (t *TxWrap) Model(model ...interface{}) orm.Query {
	return orm.NewQuery(t.tx.Model(model...))
}

func (t *TxWrap) ModelContext(c context.Context, model ...interface{}) orm.Query {
	return orm.NewQuery(t.tx.ModelContext(c, model...))
}

func (t *TxWrap) Prepare(q string) (*pg.Stmt, error) {
	return t.tx.Prepare(q)
}

func (t *TxWrap) Query(model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.Query(model, query, params...)
}

func (t *TxWrap) QueryContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.QueryContext(c, model, query, params...)
}

func (t *TxWrap) QueryOne(model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.QueryOne(model, query, params...)
}

func (t *TxWrap) QueryOneContext(c context.Context, model interface{}, query interface{}, params ...interface{}) (pg.Result, error) {
	return t.tx.QueryOneContext(c, model, query, params...)
}

func (t *TxWrap) Rollback() error {
	return t.tx.Rollback()
}

func (t *TxWrap) RunInTransaction(ctx context.Context, fn func(Tx) error) error {
	return t.tx.RunInTransaction(ctx, func(tx *pg.Tx) error { return fn(t) })
}

func (t *TxWrap) Stmt(stmt Stmt) Stmt {
	return NewStmt(t.tx.Stmt(stmt.(*StmtWrap).stmt))
}
