package main

import (
	"io"
	"log"
	"os"
	"sync"
	"time"

	"github.com/syedwshah/LearnGo/nhlApi"
)

func main() {
	//help benchmarking the request time
	now := time.Now()

	rosterFile, err := os.OpenFile("rosters.txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("error opening the file rosters.txt: %v", err)
	}
	defer rosterFile.Close()

	wrt := io.MultiWriter(os.Stdout, rosterFile)

	log.SetOutput(wrt)

	teams, err := nhlApi.GetAllTeams()
	_ = teams
	if err != nil {
		log.Fatalf("error while getting all teams: %v", err)
	}

	var wg sync.WaitGroup

	wg.Add(len(teams))

	//unbuffered channel
	results := make(chan []nhlApi.Roster)

	for _, team := range teams {
		go func(team nhlApi.Team) {
			roster, err := nhlApi.GetRosters(team.ID)
			if err != nil {
				log.Fatalf("error getting roster: %v", err)
			}

			results <- roster

			wg.Done()
		}(team)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	display(results)

	log.Printf("took %v", time.Since(now))
}

func display(results chan []nhlApi.Roster) {
	for r := range results {
		for _, ros := range r {
			log.Println("---------------")
			log.Printf("Name: %s\n", ros.Person.Fullname)
			log.Printf("ID: %d\n", ros.Person.ID)
			log.Printf("Position: %s\n", ros.Position.Abbreviation)
			log.Printf("Jersey: %s\n", ros.Jerseynumber)
			log.Println("---------------")
		}
	}
}
