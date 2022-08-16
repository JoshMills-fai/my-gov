package controllers

import (
	"html/template"

	"github.com/gin-gonic/gin"
)

func createTemplate(c *gin.Context, filepath string, data any) {
	var tmpl *template.Template
	tmpl = template.Must(template.ParseFiles(filepath))
	tmpl.Execute(c.Writer, data)
}

func HomeController(c *gin.Context) {
	////////////////////////// TODO get data from states.json to create a select instead of an input
	createTemplate(c, "views/templates/index.html", nil)
}

func MyRepresentativesController(c *gin.Context, membersMatch interface{}) {
	createTemplate(c, "views/templates/my-representatives.html", membersMatch)
}
