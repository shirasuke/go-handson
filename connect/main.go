package main

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

type Mydata struct{
	ID int
	Name string
	Mail string
	Age int
}

func (m *Mydata) Str() string{
	return "<\""+strconv.Itoa(m.ID)+":"+m.Name+"\""+m.Mail+","+strconv.Itoa(m.Age)+">"
} 

func main() {
	//.envに入れてできるようにする
	  dsn :=""
    db, err := sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal(err)
    }

    defer db.Close()

    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
}