package mymap

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return val, nil
}
