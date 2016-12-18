package hashtable

var data map[string]string

func setValue(key, value string) {
	data[key] = value
}

func getValue(key string) string {
	return data[key]
}
