package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todolist/index",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Todo Lists</h1>

<a href="{{url "todolist_new"}}">New Todolist</a>

<p>These are your todo lists:</p>
<ul class="list-group">
	{{range .Results}}
		<li class="list-group-item">
			{{.Id}} {{.Name}}
			<form method="POST" action="{{url "todolist_delete" "id" .Id}}">
				<a class="btn btn-secondary" href="{{url "todolist" "id" .Id}}">Open</a>
				<a class="btn btn-secondary" href="{{url "todolist_edit" "id" .Id}}">Edit</a>
				<button type="submit" class="btn btn-danger" onclick="return confirm('Are you sure')">Delete</button>
			</form>
		</li>
	{{end}}
</ul>
`})
}
