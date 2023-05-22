package sqldb

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Skill struct {
	Id   int
	Name string
}

func connect_to_mysql() *sql.DB {
	psqlUrl := "edwica_root:b00m5gQ40WB1@tcp(83.220.175.75:3306)/edwica"
	db, err := sql.Open("mysql", psqlUrl)
	checkErr(err)

	return db
}

func GetNoneSkills() (skills []Skill) {
	db := connect_to_mysql()
	defer db.Close()
	results, err := db.Query("SELECT id, name FROM demand WHERE is_displayed is NULL")
	checkErr(err)
	for results.Next() {
		var name string
		var id int
		err = results.Scan(&id, &name)
		skills = append(skills, Skill{Id: id, Name: name})
		checkErr(err)
	}
	return
}

func RemoveSkills(skills []Skill) {
	db := connect_to_mysql()
	defer db.Close()
	var values []interface{}

	for _, item := range skills {
		values = append(values, item)
	}
	_, err := db.Exec(createRemoveQuery(skills))
	checkErr(err)

}

func createRemoveQuery(ids []Skill) string {
	var items []string
	for _, i := range ids {
		items = append(items, strconv.Itoa(i.Id))
	}
	query := fmt.Sprintf("DELETE FROM demand WHERE id IN (%s)", strings.Join(items, ","))
	return query
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
