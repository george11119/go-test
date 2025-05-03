package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Name  string
	Teams map[string]Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1Name string, team1Score int, team2Name string, team2Score int) {
	_, team1Exists := l.Teams[team1Name]
	_, team2Exists := l.Teams[team2Name]

	if !team1Exists || !team2Exists {
		return
	}

	if team1Score > team2Score {
		l.Wins[team1Name] += 1
	} else if team2Score > team1Score {
		l.Wins[team2Name] += 1
	}
}

func (l League) Ranking() []string {
	rankings := []string{}

	for k := range l.Wins {
		rankings = append(rankings, k)
	}

	sort.Slice(rankings, func(i, j int) bool {
		return l.Wins[rankings[i]] > l.Wins[rankings[j]]
	})

	return rankings
}

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	for _, team := range r.Ranking() {
		io.WriteString(w, team+"\n")
	}
}

func main() {
	l := League{
		Name: "Big League",
		Teams: map[string]Team{
			"Italy": {
				Name:    "Italy",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"France": {
				Name:    "France",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"India": {
				Name:    "India",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
			"Nigeria": {
				Name:    "Nigeria",
				Players: []string{"Player1", "Player2", "Player3", "Player4", "Player5"},
			},
		},
		Wins: map[string]int{},
	}
	l.MatchResult("Italy", 80, "France", 70)
	l.MatchResult("India", 85, "Nigeria", 80)
	l.MatchResult("Italy", 60, "India", 55)
	l.MatchResult("France", 100, "Nigeria", 110)
	l.MatchResult("Italy", 65, "Nigeria", 70)
	l.MatchResult("France", 95, "India", 80)

	fmt.Println(l.Wins)
	RankPrinter(l, os.Stdout)
}
