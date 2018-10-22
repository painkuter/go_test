package main

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"fmt"
)

// Connect ...
func Connect() *pg.DB {
	return pg.Connect(&pg.Options{
		User:     "books",
		Password: "111",
		Addr:     ":5433",
		Database: "books",
	})
}

func CreateSchema(db *pg.DB, temp, ifNotExists bool) error {

	for _, model := range []interface{}{
		// (*Genre)(nil),
		// (*Image)(nil),
		// (*Author)(nil),
		// (*BookGenre)(nil),
		// (*Book)(nil),
		// (*BookWithCommentCount)(nil),
		// (*Translation)(nil),
		// (*Comment)(nil),

		(*Company)(nil),
		(*User)(nil),
		(*CompanyUser)(nil),

	} {
		err := db.CreateTable(model, &orm.CreateTableOptions{
			Temp:        temp,
			IfNotExists: ifNotExists,
			FKConstraints: true,
		})
		if err != nil {
			fmt.Println(err)
		}
	}
	return nil
}
