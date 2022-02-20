package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	_ "github.com/lib/pq"
)

func psql() *sql.DB {
//	connStr := "postgres://artem:1@172.18.0.1:5432/sales?sslmode=disable"
connStr := "user=artem password=1 dbname=sales sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err) //psql err
	}
	return db
}
func addPsql(mas data) {
	db := psql()
	defer db.Close()
	_, err := db.Exec("insert into sales (model, price) values ($1, $2)", mas.name, mas.price)
	if err != nil {
		fmt.Println(err) //add psql err
	}
}

type data struct {
	//id    int
	name  string
	price int
}

func Data() (mas []data) {
	a := data{name: "samsung", price: 50000}
	b := data{name: "apple", price: 70000}
	mas = []data{a, b}
	return mas
}

//create table
func Table(mas []data) (table []data) {
	for i := 1; i <= 10; i++ {
		t := mas[rand.Intn(len(mas))]
		//t.id = i
		//fmt.Println("t.id=", t)
		table = append(table, t)
	}
	return table
}

//writing to the table
func gorun() (p data) {
	t := Table(Data())
	for _, p := range t {
		addPsql(p)
		//fmt.Printf("gorun %s\n", p)
	}
	return p
}
func main() {
	http.HandleFunc("/", add)
	err := http.ListenAndServe(":9990", nil)
	if err != nil {
		log.Fatal(err) //http error
	}

}

func add(w http.ResponseWriter, r *http.Request) {
	p := gorun()
	fmt.Fprintf(w, "writed to sales")
	fmt.Println(p)
}
