package frame

import (
	"log"
	"strconv"
)

func URL(name string, vars ...string) string {
	url, err := Registry.Router.Get(name).URL(vars...)
	if (err != nil) {
		log.Fatalf("Registry.Router.Get(name).URL(): %q\n", err)
	}
	return url.String()
}

func ToString(value interface{}) string {
	switch v := value.(type) {
	case string:
		return v
	case int:
		return strconv.Itoa(v)
	// Add whatever other types you need
	default:
		return ""
	}
}
