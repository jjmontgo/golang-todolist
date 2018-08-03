package templates

import "net/http"
import "html/template"
import "fmt"
import "bytes"

type Todolist struct {
	Id int
	Name string
}

type TplIndexVars struct {
	Results []Todolist
	IndexURL string
}

const tplIndex = `
{{define "content"}}
	<h1>Index</h1>
	<p>This is the index template.</p>
	Lists from struct:
	{{range .Results}}
		<p>{{.Id}} {{.Name}}</p>
	{{end}}
{{end}}
`

func ExecuteIndexTemplate(w http.ResponseWriter, indexVars TplIndexVars) {
	Index := template.New("content")
	Index, _ = Index.Parse(tplIndex)
	var renderedContent bytes.Buffer
	if err := Index.Execute(&renderedContent, indexVars); err != nil {
		fmt.Println(err)
	}
	Layout := template.New("layout")
	Layout, _ = Layout.Parse(tplLayout)
	Layout.Execute(w, template.HTML(renderedContent.String()))
}
