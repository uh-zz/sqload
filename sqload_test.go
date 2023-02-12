package sqload

import (
	"bytes"
	"embed"
	"testing"

	_ "github.com/pingcap/tidb/parser/test_driver"
	"github.com/uh-zz/sqload/driver/mysql"
	"github.com/uh-zz/sqload/driver/postgresql"
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

func TestLoader_Parse(t *testing.T) {
	var sqls []string
	type fields struct {
		d Dialector
	}
	type args struct {
		sqlfile string
		to      *[]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "mysql/parse from valid sql string",
			fields: fields{
				d: mysql.Dialector{},
			},
			args: args{
				sqlfile: "CREATE USER 'jeffrey'@'localhost' IDENTIFIED WITH mysql_native_password;",
				to:      &sqls,
			},
			wantErr: false,
		},
		{
			name: "mysql/parse statement in postgres",
			fields: fields{
				d: mysql.Dialector{},
			},
			args: args{
				sqlfile: "CREATE USER jeffrey WITH PASSWORD 'postgres_native_password';",
				to:      &sqls,
			},
			wantErr: true,
		},
		{
			name: "postgresql/parse from valid sql string",
			fields: fields{
				d: postgresql.Dialector{},
			},
			args: args{
				sqlfile: "CREATE USER jeffrey WITH PASSWORD 'postgres_native_password';CREATE USER jeffrey WITH PASSWORD 'postgres_native_password';",
				to:      &sqls,
			},
			wantErr: false,
		},
		{
			name: "postgresql/parse from statement in mysql",
			fields: fields{
				d: postgresql.Dialector{},
			},
			args: args{
				sqlfile: "CREATE USER 'jeffrey'@'localhost' IDENTIFIED WITH mysql_native_password;",
				to:      &sqls,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Loader{
				d: tt.fields.d,
			}
			if err := l.Parse(tt.args.sqlfile, tt.args.to); (err != nil) != tt.wantErr {
				t.Errorf("Loader.Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoader_LoadFrom(t *testing.T) {
	var sqlfile bytes.Buffer

	type fields struct {
		d Dialector
	}
	type args struct {
		content   *embed.FS
		to        *bytes.Buffer
		fileNames []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "mysql/read specified sql file",
			fields: fields{
				d: mysql.Dialector{},
			},
			args: args{
				content:   &content,
				to:        &sqlfile,
				fileNames: []string{"sql/update_statement.sql"},
			},
			wantErr: false,
		},
		{
			name: "mysql/read nonexistent sql file",
			fields: fields{
				d: mysql.Dialector{},
			},
			args: args{
				content:   &content,
				to:        &sqlfile,
				fileNames: []string{"sql/nonexistent_update_statement.sql"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := Loader{
				d: tt.fields.d,
			}
			if err := l.LoadFrom(tt.args.content, tt.args.to, tt.args.fileNames...); (err != nil) != tt.wantErr {
				t.Errorf("Loader.LoadFrom() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
