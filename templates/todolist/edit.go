package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todolist/edit",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Edit Todolist</h1>

<a href="{{route "index"}}">Back to Todo Lists</a>

<form action="{{route "todolist_save"}}" method="POST">
	<input type="hidden" name="id" value="{{.Id}}" />
	<input type="text" name="name" value="{{.Name}}">
	<input type="submit" value="Save">
</form>

`})
}
