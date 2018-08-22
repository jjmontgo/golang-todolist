package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todo/index",
		HasLayout: true,
		Template: `
<h1>{{.List.Name}}</h1>
<p><a href="{{ url "index" }}" class="btn btn-link">Back to todo lists</a></p>
<ul class="list-group">
	{{range .Todos}}
		<li class="list-group-item">
			{{.Id}} {{.Name}}
		</li>
	{{end}}
</ul>
`})
}
