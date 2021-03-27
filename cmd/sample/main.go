package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Account struct {
	Id			int
	Name		string
	IsAdmin		bool
}

func main() {
	db, err := NewDB()
	if err != nil {
		panic(err)
	}

	InsertAccount("admin", true, db)
	InsertAccount("ハマの番長", false, db)

	accounts, err := AllAccount(db)
	if err != nil {
		panic(err)
	}
	for _, account := range accounts{
		fmt.Println(account)
	}

	DeleteAllAccount(db)

	db.Close()
}

func NewDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost user=sue dbname=postgres password=postgres sslmode=disable")

	if err != nil {
		return nil, err
	}
	
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func InsertAccount(user string, isAdmin bool, db *sql.DB) {
	stmt, err := db.Prepare("INSERT INTO account(name, is_admin) VALUES($1, $2)")
	if err != nil {
		panic(err)
	}
	
	_, err = stmt.Exec(user, isAdmin)
	if err != nil {
		panic(err)
	}
}

func AllAccount(db *sql.DB)([]Account, error) {
	rows, err := db.Query("SELECT * FROM account");
	if err != nil {
		return nil, err
	}

	var accounts []Account
	for rows.Next() {
		var account Account
		rows.Scan(&account.Id, &account.Name, &account.IsAdmin)
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func DeleteAllAccount(db *sql.DB) {
	_, err := db.Exec("DELETE FROM account")
	if err != nil {
		panic(err)
	}
}