package nhlApi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Team struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Link  string `json:"link"`
	Venue struct {
		Name     string `json:"name"`
		Link     string `json:"link"`
		City     string `json:"city"`
		Timezone struct {
			ID     string `json:"id"`
			Offset int    `json:"offset"`
			Tz     string `json:"tz"`
		} `json:"timeZone"`
	} `json:"venue,omitempty"`
	Abbreviation    string `json:"abbreviation"`
	Teamname        string `json:"teamName"`
	Locationname    string `json:"locationName"`
	Firstyearofplay string `json:"firstYearOfPlay"`
	Division        struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"division"`
	Conference struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	} `json:"conference"`
	Franchise struct {
		Franchiseid int    `json:"franchiseId"`
		Teamname    string `json:"teamName"`
		Link        string `json:"link"`
	} `json:"franchise"`
	Shortname       string `json:"shortName"`
	Officialsiteurl string `json:"officialSiteUrl"`
	Franchiseid     int    `json:"franchiseId"`
	Active          bool   `json:"active"`
}

type nhlTeamsResponse struct {
	Teams []Team `json:"teams"`
}

func GetAllTeams() ([]Team, error) {
	res, err := http.Get(fmt.Sprintf("%s/teams", baseURL))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var response nhlTeamsResponse
	err = json.NewDecoder(res.Body).Decode(&response)

	return response.Teams, err
}
