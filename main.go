package sqload

import (
	"bytes"
	"embed"
	"fmt"
	"io/fs"

	"github.com/pingcap/tidb/parser"
	"github.com/pingcap/tidb/parser/ast"
	_ "github.com/pingcap/tidb/parser/test_driver"
)

func Load(content *embed.FS) ([]string, error) {
	fileNames, err := getAllFilenames(content)
	if err != nil {
		fmt.Printf("getAllFileNames error: %s", err.Error())
		return nil, nil
	}

	var buf bytes.Buffer
	for _, name := range fileNames {
		file, err := content.ReadFile(name)
		if err != nil {
			return nil, err
		}
		buf.Write(file)
	}

	astNode, err := parse(buf.String())
	if err != nil {
		fmt.Printf("parse error: %v\n", err.Error())
		return nil, err
	}

	sqls := make([]string, len(astNode))
	for i, node := range astNode {
		sqls[i] = node.Text()
	}

	return sqls, nil
}

func parse(sql string) ([]ast.StmtNode, error) {
	p := parser.New()

	stmtNodes, _, err := p.Parse(sql, "", "")
	if err != nil {
		return nil, err
	}

	return stmtNodes, nil
}

func getAllFilenames(efs *embed.FS) (files []string, err error) {
	if err := fs.WalkDir(efs, ".", func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		files = append(files, path)

		return nil
	}); err != nil {
		return nil, err
	}

	return files, nil
}
