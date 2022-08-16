package models

import (
	"encoding/json"
	"fmt"
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

var members Members

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

func serveMembers(c *gin.Context) {

	///////////////////////////////////// TODO call this API maybe once a day, and set a cookie or storage item
	dataBytes := GetJson("https://api.propublica.org/congress/v1/116/senate/members.json")

	// convert JSON to struct
	e := json.Unmarshal(dataBytes, &members)

	if e != nil {
		fmt.Println(e)
	}

	c.IndentedJSON(http.StatusOK, members)
}
