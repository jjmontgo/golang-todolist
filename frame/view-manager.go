package frame

import (
	"html/template"
)

var ViewMgr ViewManager

func init() {
	ViewMgr.LayoutTemplate = "No layout"
	ViewMgr.Views = make(map[string]View)
}

type ViewManager struct {
	LayoutTemplate string
	Views map[string]View
}

func (this *ViewManager) SetLayout(layoutTemplate string) {
	this.LayoutTemplate = layoutTemplate
}

func (this *ViewManager) Add(view View) {
	tpl := template.New(view.Name)
	tpl, _ = tpl.Parse(view.Template)
	view.ParsedTemplate = tpl
	this.Views[view.Name] = view
}

func (this *ViewManager) Get(viewName string) View {
	return this.Views[viewName]
}
