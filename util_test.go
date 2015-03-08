package main

import (
	"os"
	"testing"
)

func TestIsDirEmpty(t *testing.T) {
	t.Parallel()

	emptyDir := createTmpFolder(t)
	if isEmpty, _ := IsDirEmpty(emptyDir); !isEmpty {
		t.Fatalf("Dir %s is empty, but return false", emptyDir)
	}

	nonEmptyDir := createTmpFolder(t)
	os.Mkdir(nonEmptyDir+"/_whatever", os.ModePerm)
	if isEmpty, _ := IsDirEmpty(nonEmptyDir); isEmpty {
		t.Fatalf("Dir %s is not empty, but return true", nonEmptyDir)
	}
}

func TestPrettify(t *testing.T) {
	t.Parallel()

	if result := Prettify("test space"); result != "test-space" {
		t.Fatalf("Prettify should change space to dash, but got %s", result)
	}

	if result := Prettify("  test  trim  "); result != "test-trim" {
		t.Fatalf("Prettify should trim unnecessary space, but got %s", result)
	}

	if result := Prettify("it's a test"); result != "its-a-test" {
		t.Fatalf("Prettify should remove single quote, but got %s", result)
	}

	if result := Prettify("LOWER CASE"); result != "lower-case" {
		t.Fatalf("Prettify should convert to lower case, but got %s", result)
	}
}

func TestUTF8Slice(t *testing.T) {
	t.Parallel()

	var str string = "Hello世界"
	var result string
	var err error
	var expect string
	var expectErrorMsg string

	expect = ""
	result, err = UTF8Slice(str, 0, 0)
	if result != expect || err != nil {
		t.Errorf(`UTF8Slice should return empty string when slice length is 0,
expect [%v][%v], but got [%v][%v]`, expect, nil, result, err)
	}

	expect = "世界"
	result, err = UTF8Slice(str, 5, 20)
	if result != expect || err != nil {
		t.Errorf(`UTF8Slice should return slice even if request length exceed,
expect [%v][%v], but got [%v][%v]`, expect, nil, result, err)
	}

	result, err = UTF8Slice(str, -1, 2)
	expectErrorMsg = "Start index out of bound"
	if result != "" || err == nil || err.Error() != expectErrorMsg {
		t.Errorf(`UTF8Slice should return error if start index smaller than zero`)
	}

	result, err = UTF8Slice(str, 10, 2)
	expectErrorMsg = "Start index out of bound"
	if result != "" || err == nil || err.Error() != expectErrorMsg {
		t.Errorf(`UTF8Slice should return error if start index exceed string`)
	}

	result, err = UTF8Slice(str, 0, -1)
	expectErrorMsg = "Length invalid"
	if result != "" || err == nil || err.Error() != expectErrorMsg {
		t.Errorf(`UTF8Slice should return error if length < 0`)
	}

	result, err = UTF8Slice(str, 4, 2)
	expect = "o世"
	if result != expect || err != nil {
		t.Errorf(`UTF8Slice should slice string by rune, expect [%v][%v],
but got [%v][%v]`,
			expect, nil, result, err)
	}
}
