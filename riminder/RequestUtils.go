package riminder

import (
	"fmt"
	"strconv"
)

func AddIfNotEmptyStrMap(tofill *map[string]string, input map[string]interface{}, key_to_find string) {

	if value, ok := input[key_to_find]; ok {
		if value == nil {
			return
		}
		switch val := value.(type) {
		case int:
			(*tofill)[key_to_find] = strconv.Itoa(val)
		case string:
			(*tofill)[key_to_find] = val
		case fmt.Stringer:
			(*tofill)[key_to_find] = val.String()
		default:
			panic(fmt.Sprintf("%s should be stringable.", key_to_find))
		}

	}

}

func AddIfNotEmptyMap(tofill *map[string]interface{}, input map[string]interface{}, key_to_find string) {

	if value, ok := input[key_to_find]; ok {
		if value == nil {
			return
		}
		(*tofill)[key_to_find] = value
	}

}
