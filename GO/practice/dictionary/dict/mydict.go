package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var (
	errNotFound   = errors.New("Not Found")
	errWordExists = errors.New("That word already exists")
	errCantUpdate = errors.New("Cannot update to new value")
	errNotExists  = errors.New("That word not exists")
)

// Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil // error 없으면 nil 리턴
	}
	return "", errNotFound
}

// Add new value to dictionary
func (d Dictionary) Add(key, value string) error {

	_, err := d.Search(key)

	// if err == errNotFound {
	// d[key] = value
	// } else if err == nil {
	// return errWordExists
	// }
	// return nil

	switch err {
	case errNotFound:
		d[key] = value
	case nil:
		return errWordExists
	}
	return nil
}

// Update word's definition
func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	// word not exists
	// if err != nil {
	// 	return errNotFound
	// } else {
	// 	d[word] = definition
	// }
	// return nil

	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

// Delete word from dictionary
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		delete(d, word)
	case errNotFound:
		return errNotExists
	}
	return nil
}
