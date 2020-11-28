package main

import (
	"database/sql"
	"functionsInDB/DB"
	"functionsInDB/functions"
	"functionsInDB/modules"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "Bank")
	if err != nil {
		log.Fatal("DB is not connected.", err)
	}

	defer db.Close()

	err = DB.UsersInit(db)
	if err != nil {
		log.Fatal("Error while initializing users table", err)
	}

	err = DB.AccountsInit(db)
	if err != nil {
		log.Fatal("Error while initializing accounts table", err)
	}

	err = DB.CurrencyInit(db)
	if err != nil {
		log.Fatal("Error while initializing currency table", err)
	}

	newUser := modules.Users{
		ID:       0,
		Name:     "Nekruz",
		Surname:  "Rakhimov",
		Age:      18,
		Sex:      "M",
		Login:    "Nekruz",
		Password: "Nekruz",
		Remove:   false,
	}

	err = functions.AddNewUser(db, newUser)
	if err != nil {
		log.Fatal(err)
	}

	newAccount := modules.Accounts{
		ID:       0,
		UserID:   1,
		Number:   "5555 5555 5555 5555",
		Amount:   1500,
		Currency: "1",
	}

	err = functions.AddNewAccount(db, newAccount)
	if err != nil {
		log.Fatal(err)
	}

	updatingUser := modules.Users{
		ID:       0,
		Name:     "Ilhom",
		Surname:  "Rakhimov",
		Age:      15,
		Sex:      "M",
		Login:    "Ilhom",
		Password: "Ilhom",
		Remove:   false,
	}
	err = functions.UpdateUser(db, updatingUser)
	if err != nil {
		log.Fatal(err)
	}

	updatingAccount := modules.Accounts{
		ID:       0,
		UserID:   5,
		Number:   "5558 2558 2655 2554",
		Amount:   5000,
		Currency: "2",
		Remove:   false,
	}

	err = functions.UpdateAccount(db, updatingAccount)
	if err != nil {
		log.Fatal(err)
	}

	err = functions.RemoveUser(db, 3)
	if err != nil {
		log.Fatal(err)
	}

	err = functions.RemoveAccount(db, 4)
	if err != nil {
		log.Fatal(err)
	}

	err = functions.RestoreUser(db, 2)
	if err != nil {
		log.Fatal(err)
	}

	err = functions.RestoreAccount(db, 1)
	if err != nil {
		log.Fatal(err)
	}



}



