package sqload

import (
	"bytes"
	"embed"
	"testing"

	_ "github.com/pingcap/tidb/parser/test_driver"
	"github.com/uh-zz/sqload/driver/mysql"
)

//go:embed sql/*
var content embed.FS

func TestLoader_Load(t *testing.T) {
	var sqlfile bytes.Buffer

	type fields struct {
		d Dialector
	}
	type args struct {
		content *embed.FS
		to      *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "mysql/read all sql file",
			fields: fields{
				d: mysql.Dialector{},
			},
			args: args{
				content: &content,
				to:      &sqlfile,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Loader{
				d: tt.fields.d,
			}
			if err := l.Load(tt.args.content, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Loader.Load() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}