package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "index",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Todo Lists</h1>

<a href="{{route "todolist_create"}}">New Todolist</a>

<p>These are your todo lists:</p>
<ul class="list-group">
	{{range .Results}}
		<li class="list-group-item">{{.Id}} {{.Name}}</li>
	{{end}}
</ul>
`})
}
