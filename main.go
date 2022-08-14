package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getKey() string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	key := os.Getenv("apiKey")

	return key
}

type Member struct {
	Id                      string `json:"id"`
	Title                   string `json:"title"`
	Short_title             string `json:"short_title"`
	Api_uri                 string `json:"api_uri"`
	First_name              string `json:"first_name"`
	Middle_name             string `json:"middle_name"`
	Last_name               string `json:"last_name"`
	Suffix                  string `json:"suffix"`
	Date_of_birth           string `json:"date_of_birth"`
	Gender                  string `json:"gender"`
	Party                   string `json:"party"`
	Leadership_role         string `json:"leadership_role"`
	Twitter_account         string `json:"twitter_account"`
	Facebook_account        string `json:"facebook_account"`
	Youtube_account         string `json:"youtube_account"`
	Govtrack_id             string `json:"govtrack_id"`
	Cspan_id                string `json:"cspan_id"`
	Votesmart_id            string `json:"votesmart_id"`
	Icpsr_id                string `json:"icpsr_id"`
	Crp_id                  string `json:"crp_id"`
	Google_entity_id        string `json:"google_entity_id"`
	Fec_candidate_id        string `json:"fec_candidate_id"`
	Url                     string `json:"url"`
	Rss_url                 string `json:"rss_url"`
	Contact_form            string `json:"contact_form"`
	In_office               bool   `json:"in_office"`
	Cook_pvi                string `json:"cook_pvi"`
	Dw_nominate             int    `json:"dw_nominate"`
	Ideal_point             string `json:"ideal_point"`
	Seniority               string `json:"seniority"`
	Next_election           string `json:"next_election"`
	Total_votes             int    `json:"total_votes"`
	Missed_votes            int    `json:"missed_votes"`
	Total_present           int    `json:"total_present"`
	Last_updated            string `json:"last_updated"`
	Ocd_id                  string `json:"ocd_id"`
	Office                  string `json:"office"`
	Phone                   string `json:"phone"`
	Fax                     string `json:"fax"`
	State                   string `json:"state"`
	Senate_class            string `json:"senate_class"`
	State_rank              string `json:"state_rank"`
	Lis_id                  string `json:"lis_id"`
	Missed_votes_pct        int    `json:"missed_votes_pct"`
	Votes_with_party_pct    int    `json:"votes_with_party_pct"`
	Votes_against_party_pct int    `json:"votes_against_party_pct"`
}

func GetJson(url string, target interface{}) error {

	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Key", getKey())

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("there was an error", err)
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

var m []Member

func getMembers(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, GetJson("https://api.propublica.org/congress/v1/116/senate/members.json", m))

}

func main() {

	router := gin.Default()
	router.GET("/api/members", getMembers)
	router.Run(":3000")

}
