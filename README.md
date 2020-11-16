# go-pg-wrapper

_This project was originally from https://gitlab.com/slax0rr/go-pg-wrapper_. I just forked it from gitlab and upgrade to use go-pg v10

A wrapper around [go-pg](https://github.com/go-pg/pg) providing functions that
use interfaces, for simpler unit testing.

## Usage

`go get -u github.com/cychiuae/go-pg-wrapper/v10`

Connect to your database using `go-pg` as you normally would, and pass the
`*pg.DB` to `pgwrapper.NewDB()`:

```go
package main

import (
	pgwrapper "github.com/cychiuae/go-pg-wrapper/v10"
	"github.com/go-pg/pg/v10"
)

func main() {
	db := pgwrapper.NewDB(pg.Connect(&pg.Options{
		// ...
	}))

    // use db as before
}
```

### Tests

Mocks are also provided implementing the
[testify mock](https://godoc.org/github.com/stretchr/testify/mock) package for
your tests:

```go
package foo

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/assert"
	"github.com/cychiuae/go-pg-wrapper/v10/mocks"
)

func TestFoo(t *testing.T) {
	expObj := &FooObj{
		// ...
	}

	db := new(mocks.DB)
	db.On("Model", mock.Anything).Return(db)
	db.On("Select").Return(expObj, nil)

	foo, err := Foo(db)
	assert.Nil(t, err)
	assert.Equal(t, expObj, foo)
}
```
