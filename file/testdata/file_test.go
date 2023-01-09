package file

import (
	"io/ioutil"
	"testing"
)

func TestReadFile(t *testing.T) {
	data, err := ioutil.ReadFile("./test.data")
	if err != nil {
		t.Fatal("Could not open file")
	}

	if string(data) != "hello world" {
		t.Fatal("String contents do not match expected")
	}
}
