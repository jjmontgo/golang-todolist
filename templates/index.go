package templates

var IndexView = View{
	Name: "index",
	HasLayout: true,
	Template: `
<h1>Index</h1>
<p>This is the index template.</p>
Lists from struct:
{{range .Results}}
	<p>{{.Id}} {{.Name}}</p>
{{end}}
`}
