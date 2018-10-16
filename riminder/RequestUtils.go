package riminder

func AddIfNotEmptyStrMap(tofill *map[string]string, input map[string]interface{}, key_to_find string) {

	if value, ok := input[key_to_find]; ok {
		if value == nil {
			return
		}
		(*tofill)[key_to_find] = value.(string)
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
