package mymap

func Search(dictionary map[string]string, key string) string {
	val, _ := dictionary[key]
	return val
}
