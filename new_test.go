package main

// -------------------------------------------------- tests for 'new' command
import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// New command should create layout in current directory
func TestNew(t *testing.T) {
	t.Parallel()

	dir := createTmpFolder(t)

	NewBlogCreator(dir).Create()

	posts := filepath.Join(dir, PostsDirName)
	if _, err := os.Stat(posts); os.IsNotExist(err) {
		t.Fatalf("Posts dir should be created at %s, but not", posts)
	}
	layouts := filepath.Join(dir, LayoutsDirName)
	if _, err := os.Stat(layouts); os.IsNotExist(err) {
		t.Fatalf("Layouts dir should be created at %s, but not", layouts)
	}
	assets := filepath.Join(dir, AssetsDirName)
	if _, err := os.Stat(assets); os.IsNotExist(err) {
		t.Fatalf("Assets dir should be created at %s, but not", assets)
	}

	templates := [...][2]string{
		[2]string{"base.tmpl", BaseTemplateContent},
		[2]string{"post.tmpl", PostTemplateContent},
		[2]string{"index.tmpl", IndexTemplateContent},
		[2]string{"feeds.xml", FeedsXMLContent},
	}
	for _, template := range templates {
		path := filepath.Join(layouts, template[0])

		if _, err := os.Stat(path); os.IsNotExist(err) {
			t.Fatalf("File %s should be created at %s, but not", path, template[0])
		}

		content, _ := ioutil.ReadFile(path)
		if !strings.Contains(string(content), template[1]) {
			t.Fatalf("File %s should contains \n%s\n, but not, content is \n%s\n",
				template[0], template[1], string(content))
		}
	}
}

func TestNewWhenCurrentFolderIsNotEmpty(t *testing.T) {
	t.Parallel()

	dir := createTmpFolder(t)
	os.Mkdir(filepath.Join(dir, "_whatever"), os.ModePerm)

	err := NewBlogCreator(dir).Create()

	assertErrorDueToNonEmptyDir(t, err)
}

func assertErrorDueToNonEmptyDir(t *testing.T, err error) {
	if err == nil {
		t.Fatalf("There should be an error, but no error")
	}
	if !strings.HasSuffix(err.Error(), "is not empty") {
		t.Fatalf("Error should due to non empty dir, but due to %s", err.Error())
	}
}
