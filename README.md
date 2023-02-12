# sqload

sqload reads the SQL file embedded in the Go binary and converts the statements into a string slice.

It also checks that the statements it reads are valid.

## Features

- [x] Parse statment in MySQL
- [x] Parse statment in PostgreSQL

## Install

```
go get github.com/uh-zz/sqload@latest
```

## Getting started

### Pattern 1: Reading from a single file

```
.
├── go.mod
├── go.sum
├── main.go
└── sql
    └── test.sql
```

`sql/test.sql`

```sql
INSERT INTO table001 (name,age) VALUES ('alice', 10);
```

You can also add multiple sqls in the same file.

```sql
INSERT INTO table001 (name,age) VALUES ('alice', 10);
INSERT INTO table001 (name,age) VALUES ('bob', 10);
```

`main.go`

```go
package main

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/uh-zz/sqload"
	"github.com/uh-zz/sqload/driver/mysql"
)

//go:embed sql/*
var content embed.FS

func main() {
	var (
		buf  bytes.Buffer // sql which read from file
		sqls []string // sql after parse
	)

	loader := sqload.New(mysql.Dialector{}) // for PostgreSQL: postgresql.Dialector{}

	if err := loader.Load(&content, &buf); err != nil {
		fmt.Printf("Load error: %s", err.Error())
	}

	if err := loader.Parse(buf.String(), &sqls); err != nil {
		fmt.Printf("Parse error: %s", err.Error())
	}

	fmt.Printf("%+v", sqls)
    // [INSERT INTO table001 (name,age) VALUES ('alice', 10);]
}
```

### Pattern 2: Reading from a multiple files

```
.
├── go.mod
├── go.sum
├── main.go
└── sql
    ├── test.sql
    └── test_other.sql
```

`sql/test.sql`

```sql
INSERT INTO table001 (name,age) VALUES ('alice', 10);
```

`sql/test_other.sql`

```sql
INSERT INTO table001 (name,age) VALUES ('bob', 10);
```

`main.go`

```go
package main

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/uh-zz/sqload"
	"github.com/uh-zz/sqload/driver/mysql"
)

//go:embed sql/*
var content embed.FS

func main() {
	var (
		buf  bytes.Buffer // sql which read from file
		sqls []string // sql after parse
	)

	loader := sqload.New(mysql.Dialector{}) // for PostgreSQL: postgresql.Dialector{}

	if err := loader.Load(&content, &buf); err != nil {
		fmt.Printf("Load error: %s", err.Error())
	}

	if err := loader.Parse(buf.String(), &sqls); err != nil {
		fmt.Printf("Parse error: %s", err.Error())
	}

	fmt.Printf("%+v", sqls)
    // [INSERT INTO table001 (name,age) VALUES ('alice', 10); INSERT INTO table001 (name,age) VALUES ('bob', 20);]
}
```

Or you can choose to load a file.

`main.go`

```go
package main

import (
	"bytes"
	"embed"
	"fmt"

	"github.com/uh-zz/sqload"
	"github.com/uh-zz/sqload/driver/mysql"
)

//go:embed sql/*
var content embed.FS

func main() {
	var (
		buf  bytes.Buffer // sql which read from file
		sqls []string // sql after parse
	)

	loader := sqload.New(mysql.Dialector{}) // for PostgreSQL: postgresql.Dialector{}

	if err := loader.LoadFrom(&content, &buf, "sql/test.sql"); err != nil {
		fmt.Printf("Load error: %s", err.Error())
	}

	if err := loader.Parse(buf.String(), &sqls); err != nil {
		fmt.Printf("Parse error: %s", err.Error())
	}

	fmt.Printf("%+v", sqls)
    // [INSERT INTO table001 (name,age) VALUES ('alice', 10);]
}
```

## Contributing

Guidelines are being prepared, but Contributions are welcomed and greatly appreciated.

## Todo

- [ ] other dialect
