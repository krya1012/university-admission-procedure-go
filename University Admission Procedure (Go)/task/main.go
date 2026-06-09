package main

import (
	"fmt"
	"sort"
)

type Applicant struct {
	name string
	gpa  float64
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	applicants := make([]Applicant, n)
	for i := 0; i < n; i++ {
		var first, last string
		var gpa float64
		fmt.Scan(&first, &last, &gpa)
		applicants[i] = Applicant{first + " " + last, gpa}
	}

	sort.Slice(applicants, func(i, j int) bool {
		if applicants[i].gpa != applicants[j].gpa {
			return applicants[i].gpa > applicants[j].gpa
		}
		return applicants[i].name < applicants[j].name
	})

	fmt.Println("Successful applicants:")
	for i := 0; i < m; i++ {
		fmt.Println(applicants[i].name)
	}
}
