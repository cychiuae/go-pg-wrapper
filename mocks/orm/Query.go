// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import io "io"
import mock "github.com/stretchr/testify/mock"
import orm "github.com/cychiuae/go-pg-wrapper/v10/orm"
import pg "github.com/go-pg/pg/v10"
import v9orm "github.com/go-pg/pg/v10/orm"

// Query is an autogenerated mock type for the Query type
type Query struct {
	mock.Mock
}

// AllWithDeleted provides a mock function with given fields:
func (_m *Query) AllWithDeleted() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// AppendQuery provides a mock function with given fields: fmter, b
func (_m *Query) AppendQuery(fmter v9orm.QueryFormatter, b []byte) ([]byte, error) {
	ret := _m.Called(fmter, b)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(v9orm.QueryFormatter, []byte) []byte); ok {
		r0 = rf(fmter, b)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(v9orm.QueryFormatter, []byte) error); ok {
		r1 = rf(fmter, b)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Apply provides a mock function with given fields: fn
func (_m *Query) Apply(fn func(orm.Query) (orm.Query, error)) orm.Query {
	ret := _m.Called(fn)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(func(orm.Query) (orm.Query, error)) orm.Query); ok {
		r0 = rf(fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Clone provides a mock function with given fields:
func (_m *Query) Clone() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Column provides a mock function with given fields: columns
func (_m *Query) Column(columns ...string) orm.Query {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...string) orm.Query); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// ColumnExpr provides a mock function with given fields: expr, params
func (_m *Query) ColumnExpr(expr string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, expr)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(expr, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Context provides a mock function with given fields: c
func (_m *Query) Context(c context.Context) orm.Query {
	ret := _m.Called(c)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(context.Context) orm.Query); ok {
		r0 = rf(c)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// CopyFrom provides a mock function with given fields: r, query, params
func (_m *Query) CopyFrom(r io.Reader, query interface{}, params ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, r, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(io.Reader, interface{}, ...interface{}) v9orm.Result); ok {
		r0 = rf(r, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Reader, interface{}, ...interface{}) error); ok {
		r1 = rf(r, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CopyTo provides a mock function with given fields: w, query, params
func (_m *Query) CopyTo(w io.Writer, query interface{}, params ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, w, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(io.Writer, interface{}, ...interface{}) v9orm.Result); ok {
		r0 = rf(w, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(io.Writer, interface{}, ...interface{}) error); ok {
		r1 = rf(w, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Count provides a mock function with given fields:
func (_m *Query) Count() (int, error) {
	ret := _m.Called()

	var r0 int
	if rf, ok := ret.Get(0).(func() int); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountEstimate provides a mock function with given fields: threshold
func (_m *Query) CountEstimate(threshold int) (int, error) {
	ret := _m.Called(threshold)

	var r0 int
	if rf, ok := ret.Get(0).(func(int) int); ok {
		r0 = rf(threshold)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int) error); ok {
		r1 = rf(threshold)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateTable provides a mock function with given fields: opt
func (_m *Query) CreateTable(opt *v9orm.CreateTableOptions) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v9orm.CreateTableOptions) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DB provides a mock function with given fields: db
func (_m *Query) DB(db *pg.DB) orm.Query {
	ret := _m.Called(db)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(*pg.DB) orm.Query); ok {
		r0 = rf(db)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Delete provides a mock function with given fields: values
func (_m *Query) Delete(values ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(...interface{}) v9orm.Result); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Deleted provides a mock function with given fields:
func (_m *Query) Deleted() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Distinct provides a mock function with given fields:
func (_m *Query) Distinct() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// DistinctOn provides a mock function with given fields: expr, params
func (_m *Query) DistinctOn(expr string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, expr)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(expr, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// DropTable provides a mock function with given fields: opt
func (_m *Query) DropTable(opt *v9orm.DropTableOptions) error {
	ret := _m.Called(opt)

	var r0 error
	if rf, ok := ret.Get(0).(func(*v9orm.DropTableOptions) error); ok {
		r0 = rf(opt)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Except provides a mock function with given fields: other
func (_m *Query) Except(other orm.Query) orm.Query {
	ret := _m.Called(other)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(orm.Query) orm.Query); ok {
		r0 = rf(other)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// ExceptAll provides a mock function with given fields: other
func (_m *Query) ExceptAll(other orm.Query) orm.Query {
	ret := _m.Called(other)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(orm.Query) orm.Query); ok {
		r0 = rf(other)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// ExcludeColumn provides a mock function with given fields: columns
func (_m *Query) ExcludeColumn(columns ...string) orm.Query {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...string) orm.Query); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Exec provides a mock function with given fields: query, params
func (_m *Query) Exec(query interface{}, params ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) v9orm.Result); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecOne provides a mock function with given fields: query, params
func (_m *Query) ExecOne(query interface{}, params ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, ...interface{}) v9orm.Result); ok {
		r0 = rf(query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, ...interface{}) error); ok {
		r1 = rf(query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Exists provides a mock function with given fields:
func (_m *Query) Exists() (bool, error) {
	ret := _m.Called()

	var r0 bool
	if rf, ok := ret.Get(0).(func() bool); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// First provides a mock function with given fields:
func (_m *Query) First() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// For provides a mock function with given fields: s, params
func (_m *Query) For(s string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, s)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(s, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// ForEach provides a mock function with given fields: fn
func (_m *Query) ForEach(fn interface{}) error {
	ret := _m.Called(fn)

	var r0 error
	if rf, ok := ret.Get(0).(func(interface{}) error); ok {
		r0 = rf(fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ForceDelete provides a mock function with given fields: values
func (_m *Query) ForceDelete(values ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(...interface{}) v9orm.Result); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Group provides a mock function with given fields: columns
func (_m *Query) Group(columns ...string) orm.Query {
	_va := make([]interface{}, len(columns))
	for _i := range columns {
		_va[_i] = columns[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...string) orm.Query); ok {
		r0 = rf(columns...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// GroupExpr provides a mock function with given fields: group, params
func (_m *Query) GroupExpr(group string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, group)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(group, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Having provides a mock function with given fields: having, params
func (_m *Query) Having(having string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, having)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(having, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Insert provides a mock function with given fields: values
func (_m *Query) Insert(values ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(...interface{}) v9orm.Result); ok {
		r0 = rf(values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Intersect provides a mock function with given fields: other
func (_m *Query) Intersect(other orm.Query) orm.Query {
	ret := _m.Called(other)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(orm.Query) orm.Query); ok {
		r0 = rf(other)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// IntersectAll provides a mock function with given fields: other
func (_m *Query) IntersectAll(other orm.Query) orm.Query {
	ret := _m.Called(other)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(orm.Query) orm.Query); ok {
		r0 = rf(other)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Join provides a mock function with given fields: join, params
func (_m *Query) Join(join string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, join)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(join, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// JoinOn provides a mock function with given fields: condition, params
func (_m *Query) JoinOn(condition string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, condition)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(condition, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// JoinOnOr provides a mock function with given fields: condition, params
func (_m *Query) JoinOnOr(condition string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, condition)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(condition, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Last provides a mock function with given fields:
func (_m *Query) Last() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Limit provides a mock function with given fields: n
func (_m *Query) Limit(n int) orm.Query {
	ret := _m.Called(n)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(int) orm.Query); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Model provides a mock function with given fields: model
func (_m *Query) Model(model ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, model...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...interface{}) orm.Query); ok {
		r0 = rf(model...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// New provides a mock function with given fields:
func (_m *Query) New() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Offset provides a mock function with given fields: n
func (_m *Query) Offset(n int) orm.Query {
	ret := _m.Called(n)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(int) orm.Query); ok {
		r0 = rf(n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// OnConflict provides a mock function with given fields: s, params
func (_m *Query) OnConflict(s string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, s)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(s, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Order provides a mock function with given fields: orders
func (_m *Query) Order(orders ...string) orm.Query {
	_va := make([]interface{}, len(orders))
	for _i := range orders {
		_va[_i] = orders[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...string) orm.Query); ok {
		r0 = rf(orders...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// OrderExpr provides a mock function with given fields: order, params
func (_m *Query) OrderExpr(order string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, order)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(order, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Query provides a mock function with given fields: model, query, params
func (_m *Query) Query(model interface{}, query interface{}, params ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interface{}) v9orm.Result); ok {
		r0 = rf(model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// QueryOne provides a mock function with given fields: model, query, params
func (_m *Query) QueryOne(model interface{}, query interface{}, params ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, model, query)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(interface{}, interface{}, ...interface{}) v9orm.Result); ok {
		r0 = rf(model, query, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(interface{}, interface{}, ...interface{}) error); ok {
		r1 = rf(model, query, params...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Relation provides a mock function with given fields: name, apply
func (_m *Query) Relation(name string, apply ...func(orm.Query) (orm.Query, error)) orm.Query {
	_va := make([]interface{}, len(apply))
	for _i := range apply {
		_va[_i] = apply[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, name)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...func(orm.Query) (orm.Query, error)) orm.Query); ok {
		r0 = rf(name, apply...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Returning provides a mock function with given fields: s, params
func (_m *Query) Returning(s string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, s)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(s, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Select provides a mock function with given fields: values
func (_m *Query) Select(values ...interface{}) error {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...interface{}) error); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectAndCount provides a mock function with given fields: values
func (_m *Query) SelectAndCount(values ...interface{}) (int, error) {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(...interface{}) int); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectAndCountEstimate provides a mock function with given fields: threshold, values
func (_m *Query) SelectAndCountEstimate(threshold int, values ...interface{}) (int, error) {
	var _ca []interface{}
	_ca = append(_ca, threshold)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 int
	if rf, ok := ret.Get(0).(func(int, ...interface{}) int); ok {
		r0 = rf(threshold, values...)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(int, ...interface{}) error); ok {
		r1 = rf(threshold, values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SelectOrInsert provides a mock function with given fields: values
func (_m *Query) SelectOrInsert(values ...interface{}) (bool, error) {
	var _ca []interface{}
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 bool
	if rf, ok := ret.Get(0).(func(...interface{}) bool); ok {
		r0 = rf(values...)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(values...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Set provides a mock function with given fields: set, params
func (_m *Query) Set(set string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, set)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(set, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Table provides a mock function with given fields: tables
func (_m *Query) Table(tables ...string) orm.Query {
	_va := make([]interface{}, len(tables))
	for _i := range tables {
		_va[_i] = tables[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(...string) orm.Query); ok {
		r0 = rf(tables...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// TableExpr provides a mock function with given fields: expr, params
func (_m *Query) TableExpr(expr string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, expr)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(expr, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// TableModel provides a mock function with given fields:
func (_m *Query) TableModel() v9orm.TableModel {
	ret := _m.Called()

	var r0 v9orm.TableModel
	if rf, ok := ret.Get(0).(func() v9orm.TableModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.TableModel)
		}
	}

	return r0
}

// Union provides a mock function with given fields: other
func (_m *Query) Union(other orm.Query) orm.Query {
	ret := _m.Called(other)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(orm.Query) orm.Query); ok {
		r0 = rf(other)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// UnionAll provides a mock function with given fields: other
func (_m *Query) UnionAll(other orm.Query) orm.Query {
	ret := _m.Called(other)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(orm.Query) orm.Query); ok {
		r0 = rf(other)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Update provides a mock function with given fields: scan
func (_m *Query) Update(scan ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, scan...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(...interface{}) v9orm.Result); ok {
		r0 = rf(scan...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(scan...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateNotZero provides a mock function with given fields: scan
func (_m *Query) UpdateNotZero(scan ...interface{}) (v9orm.Result, error) {
	var _ca []interface{}
	_ca = append(_ca, scan...)
	ret := _m.Called(_ca...)

	var r0 v9orm.Result
	if rf, ok := ret.Get(0).(func(...interface{}) v9orm.Result); ok {
		r0 = rf(scan...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(v9orm.Result)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(...interface{}) error); ok {
		r1 = rf(scan...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Value provides a mock function with given fields: column, value, params
func (_m *Query) Value(column string, value string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, column, value)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, string, ...interface{}) orm.Query); ok {
		r0 = rf(column, value, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// Where provides a mock function with given fields: condition, params
func (_m *Query) Where(condition string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, condition)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(condition, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereGroup provides a mock function with given fields: fn
func (_m *Query) WhereGroup(fn func(orm.Query) (orm.Query, error)) orm.Query {
	ret := _m.Called(fn)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(func(orm.Query) (orm.Query, error)) orm.Query); ok {
		r0 = rf(fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereIn provides a mock function with given fields: where, slice
func (_m *Query) WhereIn(where string, slice interface{}) orm.Query {
	ret := _m.Called(where, slice)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, interface{}) orm.Query); ok {
		r0 = rf(where, slice)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereInMulti provides a mock function with given fields: where, values
func (_m *Query) WhereInMulti(where string, values ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, where)
	_ca = append(_ca, values...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(where, values...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereNotGroup provides a mock function with given fields: fn
func (_m *Query) WhereNotGroup(fn func(orm.Query) (orm.Query, error)) orm.Query {
	ret := _m.Called(fn)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(func(orm.Query) (orm.Query, error)) orm.Query); ok {
		r0 = rf(fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereOr provides a mock function with given fields: condition, params
func (_m *Query) WhereOr(condition string, params ...interface{}) orm.Query {
	var _ca []interface{}
	_ca = append(_ca, condition)
	_ca = append(_ca, params...)
	ret := _m.Called(_ca...)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, ...interface{}) orm.Query); ok {
		r0 = rf(condition, params...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereOrGroup provides a mock function with given fields: fn
func (_m *Query) WhereOrGroup(fn func(orm.Query) (orm.Query, error)) orm.Query {
	ret := _m.Called(fn)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(func(orm.Query) (orm.Query, error)) orm.Query); ok {
		r0 = rf(fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereOrNotGroup provides a mock function with given fields: fn
func (_m *Query) WhereOrNotGroup(fn func(orm.Query) (orm.Query, error)) orm.Query {
	ret := _m.Called(fn)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(func(orm.Query) (orm.Query, error)) orm.Query); ok {
		r0 = rf(fn)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WherePK provides a mock function with given fields:
func (_m *Query) WherePK() orm.Query {
	ret := _m.Called()

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func() orm.Query); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WhereStruct provides a mock function with given fields: strct
func (_m *Query) WhereStruct(strct interface{}) orm.Query {
	ret := _m.Called(strct)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(interface{}) orm.Query); ok {
		r0 = rf(strct)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// With provides a mock function with given fields: name, subq
func (_m *Query) With(name string, subq orm.Query) orm.Query {
	ret := _m.Called(name, subq)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, orm.Query) orm.Query); ok {
		r0 = rf(name, subq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WithDelete provides a mock function with given fields: name, subq
func (_m *Query) WithDelete(name string, subq orm.Query) orm.Query {
	ret := _m.Called(name, subq)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, orm.Query) orm.Query); ok {
		r0 = rf(name, subq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WithInsert provides a mock function with given fields: name, subq
func (_m *Query) WithInsert(name string, subq orm.Query) orm.Query {
	ret := _m.Called(name, subq)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, orm.Query) orm.Query); ok {
		r0 = rf(name, subq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WithUpdate provides a mock function with given fields: name, subq
func (_m *Query) WithUpdate(name string, subq orm.Query) orm.Query {
	ret := _m.Called(name, subq)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string, orm.Query) orm.Query); ok {
		r0 = rf(name, subq)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}

// WrapWith provides a mock function with given fields: name
func (_m *Query) WrapWith(name string) orm.Query {
	ret := _m.Called(name)

	var r0 orm.Query
	if rf, ok := ret.Get(0).(func(string) orm.Query); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(orm.Query)
		}
	}

	return r0
}
