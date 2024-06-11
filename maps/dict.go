package main

import "errors"

var (
	ErrNotFound   = errors.New("could not find the work you were looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}

	d[key] = value

	return nil
}

func (d Dictionary) Update(word, definition string) {
	d[word] = definition
}
func main() {

	title, page, again := getBookInfo()

	println(title, page, again)
}

func getBookInfo() (string, int, string) {
	return "hello", 100, "hel"
}
