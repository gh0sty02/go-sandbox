package main

import (
	"testing"
)

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}

		got, _ := dict.Search("test")
		want := "this is a test"

		assert_string(t, got, want)

	})
	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is a test"}
		want := "could not find the work you were looking for"
		_, err := dict.Search("unknown")

		if err == nil {
			t.Fatal("expected to get a error")
		}

		assert_string(t, err.Error(), want)
	})
}
func assert_string(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q ", got, want)
	}
}
func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)

	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, definition)
		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)

	})

}
func TestUpdate(t *testing.T) {
	word := "test"
	definiton := "this is just a test"

	dictionary := Dictionary{word: definiton}

	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}
func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("shoudl find added word", err)

	}
	assert_string(t, got, definition)
}
func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q, want %q", got, want)
	}
}
