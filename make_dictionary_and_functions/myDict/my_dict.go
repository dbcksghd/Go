package myDict

import "errors"

var errNotFound = errors.New("찾지 못했습니다.")

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, exi := d[word]
	if exi {
		return val, nil
	}
	return "", errNotFound
}
