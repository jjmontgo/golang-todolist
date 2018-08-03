package templates

import (
	"net/http"
	"html/template"
	"fmt"
	"bytes"
)

type View struct {
	Name string
	HasLayout bool
	Vars interface{}
	Template string
}

func (this View) Execute(w http.ResponseWriter, vars interface{}) {
	this.Vars = vars
	contentTemplate := template.New(this.Name)
	contentTemplate, _ = contentTemplate.Parse(this.Template)
	var renderedContent bytes.Buffer
	if err := contentTemplate.Execute(&renderedContent, this.Vars); err != nil {
		fmt.Println(err)
	}
	if this.HasLayout {
		layoutTemplate := template.New("layout")
		layoutTemplate, _ = layoutTemplate.Parse(TplLayout)
		layoutTemplate.Execute(w, template.HTML(renderedContent.String()))
	} else {
		contentTemplate.Execute(w, template.HTML(renderedContent.String()))
	}
}
