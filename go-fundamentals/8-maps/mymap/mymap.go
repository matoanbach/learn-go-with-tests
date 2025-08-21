package mymap

type Dictionary map[string]string

func (d Dictionary) Search(key string) string {
	val, _ := d[key]
	return val
}
