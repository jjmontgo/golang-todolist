package frame

type Controller struct {
	Name string
	Actions map[string]func()
}

func NewController(name string) *Controller {
	newController := &Controller{Name: name}
	newController.Actions = make(map[string]func())
	ControllerMap[name] = newController
	return newController
}

func (this *Controller) Render(templateName string, vars interface{}) {
	view := ViewMgr.Get(templateName)
	view.Render(vars)
}
