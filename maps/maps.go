package maps

const (
	ErrNotFound                = DictionaryErr("could not find the word you were looking for")
	ErrWordExists              = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist        = DictionaryErr("cannot update because word does not exist")
	ErrCannotDeleteWordMissing = DictionaryErr("cannot delete the word because it is not present")
)

type DictionaryErr string

type Dictionary map[string]string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}
	return nil
}
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrCannotDeleteWordMissing
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

/*
func Search(dictionary map[string]string, word string) string {
	return dictionary[word]
}
*/
