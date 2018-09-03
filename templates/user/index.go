package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "user/index",
		HasLayout: true,
		Template: `
<h1>Users</h1>

<a href="{{url "user_new"}}">New User</a>

<table class="table">
	<thead>
		<tr>
			<th>Username</th>
			<th>Email</th>
			<th>Actions</th>
		</tr>
	</thead>
	{{range .Users}}
		<tr>
			<td>{{.Username}}</td>
			<td>{{.Email}}</td>
			<td>
				{{$id := uint_to_string .Id}}
				<form method="POST" action="{{url "user_delete" "id" $id}}">
					<a class="btn btn-secondary" href="{{url "user_edit" "id" $id}}">Edit</a>
					<button type="submit" class="btn btn-danger" onclick="return confirm('Are you sure')">Delete</button>
				</form>
			</td>
		</tr>
	{{end}}
</ul>
`})
}
