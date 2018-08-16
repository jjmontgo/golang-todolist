package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todolist/index",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Todo Lists</h1>

<a href="{{route "todolist_edit"}}">New Todolist</a>

<p>These are your todo lists:</p>
<ul class="list-group">
	{{range .Results}}
		<li class="list-group-item">
			{{.Id}} {{.Name}}

			{{with $id := .Id | tostring}}
				<form method="POST" action="{{route "todolist_delete" "id" $id}}">
					<button type="submit" onclick="return confirm('Are you sure')">Delete</button>
				</form>
			{{end}}
		</li>
	{{end}}
</ul>
`})
}
