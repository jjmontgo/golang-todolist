package frame

import (
	"net/http"
	"html/template"
	"fmt"
	"bytes"
)

type View struct {
	Name string
	HasLayout bool
	Vars interface{}	// gets set when View.Execute() is called
	Template string
	ParsedTemplate *template.Template
}

func (this View) Render(w http.ResponseWriter, vars interface{}) {
	this.Vars = vars
	var renderedContent bytes.Buffer
	if err := this.ParsedTemplate.Execute(&renderedContent, this.Vars); err != nil {
		fmt.Println(err)
	}
	if this.HasLayout {
		layoutTemplate := template.New("layout")
		layoutTemplate, _ = layoutTemplate.Parse(ViewMgr.LayoutTemplate)
		layoutTemplate.Execute(w, template.HTML(renderedContent.String()))
	} else {
		this.ParsedTemplate.Execute(w, template.HTML(renderedContent.String()))
	}
}
