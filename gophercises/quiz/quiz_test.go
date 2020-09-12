package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	want := "Hello Github"
	got := SayHello()

	if got != want {
		t.Errorf("\ngot : %v\nwant: %v", got, want)
	}

}

func TestReadFile(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("dummy, csv, data")
	content, err := readFile(&buf)
	if err != nil {
		t.Errorf("Failed to read csv data")
	}
	fmt.Print(content)
}
