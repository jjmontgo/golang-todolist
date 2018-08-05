package templates

import "golang-todolist/frame"

func init() {
	frame.ViewMgr.SetLayout(
`<html>
<head>
	<title></title>
</head>
<body>
{{.}}
</body>
</html>`)
}
