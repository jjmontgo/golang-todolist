package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todo/index",
		HasLayout: true,
		Template: `
<h1>{{.List.Name}}</h1>
<p>
	<a href="{{ url "todo_new" "todo_list_id" .List.Id }}" class="btn btn-primary">Add todo</a>
	<a href="{{ url "index" }}" class="btn btn-link">Back to todo lists</a>
</p>
<ul class="list-group">
	{{range .Todos}}
		<li class="list-group-item">
			{{.Id}} {{.Name}}
			<form method="POST" action="{{url "todo_delete" "id" .Id}}">
				<a class="btn btn-secondary" href="{{url "todo_edit" "id" .Id}}">Edit</a>
				<button type="submit" class="btn btn-danger" onclick="return confirm('Are you sure')">Delete</button>
			</form>
		</li>
	{{end}}
</ul>
`})
}
