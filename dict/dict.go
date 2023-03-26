package main

type Dict map[string]string

const (
	NoKey = ErrDict("Has a error: word do'n exist")
	HasKey = ErrDict("Has a error: Word exist")
)

type ErrDict string

func (e ErrDict) Error() string {
	return string(e)
}

/*
Uma propriedade interessante dos maps é que você pode
modificá-los sem passá-los como ponteiro. Isso é porque
o map é um tipo referência. Isso significa que ele contém
uma referência à estrutura de dado que estamos utilizando,
assim como um ponteiro.
*/
func (d Dict) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case NoKey:
		d[key] = value
	case nil:
		return HasKey
	default:
		return err
	}
	return nil
}

func (d Dict) Search(word string) (string, error) {
	value, key := d[word]
	if !key {
		return "", NoKey
	}
	return value, nil
}

func (d Dict) Update(key , value string) error {
	_, err := d.Search(key)
	switch err {
	case NoKey:
		return HasKey
	case nil:
		d[key] = value
	default:
		return err
	}
	return nil
}

func (d Dict) Delete(key string) {
	delete(d, key)
}