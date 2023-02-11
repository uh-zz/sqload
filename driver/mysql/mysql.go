package mysql

import "github.com/pingcap/tidb/parser"

type Dialector struct{}

func (d Dialector) Name() string {
	return "mysql"
}

func (d Dialector) Parse(sqlfile string) ([]string, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sqlfile, "", "")
	if err != nil {
		return nil, err
	}

	sqls := make([]string, len(stmtNodes))
	for i, node := range stmtNodes {
		sqls[i] = node.Text()
	}

	return sqls, nil
}
