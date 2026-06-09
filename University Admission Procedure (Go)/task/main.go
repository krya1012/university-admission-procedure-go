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
	exams      [4]float64
	priorities [3]string
}

var departments = []string{"Biotech", "Chemistry", "Engineering", "Mathematics", "Physics"}

var deptExams = map[string][]int{
	"Physics":     {0, 2},
	"Biotech":     {1, 0},
	"Chemistry":   {1},
	"Mathematics": {2},
	"Engineering": {3, 2},
}

func score(a Applicant, dept string) float64 {
	idx := deptExams[dept]
	sum := 0.0
	for _, i := range idx {
		sum += a.exams[i]
	}
	return sum / float64(len(idx))
}

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
		var exams [4]float64
		for i := 0; i < 4; i++ {
			exams[i], _ = strconv.ParseFloat(parts[2+i], 64)
		}
		pool = append(pool, Applicant{
			name:       parts[0] + " " + parts[1],
			exams:      exams,
			priorities: [3]string{parts[6], parts[7], parts[8]},
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
				si, sj := score(candidates[i], dept), score(candidates[j], dept)
				if si != sj {
					return si > sj
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

	for _, dept := range departments {
		list := admitted[dept]
		sort.Slice(list, func(i, j int) bool {
			si, sj := score(list[i], dept), score(list[j], dept)
			if si != sj {
				return si > sj
			}
			return list[i].name < list[j].name
		})
		f, _ := os.Create(strings.ToLower(dept) + ".txt")
		for _, a := range list {
			fmt.Fprintf(f, "%s %s\n", a.name, formatScore(score(a, dept)))
		}
		f.Close()
	}
}
