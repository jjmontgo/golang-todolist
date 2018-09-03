package frame

import (
	"html/template"
	"bytes"
	"log"
	// "fmt"
)

func NewView(view *View) {
	tpl := template.New(view.Name)
	tpl.Funcs(template.FuncMap{
		"url": URL, // frame/helpers.go
		"user_is_logged_in": UserIsLoggedIn, // frame/auth.go
		"to_string": ToString, // frame/helpers.go
		"uint_to_string": UintToString,
	})
	if view.LayoutTemplateName == "" {
		view.LayoutTemplateName = "layout"
	}
	tpl, _ = tpl.Parse(view.Template)
	view.ParsedTemplate = tpl
	Registry.Views[view.Name] = view
}

type View struct {
	Name string
	Template string
	HasLayout bool
	LayoutTemplateName string // default to "layout"
	ParsedTemplate *template.Template
}

func (this *View) Render(params ...interface{}) {
	vars := make(map[string]interface{})
	if len(params) % 2 == 0 {
		// populate the vars map; vars[params[0]] = vars[params[1]]
		var key string
		for i, v := range params {
			if i % 2 == 0 {
				key = ToString(v) // frame/helpers.go
			} else {
				vars[key] = v
			}
		}
	}

	if this.HasLayout {
		// buffer the template to renderedContent
		var renderedContent bytes.Buffer
		if err := this.ParsedTemplate.Execute(&renderedContent, vars); err != nil {
			log.Fatalf("this.ParsedTemplate.Execute(): %q\n", err)
		}
		// render the content to the layout as html
		layoutView := Registry.Views[this.LayoutTemplateName]
		html := template.HTML(renderedContent.String())
		layoutView.Render("Content", html)
	} else {
		// render the template to the response with no layout
		this.ParsedTemplate.Execute(Registry.Response, vars)
	}
}
