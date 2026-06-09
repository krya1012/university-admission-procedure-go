package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Applicant struct {
	name       string
	gpa        float64
	priorities [3]string
}

var departments = []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}

func formatScore(f float64) string {
	s := strconv.FormatFloat(f, 'f', -1, 64)
	if !strings.Contains(s, ".") {
		s += ".0"
	}
	return s
}

func main() {
	var n int
	fmt.Scan(&n)

	file, _ := os.Open("applicants.txt")
	defer file.Close()

	var pool []Applicant
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		gpa, _ := strconv.ParseFloat(parts[2], 64)
		pool = append(pool, Applicant{
			name:       parts[0] + " " + parts[1],
			gpa:        gpa,
			priorities: [3]string{parts[3], parts[4], parts[5]},
		})
	}

	admitted := make(map[string][]Applicant)

	for wave := 0; wave < 3; wave++ {
		acceptedSet := make(map[string]bool)
		for _, dept := range departments {
			slots := n - len(admitted[dept])
			if slots <= 0 {
				continue
			}
			var candidates []Applicant
			for _, a := range pool {
				if a.priorities[wave] == dept {
					candidates = append(candidates, a)
				}
			}
			sort.Slice(candidates, func(i, j int) bool {
				if candidates[i].gpa != candidates[j].gpa {
					return candidates[i].gpa > candidates[j].gpa
				}
				return candidates[i].name < candidates[j].name
			})
			if len(candidates) > slots {
				candidates = candidates[:slots]
			}
			admitted[dept] = append(admitted[dept], candidates...)
			for _, c := range candidates {
				acceptedSet[c.name] = true
			}
		}
		var next []Applicant
		for _, a := range pool {
			if !acceptedSet[a.name] {
				next = append(next, a)
			}
		}
		pool = next
	}

	for i, dept := range departments {
		if i > 0 {
			fmt.Println()
		}
		fmt.Println(dept)
		list := admitted[dept]
		sort.Slice(list, func(i, j int) bool {
			if list[i].gpa != list[j].gpa {
				return list[i].gpa > list[j].gpa
			}
			return list[i].name < list[j].name
		})
		for _, a := range list {
			fmt.Printf("%s %s\n", a.name, formatScore(a.gpa))
		}
	}
}
