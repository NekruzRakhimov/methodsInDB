package DB

import (
	"database/sql"
	"fmt"
)

func init() {
	fmt.Println("Hello")
}

func UsersInit(db *sql.DB) (err error) {//initializing users table
	const usersDB = `CREATE TABLE if not exists users (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text not null,
    surname TEXT NOT NULL,
    age INTEGER NOT NULL,
    sex TEXT,
    login TEXT,
    password TEXT,
    remove BOOLEAN NOT NULL DEFAULT FALSE
	)`

	_, err = db.Exec(usersDB)

	if err != nil {
		return err
	}
	return nil
}

func AccountsInit(db *sql.DB) (err error) {//initializing accounts table
	const accaunts  = `CREATE TABLE if not exists accaunts (
   	id integer PRIMARY KEY AUTOINCREMENT,
   	userId integer references users(id) not null ,
   	number integer,
   	amount integer,
   	currency integer references currency(id),
   	remove BOOLEAN NOT NULL DEFAULT FALSE
)`
	_, err = db.Exec(accaunts)

	if err != nil {
		return err
	}
	return nil
}

func CurrencyInit(db *sql.DB) (err error) {//initializing currency table
	const currency = `CREATE TABLE if not exists currency (
    id integer PRIMARY KEY AUTOINCREMENT,
    name text
	)`
	_, err = db.Exec(currency)

	if err != nil {
		return err
	}
	return nil
}



