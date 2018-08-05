package templates

import "golang-todolist/frame"

var TestView = frame.View{
	Name: "test",
	HasLayout: true,
	Template: `
<h1>Test Template</h1>
<p>This is the test template.</p>
`}
