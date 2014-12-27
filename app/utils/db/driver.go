package db

import (
	"database/sql"

	schm "github.com/kaneshin/snacky-go/app/entities/schemata"

	"github.com/coopernurse/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type Driver struct {
	Dbm *gorp.DbMap
}

func InitDB() *Driver {
	// initialize the DbMap
	dbm, err := initDB()
	if err != nil {
		panic(err)
	}
	driver := Driver{}
	driver.Dbm = dbm
	return &driver
}

func initDB() (*gorp.DbMap, error) {
	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/dev_snacky")
	if err != nil {
		return nil, err
	}

	// construct a gorp DbMap
	dbm := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{Engine: "InnoDB", Encoding: "UTF8"}}

	// add a table, setting the table name to 'posts' and
	// specifying that the Id property is an auto incrementing PK
	dbm.AddTableWithName(schm.Snack{}, "snacks")
	dbm.AddTableWithName(schm.Tag{}, "tags")
	dbm.AddTableWithName(schm.Kind{}, "kinds")
	dbm.AddTableWithName(schm.Maker{}, "makers")
	dbm.AddTableWithName(schm.Area{}, "areas")
	dbm.AddTableWithName(schm.SnackTag{}, "snack_tag_relations").SetKeys(true, "Id")

	// create the table. in a production system you'd generally
	// use a migration tool, or create the tables via scripts
	err = dbm.CreateTablesIfNotExists()
	if err != nil {
		return nil, err
	}

	return dbm, nil
}
