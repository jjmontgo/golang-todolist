package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "user/edit",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Edit User</h1>

<a href="{{url "users"}}">Back to Users</a>

<form action="{{url "user_save"}}" method="POST">
	<input type="hidden" name="id" value="{{.User.Id}}" />
	<div class="form-group">
		<label for="username">Username</label>
		<input class="form-control" type="text" id="username" name="username" value="{{.User.Username}}">
	</div>
	<div class="form-group">
		<label for="email">Email</label>
		<input class="form-control" type="text" id="email" name="email" value="{{.User.Email}}">
	</div>
	<div class="form-group">
		<label for="password">Password</label>
		<input class="form-control" type="text" id="password" name="password" value="">
	</div>
	<input type="submit" value="Save">
</form>
`})
}
