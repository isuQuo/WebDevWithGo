package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}

func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected!")

	// Create a table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name TEXT,
			email TEXT UNIQUE NOT NULL
		);

		CREATE TABLE IF NOT EXISTS orders (
			id SERIAL PRIMARY KEY,
			user_id INT NOT NULL,
			amount INT,
			description TEXT
		);
	`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created")

	// Insert some data...
	/* name := "New user"
	email := "new@user2.com"
	var id int

	row := db.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		RETURNING id;
	`, name, email)

	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}

	fmt.Println("User created. id=", id) */

	// Query for a user id
	/* id := 2
	row := db.QueryRow(`
		SELECT name, email
		FROM users
		WHERE id=$1;`, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User information: user=%s, email=%s\n", name, email) */

	/* userId := 2
	for i := 1; i <= 5; i++ {
		amount := i * 100
		desc := fmt.Sprintf("Fake order #%d", i)
		_, err = db.Exec(`
			INSERT INTO orders(user_id, amount, description)
			VALUES($1, $2, $3)`, userId, amount, desc)
		if err != nil {
			panic(err)
		}

		fmt.Println("Created fake orders")
	} */

	type Order struct {
		ID     int
		UserId int
		Amount int
		Desc   string
	}

	var orders []Order
	userId := 2
	rows, err := db.Query(`
		SELECT id, amount, description
		FROM orders
		WHERE user_id=$1`, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserId = userId
		err := rows.Scan(&order.ID, &order.Amount, &order.Desc)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Orders:", orders)
}
