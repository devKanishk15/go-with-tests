package maps

import "errors"

type Dictionary map[string]string

var errorFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	defination, ok := d[word]
	if !ok {
		return "", errorFound
	}	
	
	return defination, nil
}

func (d Dictionary) Add(word, defination string) {
	d[word] = defination
}
