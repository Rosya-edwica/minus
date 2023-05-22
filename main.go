package main

import (
	"fmt"
	"minus/sqldb"
	"strings"
	"sync"
)

func main() {
	minusSkills := sqldb.GetMinusSkills()
	noneSkills := sqldb.GetNoneSkills()
	groups := groupSkills(noneSkills)

	var wg sync.WaitGroup
	wg.Add(len(groups))
	for _, group := range groups {
		go findDuplicates(group, minusSkills, &wg)
	}
	wg.Wait()

}

func findDuplicates(skills []sqldb.Skill, minusSkills []string, wg *sync.WaitGroup) {
	var duplicates []sqldb.Skill

	for _, skill := range skills {
		for _, minus := range minusSkills {
			if contains(strings.Split(strings.ToLower(skill.Name), " "), strings.ToLower(minus)) {
				duplicates = append(duplicates, skill)
			}
		}

	}
	if len(duplicates) != 0 {
		sqldb.SaveMinusSkills(duplicates)
		sqldb.RemoveSkills(duplicates)
		fmt.Println("Duplicates: ", len(duplicates))
		duplicates = []sqldb.Skill{}

	}
	wg.Done()
}

func groupSkills(skills []sqldb.Skill) (groups [][]sqldb.Skill) {
	for i := 0; i < len(skills); i += 1000 {
		group := skills[i:][:1000]
		groups = append(groups, group)
	}
	return
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
