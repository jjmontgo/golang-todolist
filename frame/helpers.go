package frame

import (
	"log"
	"strconv"
	"strings"
	"unicode"
)

/**
 * Takes an even, indefinite number of function arguments and translates
 * it into a string-indexed map.
 *
 * eg. Function("var1", value, "var2", value) ...
 * 	map["var1"] = value; map["var2"] = value;
 */
func BuildParameterMap(params ...interface{}) map[string]interface{} {
	vars := make(map[string]interface{})
	if len(params) % 2 == 0 {
		// populate the vars map; vars[params[0]] = vars[params[1]]
		var key string
		for i, v := range params {
			if i % 2 == 0 {
				key = ToString(v) // frame/helpers.go
			} else {
				vars[key] = v
			}
		}
	} else {
		log.Printf("BuildParameterMap received uneven number of arguments: %v", params)
	}
	return vars
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
