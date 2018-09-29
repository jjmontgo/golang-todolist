package frame

import (
	"os"
	"log"
	"strconv"
	"strings"
	"unicode"
)

func URL(name string, vars ...string) string {
	url, err := Registry.Router.Get(name).URL(vars...)
	if (err != nil) {
		log.Fatalf("Registry.Router.Get(name).URL(): %q\n", err)
	}
	apiGatewayPathPrefix := os.Getenv("API_GATEWAY_PATH_PREFIX")
	return apiGatewayPathPrefix + url.String()
}

func AbsoluteURL(name string, vars ...string) string {
	relativeURL := URL(name, vars...)
	var httpOrHttps string
	isUsingTLS := Registry.Request.TLS != nil
	isProxiedHttps := Registry.Request.Header.Get("X-Forwarded-Proto") == "https"
	if isUsingTLS || isProxiedHttps {
		httpOrHttps = "https://"
	} else {
		httpOrHttps = "http://"
	}
	return httpOrHttps + Registry.Request.Host + relativeURL
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

func StringToUint(value string) uint {
	if value == "" {
		return 0
	}
	u64, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		log.Fatal(err)
	}
	return uint(u64)
}

func UintToString(value uint) string {
	return strconv.FormatUint(uint64(value), 10)
}

func RemoveWhitespace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}
