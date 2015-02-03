package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
)

func NewBlogBuilder(dir string) BlogBuilder {
	return BlogBuilder{dir, NewPostParser(), map[string]*template.Template{}}
}

// A BlogBuilder that generate static files from posts/layouts
type BlogBuilder struct {
	// where posts, layouts directory exist
	dir string

	// Parser that parse Post from post files
	postParser PostParser

	// map contains compiled template for later use
	templates map[string]*template.Template
}

// Generate static files to specified directory
func (b BlogBuilder) Build(output string) error {
	os.RemoveAll(output)
	os.Mkdir(output, os.ModePerm)

	if err := b.compileTemplates(); err != nil {
		return fmt.Errorf("Fail to compile templates due to %v", err)
	}

	posts := []Post{}
	for _, path := range b.getPostPaths() {
		post := b.postParser.Parse(path)
		posts = append(posts, post)
		err := b.generatePost(post, output)
		if err != nil {
			return fmt.Errorf("Fail to generate post due to %v", err)
		}
	}

	if err := b.generateBlogIndex(posts, output); err != nil {
		return fmt.Errorf("Fail to generate blog index due to %v", err)
	}

	return nil
}

func (b BlogBuilder) getPostPaths() []string {
	postDir := filepath.Join(b.dir, PostsDirName)
	paths, _ := filepath.Glob(postDir + string(os.PathSeparator) + "*")
	return paths
}

func (b BlogBuilder) compileTemplates() error {
	templates := map[string]string{
		"index": IndexTemplateName,
		"post":  PostTemplateName,
	}

	baseTemplatePath := filepath.Join(b.dir, LayoutsDirName, BaseTemplateName)
	for name, filename := range templates {
		templatePath := filepath.Join(b.dir, LayoutsDirName, filename)
		template, err := template.ParseFiles(baseTemplatePath, templatePath)
		if err != nil {
			return fmt.Errorf("Fail to parse template %s due to %v", filename, err)
		}
		b.templates[name] = template
	}
	return nil
}

func (b BlogBuilder) generatePost(post Post, output string) error {
	postDir := filepath.Join(output, Prettify(post.Title()))

	err := os.Mkdir(postDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Unable to create post folder %s", postDir)
	}

	index := filepath.Join(postDir, "index.html")
	file, err := os.Create(index)
	if err != nil {
		return fmt.Errorf("Unable to create file %s", index)
	}

	data := map[string]interface{}{
		"Content": template.HTML(post.HtmlContent()),
	}

	if err := b.templates["post"].Execute(file, data); err != nil {
		return err
	}
	return nil
}

func (b BlogBuilder) generateBlogIndex(posts []Post, output string) error {
	index := filepath.Join(output, "index.html")
	file, err := os.Create(index)
	if err != nil {
		return fmt.Errorf("Unable to create file %s", index)
	}

	data := map[string]interface{}{
		"Posts": posts,
	}

	if err := b.templates["index"].Execute(file, data); err != nil {
		return err
	}
	return nil
}
