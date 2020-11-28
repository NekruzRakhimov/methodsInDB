package DB

import (
	"database/sql"
)

func UsersInit(db *sql.DB) (err error) {//initializing users table
	const DDLusersDB = `CREATE TABLE if not exists users (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text not null,
    surname TEXT NOT NULL,
    age INTEGER NOT NULL,
    sex TEXT,
    login TEXT,
    password TEXT,
    remove BOOLEAN NOT NULL DEFAULT FALSE
	)`

	_, err = db.Exec(DDLusersDB)

	if err != nil {
		return err
	}
	return nil
}

func AccountsInit(db *sql.DB) (err error) {//initializing accounts table
	const DDLaccaunts  = `CREATE TABLE if not exists accaunts (
   	id integer PRIMARY KEY AUTOINCREMENT,
   	userId integer references users(id) not null ,
   	number integer,
   	amount integer,
   	currency integer references currency(id),
   	remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(DDLaccaunts)

	if err != nil {
		return err
	}
	return nil
}

func CurrencyInit(db *sql.DB) (err error) {//initializing currency table
	const DDLcurrency = `CREATE TABLE if not exists currency (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text
	)`
	_, err = db.Exec(DDLcurrency)

	if err != nil {
		return err
	}
	return nil
}



