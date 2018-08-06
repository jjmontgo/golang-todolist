package templates

import "golang-todolist/frame"

type Todolist struct {
	Id int
	Name string
}

type IndexVars struct {
	Results []Todolist
	IndexURL string
}

func init() {
	frame.ViewMgr.Add(frame.View{
		Name: "index",
		HasLayout: true,
		// Vars,
		Template: `
	<h1>Index</h1>
	<p>This is the index template.</p>
	Lists from struct:
	{{range .Results}}
		<p>{{.Id}} {{.Name}}</p>
	{{end}}
	`})
}
