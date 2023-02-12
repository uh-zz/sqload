package sqload

import (
	"bytes"
	"embed"
	"io/fs"

	_ "github.com/pingcap/tidb/parser/test_driver"
)

type Loader struct {
	d Dialector
}

func New(dialector Dialector) Loader {
	return Loader{d: dialector}
}

func (l Loader) Load(content *embed.FS, to *bytes.Buffer) error {
	fileNames, err := getAllFilenames(content)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	for _, name := range fileNames {
		file, err := content.ReadFile(name)
		if err != nil {
			return err
		}
		buf.Write(file)
	}

	*to = buf

	return nil
}

func (l Loader) Parse(sqlfile string) ([]string, error) {
	var (
		sqls []string
		err  error
	)

	if dialector, ok := l.d.(interface {
		Parse(string) ([]string, error)
	}); ok {
		if sqls, err = dialector.Parse(sqlfile); err != nil {
			return nil, err
		}
	}
	return sqls, nil
}

func getAllFilenames(efs *embed.FS) ([]string, error) {
	var files []string
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
