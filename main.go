package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
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

type Members struct {
	Status    string `json:"status"`
	Copyright string `json:"copyright"`
	Results   []struct {
		Congress    string        `json:"congress"`
		Chamber     string        `json:"chamber"`
		NumResults  int           `json:"num_results"`
		Offset      int           `json:"offset"`
		Memberslist []Memberslist `json:"members"`
	} `json:"results"`
}

type Memberslist struct {
	ID                   string      `json:"id"`
	Title                string      `json:"title"`
	ShortTitle           string      `json:"short_title"`
	APIURI               string      `json:"api_uri"`
	FirstName            string      `json:"first_name"`
	MiddleName           interface{} `json:"middle_name"`
	LastName             string      `json:"last_name"`
	Suffix               interface{} `json:"suffix"`
	DateOfBirth          string      `json:"date_of_birth"`
	Gender               string      `json:"gender"`
	Party                string      `json:"party"`
	LeadershipRole       interface{} `json:"leadership_role"`
	TwitterAccount       string      `json:"twitter_account"`
	FacebookAccount      string      `json:"facebook_account"`
	YoutubeAccount       string      `json:"youtube_account"`
	GovtrackID           string      `json:"govtrack_id"`
	CspanID              string      `json:"cspan_id"`
	VotesmartID          string      `json:"votesmart_id"`
	IcpsrID              string      `json:"icpsr_id"`
	CrpID                string      `json:"crp_id"`
	GoogleEntityID       string      `json:"google_entity_id"`
	FecCandidateID       string      `json:"fec_candidate_id"`
	URL                  string      `json:"url"`
	RssURL               string      `json:"rss_url"`
	ContactForm          string      `json:"contact_form"`
	InOffice             bool        `json:"in_office"`
	CookPvi              interface{} `json:"cook_pvi"`
	DwNominate           float64     `json:"dw_nominate"`
	IdealPoint           interface{} `json:"ideal_point"`
	Seniority            string      `json:"seniority"`
	NextElection         string      `json:"next_election"`
	TotalVotes           int         `json:"total_votes"`
	MissedVotes          int         `json:"missed_votes"`
	TotalPresent         int         `json:"total_present"`
	LastUpdated          string      `json:"last_updated"`
	OcdID                string      `json:"ocd_id"`
	Office               string      `json:"office"`
	Phone                string      `json:"phone"`
	Fax                  string      `json:"fax"`
	State                string      `json:"state"`
	SenateClass          string      `json:"senate_class"`
	StateRank            string      `json:"state_rank"`
	LisID                string      `json:"lis_id"`
	MissedVotesPct       float64     `json:"missed_votes_pct"`
	VotesWithPartyPct    float64     `json:"votes_with_party_pct"`
	VotesAgainstPartyPct float64     `json:"votes_against_party_pct"`
}

type MyState struct {
	Abbreviation string
}

// ///////////////////////////////////////////////TODO Rename these
var members Members
var mymembers Members

func GetJson(url string) []byte {

	client := &http.Client{Timeout: 10 * time.Second}
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("X-API-Key", getKey())

	res, err := client.Do(req)

	if err != nil {
		fmt.Println("there was an error", err)
	}

	// recieve JSON as []byte
	data, _ := ioutil.ReadAll(res.Body)

	return data

}

// ///////////////////////////////TODO differentiate between calling external api and ours
func getMembers(c *gin.Context) {

	// return []byte
	dataBytes := GetJson("https://api.propublica.org/congress/v1/116/senate/members.json")

	// convert JSON to struct
	e := json.Unmarshal(dataBytes, &members)

	if e != nil {
		fmt.Println(e)
	}

	c.IndentedJSON(http.StatusOK, members)
}

func home(c *gin.Context) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(c.Writer, nil)
}

func myrepresentatives(c *gin.Context) {
	//Consume from our own API
	////////////////////////////////////////////////Todo, make this url dynamic
	res, _ := http.Get("http://localhost:3000/api/members")
	data, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(data, &mymembers)

	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles("templates/my-representatives.html"))

	c.Request.ParseForm()

	var UserState MyState = MyState{c.Request.Form.Get("state")}

	if len(UserState.Abbreviation) == 0 {
		fmt.Println("no state")
		c.Redirect(301, "/")
	} else {

		var memberMatch []Memberslist
		//filter by user state
		for _, member := range mymembers.Results[0].Memberslist {
			if member.State == UserState.Abbreviation {
				memberMatch = append(memberMatch, member)
			}
		}

		fmt.Printf("%+v\n", memberMatch)

		tmpl.Execute(c.Writer, memberMatch)
	}

}

func main() {

	router := gin.Default()
	router.GET("/", home)
	router.GET("/api/members", getMembers)
	router.POST("/my-representatives", myrepresentatives)
	router.Run(":3000")

}
