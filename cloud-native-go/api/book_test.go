package api

import (
	"testing"
)

func TestBookToJSON(t *testing.T) {

	book := Book{Title: "Cloud Native Go", Author: "M,-L. Reimer", ISBN: "0123456789"}
	json := book.ToJSON()

	got := string(json)

	want := `{"title":"Cloud Native Go","author":"M,-L. Reimer","isbn":"0123456789"}`

	if got != want {
		t.Errorf("got\n%q\nwant\n%q\n --- Book JSON marshalling wrong.", got, want)
	}
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"M,-L. Reimer","isbn":"0123456789"}`)
	got := FromJson(json)

	want := Book{Title: "Cloud Native Go", Author: "M,-L. Reimer", ISBN: "0123456789"}

	if got != want {
		t.Errorf("got:  %q\nwant: %q\nBook JSON Unmarshaling wrong.", got, want)
	}

}
