package frame

import (
	"html/template"
	"bytes"
	"log"
)

type View struct {
	Name string
	HasLayout bool
	// Vars interface{}	// gets set when View.Execute() is called
	Template string
	ParsedTemplate *template.Template
}

func (this *View) Render(vars interface{}) {
	var renderedContent bytes.Buffer
	if err := this.ParsedTemplate.Execute(&renderedContent, vars); err != nil {
		log.Fatalf("this.ParsedTemplate.Execute(): %q\n", err)
	}
	if this.HasLayout {
		layoutTemplate := template.New("layout")
		layoutTemplate, _ = layoutTemplate.Parse(ViewMgr.LayoutTemplate)
		layoutTemplate.Execute(Registry.Response, template.HTML(renderedContent.String()))
	} else {
		this.ParsedTemplate.Execute(Registry.Response, template.HTML(renderedContent.String()))
	}
}

func (this *View) Route(name string) string {
	url, err := Registry.Router.Get(name).URL()
	if (err != nil) {
		log.Fatalf("Registry.Router.Get(name).URL(): %q\n", err)
	}
	return url.String()
}
