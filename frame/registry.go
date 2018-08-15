package frame

import (
	"net/http"
	"github.com/gorilla/mux"
)

type registry struct{
	Router *mux.Router
	Request *http.Request
	Response http.ResponseWriter
	Controllers map[string]*Controller
	Views map[string]*View
}

var Registry registry

func init() {
	Registry.Controllers = make(map[string]*Controller)
	Registry.Views = make(map[string]*View)
}
