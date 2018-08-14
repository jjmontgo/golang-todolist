package frame

import (
	"net/http"
	"html/template"
	"bytes"
	"log"
	"github.com/gorilla/mux"
)

type View struct {
	Name string
	HasLayout bool
	Vars interface{}	// gets set when View.Execute() is called
	Template string
	ParsedTemplate *template.Template
	Router *mux.Router
}

func (this *View) Render(w http.ResponseWriter, vars interface{}) {
	this.Vars = vars
	var renderedContent bytes.Buffer
	if err := this.ParsedTemplate.Execute(&renderedContent, this.Vars); err != nil {
		log.Fatalf("this.ParsedTemplate.Execute(): %q\n", err)
	}
	if this.HasLayout {
		layoutTemplate := template.New("layout")
		layoutTemplate, _ = layoutTemplate.Parse(ViewMgr.LayoutTemplate)
		layoutTemplate.Execute(w, template.HTML(renderedContent.String()))
	} else {
		this.ParsedTemplate.Execute(w, template.HTML(renderedContent.String()))
	}
}

func (this *View) Route(name string) string {
	url, err := this.Router.Get(name).URL()
	if (err != nil) {
		log.Fatalf("this.Router.Get(name).URL(): %q\n", err)
	}
	return url.String()
}
