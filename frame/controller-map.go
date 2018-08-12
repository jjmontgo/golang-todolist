package frame

var ControllerMap map[string]*Controller

func init() {
	ControllerMap = make(map[string]*Controller)
}
