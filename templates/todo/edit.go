package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todo/edit",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Edit Todo</h1>

{{$list_id := uint_to_string .List.Id}}
<a href="{{url "todolist" "id" $list_id}}">Back to Todo List</a>

<form action="{{url "todo_save"}}" method="POST">
	<input type="hidden" name="id" value="{{.Todo.Id}}" />
	<input type="hidden" name="todo_list_id" value="{{.List.Id}}" />
	<div class="form-group">
		<input class="form-control {{if .Error}}is-invalid{{end}}" type="text" name="name" value="{{.Todo.Name}}">
		{{if .Error}}
			<div class="invalid-feedback">{{.Error}}</div>
		{{end}}
	</div>
	<input type="submit" value="Save">
</form>

`})
}
