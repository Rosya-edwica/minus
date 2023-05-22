package sqldb

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

func connect_to_postgres() *sql.DB {
	psqlUrl := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "94.250.253.88", 5432, "edwica_root", "9k35XQ&s", "edwica")
	db, err := sql.Open("postgres", psqlUrl)
	checkErr(err)
	return db
}

// Здесь хранится основная база минусов слов. Сравниваем хранилище минус слов с необработанными навыками
func GetMinusSkillsFromPostgres() (skills []string) {
	db := connect_to_postgres()
	rows, err := db.Query("SELECT name FROM minus_skill")
	checkErr(err)
	defer rows.Close()

	for rows.Next() {
		var name string

		err = rows.Scan(&name)
		checkErr(err)
		skills = append(skills, name)
	}
	return
}

func SaveMinusSkills(skills []Skill) {
	db := connect_to_postgres()
	defer db.Close()
	_, err := db.Exec(createInsertQuery(skills))
	if err != nil {
		fmt.Println(err)
	}
}

func createInsertQuery(skills []Skill) string {
	names := []string{}
	for _, skill := range skills {
		names = append(names, "('"+strings.ToLower(skill.Name)+"')")
	}
	query := fmt.Sprintf("INSERT INTO minus_skill(name) VALUES %s ON CONFLICT DO NOTHING;", strings.Join(names, ","))
	return query
}
