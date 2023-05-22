package main

import (
	"fmt"
	"minus/sqldb"
	"strings"
)

func main() {
	minusSkills := sqldb.GetMinusSkills()
	noneSkills := sqldb.GetNoneSkills()

	var duplicates []sqldb.Skill

	for i, skill := range noneSkills {
		for _, minus := range minusSkills {
			if contains(strings.Split(strings.ToLower(skill.Name), " "), strings.ToLower(minus)) {
				duplicates = append(duplicates, skill)
			}
		}
		fmt.Println(i, ":", len(noneSkills))
		if len(duplicates) > 500 {
			sqldb.SaveMinusSkills(duplicates)
			sqldb.RemoveSkills(duplicates)
			fmt.Println("Duplicates: ", len(duplicates))
			duplicates = []sqldb.Skill{}
		}

	}

}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
