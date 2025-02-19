package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

type User struct {
	Username   string
	Email      string
	Password   string
	Avatar     string
	Bio        string
	Created_at time.Time
}

func main() {
	connStr := "postgres://postgres:Smetana18@localhost:5432/gopg?sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	createUser(db)
	user := User{"klepa", "bankaiof@gmail.com", "Smetana18", "", "", time.Now()}
	pk := insertUser(db, user)
	fmt.Printf("ID = %d\n", pk)
}

func createUser(db *sql.DB) {
	/*
	   Table users Twitch
	   id	            UUID (PK)	          Уникальный идентификатор
	   username	    VARCHAR(50)	          Уникальное имя пользователя
	   email	        VARCHAR(255)	      Почта (уникальная)
	   password	    VARCHAR(255)	      Захешированный пароль
	   avatar	        TEXT	              Ссылка на аватар
	   bio	            TEXT	              Описание профиля
	   created_at	    TIMESTAMP	          Дата регистрации
	   updated_at	    TIMESTAMP	          Дата обновления профиля

	*/
	query := `CREATE TABLE IF NOT EXISTS users (
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(50) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    avatar      TEXT,
    bio         TEXT,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertUser(db *sql.DB, user User) int {
	query := `INSERT INTO users(username, email, password, avatar, bio, created_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`

	var pk int
	err := db.QueryRow(query, user.Username, user.Email, user.Password, user.Avatar, user.Bio, user.Created_at).Scan(&pk)
	if err != nil {
		log.Fatal(err)
	}
	return pk
}
