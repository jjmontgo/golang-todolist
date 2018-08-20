package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "todo/index",
		HasLayout: true,
		Template: `
<h1>{{.List.Name}}</h1>
<p><a href="{{ url "index" }}" class="btn btn-link">Back to todo lists</a></p>

`})
}
