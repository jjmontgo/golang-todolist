package templates

import "html/template"

const tplTest = `
<h1>Test Template</h1>
<p>This is the test template.</p>
`

func LoadTestTemplate() *template.Template {
	Test := template.New("test")
	Test, _ = Test.Parse(tplTest)
	return Test
}
