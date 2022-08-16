package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"my-gov/controllers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func home(c *gin.Context) {

	controllers.HomeController(c)

}

func myRepresentatives(c *gin.Context) {
	var mymembers Members
	//Consume from our own API
	////////////////////////////////////////////////Todo, make this url dynamic
	res, _ := http.Get("http://localhost:3000/api/members")
	data, _ := ioutil.ReadAll(res.Body)

	json.Unmarshal(data, &mymembers)

	c.Request.ParseForm()

	var UserState MyState = MyState{strings.ToUpper(c.Request.Form.Get("state"))}

	if len(UserState.Abbreviation) == 0 {
		fmt.Println("no state")
		c.Redirect(301, "/")
	}

	var memberMatch []Memberslist

	//filter by user state
	for _, member := range mymembers.Results[0].Memberslist {
		if member.State == UserState.Abbreviation {
			memberMatch = append(memberMatch, member)
		}
	}

	controllers.MyRepresentativesController(c, &memberMatch)

}
