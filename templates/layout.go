package templates

import "golang-todolist/frame"

func init() {
	frame.TplLayout =
	`<html>
	<head>
		<title></title>
	</head>
	<body>
	{{.}}
	</body>
	</html>`
}
