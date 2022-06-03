package main

import "fmt"

type Team struct {
	name string
	wins int
	ties int
}

// Write your code here.
func GetBestTeam(teams []*Team) *Team {
	mostPoints := -1
	var bestTeam *Team = nil

	for _, team := range teams {
		points := team.wins + team.ties
		if points > mostPoints {
			mostPoints = points
			bestTeam = team
		}
	}

	return bestTeam
}

func main() {
	// teams := []*Team{&Team{name: "A", wins: 5, ties: 2}, &Team{name: "B", wins: 2, ties: 1}, &Team{name: "C", wins: 2, ties: 4}}
	teams := []*Team{{name: "A", wins: 5, ties: 2}, {name: "B", wins: 2, ties: 1}, {name: "C", wins: 2, ties: 4}}
	bestTeam := *GetBestTeam(teams)
	bestTeamName := bestTeam.name // "A" - the best team is Team "A"
	fmt.Println(bestTeamName)
}