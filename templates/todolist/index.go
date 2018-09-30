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
		{{$id := uint_to_string .Id}}
		<li class="list-group-item">
			<div class="row">
				<div class="col-sm">
					{{$imgSrc := .GetImgSrc}}
					<a href="{{url "todolist_image_form" "id" $id}}">
						{{if ne $imgSrc ""}}
							<img src="{{$imgSrc}}" class="img-fluid" />
						{{else}}
							No Image
						{{end}}
					</a>
				</div>
				<div class="col-sm">
					{{.Id}} {{.Name}}
				</div>
				<div class="col-sm">
					<form method="POST" action="{{url "todolist_delete" "id" $id}}">
						<a class="btn btn-secondary" href="{{url "todolist" "id" $id}}">Open</a>
						<a class="btn btn-secondary" href="{{url "todolist_edit" "id" $id}}">Edit</a>
						<a class="btn btn-secondary" href="{{url "todolist_email" "id" $id}}">Email</a>
						<button type="submit" class="btn btn-danger" onclick="return confirm('Are you sure')">Delete</button>
					</form>
				</div>
			</div>
		</li>
	{{end}}
</ul>
`})
}
