package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "layout",
		Template:
`<!doctype html>
<html lang="en">
<head>
	<meta charset="utf-8">
	<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
	<link rel="stylesheet" href="https://stackpath.bootstrapcdn.com/bootstrap/4.1.3/css/bootstrap.min.css" integrity="sha384-MCw98/SFnGE8fJT3GXwEOngsV7Zt27NXFoaoApmYm81iuXoPkFOJwJ8ERdknLPMO" crossorigin="anonymous">
	<title>Todolist App</title>
</head>
<body>
	<nav class="navbar navbar-expand-lg navbar-light bg-light">
		<a class="navbar-brand" href="#">Todo Lists</a>
		<button class="navbar-toggler" type="button" data-toggle="collapse" data-target="#navbarSupportedContent" aria-controls="navbarSupportedContent" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
		<div class="collapse navbar-collapse" id="navbarSupportedContent">
			<ul class="navbar-nav mr-auto">
				<li class="nav-item active">
					<a class="nav-link" href="{{url "index"}}">Home</a>
				</li>
				{{if user_is_logged_in}}
					<li class="nav-item">
						<a class="nav-link" href="{{url "logout"}}">Logout</a>
					</li>
				{{else}}
					<li class="nav-item">
						<a class="nav-link" href="{{url "login"}}">Login</a>
					</li>
				{{end}}
			</ul>
		</div>
	</nav>
	<div class="container">
		{{.Content}}
	</div>
</body>
</html>`})
}
