package storage

import (
	"github.com/SteMak/vanilla/out"
)

type table interface {
	name() string
	createTable() error
}

var tables []table

func init() {
	tables = append(tables,
		&Users,
	)
}

func Migrate() {
	out.Infoln("Migrations started")
	for _, table := range tables {
		out.Infof("Table %s... ", table.name())
		if err := table.createTable(); err != nil {
			out.Infoln("[FAIL]")
			out.Fatal(err)
		}
		out.Infoln("[OK]")
	}
	out.Infoln("Migrations ended")
}