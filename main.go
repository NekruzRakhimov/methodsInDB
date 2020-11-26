package main

import (
	"database/sql"
	"fmt"
	"functionsInDB/DB"
	"functionsInDB/modules"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "Bank")
	if err != nil {
		log.Fatal("DB is not connected.", err)
	}

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

	newAccount := modules.Accounts{
		ID:       0,
		UserID:   1,
		Number:   "5555 5555 5555 5555",
		Amount:   1500,
		Currency: "1",
	}

	updatingAccount := modules.Accounts{
		ID:       0,
		UserID:   5,
		Number:   "5558 2558 2655 2554",
		Amount:   5000,
		Currency: "2",
		Remove:   false,
	}

	err = AddNewUser(db, newUser)
	if err != nil {
		log.Fatal(err)
	}

	err = AddNewAccount(db, newAccount)
	if err != nil {
		log.Fatal(err)
	}

	err = UpdateUser(db, updatingUser)
	if err != nil {
		log.Fatal(err)
	}

	err = RemoveUser(db)
	if err != nil {
		log.Fatal(err)
	}

	err = RemoveAccount(db)
	if err != nil {
		log.Fatal(err)
	}

	err = RestoreUser(db)
	if err != nil {
		log.Fatal(err)
	}

	err = UpdateAccount(db, updatingAccount)
	if err != nil {
		log.Fatal(err)
	}

}

func AddNewUser(db *sql.DB, user modules.Users) (err error) {
	_, err = db.Exec("INSERT INTO users(name, surname, age, sex, login, password, remove) values(($1), ($2), ($3), ($4), ($5), ($6), ($7))", user.Name, user.Surname, user.Age, user.Sex, user.Login, user.Password, user.Remove)

	if err != nil {
		fmt.Println("Can't insert to DB", err)
	}

	return err
}

func AddNewAccount(db *sql.DB, account modules.Accounts) (err error) {
	_, err = db.Exec("INSERT INTO accaunts(userId, number, amount, currency, remove) values(($1), ($2), ($3), ($4), ($5))", account.UserID, account.Number, account.Amount, account.Currency, account.Remove)

	if err != nil {
		fmt.Println("Can't insert to DB", err)
	}
	return err
}

func RemoveUser(db *sql.DB) (err error) {
	_, err = db.Exec(`UPDATE users SET remove = TRUE WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func RemoveAccount(db *sql.DB) (err error) {
	_, err = db.Exec(`UPDATE accaunts SET remove = TRUE WHERE id = 1`)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func RestoreUser(db *sql.DB) (err error) {
	_, err = db.Exec(`UPDATE users SET remove = FALSE WHERE id = 5`)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func RestoreAccount(db *sql.DB) (err error) {
	_, err = db.Exec(`UPDATE accaunts SET remove = FALSE WHERE id = 5`)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UpdateUser(db *sql.DB, user modules.Users) (err error) {
	_, err = db.Exec("UPDATE users SET (name, surname, age, sex, login, password, remove) = (($1), ($2), ($3), ($4), ($5), ($6), ($7)) WHERE id = 1", user.Name, user.Surname, user.Age, user.Sex, user.Login, user.Password, user.Remove)

	if err != nil {
		fmt.Println("Can't update DB", err)
	}

	return err
}

func UpdateAccount(db *sql.DB, account modules.Accounts) (err error) {
	_, err = db.Exec("UPDATE accaunts SET (userId, number, amount, currency, remove) = (($1), ($2), ($3), ($4), ($5)) WHERE id = 1", account.UserID, account.Number, account.Amount, account.Currency, account.Remove)

	if err != nil {
		fmt.Println("Can't update DB", err)
	}

	return err
}

