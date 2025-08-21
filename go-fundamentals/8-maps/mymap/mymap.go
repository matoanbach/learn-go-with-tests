package mymap

type Dictionary map[string]string

var (
	ErrNotFound      = DictionaryErr("could not find the word you were looking for")
	ErrAlreadyExists = DictionaryErr("key val pair already existed")
)

func (d Dictionary) Search(key string) (string, error) {
	val, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return val, nil
}

func (d Dictionary) Add(key, val string) error {
	_, err := d.Search(key)
	if err == nil {
		return ErrAlreadyExists
	}
	d[key] = val

	return nil
}

func (d Dictionary) Update(key, newVal string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}
	d[key] = newVal

	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	if err != nil {
		return err
	}

	delete(d, key)
	return nil
}

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
