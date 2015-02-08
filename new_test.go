package main

// -------------------------------------------------- tests for 'new' command
import (
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
	layouts := filepath.Join(dir, LayoutsDirName)
	assets := filepath.Join(dir, AssetsDirName)
	assertFilePathExist(t, posts)
	assertFilePathExist(t, layouts)
	assertFilePathExist(t, assets)

	templates := [...][2]string{
		[2]string{"base.tmpl", BaseTemplateContent},
		[2]string{"post.tmpl", PostTemplateContent},
		[2]string{"index.tmpl", IndexTemplateContent},
	}
	for _, template := range templates {
		path := filepath.Join(layouts, template[0])
		assertFilePathExist(t, path)
		assertFileContains(t, path, template[1])
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
