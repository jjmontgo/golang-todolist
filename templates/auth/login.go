package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "auth/login",
		HasLayout: true,
		// LayoutTemplateName: "layout",
		Template: `
<h1>Login</h1>

<form action="{{url "login_validate"}}" method="POST">
	{{if .Error}}
		<p>{{.Error}}</p>
	{{end}}
	<div class="form-group">
		<label for="username">Username</label>
		<input class="form-control" type="text" id="username" name="username" value="">
	</div>
	<div class="form-group">
		<label for="password">Password</label>
		<input class="form-control" type="text" id="password" name="password" value="">
	</div>
	<input type="submit" value="Login">
</form>
`})
}
