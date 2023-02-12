package postgresql

import (
	"github.com/auxten/postgresql-parser/pkg/sql/parser"
)

type Dialector struct{}

func (d Dialector) Name() string {
	return "postgresql"
}

func (d Dialector) Parse(sqlfile string, to *[]string) error {
	statements, err := parser.Parse(sqlfile)
	if err != nil {
		return err
	}

	sqls := make([]string, len(statements))
	for i, statment := range statements {
		sqls[i] = statment.SQL
	}

	*to = sqls

	return nil
}
