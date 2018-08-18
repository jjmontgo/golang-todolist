package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todolist/edit",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Edit Todolist</h1>

<a href="{{url "index"}}">Back to Todo Lists</a>

<form action="{{url "todolist_save"}}" method="POST">
	<input type="hidden" name="id" value="{{.List.Id}}" />
	<div class="form-group">
		<input class="form-control {{if .Error}}is-invalid{{end}}" type="text" name="name" value="{{.List.Name}}">
		{{if .Error}}
			<div class="invalid-feedback">{{.Error}}</div>
		{{end}}
	</div>
	<input type="submit" value="Save">
</form>

`})
}
