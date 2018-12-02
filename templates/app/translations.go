package templates

import "golang-todolist/frame"

func init() {
	frame.NewView(&frame.View{
		Name: "app/translations",
		HasLayout: false,
		Template: `
window.t = {{ json_encode .Translations }};
`})
}
