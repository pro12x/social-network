package utils

type Map struct {
	Name string
	Data interface{}
}

func JsonToString(data map[string]interface{}) string {
	var str string
	for key, value := range data {
		str += key + ": " + value.(string) + "\n"
	}
	return str
}
