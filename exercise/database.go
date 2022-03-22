package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"time"
)

func main() {
	fmt.Println("Go SQL")
	dbHost := "159.223.42.164"
	dbPort := "5432"
	dbName := "enigma"
	dbUser := "postgres"
	dbPassword := "P@ssw0rd"

	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	println(dataSourceName)
	conn, err := sqlx.Connect("pgx", dataSourceName)
	if err != nil {
		fmt.Println("Unable to connect to database: %v\n", err)
		panic(err)
	}
	defer conn.DB.Close()

	customers := []Customers{}
	err = conn.Select(&customers, "SELECT * FROM m_customers order by id")
	if err != nil {
		fmt.Println(err)
	}
	for _, customer := range customers {
		fmt.Println(customer)
	}
	//rows, err := conn.Query(context.Background(), "SELECT id,first_name FROM m_customers order by id")
	//if err != nil {
	//	log.Println(err)
	//}
	//for rows.Next() {
	//	customers := Customers{}
	//	err := rows.Scan(&customers.Id, &customers.FirstName)
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Println(customers)
	//}

	tx := conn.MustBegin()
	_, err = tx.Exec("insert into m_customers(id) values($1)", "C003")
	if err != nil {
		panic(err)
	}
	id := "coo4"
	nama := "joko"
	_, err = tx.Exec("insert into m_customers(id,first_name) values($1,$2)", id, nama)
	if err != nil {
		panic(err)
	}
	panic("Disconnected")
	tx.Commit()

	//_, err = conn.Exec(context.Background(), "insert into m_customers(id) values($1)", "C001")
	//if err != nil {
	//	panic(err)
	//}

	//rows, err := conn.Query(context.Background(), "SELECT count(id) FROM m_customers")
	//if err != nil {
	//	log.Println(err)
	//}
	//for rows.Next() {
	//	totalCustomer := TotalCustomer{}
	//	err := rows.Scan(&totalCustomer.Total)
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Println(totalCustomer)
	//}
}

type TotalCustomer struct {
	Total int `db:"total"`
}

//struct tag
type Customers struct {
	Id        string
	FirstName string    `db:"first_name"`
	LastName  string    `db:"last_name"`
	BirtDate  time.Time `db:"birth_date"`
	Address   string
	Status    int
	Username  sql.NullString
	Password  sql.NullString
	Email     sql.NullString
}
