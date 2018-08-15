package frame

import (
	"net/http"
	"github.com/gorilla/mux"
)

type registry struct{
	Router *mux.Router
	Request *http.Request
	Response http.ResponseWriter
}

var Registry registry
