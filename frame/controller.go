package frame

type Controller struct {
	Name string
	Actions map[string]func()
}

func NewController(name string) *Controller {
	newController := &Controller{Name: name}
	newController.Actions = make(map[string]func())
	Registry.Controllers[name] = newController
	return newController
}

func (this *Controller) Render(templateName string, vars interface{}) {
	view := Registry.Views[templateName]
	view.Render(vars)
}
