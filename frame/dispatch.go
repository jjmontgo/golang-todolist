package frame

import "net/http"

func Dispatch(controllerName string, actionName string) http.HandlerFunc {
	controller, controllerExists := Registry.Controllers[controllerName]
	if (!controllerExists) {
		panic("Controller not set on Registry: " + controllerName)
	}

	action, actionExists := controller.Actions[actionName]
	if (!actionExists) {
		panic("Action not set on controller '" + controllerName + "': " + actionName)
	}

	return func (w http.ResponseWriter, r *http.Request) {
		Registry.Request = r
		Registry.Response = w
		if (!controller.IsAccessible(actionName)) {
			controller.Redirect(URL("login")) // or a 403 error?
			return
		}
		action()
		controller.AfterAction()
	}
}
