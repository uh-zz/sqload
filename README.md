# sqload

sqload reads the SQL file embedded in the Go binary and converts the statements into a string slice.

It also checks that the statements it reads are valid.

## usage

```go
//go:embed sql/*
var content embed.FS

loader := sqload.New(mysql.Dialector{})
sqlfile, _ := loader.Load(&content)
sqls, _ := loader.Parse(sqlfile)

fmt.Printf("debug: %+v", sqls)
```
