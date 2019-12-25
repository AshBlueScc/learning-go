package mysqlx

import (
	"database/sql"
	"fmt"
	_ "github.com/jmoiron/sqlx"
	"log"
)

var schema =
`CREATE TABLE if not EXISTS person (
	id int(20) auto_increment primary key,
    first_name varchar(20),
    last_name varchar(10),
    email varchar(20)
);`

var schema1 =
`CREATE TABLE if not EXISTS place (
	id int(20) auto_increment primary key,
    country varchar(20),
    city varchar(20),
    telcode int(10)
);`

type Person struct {
	Id int64
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

type Place struct {
	Id int64
	Country string
	City    sql.NullString
	TelCode int
}

func CreateTables()(){
	//MustExec一次只能执行一个语句，所以两个建表语句不能写在一个schema里面
	Db.MustExec(schema)
	Db.MustExec(schema1)
}

func InsertData()(){
	//事务
	tx := Db.MustBegin()
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "Jason", "Moiron", "jmoiron@jmoiron.net")
	tx.MustExec("INSERT INTO person (first_name, last_name, email) VALUES (?, ?, ?)", "John", "Doe", "johndoeDNE@gmail.net")
	tx.MustExec("INSERT INTO place (country, city, telcode) VALUES (?, ?, ?)", "United States", "New York", "1")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES (?, ?)", "Hong Kong", "852")
	tx.MustExec("INSERT INTO place (country, telcode) VALUES (?, ?)", "Singapore", "65")
	// Named queries can use structs, so if you have an existing struct (i.e. person := &Person{}) that you have populated, you can pass it in as &person
	tx.NamedExec("INSERT INTO person (id, first_name, last_name, email) VALUES (:id, :first_name, :last_name, :email)", &Person{10, "Jane", "Citizen", "jane.citzen@example.com"})
	tx.Commit()

	//查询数据库，用[]Person{}结构体来存
	people := []Person{}
	Db.Select(&people, "SELECT * FROM person ORDER BY first_name ASC")
	jason, john := people[0], people[1]

	fmt.Printf("%#v\n%#v", jason, john)
}

func SelectData()(){
	// You can also get a single result, a la QueryRow
	jason := Person{}
	err := Db.Get(&jason, "SELECT * FROM person WHERE first_name=?", "Jason")
	fmt.Printf("%#v\n", jason)
	// Person{FirstName:"Jason", LastName:"Moiron", Email:"jmoiron@jmoiron.net"}

	// if you have null fields and use SELECT *, you must use sql.Null* in your struct
	places := []Place{}
	err = Db.Select(&places, "SELECT * FROM place ORDER BY telcode ASC")
	if err != nil {
		fmt.Println(err)
		return
	}
	usa, singsing, honkers := places[0], places[1], places[2]

	fmt.Printf("%#v\n%#v\n%#v\n", usa, singsing, honkers)
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
	// Loop through rows using only one struct
	place := Place{}
	person := Person{}
	rows, _ := Db.Queryx("SELECT * FROM place")
	for rows.Next() {
		err := rows.StructScan(&place)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", place)
	}
	// Place{Country:"United States", City:sql.NullString{String:"New York", Valid:true}, TelCode:1}
	// Place{Country:"Hong Kong", City:sql.NullString{String:"", Valid:false}, TelCode:852}
	// Place{Country:"Singapore", City:sql.NullString{String:"", Valid:false}, TelCode:65}

	// Named queries, using `:name` as the bindvar.  Automatic bindvar support
	// which takes into account the dbtype based on the driverName on sqlx.Open/Connect
	_, _ = Db.NamedExec(`INSERT INTO person (id, first_name,last_name,email) VALUES (:id, :first,:last,:email)`,
		map[string]interface{}{
			"id": "20",
			"first": "Bin",
			"last": "Smuth",
			"email": "bensmith@allblacks.nz",
		})

	// Selects Mr. Smith from the database
	rows, _ = Db.NamedQuery(`SELECT * FROM person WHERE first_name=:fn`, map[string]interface{}{"fn": "Bin"})
	for rows.Next() {
		err := rows.StructScan(&person)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", person)
	}

	// Named queries can also use structs.  Their bind names follow the same rules
	// as the name -> db mapping, so struct fields are lowercased and the `db` tag
	// is taken into consideration.
	rows, _ = Db.NamedQuery(`SELECT * FROM person WHERE first_name=:first_name`, jason)
	for rows.Next() {
		err := rows.StructScan(&person)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("%#v\n", person)
	}
}