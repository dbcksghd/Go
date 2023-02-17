package myDict

import "errors"

var (
	errNotFound     = errors.New("찾지 못했습니다.")
	errWordExists   = errors.New("이 단어는 이미 있습니다.")
	errCannotUpdate = errors.New("없는 단어는 업뎃할 수 없습니다.")
)

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	val, exi := d[word]
	if exi {
		return val, nil
	}
	return "", errNotFound
}

func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
}

func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = def
	case errNotFound:
		return errCannotUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
