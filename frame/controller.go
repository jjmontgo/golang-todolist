package frame

import "net/http"

type Controller struct {}

func (this *Controller) Render(w http.ResponseWriter, templateName string, vars interface{}) {
	ViewMgr.Get(templateName).Render(w, vars)
}
