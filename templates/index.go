package templates

import (
	"golang-todolist/frame"
	"golang-todolist/model/todolist"
)

type IndexVars struct {
	Results []todolist.Todolist
	IndexURL string
}

func init() {
	frame.ViewMgr.Add(frame.View{
		Name: "index",
		HasLayout: true,
		// Vars,
		Template: `
	<h1>Todo Lists</h1>
	<p>These are your todo lists:</p>
	<ul class="list-group">
		{{range .Results}}
			<li class="list-group-item">{{.Id}} {{.Name}}</li>
		{{end}}
	</ul>
	`})
}
