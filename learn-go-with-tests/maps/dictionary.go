package maps

type Dictionary map[string]string
type DictionaryError string

func (d DictionaryError) Error() string {
	return string(d)
}

const (
	ErrorNotFound    = DictionaryError("can't search because key not found")
	ErrorKeyExists   = DictionaryError("can't add beacuse key already existed")
	ErrorKeyNotExist = DictionaryError("can't update because key not existed")
)

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		return ErrorKeyExists
	case ErrorNotFound:
		d[key] = value
	}
	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrorNotFound:
		return ErrorKeyNotExist
	case nil:
		d[key] = value
	}
	return nil
}
