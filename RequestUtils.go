package riminder

import (
	"fmt"
	"strconv"
)

// AddIfNotEmptyStrMap add the value un the input map with the key (keyTofind) if it exist and is not null to tofill map(string version).
func AddIfNotEmptyStrMap(tofill *map[string]string, input map[string]interface{}, keyToFind string) {

	if value, ok := input[keyToFind]; ok {
		if value == nil {
			return
		}
		switch val := value.(type) {
		case int:
			(*tofill)[keyToFind] = strconv.Itoa(val)
		case string:
			(*tofill)[keyToFind] = val
		case fmt.Stringer:
			(*tofill)[keyToFind] = val.String()
		default:
			panic(fmt.Sprintf("%s should be stringable.", keyToFind))
		}

	}

}

// AddIfNotEmptyMap add the value un the input map with the key (keyTofind) if it exist and is not null to tofill map.
func AddIfNotEmptyMap(tofill *map[string]interface{}, input map[string]interface{}, keyToFind string) {

	if value, ok := input[keyToFind]; ok {
		if value == nil {
			return
		}
		(*tofill)[keyToFind] = value
	}

}
