package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todo/index",
		HasLayout: true,
		Template: `
<h1>{{.List.Name}}</h1>
<p>
	{{$list_id := uint_to_string .List.Id}}
	<a href="{{ url "todo_new" "todo_list_id" $list_id }}" class="btn btn-primary">Add todo</a>
	<a href="{{ url "index" }}" class="btn btn-link">Back to todo lists</a>
</p>
<ul class="list-group">
	{{range .Todos}}
		<li class="list-group-item">
			{{.Id}} {{.Name}}
			{{$id := uint_to_string .Id}}
			<form method="POST" action="{{url "todo_delete" "id" $id}}">
				<a class="btn btn-secondary" href="{{url "todo_edit" "id" $id}}">Edit</a>
				<button type="submit" class="btn btn-danger" onclick="return confirm('Are you sure')">Delete</button>
			</form>
		</li>
	{{end}}
</ul>
`})
}
