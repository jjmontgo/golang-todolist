package frame

import (
	"net/http"
)

func Dispatch(controllerName string, actionName string) http.HandlerFunc {
	controller, controllerExists := ControllerMap[controllerName]
	if (!controllerExists) {
		panic("Controller not set on ControllerMgr: " + controllerName)
	}
	action, actionExists := controller.Actions[actionName]
	if (!actionExists) {
		panic("Action not set on controller '" + controllerName + "': " + actionName)
	}

	return func (w http.ResponseWriter, r *http.Request) {
		controller.Request = r
		controller.Response = w
		action()
	}
}
