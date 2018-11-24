package frame

import (
	"net/http"
	"os"
	"runtime"
	"fmt"
)

func Dispatch(controllerName string, actionName string) http.HandlerFunc {
	registeredController, controllerExists := Registry.Controllers[controllerName]
	if (!controllerExists) {
		panic("Controller not set on Registry: " + controllerName)
	}

	action, actionExists := registeredController.Actions[actionName]
	if (!actionExists) {
		panic("Action not set on controller '" + controllerName + "': " + actionName)
	}

	return func (w http.ResponseWriter, r *http.Request) {
		instantiatedController := registeredController
		instantiatedController.Request = r
		instantiatedController.Response = w

		// Allow cross origin requests from frontend
		frontendOrigin := os.Getenv("FRONTEND_ORIGIN")
		w.Header().Set("Access-Control-Allow-Origin", frontendOrigin)
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    w.Header().Set("Access-Control-Allow-Headers", "*")

		// Set a default content type;
		// AWS API Gateway sets it to JSON if I don't set one, which causes an
		// error on empty responses
		w.Header().Set("Content-Type", "text/html;charset=UTF-8")

		// Check if action is accessible
		if (!instantiatedController.IsAccessible(actionName)) {
			instantiatedController.HttpStatus(http.StatusForbidden)
			instantiatedController.RenderJSON("message", "Unauthorized; must sign in first")
		} else {
			action()
		}
		instantiatedController.AfterAction() // store session data

		if os.Getenv("MODE") == "dev" {
			PrintMemUsage()
		}
	}
}

func PrintMemUsage() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	// For info on each, see: https://golang.org/pkg/runtime/#MemStats
	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
