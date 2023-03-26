package main

import "testing"
func TestSearch(t *testing.T) {
	dict := Dict{"test": "This is only a test"}

	t.Run("Valid word", func(t *testing.T) {
		result, _ := dict.Search("test")
		expected := "This is only a test"
		compareStrings(t, result, expected)
	})

	t.Run("Invalid word", func(t *testing.T) {
		_, err := dict.Search("value")
		compareError(t, err, NoKey)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New Key", func(t *testing.T) {
		dict := Dict{}
		key := "test"
		value := "This is only a test"
		err := dict.Add(key, value)
		compareError(t, err, nil)
		compareValue(t, dict, key, value)
	})

	t.Run("Existing key", func(t *testing.T) {
		key := "test"
		value := "This is only a test"
		dict := Dict{key: value}
		err := dict.Add(key, value)
		compareError(t, err, HasKey)
		compareValue(t, dict, key, value)
	})
}

func TestDelete(t *testing.T) {
	key := "test"
	dict := Dict{key: "This is only a test"}
	dict.Delete(key)
	_, err := dict.Search(key)
	if err != NoKey {
		t.Errorf("expected '%s' should be deleted", err)
	}
}

func TestUpdate(t *testing.T) {
	t.Run("Key exist", func(t *testing.T) {
		key := "test"
		value := "This is only a test"
		dict := Dict{key: value}
		newValue := "This is only a another test"
		err := dict.Update(key, newValue)
		compareError(t, err, nil)
		compareValue(t, dict, key, newValue)
	})
	t.Run("Key don't exist", func(t *testing.T) {
		key := "test"
		value := "This is only a test"
		dict := Dict{}
		err := dict.Update(key, value)
		compareError(t, err, HasKey)
	})
}

func compareValue(t *testing.T, dict Dict, key, expected string) {
	t.Helper()
	result, err := dict.Search(key)
	if err != nil {
		t.Fatal("word not found")
	}
	if expected != result {
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}

func compareStrings(t *testing.T, result, expected string) {
	t.Helper()
	if result != expected {
		t.Errorf("result '%s', expected '%s', given '%s'", result, expected, "test")
	}
}

func compareError(t *testing.T, err, expected error) {
	t.Helper()
	if err != expected {
		t.Errorf("result err '%s', expected '%s'", err, expected)
	}
}