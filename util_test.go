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
