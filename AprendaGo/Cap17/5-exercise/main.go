package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type sortByAge []user

func (s sortByAge) Len() int           { return len(s) }
func (s sortByAge) Less(i, j int) bool { return s[i].Age < s[j].Age }
func (s sortByAge) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

type sortByLast []user

func (s sortByLast) Len() int           { return len(s) }
func (s sortByLast) Less(i, j int) bool { return s[i].Last < s[j].Last }
func (s sortByLast) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}

	printUsers(users)
	sort.Sort(sortByAge(users))
	printUsers(users)
	sort.Sort(sortByLast(users))
	printUsers(users)

}

func printUsers(users []user) {
	fmt.Println()
	for _, u := range users {
		fmt.Printf("Name: %v %v\tAge: %v\n", u.First, u.Last, u.Age)
		sort.Strings(u.Sayings)
		for _, say := range u.Sayings {
			fmt.Println("\t", say)
		}
	}
}
