package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id    int
	Name  string
	Pass  string
	Mail  string
	Phone string
}

var db *sql.DB

// Открывает соединение с базой данных.
func OpenDB() {
	var err error
	db, err = sql.Open("sqlite3", "./database/db.db")
	if err != nil {
		log.Fatalf("Failed to open db: %s", err)
	}
}

// Функция задает sql-запрос к базе данных о наличии пользователя с таким логином и паролем, а затем возвращает булево значение.
func FindUserFromDB(username, password string) bool {
	u := User{}
	row := db.QueryRow("SELECT * FROM users WHERE name = ? AND pass = ?", username, password)
	err := row.Scan(&u.Id, &u.Name, &u.Pass, &u.Mail, &u.Phone)
	if err != nil {
		return false
	}
	return true
}

// Получаем срез пользователей из базы данных
func GetUsersFromDB() []User {
	var users []User

	rows, err := db.Query("SELECT * FROM Users")
	if err != nil {
		log.Fatalf("Failed to exec query: %s", err)
	}

	for rows.Next() {
		u := User{}
		err := rows.Scan(&u.Id, &u.Name, &u.Pass, &u.Mail, &u.Phone)
		if err != nil {
			log.Fatalf("Failed to scan row: %s", err)
		}
		users = append(users, u)
	}
	return users
}

// Получаем данные о пользователе по имени
func GetUserByNameFromDB(username string) (User, error) {
	u := User{}
	row := db.QueryRow("SELECT * FROM users WHERE name = ?", username)
	err := row.Scan(&u.Id, &u.Name, &u.Pass, &u.Mail, &u.Phone)
	if err != nil {
		return u, fmt.Errorf("Failed to scan user: %s", err)
	}
	return u, nil
}
