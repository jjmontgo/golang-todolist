package templates

import "html/template"

const tplIndex = `
<h1>Index</h1>
<p>This is the index template.</p>
Lists:
{{range .}}
	<p>{{.Id}} {{.Name}}</p>
{{end}}
`

func LoadIndexTemplate() *template.Template {
	Index := template.New("index")
	Index, _ = Index.Parse(tplIndex)
	return Index
}
