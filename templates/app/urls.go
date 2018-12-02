package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "app/urls",
		HasLayout: false,
		Template: `
window.urls = {{ json_encode .Routes }};
`})
}
