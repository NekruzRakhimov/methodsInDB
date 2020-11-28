package functions

import (
	"database/sql"
	"functionsInDB/modules"
	"log"
)

func AddNewUser(db *sql.DB, user modules.Users) (err error) {
	_, err = db.Exec("INSERT INTO users(name, surname, age, sex, login, password, remove) values(($1), ($2), ($3), ($4), ($5), ($6), ($7))", user.Name, user.Surname, user.Age, user.Sex, user.Login, user.Password, user.Remove)

	if err != nil {
		log.Fatal("Can't insert to DB", err)
	}
	return err
}

func AddNewAccount(db *sql.DB, account modules.Accounts) (err error) {
	_, err = db.Exec("INSERT INTO accaunts(userId, number, amount, currency, remove) values(($1), ($2), ($3), ($4), ($5))", account.UserID, account.Number, account.Amount, account.Currency, account.Remove)

	if err != nil {
		log.Fatal("Can't insert to DB", err)
	}
	return err
}

func RemoveUser(db *sql.DB, id int64) (err error) {
	_, err = db.Exec(`UPDATE users SET remove = TRUE WHERE id = ($1)`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func RemoveAccount(db *sql.DB, id int64) (err error) {
	_, err = db.Exec(`UPDATE accaunts SET remove = TRUE WHERE id = ($1)`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func RestoreUser(db *sql.DB, id int64) (err error) {
	_, err = db.Exec(`UPDATE users SET remove = FALSE WHERE id = ($1)`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func RestoreAccount(db *sql.DB, id int64) (err error) {
	_, err = db.Exec(`UPDATE accaunts SET remove = FALSE WHERE id = ($1)`, id)
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func UpdateUser(db *sql.DB, user modules.Users) (err error) {
	_, err = db.Exec("UPDATE users SET (name, surname, age, sex, login, password, remove) = (($1), ($2), ($3), ($4), ($5), ($6), ($7)) WHERE id = 1", user.Name, user.Surname, user.Age, user.Sex, user.Login, user.Password, user.Remove)

	if err != nil {
		log.Fatal("Can't update DB", err)
	}
	return err
}

func UpdateAccount(db *sql.DB, account modules.Accounts) (err error) {
	_, err = db.Exec("UPDATE accaunts SET (userId, number, amount, currency, remove) = (($1), ($2), ($3), ($4), ($5)) WHERE id = 1", account.UserID, account.Number, account.Amount, account.Currency, account.Remove)

	if err != nil {
		log.Fatal("Can't update DB", err)
	}

	return err
}
